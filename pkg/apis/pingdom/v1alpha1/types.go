package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPCheck is a specification for a HTTPCheck resource
type HTTPCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HTTPCheckSpec   `json:"spec"`
	Status HTTPCheckStatus `json:"status"`
}

// HTTPCheckSpec is the spec for a HTTPCheck resource
type HTTPCheckSpec struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// HTTPCheckStatus is the status for a HTTPCheck resource
type HTTPCheckStatus struct {
	PingdomStatus string `json:"pingdomStatus"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HTTPCheckList is a list of HTTPCheck resources
type HTTPCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []HTTPCheck `json:"items"`
}
