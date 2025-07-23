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

package controllers

import (
	gocontext "context"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"
	kubevirtv1 "kubevirt.io/api/core/v1"
	infrav1 "sigs.k8s.io/cluster-api-provider-kubevirt/api/v1alpha1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/cluster-api/util/predicates"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// KubevirtVirtualMachineReconciler reconciles a KubeVirtMachine object
type KubevirtVirtualMachineReconciler struct {
	client.Client
	Log logr.Logger
}

func (r *KubevirtVirtualMachineReconciler) SetupWithManager(goctx gocontext.Context, mgr ctrl.Manager, options controller.Options) error {
	return ctrl.NewControllerManagedBy(mgr).
		WithOptions(options).
		WithEventFilter(predicates.ResourceNotPaused(r.Scheme(), ctrl.LoggerFrom(goctx))).
		For(&kubevirtv1.VirtualMachine{}).
		Watches( // at least one for is needed, but perhaps we could skip
			// builder as per: https://github.com/kubernetes-sigs/controller-runtime/issues/1330
			&kubevirtv1.VirtualMachine{},
			handler.EnqueueRequestsFromMapFunc(r.FilterVirtualMachines)). // Owns equivalent
		Complete(r)
}

// FilterVirtualMachines is a handler.ToRequestsFunc to be used to enqueue requests for reconciliation
// of KubeVirt VirtualMachines, but only enqueues requests for VMs created by CAPK.
func (r *KubevirtVirtualMachineReconciler) FilterVirtualMachines(goctx gocontext.Context, o client.Object) []reconcile.Request {
	result := []reconcile.Request{}

	kvVirtualMachine, ok := o.(*kubevirtv1.VirtualMachine)
	if !ok {
		r.Log.Error(errors.Errorf("expected a VirtualMachine but got a %T", o), "failed to get VirtualMachine")
		return nil
	}

	_, exists := kvVirtualMachine.Labels[clusterv1.ClusterNameLabel]
	if !exists {
		// Not CAPK created VM
		return result
	}

	namespacedName := client.ObjectKey{Namespace: kvVirtualMachine.Namespace, Name: kvVirtualMachine.Name}
	result = append(result, reconcile.Request{NamespacedName: namespacedName})

	return result
}

// TODO cleanup
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=kubevirtmachines,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=infrastructure.cluster.x-k8s.io,resources=kubevirtmachines/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="kubevirt.io",resources=virtualmachines,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups="kubevirt.io",resources=virtualmachines/status;,verbs=get

func (r *KubevirtVirtualMachineReconciler) Reconcile(goctx gocontext.Context, req ctrl.Request) (_ ctrl.Result, reterr error) {
	log := ctrl.LoggerFrom(goctx)

	// Fetch the KubeVirt VirtualMachine
	kvVirtualMachine := &kubevirtv1.VirtualMachine{}
	err := r.Client.Get(goctx, req.NamespacedName, kvVirtualMachine)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return reconcile.Result{}, nil
		}
		return reconcile.Result{}, err // TODO ignore no permission failures
	}

	deleteRequested := !kvVirtualMachine.ObjectMeta.DeletionTimestamp.IsZero()

	capkvKubevirtMachine := &infrav1.KubevirtMachine{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: kvVirtualMachine.Labels[infrav1.KubevirtMachineNamespaceLabel],
			Name:      kvVirtualMachine.Labels[infrav1.KubevirtMachineNameLabel],
		},
	}

	log.Info("Emitting GenericEvent to machineControllerChan",
		"namespace", capkvKubevirtMachine.Namespace,
		"name", capkvKubevirtMachine.Name,
	)

	machineControllerChan <- event.GenericEvent{Object: capkvKubevirtMachine}

	if deleteRequested && controllerutil.ContainsFinalizer(kvVirtualMachine, infrav1.MachineFinalizer) {
		log.Info("Handling delete request of Virtual Machine")
		// At this point we have successfully updated kubevirtmachine
		// controller about the deletion, so we can allow deletion to continue
		controllerutil.RemoveFinalizer(kvVirtualMachine, infrav1.MachineFinalizer)
		err = r.Client.Update(goctx, kvVirtualMachine)
		if err != nil {
			// if it is not found, we can ignore it
			err2 := r.Client.Get(goctx, req.NamespacedName, kvVirtualMachine)
			if err2 != nil && apierrors.IsNotFound(err2) {
				return reconcile.Result{}, nil
			}
		}
		return reconcile.Result{}, err
	}

	return reconcile.Result{}, nil
}
