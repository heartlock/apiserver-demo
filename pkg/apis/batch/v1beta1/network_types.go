
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


package v1beta1

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/heartlock/apiserver-demo/pkg/apis/batch"
)

// +genclient=true

// Network
// +k8s:openapi-gen=true
// +resource:path=networks,strategy=NetworkStrategy
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSpec   `json:"spec,omitempty"`
	Status NetworkStatus `json:"status,omitempty"`
}

// NetworkSpec defines the desired state of Network
type NetworkSpec struct {
}

// NetworkStatus defines the observed state of Network
type NetworkStatus struct {
}

// Validate checks that an instance of Network is well formed
func (NetworkStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*batch.Network)
	log.Printf("Validating fields for Network %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	return errors
}

// DefaultingFunction sets default Network field values
func (NetworkSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Network)
	// set default field values here
	log.Printf("Defaulting fields for Network %s\n", obj.Name)
}
