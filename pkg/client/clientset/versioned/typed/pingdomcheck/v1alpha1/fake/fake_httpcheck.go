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

package fake

import (
	v1alpha1 "github.com/nalum/pingdom-operator/pkg/apis/pingdomcheck/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeHTTPChecks implements HTTPCheckInterface
type FakeHTTPChecks struct {
	Fake *FakePingdomcheckV1alpha1
	ns   string
}

var httpchecksResource = schema.GroupVersionResource{Group: "pingdomcheck.mallon.io", Version: "v1alpha1", Resource: "httpchecks"}

var httpchecksKind = schema.GroupVersionKind{Group: "pingdomcheck.mallon.io", Version: "v1alpha1", Kind: "HTTPCheck"}

// Get takes name of the hTTPCheck, and returns the corresponding hTTPCheck object, and an error if there is any.
func (c *FakeHTTPChecks) Get(name string, options v1.GetOptions) (result *v1alpha1.HTTPCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(httpchecksResource, c.ns, name), &v1alpha1.HTTPCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPCheck), err
}

// List takes label and field selectors, and returns the list of HTTPChecks that match those selectors.
func (c *FakeHTTPChecks) List(opts v1.ListOptions) (result *v1alpha1.HTTPCheckList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(httpchecksResource, httpchecksKind, c.ns, opts), &v1alpha1.HTTPCheckList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.HTTPCheckList{}
	for _, item := range obj.(*v1alpha1.HTTPCheckList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested hTTPChecks.
func (c *FakeHTTPChecks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(httpchecksResource, c.ns, opts))

}

// Create takes the representation of a hTTPCheck and creates it.  Returns the server's representation of the hTTPCheck, and an error, if there is any.
func (c *FakeHTTPChecks) Create(hTTPCheck *v1alpha1.HTTPCheck) (result *v1alpha1.HTTPCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(httpchecksResource, c.ns, hTTPCheck), &v1alpha1.HTTPCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPCheck), err
}

// Update takes the representation of a hTTPCheck and updates it. Returns the server's representation of the hTTPCheck, and an error, if there is any.
func (c *FakeHTTPChecks) Update(hTTPCheck *v1alpha1.HTTPCheck) (result *v1alpha1.HTTPCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(httpchecksResource, c.ns, hTTPCheck), &v1alpha1.HTTPCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPCheck), err
}

// Delete takes name of the hTTPCheck and deletes it. Returns an error if one occurs.
func (c *FakeHTTPChecks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(httpchecksResource, c.ns, name), &v1alpha1.HTTPCheck{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeHTTPChecks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(httpchecksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.HTTPCheckList{})
	return err
}

// Patch applies the patch and returns the patched hTTPCheck.
func (c *FakeHTTPChecks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.HTTPCheck, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(httpchecksResource, c.ns, name, data, subresources...), &v1alpha1.HTTPCheck{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.HTTPCheck), err
}
