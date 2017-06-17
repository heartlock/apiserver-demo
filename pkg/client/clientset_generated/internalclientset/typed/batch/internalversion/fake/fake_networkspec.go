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
	batch "github.com/heartlock/apiserver-demo/pkg/apis/batch"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworkSpecs implements NetworkSpecInterface
type FakeNetworkSpecs struct {
	Fake *FakeBatch
	ns   string
}

var networkspecsResource = schema.GroupVersionResource{Group: "batch", Version: "", Resource: "networkspecs"}

var networkspecsKind = schema.GroupVersionKind{Group: "batch", Version: "", Kind: "NetworkSpec"}

func (c *FakeNetworkSpecs) Create(networkSpec *batch.NetworkSpec) (result *batch.NetworkSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkspecsResource, c.ns, networkSpec), &batch.NetworkSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkSpec), err
}

func (c *FakeNetworkSpecs) Update(networkSpec *batch.NetworkSpec) (result *batch.NetworkSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkspecsResource, c.ns, networkSpec), &batch.NetworkSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkSpec), err
}

func (c *FakeNetworkSpecs) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkspecsResource, c.ns, name), &batch.NetworkSpec{})

	return err
}

func (c *FakeNetworkSpecs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkspecsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &batch.NetworkSpecList{})
	return err
}

func (c *FakeNetworkSpecs) Get(name string, options v1.GetOptions) (result *batch.NetworkSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkspecsResource, c.ns, name), &batch.NetworkSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkSpec), err
}

func (c *FakeNetworkSpecs) List(opts v1.ListOptions) (result *batch.NetworkSpecList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkspecsResource, networkspecsKind, c.ns, opts), &batch.NetworkSpecList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkSpecList), err
}

// Watch returns a watch.Interface that watches the requested networkSpecs.
func (c *FakeNetworkSpecs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkspecsResource, c.ns, opts))

}

// Patch applies the patch and returns the patched networkSpec.
func (c *FakeNetworkSpecs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.NetworkSpec, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkspecsResource, c.ns, name, data, subresources...), &batch.NetworkSpec{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkSpec), err
}
