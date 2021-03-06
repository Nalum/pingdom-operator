package main

import (
	"fmt"
	"log"
	"time"

	"github.com/golang/glog"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/scheme"
	typedcorev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/workqueue"

	pingdomV1Alpha1 "github.com/nalum/pingdom-operator/pkg/apis/pingdom/v1alpha1"
	clientset "github.com/nalum/pingdom-operator/pkg/client/clientset/versioned"
	pingdomScheme "github.com/nalum/pingdom-operator/pkg/client/clientset/versioned/scheme"
	informers "github.com/nalum/pingdom-operator/pkg/client/informers/externalversions"
	listers "github.com/nalum/pingdom-operator/pkg/client/listers/pingdom/v1alpha1"
	"github.com/nalum/pingdom-operator/pkg/pingdomclient"
)

const controllerAgentName = "pingdom-controller"

const (
	// SuccessSynced is used as part of the Event 'reason' when a Resource is synced
	SuccessSynced = "Synced"
	// ErrResourceExists is used as part of the Event 'reason' when a Resource fails
	// to sync due to an already existing Pingdom Resource
	ErrResourceExists = "ErrResourceExists"

	// MessageResourceExists is the message used for Events when a resource
	// fails to sync due to a Deployment already existing
	MessageResourceExists = "Resource %q already exists and is not managed by HTTPCheck"
	// MessageResourceSynced is the message used for an Event fired when a HTTPCheck
	// is synced successfully
	MessageResourceSynced = "HTTPCheck synced successfully"
)

// Controller is the controller implementation for HTTPCheck resources
type Controller struct {
	// kubeClientSet is a standard kubernetes clientset
	kubeClientSet kubernetes.Interface
	// pingdomClientSet is a clientset for our own API group
	pingdomClientSet clientset.Interface

	pingdomLister  listers.HTTPCheckLister
	pingdomsSynced cache.InformerSynced

	// workqueue is a rate limited work queue. This is used to queue work to be
	// processed instead of performing it as soon as a change happens. This
	// means we can ensure we only process a fixed amount of resources at a
	// time, and makes it easy to ensure we are never processing the same item
	// simultaneously in two different workers.
	workqueue workqueue.RateLimitingInterface
	// recorder is an event recorder for recording Event resources to the
	// Kubernetes API.
	recorder record.EventRecorder
	// pingdomAPIClient is a client that will manage checks in the Pingdom API based
	// on the CRDs on the Kubernetes cluster
	pingdomAPIClient *pingdomclient.Client
}

// NewController returns a new sample controller
func NewController(
	kubeclientset kubernetes.Interface,
	pingdomAPIClient *pingdomclient.Client,
	pingdomclientset clientset.Interface,
	pingdomInformerFactory informers.SharedInformerFactory) *Controller {

	// obtain references to shared index informers for the Deployment and HTTPCheck
	// types.
	pingdomInformer := pingdomInformerFactory.Pingdom().V1alpha1().HTTPChecks()

	// Create event broadcaster
	// Add sample-controller types to the default Kubernetes Scheme so Events can be
	// logged for sample-controller types.
	pingdomScheme.AddToScheme(scheme.Scheme)
	glog.V(4).Info("Creating event broadcaster")
	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartLogging(glog.Infof)
	eventBroadcaster.StartRecordingToSink(&typedcorev1.EventSinkImpl{Interface: kubeclientset.CoreV1().Events("")})
	recorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: controllerAgentName})

	controller := &Controller{
		kubeClientSet:    kubeclientset,
		pingdomClientSet: pingdomclientset,
		pingdomLister:    pingdomInformer.Lister(),
		pingdomsSynced:   pingdomInformer.Informer().HasSynced,
		workqueue:        workqueue.NewNamedRateLimitingQueue(workqueue.DefaultControllerRateLimiter(), "Pingdom Resources"),
		recorder:         recorder,
		pingdomAPIClient: pingdomAPIClient,
	}

	glog.Info("Setting up event handlers")
	// Set up an event handler for when HTTPCheck resources change
	pingdomInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: controller.enqueueHTTPCheck,
		UpdateFunc: func(old, new interface{}) {
			controller.enqueueHTTPCheck(new)
		},
		DeleteFunc: func(obj interface{}) {
			controller.deleteHTTPCheck(obj)
		},
	})

	return controller
}

// Run will set up the event handlers for types we are interested in, as well
// as syncing informer caches and starting workers. It will block until stopCh
// is closed, at which point it will shutdown the workqueue and wait for
// workers to finish processing their current work items.
func (c *Controller) Run(threadiness int, stopCh <-chan struct{}) error {
	defer runtime.HandleCrash()
	defer c.workqueue.ShutDown()

	// Start the informer factories to begin populating the informer caches
	glog.Info("Starting Pingdom Resource controller")

	// Wait for the caches to be synced before starting workers
	glog.Info("Waiting for informer caches to sync")
	if ok := cache.WaitForCacheSync(stopCh, c.pingdomsSynced); !ok {
		return fmt.Errorf("failed to wait for caches to sync")
	}

	glog.Info("Starting workers")
	// Launch two workers to process HTTPCheck resources
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}

	glog.Info("Started workers")
	<-stopCh
	glog.Info("Shutting down workers")

	return nil
}

// runWorker is a long-running function that will continually call the
// processNextWorkItem function in order to read and process a message on the
// workqueue.
func (c *Controller) runWorker() {
	for c.processNextWorkItem() {
	}
}

// processNextWorkItem will read a single work item off the workqueue and
// attempt to process it, by calling the syncHandler.
func (c *Controller) processNextWorkItem() bool {
	obj, shutdown := c.workqueue.Get()

	if shutdown {
		return false
	}

	// We wrap this block in a func so we can defer c.workqueue.Done.
	err := func(obj interface{}) error {
		// We call Done here so the workqueue knows we have finished
		// processing this item. We also must remember to call Forget if we
		// do not want this work item being re-queued. For example, we do
		// not call Forget if a transient error occurs, instead the item is
		// put back on the workqueue and attempted again after a back-off
		// period.
		defer c.workqueue.Done(obj)
		var key string
		var ok bool

		// We expect strings to come off the workqueue. These are of the
		// form namespace/name. We do this as the delayed nature of the
		// workqueue means the items in the informer cache may actually be
		// more up to date that when the item was initially put onto the
		// workqueue.
		if key, ok = obj.(string); !ok {
			// As the item in the workqueue is actually invalid, we call
			// Forget here else we'd go into a loop of attempting to
			// process a work item that is invalid.
			c.workqueue.Forget(obj)
			runtime.HandleError(fmt.Errorf("expected string in workqueue but got %#v", obj))
			return nil
		}

		// Run the syncHandler, passing it the namespace/name string of the
		// HTTPCheck resource to be synced.
		if err := c.syncHandler(key); err != nil {
			return fmt.Errorf("error syncing '%s': %s", key, err.Error())
		}

		// Finally, if no error occurs we Forget this item so it does not
		// get queued again until another change happens.
		c.workqueue.Forget(obj)
		glog.Infof("Successfully synced '%s'", key)
		return nil
	}(obj)

	if err != nil {
		runtime.HandleError(err)
		return true
	}

	return true
}

// syncHandler compares the actual state with the desired, and attempts to
// converge the two. It then updates the Status block of the HTTPCheck resource
// with the current status of the resource.
func (c *Controller) syncHandler(key string) error {
	// Convert the namespace/name string into a distinct namespace and name
	namespace, name, err := cache.SplitMetaNamespaceKey(key)

	if err != nil {
		runtime.HandleError(fmt.Errorf("invalid resource key: %s", key))
		return nil
	}

	// Get the HTTPCheck resource with this namespace/name
	check, err := c.pingdomLister.HTTPChecks(namespace).Get(name)

	if err != nil {
		// The HTTPCheck resource may no longer exist, in which case we stop
		// processing.
		if errors.IsNotFound(err) {
			runtime.HandleError(fmt.Errorf("HTTPCheck '%s' in work queue no longer exists", key))
			return nil
		}

		return err
	}

	pCheck, err := pingdomclient.NewHTTPCheck(check.Spec.Name, check.Spec.URL)

	if err != nil {
		return err
	}

	pCheck.SetID(check.Status.PingdomID)

	if check.Status.PingdomID != 0 {
		oCheck, err := c.pingdomAPIClient.GetCheck(check.Status.PingdomID)

		// If an error occurs during Get, we'll requeue the item so we can
		// attempt processing again later. This could have been caused by a
		// temporary network failure, or any other transient reason.
		if err != nil {
			return err
		}

		if pCheck.Compare(oCheck) {
			c.recorder.Event(check, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
			return nil
		}

		err = c.pingdomAPIClient.UpdateCheck(pCheck)

		// If an error occurs during Update, we'll requeue the item so we can
		// attempt processing again later. This could have been caused by a
		// temporary network failure, or any other transient reason.
		if err != nil {
			return err
		}
	} else {
		err = c.pingdomAPIClient.CreateCheck(pCheck)

		// If an error occurs during Create, we'll requeue the item so we can
		// attempt processing again later. This could have been caused by a
		// temporary network failure, or any other transient reason.
		if err != nil {
			return err
		}
	}

	// Finally, we update the status block of the HTTPCheck resource to reflect the
	// current state of the world
	err = c.updateHTTPCheckStatus(check, pCheck.GetID())

	if err != nil {
		return err
	}

	c.recorder.Event(check, corev1.EventTypeNormal, SuccessSynced, MessageResourceSynced)
	return nil
}

func (c *Controller) updateHTTPCheckStatus(check *pingdomV1Alpha1.HTTPCheck, checkID int) error {
	// NEVER modify objects from the store. It's a read-only, local cache.
	// You can use DeepCopy() to make a deep copy of original object and modify this copy
	// Or create a copy manually for better performance
	checkCopy := check.DeepCopy()
	checkCopy.Status.PingdomStatus = SuccessSynced
	checkCopy.Status.PingdomID = checkID
	// Until #38113 is merged, we must use Update instead of UpdateStatus to
	// update the Status block of the HTTPCheck resource. UpdateStatus will not
	// allow changes to the Spec of the resource, which is ideal for ensuring
	// nothing other than resource status has been updated.
	_, err := c.pingdomClientSet.Pingdom().HTTPChecks(check.Namespace).Update(checkCopy)
	return err
}

// enqueueHTTPCheck takes a HTTPCheck resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than HTTPCheck.
func (c *Controller) enqueueHTTPCheck(obj interface{}) {
	var key string
	var err error

	if key, err = cache.MetaNamespaceKeyFunc(obj); err != nil {
		runtime.HandleError(err)
		return
	}

	c.workqueue.AddRateLimited(key)
}

// deleteHTTPCheck takes a HTTPCheck resource and converts it into a namespace/name
// string which is then put onto the work queue. This method should *not* be
// passed resources of any type other than HTTPCheck.
func (c *Controller) deleteHTTPCheck(obj interface{}) {
	log.Printf("Deleteing: %#v", obj)
	if httpCheck, ok := obj.(*pingdomV1Alpha1.HTTPCheck); ok {
		check, err := pingdomclient.NewHTTPCheck(httpCheck.Spec.Name, httpCheck.Spec.URL)

		if err != nil {
			runtime.HandleError(err)
			return
		}

		check.SetID(httpCheck.Status.PingdomID)
		c.pingdomAPIClient.DeleteCheck(check)
	} else {
		log.Println("There was an issue type checking the Object")
	}
}
