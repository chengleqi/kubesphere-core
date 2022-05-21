package controllers

import (
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	coreV1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
	"time"
)

type PodCtl struct {
	K8sClient *kubernetes.Clientset
	lister    coreV1.PodLister
	informer  cache.SharedIndexInformer
}

func (c *PodCtl) sync(stopChan chan struct{}) {
	c.initListerAndInformer()
	list, err := c.lister.List(labels.Everything())
	if err != nil {
		klog.Error(err)
		return
	}
	// get all item.Name
	for _, item := range list {
		fmt.Println(item.Name)
	}
	c.informer.Run(stopChan)
}

func (c *PodCtl) initListerAndInformer() {
	informerFactory := informers.NewSharedInformerFactory(c.K8sClient, 30*time.Second)
	// create lister
	c.lister = informerFactory.Core().V1().Pods().Lister()

	// create informer
	informer := informerFactory.Core().V1().Pods().Informer()
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			object := obj.(*v1.Pod)
			fmt.Println(object.Name)
			fmt.Println("detect add event")
		},
		UpdateFunc: func(old, new interface{}) {
			oldObj, newObj := old.(*v1.Pod), new.(*v1.Pod)
			fmt.Println(oldObj.Name, newObj.Name)
			fmt.Println("detect update event")
		},
		DeleteFunc: func(obj interface{}) {
			object := obj.(*v1.Pod)
			fmt.Println(object.Name)
			fmt.Println("detect delete event")
		},
	})
	c.informer = informer
}
