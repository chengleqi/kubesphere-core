package client

import (
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog/v2"
)

var k8sClient *kubernetes.Clientset

func NewK8sClient() *kubernetes.Clientset {
	if k8sClient != nil {
		return k8sClient
	}

	kubeConfig, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		klog.Fatal(err)
	}

	k8sClient, err = kubernetes.NewForConfig(kubeConfig)
	if err != nil {
		klog.Fatal(err)
	}
	return k8sClient
}
