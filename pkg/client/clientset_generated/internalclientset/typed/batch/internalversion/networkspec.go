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

// NetworkSpecsGetter has a method to return a NetworkSpecInterface.
// A group's client should implement this interface.
type NetworkSpecsGetter interface {
	NetworkSpecs(namespace string) NetworkSpecInterface
}

// NetworkSpecInterface has methods to work with NetworkSpec resources.
type NetworkSpecInterface interface {
	Create(*batch.NetworkSpec) (*batch.NetworkSpec, error)
	Update(*batch.NetworkSpec) (*batch.NetworkSpec, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*batch.NetworkSpec, error)
	List(opts v1.ListOptions) (*batch.NetworkSpecList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.NetworkSpec, err error)
	NetworkSpecExpansion
}

// networkSpecs implements NetworkSpecInterface
type networkSpecs struct {
	client rest.Interface
	ns     string
}

// newNetworkSpecs returns a NetworkSpecs
func newNetworkSpecs(c *BatchClient, namespace string) *networkSpecs {
	return &networkSpecs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Create takes the representation of a networkSpec and creates it.  Returns the server's representation of the networkSpec, and an error, if there is any.
func (c *networkSpecs) Create(networkSpec *batch.NetworkSpec) (result *batch.NetworkSpec, err error) {
	result = &batch.NetworkSpec{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("networkspecs").
		Body(networkSpec).
		Do().
		Into(result)
	return
}

// Update takes the representation of a networkSpec and updates it. Returns the server's representation of the networkSpec, and an error, if there is any.
func (c *networkSpecs) Update(networkSpec *batch.NetworkSpec) (result *batch.NetworkSpec, err error) {
	result = &batch.NetworkSpec{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("networkspecs").
		Name(networkSpec.Name).
		Body(networkSpec).
		Do().
		Into(result)
	return
}

// Delete takes name of the networkSpec and deletes it. Returns an error if one occurs.
func (c *networkSpecs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkspecs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *networkSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("networkspecs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Get takes name of the networkSpec, and returns the corresponding networkSpec object, and an error if there is any.
func (c *networkSpecs) Get(name string, options v1.GetOptions) (result *batch.NetworkSpec, err error) {
	result = &batch.NetworkSpec{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkspecs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of NetworkSpecs that match those selectors.
func (c *networkSpecs) List(opts v1.ListOptions) (result *batch.NetworkSpecList, err error) {
	result = &batch.NetworkSpecList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("networkspecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested networkSpecs.
func (c *networkSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("networkspecs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Patch applies the patch and returns the patched networkSpec.
func (c *networkSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.NetworkSpec, err error) {
	result = &batch.NetworkSpec{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("networkspecs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
