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

type KubevirtClusterPhase string

const (
	// KubeVirtClusterPhasePending is the first state a Cluster is assigned by
	// Cluster API Cluster controller after being created.
	KubeVirtClusterPhasePending = KubevirtClusterPhase("pending")

	// KubeVirtClusterPhaseProvisioning is the state when the Cluster has a provider infrastructure
	// object associated and can start provisioning.
	KubeVirtClusterPhaseProvisioning = KubevirtClusterPhase("provisioning")

	// KubeVirtClusterPhaseProvisioned is the state when its
	// infrastructure has been created and configured.
	KubeVirtClusterPhaseProvisioned = KubevirtClusterPhase("provisioned")

	// KubeVirtClusterPhaseDeleting is the Cluster state when a delete
	// request has been sent to the API Server,
	// but its infrastructure has not yet been fully deleted.
	KubeVirtClusterPhaseDeleting = KubevirtClusterPhase("deleting")

	// KubeVirtClusterPhaseFailed is the Cluster state when the system
	// might require user intervention.
	KubeVirtClusterPhaseFailed = KubevirtClusterPhase("failed")

	// KubeVirtClusterPhaseUpgrading is the Cluster state when the system
	// is in the middle of a update.
	KubeVirtClusterPhaseUpgrading = KubevirtClusterPhase("upgrading")

	// KubeVirtClusterPhaseUnknown is returned if the Cluster state cannot be determined.
	KubeVirtClusterPhaseUnknown = KubevirtClusterPhase("")
)

type KubevirtLoadBalancerPhase string

const (
	// KubeVirtLoadBalancerPhasePending is the first state a LoadBalancer is assigned by
	// the controller after being created.
	KubeVirtLoadBalancerPhasePending = KubevirtLoadBalancerPhase("pending")

	// KubeVirtLoadBalancerPhaseProvisioning is the state when the LoadBalancer is waiting for the
	// first replica to be ready.
	KubeVirtLoadBalancerPhaseProvisioning = KubevirtLoadBalancerPhase("provisioning")

	// KubeVirtLoadBalancerPhaseProvisioned is the state when its infrastructure has been created
	// and configured. All replicas are ready and we have the desired number of replicas.
	KubeVirtLoadBalancerPhaseProvisioned = KubevirtLoadBalancerPhase("provisioned")

	// KubeVirtLoadBalancerPhaseScaling is the state when replicas are being scaled.
	KubeVirtLoadBalancerPhaseScaling = KubevirtLoadBalancerPhase("scaling")

	// KubeVirtLoadBalancerPhaseUpgrading is the state when the system is in the middle of a update.
	KubeVirtLoadBalancerPhaseUpgrading = KubevirtLoadBalancerPhase("upgrading")

	// KubeVirtLoadBalancerPhaseDeleting is the state when a delete request has been sent to
	// the API Server, but its infrastructure has not yet been fully deleted.
	KubeVirtLoadBalancerPhaseDeleting = KubevirtLoadBalancerPhase("deleting")

	// KubeVirtLoadBalancerPhaseFailed is the state when the system might require user intervention.
	KubeVirtLoadBalancerPhaseFailed = KubevirtLoadBalancerPhase("failed")

	// KubeVirtLoadBalancerPhaseUnknown is returned if the state cannot be determined.
	KubeVirtLoadBalancerPhaseUnknown = KubevirtLoadBalancerPhase("")
)
