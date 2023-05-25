/*
Copyright (c) 2023 OceanBase
ob-operator is licensed under Mulan PSL v2.
You can use this software according to the terms and conditions of the Mulan PSL v2.
You may obtain a copy of Mulan PSL v2 at:
         http://license.coscl.org.cn/MulanPSL2
THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
See the Mulan PSL v2 for more details.
*/

package resource

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	kubeerrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	apipod "k8s.io/kubernetes/pkg/api/v1/pod"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/go-logr/logr"
	"github.com/pkg/errors"

	cloudv2alpha1 "github.com/oceanbase/ob-operator/api/v2alpha1"
	clusterstatus "github.com/oceanbase/ob-operator/pkg/const/status/obcluster"
	serverstatus "github.com/oceanbase/ob-operator/pkg/const/status/observer"
	"github.com/oceanbase/ob-operator/pkg/task"
	flowname "github.com/oceanbase/ob-operator/pkg/task/const/flow/name"
	taskname "github.com/oceanbase/ob-operator/pkg/task/const/task/name"
)

type OBServerManager struct {
	ResourceManager
	Ctx      context.Context
	OBServer *cloudv2alpha1.OBServer
	Client   client.Client
	Recorder record.EventRecorder
	Logger   *logr.Logger
}

func (m *OBServerManager) GetTaskFunc(name string) (func() error, error) {
	switch name {
	case taskname.CreateOBPVC:
		return m.CreateOBPVC, nil
	case taskname.CreateOBPod:
		return m.CreateOBPod, nil
	case taskname.WaitOBPodReady:
		return m.WaitOBPodReady, nil
	case taskname.StartOBServer:
		return m.StartOBServer, nil
	case taskname.AddServer:
		return m.AddServer, nil
	default:
		return nil, errors.New("Can not find an function for task")
	}
}

func (m *OBServerManager) IsNewResource() bool {
	return m.OBServer.Status.Status == ""
}

func (m *OBServerManager) InitStatus() {
	m.Logger.Info("newly created server, init status")
	status := cloudv2alpha1.OBServerStatus{
		Image:  m.OBServer.Spec.OBServerTemplate.Image,
		Status: serverstatus.New,
	}
	m.OBServer.Status = status
}

func (m *OBServerManager) SetOperationContext(c *cloudv2alpha1.OperationContext) {
	m.OBServer.Status.OperationContext = c
}

func (m *OBServerManager) UpdateStatus() error {
	// get Pod status and update
	pod, err := m.getPod()
	if err != nil {
		if kubeerrors.IsNotFound(err) {
			m.Logger.Error(err, "pod not found")
		} else {
			return errors.Wrap(err, "get pod when update status")
		}
	} else {
		m.Logger.Info(">>>>>>>>>>>>>>get pod<<<<<<<<<<<<", "pod ip", pod.Status.PodIP, "pod phase", pod.Status.Phase)
		m.OBServer.Status.Ready = apipod.IsPodReady(pod)
		m.OBServer.Status.PodPhase = pod.Status.Phase
		m.OBServer.Status.PodIp = pod.Status.PodIP
		m.OBServer.Status.NodeIp = pod.Status.HostIP
	}

	m.Logger.Info("update observer status", "status", m.OBServer.Status)
	m.Logger.Info("update observer status", "operation context", m.OBServer.Status.OperationContext)
	err = m.Client.Status().Update(m.Ctx, m.OBServer)
	if err != nil {
		m.Logger.Error(err, "Got error when update observer status")
	}
	return err
}

func (m *OBServerManager) GetTaskFlow() (*task.TaskFlow, error) {
	// exists unfinished task flow, return the last task flow
	if m.OBServer.Status.OperationContext != nil {
		m.Logger.Info("get task flow from observer status")
		return task.NewTaskFlow(m.OBServer.Status.OperationContext), nil
	}
	// newly created observer
	var taskFlow *task.TaskFlow
	var err error
	var obcluster *cloudv2alpha1.OBCluster

	m.Logger.Info("create task flow according to observer status")
	if m.OBServer.Status.Status == serverstatus.New {
		obcluster, err = m.getOBCluster()
		if err != nil {
			return nil, errors.Wrap(err, "Get obcluster")
		}
		if obcluster.Status.Status == clusterstatus.New {
			// created when create obcluster
			m.Logger.Info("Create observer when create obcluster")
			taskFlow, err = task.GetRegistry().Get(flowname.CreateServerForBootstrap)
		} else {
			// created normally
			m.Logger.Info("Create observer when obcluster already exists")
			taskFlow, err = task.GetRegistry().Get(flowname.CreateServer)
		}
		if err != nil {
			return nil, errors.Wrap(err, "Get create observer task flow")
		}
		return taskFlow, nil
	}
	// scale observer
	// upgrade

	// no need to execute task flow
	return nil, nil
}

func (m *OBServerManager) ClearTaskInfo() {
	m.OBServer.Status.Status = serverstatus.Running
	m.OBServer.Status.OperationContext = nil
}

func (m *OBServerManager) FinishTask() {
	m.OBServer.Status.Status = m.OBServer.Status.OperationContext.TargetStatus
	m.OBServer.Status.OperationContext = nil
}

func (m *OBServerManager) generateNamespacedName(name string) types.NamespacedName {
	var namespacedName types.NamespacedName
	namespacedName.Namespace = m.OBServer.Namespace
	namespacedName.Name = name
	return namespacedName
}

func (m *OBServerManager) getPod() (*corev1.Pod, error) {
	// this label always exists
	pod := &corev1.Pod{}
	err := m.Client.Get(m.Ctx, m.generateNamespacedName(m.OBServer.Name), pod)
	if err != nil {
		return nil, errors.Wrap(err, "get pod")
	}
	return pod, nil
}

func (m *OBServerManager) getOBCluster() (*cloudv2alpha1.OBCluster, error) {
	// this label always exists
	clusterName, _ := m.OBServer.Labels["reference-cluster"]
	obcluster := &cloudv2alpha1.OBCluster{}
	err := m.Client.Get(m.Ctx, m.generateNamespacedName(clusterName), obcluster)
	if err != nil {
		return nil, errors.Wrap(err, "get obcluster")
	}
	return obcluster, nil
}

// get observer from K8s api server
func (m *OBServerManager) getOBServer() (*cloudv2alpha1.OBServer, error) {
	// this label always exists
	observer := &cloudv2alpha1.OBServer{}
	err := m.Client.Get(m.Ctx, m.generateNamespacedName(m.OBServer.Name), observer)
	if err != nil {
		return nil, errors.Wrap(err, "get observer")
	}
	return observer, nil
}

func (m *OBServerManager) getOBZone() (*cloudv2alpha1.OBZone, error) {
	// this label always exists
	zoneName, _ := m.OBServer.Labels["reference-zone"]
	obzone := &cloudv2alpha1.OBZone{}
	err := m.Client.Get(m.Ctx, m.generateNamespacedName(zoneName), obzone)
	if err != nil {
		return nil, errors.Wrap(err, "get obzone")
	}
	return obzone, nil
}