package client

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

var serviceClient corev1.ServiceInterface

func init() {
	serviceClient = k8sclient.CoreV1().Services(apiv1.NamespaceDefault)
}

var serviceExample = &apiv1.Service{}

// CreateService create service
func CreateService() {
	serviceClient.Create(serviceExample)
}

// ListService list service
func ListService() {
	serviceClient.List(metav1.ListOptions{})
}

// UpdateService update service
func UpdateService() {
	serviceClient.Update(serviceExample)
}

// DeleteService delete service
func DeleteService() {
	serviceClient.Delete("", &metav1.DeleteOptions{})
}
