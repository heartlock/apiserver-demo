
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


package network

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/controller"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/util/workqueue"

	"github.com/heartlock/apiserver-demo/pkg/apis/batch/v1beta1"
	"github.com/heartlock/apiserver-demo/pkg/controller/sharedinformers"
	listers "github.com/heartlock/apiserver-demo/pkg/client/listers_generated/batch/v1beta1"
)

// +controller:group=batch,version=v1beta1,kind=Network,resource=networks
type NetworkControllerImpl struct {
	// informer listens for events about Network
	informer cache.SharedIndexInformer

	// lister indexes properties about Network
	lister listers.NetworkLister
}

// Init initializes the controller and is called by the generated code
// Registers eventhandlers to enqueue events
// config - client configuration for talking to the apiserver
// si - informer factory shared across all controllers for listening to events and indexing resource properties
// queue - message queue for handling new events.  unique to this controller.
func (c *NetworkControllerImpl) Init(
	config *rest.Config,
	si *sharedinformers.SharedInformers,
	queue workqueue.RateLimitingInterface) {

	// Set the informer and lister for subscribing to events and indexing networks labels
	i := si.Factory.Batch().V1beta1().Networks()
	c.informer = i.Informer()
	c.lister = i.Lister()

	// Add an event handler to enqueue a message for networks adds / updates
	c.informer.AddEventHandler(&controller.QueueingEventHandler{queue})
}

// Reconcile handles enqueued messages
func (c *NetworkControllerImpl) Reconcile(u *v1beta1.Network) error {
	// Implement controller logic here
	log.Printf("Running reconcile Network for %s\n", u.Name)
	return nil
}

func (c *NetworkControllerImpl) Get(namespace, name string) (*v1beta1.Network, error) {
	return c.lister.Networks(namespace).Get(name)
}
