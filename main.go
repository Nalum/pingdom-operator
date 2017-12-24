package main

import (
	"flag"
	"time"

	"github.com/golang/glog"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	// Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"

	clientset "github.com/nalum/pingdom-operator/pkg/client/clientset/versioned"
	informers "github.com/nalum/pingdom-operator/pkg/client/informers/externalversions"
	"k8s.io/sample-controller/pkg/signals"
)

var (
	masterURL   string
	kubeconfig  string
	pingdomUser string
	pingdomPass string
)

func main() {
	flag.Parse()

	if pingdomPass == "" || pingdomUser == "" {
		glog.Exitln("You must provide the user details so that we can authenticate against the Pingdom API.")
	}

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		glog.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	pingdomCheckclientset, err := clientset.NewForConfig(cfg)
	if err != nil {
		glog.Fatalf("Error building example clientset: %s", err.Error())
	}

	pingdomCheckInformerFactory := informers.NewSharedInformerFactory(pingdomCheckclientset, time.Second*30)

	controller := NewController(kubeClient, pingdomCheckclientset, pingdomCheckInformerFactory)

	go pingdomCheckInformerFactory.Start(stopCh)

	if err = controller.Run(2, stopCh); err != nil {
		glog.Fatalf("Error running controller: %s", err.Error())
	}
}

func init() {
	flag.StringVar(&kubeconfig, "kubeconfig", "", "Path to a kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&masterURL, "master", "", "The address of the Kubernetes API server. Overrides any value in kubeconfig. Only required if out-of-cluster.")
	flag.StringVar(&pingdomUser, "pingdom-user", "", "The user used to authenticate against the Pingdom API. Required")
	flag.StringVar(&pingdomPass, "pingdom-password", "", "The password used to authenticate against the Pingdom API. Required")
}
