/*
Copyright (c) 2021 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

package observer

import (
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"

	cloudv1 "github.com/oceanbase/ob-operator/apis/cloud/v1"
	observerconst "github.com/oceanbase/ob-operator/pkg/controllers/observer/const"
	"github.com/oceanbase/ob-operator/pkg/controllers/observer/core"
	"github.com/oceanbase/ob-operator/pkg/infrastructure/kube"
	"github.com/oceanbase/ob-operator/pkg/kubeclient"
)

var (
	controllerKind = cloudv1.SchemeGroupVersion.WithKind("OBCluster")
)

// Add creates a new Controller and adds it to the Manager with default RBAC.
// The Manager will set fields on the Controller and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	if !kube.DiscoverGVK(controllerKind) {
		return nil
	}
	return add(mgr, newReconciler(mgr))
}

func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	return &core.OBClusterReconciler{
		CRClient: kubeclient.NewClientFromManager(mgr),
		Scheme:   mgr.GetScheme(),
		Recorder: mgr.GetEventRecorderFor(observerconst.ControllerName),
	}
}

// add a new Controller to mgr with r
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New(
		observerconst.ControllerName,
		mgr,
		controller.Options{
			Reconciler:              r,
			MaxConcurrentReconciles: observerconst.ConcurrentReconciles,
		},
	)
	if err != nil {
		klog.Errorln(err)
		return err
	}

	// Watch for changes to OBCluster
	err = c.Watch(
		&source.Kind{Type: &cloudv1.OBCluster{}},
		&handler.EnqueueRequestForObject{},
	)
	if err != nil {
		klog.Errorln(err)
		return err
	}

	// Watch for changes to StatefulApp
	err = c.Watch(
		&source.Kind{Type: &cloudv1.StatefulApp{}},
		&statefulAppEventHandler{
			enqueueHandler: handler.EnqueueRequestForOwner{
				IsController: true,
				OwnerType:    &cloudv1.OBCluster{},
			},
		},
		&statefulAppPredicate{},
	)
	if err != nil {
		klog.Errorln(err)
		return err
	}

	// Watch for changes to tenant
	err = c.Watch(
		&source.Kind{Type: &cloudv1.Tenant{}},
		&tenantEventHandler{
			enqueueHandler: handler.EnqueueRequestForOwner{
				IsController: true,
				OwnerType:    &cloudv1.OBCluster{},
			},
		},
		&tenantPredicate{},
	)
	if err != nil {
		klog.Errorln(err)
		return err
	}

	return nil
}
