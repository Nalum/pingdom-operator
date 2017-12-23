/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	v1alpha1 "github.com/nalum/pingdom-operator/pkg/apis/pingdomcheck/v1alpha1"
	scheme "github.com/nalum/pingdom-operator/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// HTTPChecksGetter has a method to return a HTTPCheckInterface.
// A group's client should implement this interface.
type HTTPChecksGetter interface {
	HTTPChecks(namespace string) HTTPCheckInterface
}

// HTTPCheckInterface has methods to work with HTTPCheck resources.
type HTTPCheckInterface interface {
	Create(*v1alpha1.HTTPCheck) (*v1alpha1.HTTPCheck, error)
	Update(*v1alpha1.HTTPCheck) (*v1alpha1.HTTPCheck, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.HTTPCheck, error)
	List(opts v1.ListOptions) (*v1alpha1.HTTPCheckList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HTTPCheck, err error)
	HTTPCheckExpansion
}

// hTTPChecks implements HTTPCheckInterface
type hTTPChecks struct {
	client rest.Interface
	ns     string
}

// newHTTPChecks returns a HTTPChecks
func newHTTPChecks(c *PingdomcheckV1alpha1Client, namespace string) *hTTPChecks {
	return &hTTPChecks{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the hTTPCheck, and returns the corresponding hTTPCheck object, and an error if there is any.
func (c *hTTPChecks) Get(name string, options v1.GetOptions) (result *v1alpha1.HTTPCheck, err error) {
	result = &v1alpha1.HTTPCheck{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpchecks").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HTTPChecks that match those selectors.
func (c *hTTPChecks) List(opts v1.ListOptions) (result *v1alpha1.HTTPCheckList, err error) {
	result = &v1alpha1.HTTPCheckList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("httpchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested hTTPChecks.
func (c *hTTPChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("httpchecks").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a hTTPCheck and creates it.  Returns the server's representation of the hTTPCheck, and an error, if there is any.
func (c *hTTPChecks) Create(hTTPCheck *v1alpha1.HTTPCheck) (result *v1alpha1.HTTPCheck, err error) {
	result = &v1alpha1.HTTPCheck{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("httpchecks").
		Body(hTTPCheck).
		Do().
		Into(result)
	return
}

// Update takes the representation of a hTTPCheck and updates it. Returns the server's representation of the hTTPCheck, and an error, if there is any.
func (c *hTTPChecks) Update(hTTPCheck *v1alpha1.HTTPCheck) (result *v1alpha1.HTTPCheck, err error) {
	result = &v1alpha1.HTTPCheck{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("httpchecks").
		Name(hTTPCheck.Name).
		Body(hTTPCheck).
		Do().
		Into(result)
	return
}

// Delete takes name of the hTTPCheck and deletes it. Returns an error if one occurs.
func (c *hTTPChecks) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpchecks").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *hTTPChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("httpchecks").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched hTTPCheck.
func (c *hTTPChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HTTPCheck, err error) {
	result = &v1alpha1.HTTPCheck{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("httpchecks").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}