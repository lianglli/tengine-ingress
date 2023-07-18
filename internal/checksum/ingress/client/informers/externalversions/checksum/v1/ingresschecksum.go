/*
Copyright 2021 The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	checksumv1 "k8s.io/ingress-nginx/internal/checksum/ingress/apis/checksum/v1"
	versioned "k8s.io/ingress-nginx/internal/checksum/ingress/client/clientset/versioned"
	internalinterfaces "k8s.io/ingress-nginx/internal/checksum/ingress/client/informers/externalversions/internalinterfaces"
	v1 "k8s.io/ingress-nginx/internal/checksum/ingress/client/listers/checksum/v1"
)

// IngressCheckSumInformer provides access to a shared informer and lister for
// IngressCheckSums.
type IngressCheckSumInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.IngressCheckSumLister
}

type ingressCheckSumInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewIngressCheckSumInformer constructs a new informer for IngressCheckSum type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewIngressCheckSumInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredIngressCheckSumInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredIngressCheckSumInformer constructs a new informer for IngressCheckSum type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredIngressCheckSumInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TengineV1().IngressCheckSums(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TengineV1().IngressCheckSums(namespace).Watch(context.TODO(), options)
			},
		},
		&checksumv1.IngressCheckSum{},
		resyncPeriod,
		indexers,
	)
}

func (f *ingressCheckSumInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredIngressCheckSumInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ingressCheckSumInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&checksumv1.IngressCheckSum{}, f.defaultInformer)
}

func (f *ingressCheckSumInformer) Lister() v1.IngressCheckSumLister {
	return v1.NewIngressCheckSumLister(f.Informer().GetIndexer())
}
