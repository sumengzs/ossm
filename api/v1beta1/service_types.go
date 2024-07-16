/*
Copyright 2024 OSSM Author.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ServiceSpec defines the desired state of Service
type ServiceSpec struct {
	// External Represents the configuration of the object storage service for external network access
	//
	// +kubebuilder:validation:Required
	External ServiceConnectConfig `json:"external"`

	// Internal Represents the configuration of the object storage service for internal network access
	Internal *ServiceConnectConfig `json:"internal,omitempty"`

	// Disabled indicates whether the service is disabled
	//
	// If it is Disabled, user cannot do any operate for the service util it is enabled again.
	Disabled bool `json:"disabled,omitempty"`

	// HealthCheckInterval Indicates the interval of health check
	HealthCheckInterval int64 `json:"healthCheckInterval,omitempty"`

	// BucketUsageSyncInterval Indicates the interval of bucket usage synchronization
	BucketUsageSyncInterval int64 `json:"bucketUsageSyncInterval,omitempty"`
}

type ServiceConnectConfig struct {
	// Endpoint Indicates the access address of the service
	//
	// Does not include protocols, such as: s 3, http, https, etc.
	// it consists only of host:port. such as: 192.168.0.1:8080
	// +kubebuilder:validation:Required
	Endpoint string `json:"endpoint,omitempty"`

	// Secure indicates whether the service is accessed through a secure protocol
	Secure bool `json:"secure,omitempty"`

	// AccessKey indicates the access key of the service
	//
	// Together with SecretKey, it forms an access credential for accessing services.
	// it and Token can not be empty as the same time.
	AccessKey string `json:"accessKey"`

	// SecretKey indicates the secret key of the service
	//
	// It cannot be empty when AccessKey is not empty.
	SecretKey string `json:"secretKey"`

	// Token is the credential used to access the service
	//
	// It and AccessKey can not be empty as the same time.
	Token string `json:"token"`
}

// ServiceStatus defines the observed state of Service
type ServiceStatus struct {
	Phase      string             `json:"phase,omitempty"`
	Online     bool               `json:"online,omitempty"`
	Message    string             `json:"message,omitempty"`
	Summary    Summary            `json:"summary,omitempty"`
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type Summary struct {
	Usage       int64 `json:"usage,omitempty"`
	BucketTotal int64 `json:"bucketTotal,omitempty"`
	ObjectTotal int64 `json:"objectTotal,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// Service is the Schema for the services API
type Service struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ServiceSpec   `json:"spec,omitempty"`
	Status ServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServiceList contains a list of Service
type ServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Service `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Service{}, &ServiceList{})
}
