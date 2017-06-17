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

// This file was automatically generated by lister-gen

package internalversion

import (
	batch "github.com/heartlock/apiserver-demo/pkg/apis/batch"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkStatusLister helps list NetworkStatuses.
type NetworkStatusLister interface {
	// List lists all NetworkStatuses in the indexer.
	List(selector labels.Selector) (ret []*batch.NetworkStatus, err error)
	// NetworkStatuses returns an object that can list and get NetworkStatuses.
	NetworkStatuses(namespace string) NetworkStatusNamespaceLister
	NetworkStatusListerExpansion
}

// networkStatusLister implements the NetworkStatusLister interface.
type networkStatusLister struct {
	indexer cache.Indexer
}

// NewNetworkStatusLister returns a new NetworkStatusLister.
func NewNetworkStatusLister(indexer cache.Indexer) NetworkStatusLister {
	return &networkStatusLister{indexer: indexer}
}

// List lists all NetworkStatuses in the indexer.
func (s *networkStatusLister) List(selector labels.Selector) (ret []*batch.NetworkStatus, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*batch.NetworkStatus))
	})
	return ret, err
}

// NetworkStatuses returns an object that can list and get NetworkStatuses.
func (s *networkStatusLister) NetworkStatuses(namespace string) NetworkStatusNamespaceLister {
	return networkStatusNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkStatusNamespaceLister helps list and get NetworkStatuses.
type NetworkStatusNamespaceLister interface {
	// List lists all NetworkStatuses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*batch.NetworkStatus, err error)
	// Get retrieves the NetworkStatus from the indexer for a given namespace and name.
	Get(name string) (*batch.NetworkStatus, error)
	NetworkStatusNamespaceListerExpansion
}

// networkStatusNamespaceLister implements the NetworkStatusNamespaceLister
// interface.
type networkStatusNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkStatuses in the indexer for a given namespace.
func (s networkStatusNamespaceLister) List(selector labels.Selector) (ret []*batch.NetworkStatus, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*batch.NetworkStatus))
	})
	return ret, err
}

// Get retrieves the NetworkStatus from the indexer for a given namespace and name.
func (s networkStatusNamespaceLister) Get(name string) (*batch.NetworkStatus, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(batch.Resource("networkstatus"), name)
	}
	return obj.(*batch.NetworkStatus), nil
}
