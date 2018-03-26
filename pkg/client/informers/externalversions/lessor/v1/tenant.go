// This file was automatically generated by informer-gen

package v1

import (
	lessor_io_v1 "github.com/lessor/lessor/pkg/apis/lessor.io/v1"
	versioned "github.com/lessor/lessor/pkg/client/clientset/versioned"
	internalinterfaces "github.com/lessor/lessor/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/lessor/lessor/pkg/client/listers/lessor/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	time "time"
)

// TenantInformer provides access to a shared informer and lister for
// Tenants.
type TenantInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TenantLister
}

type tenantInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewTenantInformer constructs a new informer for Tenant type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTenantInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options meta_v1.ListOptions) (runtime.Object, error) {
				return client.LessorV1().Tenants(namespace).List(options)
			},
			WatchFunc: func(options meta_v1.ListOptions) (watch.Interface, error) {
				return client.LessorV1().Tenants(namespace).Watch(options)
			},
		},
		&lessor_io_v1.Tenant{},
		resyncPeriod,
		indexers,
	)
}

func defaultTenantInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewTenantInformer(client, meta_v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *tenantInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&lessor_io_v1.Tenant{}, defaultTenantInformer)
}

func (f *tenantInformer) Lister() v1.TenantLister {
	return v1.NewTenantLister(f.Informer().GetIndexer())
}
