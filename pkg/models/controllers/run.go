package controllers

import (
	"github.com/chengleqi/kubesphere-core/pkg/client"
	"k8s.io/client-go/kubernetes"
	"sync"
)

type resourceControllers struct {
	Controllers map[string]Controller
	k8sClient   *kubernetes.Clientset
}

var ResourceControllers resourceControllers

func Run(stopChan chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()

	k8sClient := client.NewK8sClient()
	ResourceControllers = resourceControllers{k8sClient: k8sClient, Controllers: make(map[string]Controller)}

	for _, item := range []string{"Pod"} {
		ResourceControllers.runController(item, stopChan, wg)
	}

	for {
		select {
		case <-stopChan:
			return
		default:
		}
	}
}

func (rec *resourceControllers) runController(name string, stopChan chan struct{}, wg *sync.WaitGroup) {
	var ctl Controller
	switch name {
	case "Pod":
		ctl = &PodCtl{K8sClient: rec.k8sClient}
	default:
		return
	}
	rec.Controllers[name] = ctl
	wg.Add(1)
	go listAndWatch(ctl, stopChan, wg)
}
