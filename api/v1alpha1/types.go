/*
Copyright 2020 The Kubernetes Authors.
Portions Copyright © Microsoft Corporation.

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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KubeVirtMachineProviderConditionType is a valid value for KubeVirtMachineProviderCondition.Type
type KubevirtMachineProviderConditionType string

// Valid conditions for an KubeVirt machine instance
const (
	// MachineCreated indicates whether the machine has been created or not. If not,
	// it should include a reason and message for the failure.
	MachineCreated KubevirtMachineProviderConditionType = "MachineCreated"
)

// KubeVirtMachineProviderCondition is a condition in a KubeVirtMachineProviderStatus
type KubevirtMachineProviderCondition struct {
	// Type is the type of the condition.
	Type KubevirtMachineProviderConditionType `json:"type"`
	// Status is the status of the condition.
	Status corev1.ConditionStatus `json:"status"`
	// LastProbeTime is the last time we probed the condition.
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime"`
	// LastTransitionTime is the last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// Reason is a unique, one-word, CamelCase reason for the condition's last transition.
	// +optional
	Reason string `json:"reason"`
	// Message is a human-readable message indicating details about last transition.
	// +optional
	Message string `json:"message"`
}

const (
	// ControlPlane machine label
	ControlPlane string = "control-plane"
	// Node machine label
	Node string = "node"
)

const (
	// Value contains namespace holding KubeVirtMachine CRs
	CapKVMachineNameLabel      = "capkv.cluster.x-k8s.io/kubevirt-machine-name"
	CapKVMachineNamespaceLabel = "capkv.cluster.x-k8s.io/kubevirt-machine-namespace"

	ClusterRoleLabel = "cluster.x-k8s.io/role"
)

// LoadBalancerType describes the type of load balancer to be used, e.g. for
// control plane endpoint
type LoadBalancerType string

// A list of supported load balancer types
const (
	// LoadBalancerKubernetesServiceNodePort indicates using Kubernetes Service
	// resource of type NodePort running on the fabric cluster to load balance
	// pods
	LoadBalancerKubernetesServiceNodePort LoadBalancerType = "KubernetesServiceNodePort"
	// LoadBalancerKubernetesServiceClusterIP indicates using Kubernetes Service
	// resource of type ClusterIP running on the fabric cluster to load balance
	// pods
	LoadBalancerKubernetesServiceClusterIP LoadBalancerType = "KubernetesServiceClusterIP"
	// LoadBalancerKubernetesServiceLoadBalancer indicates using Kubernetes Service
	// resource of type LoadBalancer running on the fabric cluster to load balance
	// pods
	LoadBalancerKubernetesServiceLoadBalancer LoadBalancerType = "KubernetesServiceLoadBalancer"
	// LoadBalancerNone indicates that no load balancer should be automatically
	// deployed and instead load balancer is deployed as part of cluster template
	LoadBalancerNone LoadBalancerType = "None"
)

// VMState describes the state of an KubeVirt virtual machine.
type VMState string

// A list of statuses defined for virtual machines; at the moment taken from kubevirt
const (
	// VirtualMachineStatusStopped indicates that the virtual machine is currently stopped and isn't expected to start.
	VirtualMachineStatusStopped VMState = "Stopped"
	// VirtualMachineStatusProvisioning indicates that cluster resources associated with the virtual machine
	// (e.g., DataVolumes) are being provisioned and prepared.
	VirtualMachineStatusProvisioning VMState = "Provisioning"
	// VirtualMachineStatusStarting indicates that the virtual machine is being prepared for running.
	VirtualMachineStatusStarting VMState = "Starting"
	// VirtualMachineStatusRunning indicates that the virtual machine is running.
	VirtualMachineStatusRunning VMState = "Running"
	// VirtualMachineStatusPaused indicates that the virtual machine is paused.
	VirtualMachineStatusPaused VMState = "Paused"
	// VirtualMachineStatusStopping indicates that the virtual machine is in the process of being stopped.
	VirtualMachineStatusStopping VMState = "Stopping"
	// VirtualMachineStatusTerminating indicates that the virtual machine is in the process of deletion,
	// as well as its associated resources (VirtualMachineInstance, DataVolumes, …).
	VirtualMachineStatusTerminating VMState = "Terminating"
	// VirtualMachineStatusCrashLoopBackOff indicates that the virtual machine is currently in a crash loop waiting to be retried.
	VirtualMachineStatusCrashLoopBackOff VMState = "CrashLoopBackOff"
	// VirtualMachineStatusMigrating indicates that the virtual machine is in the process of being migrated
	// to another host.
	VirtualMachineStatusMigrating VMState = "Migrating"
	// VirtualMachineStatusUnschedulable indicates that the virtual machine could not be scheduled onto a
	// baremetal node, typically due to an error in resources requested.
	VirtualMachineStatusUnschedulable VMState = "ErrorUnschedulable"
	// VirtualMachineStatusUnknown indicates that the state of the virtual machine could not be obtained,
	// typically due to an error in communicating with the host on which it's running.
	VirtualMachineStatusUnknown VMState = "Unknown"
	// VirtualMachineStatusErrImagePull indicates that an error has occurred while pulling an image for
	// a containerDisk VM volume.
	VirtualMachineStatusErrImagePull VMState = "ErrImagePull"
	// VirtualMachineStatusImagePullBackOff indicates that an error has occurred while pulling an image for
	// a containerDisk VM volume, and that kubelet is backing off before retrying.
	VirtualMachineStatusImagePullBackOff VMState = "ImagePullBackOff"
	// VirtualMachineStatusPvcNotFound indicates that the virtual machine references a PVC volume which doesn't exist.
	VirtualMachineStatusPvcNotFound VMState = "ErrorPvcNotFound"
	// VirtualMachineStatusDataVolumeError indicates that an error has been reported by one of the DataVolumes
	// referenced by the virtual machines.
	VirtualMachineStatusDataVolumeError VMState = "DataVolumeError"
	// VirtualMachineStatusWaitingForVolumeBinding indicates that some PersistentVolumeClaims backing
	// the virtual machine volume are still not bound.
	VirtualMachineStatusWaitingForVolumeBinding VMState = "WaitingForVolumeBinding"
)

// Image defines information about the image to use for VM creation.
// There are three ways to specify an image: by ID, by publisher, or by Shared Image Gallery.
// If specifying an image by ID, only the ID field needs to be set.
// If specifying an image by publisher, the Publisher, Offer, SKU, and Version fields must be set.
// If specifying an image from a Shared Image Gallery, the SubscriptionID, ResourceGroup,
// Gallery, Name, and Version fields must be set.
type Image struct {
	Publisher *string `json:"publisher,omitempty"`
	Offer     *string `json:"offer,omitempty"`
	SKU       *string `json:"sku,omitempty"`

	ID *string `json:"id,omitempty"`

	SubscriptionID *string `json:"subscriptionID,omitempty"`
	ResourceGroup  *string `json:"resourceGroup,omitempty"`
	Gallery        *string `json:"gallery,omitempty"`
	Name           *string `json:"name,omitempty"`

	Version *string `json:"version,omitempty"`
	OSType  OSType  `json:"osType"`
}

// OSType describes the OS type of a disk.
type OSType string

const (
	// OSTypeLinux
	OSTypeLinux = OSType("Linux")
	// OSTypeWindows
	OSTypeWindows = OSType("Windows")
)

const (
	AnnotationClusterInfrastructureReady = "kubevirt.cluster.sigs.k8s.io/infrastructure-ready"
	ValueReady                           = "true"
	AnnotationControlPlaneReady          = "kubevirt.cluster.sigs.k8s.io/control-plane-ready"
)
