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
package internalversion

import (
	batch "github.com/heartlock/apiserver-demo/pkg/apis/batch"
	scheme "github.com/heartlock/apiserver-demo/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// NetworkStatusesGetter has a method to return a NetworkStatusInterface.
// A group's client should implement this interface.
type NetworkStatusesGetter interface {
	NetworkStatuses(namespace string) NetworkStatusInterface
}

// NetworkStatusInterface has methods to work with NetworkStatus resources.
type NetworkStatusInterface interface {
	Create(*batch.NetworkStatus) (*batch.NetworkStatus, error)
	Update(*batch.NetworkStatus) (*batch.NetworkStatus, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*batch.NetworkStatus, error)
	List(opts v1.ListOptions) (*batch.NetworkStatusList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.NetworkStatus, err error)
	NetworkStatusExpansion
}

// networkStatuses implements NetworkStatusInterface
type networkStatuses struct {
	client rest.Interface
	ns     string
}

// newNetworkStatuses returns a NetworkStatuses
func newNetworkStatuses(c *BatchClient, namespace string) *networkStatuses {
	return &networkStatuses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a networkStatus and creates it.  Returns the server's representation of the networkStatus, and an error, if there is any.
func (c *networkStatuses) Create(networkStatus *batch.NetworkStatus) (result *batch.NetworkStatus, err error) {
	result = &batch.NetworkStatus{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkstatuses").
		Body(networkStatus).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkStatus and updates it. Returns the server's representation of the networkStatus, and an error, if there is any.
func (c *networkStatuses) Update(networkStatus *batch.NetworkStatus) (result *batch.NetworkStatus, err error) {
	result = &batch.NetworkStatus{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkstatuses").
		Name(networkStatus.Name).
		Body(networkStatus).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkStatus and deletes it. Returns an error if one occurs.
func (c *networkStatuses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkstatuses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkStatuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkstatuses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the networkStatus, and returns the corresponding networkStatus object, and an error if there is any.
func (c *networkStatuses) Get(name string, options v1.GetOptions) (result *batch.NetworkStatus, err error) {
	result = &batch.NetworkStatus{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkstatuses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkStatuses that match those selectors.
func (c *networkStatuses) List(opts v1.ListOptions) (result *batch.NetworkStatusList, err error) {
	result = &batch.NetworkStatusList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkstatuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkStatuses.
func (c *networkStatuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkstatuses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched networkStatus.
func (c *networkStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.NetworkStatus, err error) {
	result = &batch.NetworkStatus{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkstatuses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
