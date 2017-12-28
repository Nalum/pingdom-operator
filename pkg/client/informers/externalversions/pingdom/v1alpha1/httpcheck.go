/*
Copyright 2017 The Kubernetes sample-controller Authors.

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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	time "time"

	pingdom_v1alpha1 "github.com/nalum/pingdom-operator/pkg/apis/pingdom/v1alpha1"
	versioned "github.com/nalum/pingdom-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nalum/pingdom-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nalum/pingdom-operator/pkg/client/listers/pingdom/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HTTPCheckInformer provides access to a shared informer and lister for
// HTTPChecks.
type HTTPCheckInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.HTTPCheckLister
}

type hTTPCheckInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHTTPCheckInformer constructs a new informer for HTTPCheck type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHTTPCheckInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHTTPCheckInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHTTPCheckInformer constructs a new informer for HTTPCheck type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHTTPCheckInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PingdomV1alpha1().HTTPChecks(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.PingdomV1alpha1().HTTPChecks(namespace).Watch(options)
			},
		},
		&pingdom_v1alpha1.HTTPCheck{},
		resyncPeriod,
		indexers,
	)
}

func (f *hTTPCheckInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHTTPCheckInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *hTTPCheckInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&pingdom_v1alpha1.HTTPCheck{}, f.defaultInformer)
}

func (f *hTTPCheckInformer) Lister() v1alpha1.HTTPCheckLister {
	return v1alpha1.NewHTTPCheckLister(f.Informer().GetIndexer())
}
