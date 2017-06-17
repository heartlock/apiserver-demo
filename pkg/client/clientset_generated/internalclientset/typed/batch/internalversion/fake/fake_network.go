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
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNetworks implements NetworkInterface
type FakeNetworks struct {
	Fake *FakeBatch
	ns   string
}

var networksResource = schema.GroupVersionResource{Group: "batch", Version: "", Resource: "networks"}

var networksKind = schema.GroupVersionKind{Group: "batch", Version: "", Kind: "Network"}

func (c *FakeNetworks) Create(network *batch.Network) (result *batch.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networksResource, c.ns, network), &batch.Network{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.Network), err
}

func (c *FakeNetworks) Update(network *batch.Network) (result *batch.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networksResource, c.ns, network), &batch.Network{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.Network), err
}

func (c *FakeNetworks) UpdateStatus(network *batch.Network) (*batch.Network, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(networksResource, "status", c.ns, network), &batch.Network{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.Network), err
}

func (c *FakeNetworks) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networksResource, c.ns, name), &batch.Network{})

	return err
}

func (c *FakeNetworks) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networksResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &batch.NetworkList{})
	return err
}

func (c *FakeNetworks) Get(name string, options v1.GetOptions) (result *batch.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networksResource, c.ns, name), &batch.Network{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.Network), err
}

func (c *FakeNetworks) List(opts v1.ListOptions) (result *batch.NetworkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networksResource, networksKind, c.ns, opts), &batch.NetworkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &batch.NetworkList{}
	for _, item := range obj.(*batch.NetworkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested networks.
func (c *FakeNetworks) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networksResource, c.ns, opts))

}

// Patch applies the patch and returns the patched network.
func (c *FakeNetworks) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.Network, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networksResource, c.ns, name, data, subresources...), &batch.Network{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.Network), err
}
