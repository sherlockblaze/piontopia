package client

import (
	k8s "k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

var k8sclient *k8s.Clientset

func init() {
	k8sclient = newClient()
}

func newClient() *k8s.Clientset {
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err)
	}

	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		panic(err)
	}
	return clientset
}
