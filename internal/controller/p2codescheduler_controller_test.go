/*
Copyright 2024.

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

package controller

import (
	"context"
	"fmt"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	cachev1alpha1 "github.com/PoolPooer/p2code-scheduler/api/v1alpha1"
)

type P2codeSchedulerReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// Reconcile watches P2codeScheduler resources and reacts when they are created or updated
func (r *P2codeSchedulerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	// Fetch the P2codeScheduler resource
	var scheduler cachev1alpha1.P2codeScheduler
	if err := r.Get(ctx, req.NamespacedName, &scheduler); err != nil {
		logger.Error(err, "unable to fetch P2codeScheduler")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Log a simple "Hello World" message
	fmt.Println("Hello World: Detected P2codeScheduler resource", "name", scheduler.Name)

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *P2codeSchedulerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&cachev1alpha1.P2codeScheduler{}).
		Complete(r)
}
