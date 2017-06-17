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

// FakeNetworkStatuses implements NetworkStatusInterface
type FakeNetworkStatuses struct {
	Fake *FakeBatch
	ns   string
}

var networkstatusesResource = schema.GroupVersionResource{Group: "batch", Version: "", Resource: "networkstatuses"}

var networkstatusesKind = schema.GroupVersionKind{Group: "batch", Version: "", Kind: "NetworkStatus"}

func (c *FakeNetworkStatuses) Create(networkStatus *batch.NetworkStatus) (result *batch.NetworkStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(networkstatusesResource, c.ns, networkStatus), &batch.NetworkStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkStatus), err
}

func (c *FakeNetworkStatuses) Update(networkStatus *batch.NetworkStatus) (result *batch.NetworkStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(networkstatusesResource, c.ns, networkStatus), &batch.NetworkStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkStatus), err
}

func (c *FakeNetworkStatuses) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(networkstatusesResource, c.ns, name), &batch.NetworkStatus{})

	return err
}

func (c *FakeNetworkStatuses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(networkstatusesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &batch.NetworkStatusList{})
	return err
}

func (c *FakeNetworkStatuses) Get(name string, options v1.GetOptions) (result *batch.NetworkStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(networkstatusesResource, c.ns, name), &batch.NetworkStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkStatus), err
}

func (c *FakeNetworkStatuses) List(opts v1.ListOptions) (result *batch.NetworkStatusList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(networkstatusesResource, networkstatusesKind, c.ns, opts), &batch.NetworkStatusList{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkStatusList), err
}

// Watch returns a watch.Interface that watches the requested networkStatuses.
func (c *FakeNetworkStatuses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(networkstatusesResource, c.ns, opts))

}

// Patch applies the patch and returns the patched networkStatus.
func (c *FakeNetworkStatuses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.NetworkStatus, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(networkstatusesResource, c.ns, name, data, subresources...), &batch.NetworkStatus{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.NetworkStatus), err
}
