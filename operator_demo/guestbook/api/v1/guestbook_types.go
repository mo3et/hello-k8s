/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GuestbookSpec defines the desired state of Guestbook.
type GuestbookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of Guestbook. Edit guestbook_types.go to remove/update
	// Foo      string `json:"foo,omitempty"`

	ReplicaCount int32  `json:"replicaCount"`
	Replicas     int32  `json:"replicas,omitempty"`
	Image        string `json:"image,omitempty"`

	// Hello string `json:"hello,omitempty"`
	// IsMe  bool   `json:"isMe,omitempty"`

	// Quantity of instances
	// +kubebuilder:validation:Maximum=10
	// +kubebuilder:validation:Minimum=1
	Size int32 `json:"size"`

	// Name of the ConfigMap for GuestbookSpec's configuration
	// +kubebuilder:validation:MaxLength=15
	// +kubebuilder:validation:MinLength=1
	ConfigMapName string `json:"configMapName"`

	// +kubebuilder:validation:Enum=Phone;Address;Name
	Type string `json:"alias,omitempty"`
}

// GuestbookStatus defines the observed state of Guestbook.
type GuestbookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	AvailableReplicas int32 `json:"availableReplicas,omitempty"`
	AvailablePods     int32 `json:"availablePods,omitempty"`
	AvailableServices int32 `json:"availableServices,omitempty"`

	// PodName of the active Guestbook node.
	Active string `json:"active"`

	// PodName of standby Guestbook nodes.
	Standby string `json:"standby"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=cluster

// Guestbook is the Schema for the guestbooks API.
type Guestbook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuestbookSpec   `json:"spec,omitempty"`
	Status GuestbookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuestbookList contains a list of Guestbook.
type GuestbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Guestbook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Guestbook{}, &GuestbookList{})
}
