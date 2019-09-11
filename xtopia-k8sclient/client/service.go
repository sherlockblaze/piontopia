package client

import (
	"encoding/json"
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/util/retry"
)

/*
createServiceClient create service client by project
@param project
**/
func createServiceClient(project string) corev1.ServiceInterface {
	return k8sclient.CoreV1().Services(project)
}

/*
CreateService create a service
@param project
@param service
**/
func CreateService(project string, service *apiv1.Service) (*apiv1.Service, error) {
	log.Printf("creating service [%s] in project [%s]", service.Name, project)
	serviceClient := createServiceClient(project)
	service, err := serviceClient.Create(service)
	if err != nil {
		log.Printf("failed to create service [%s] in project [%s], error: [%s]", service.Name, project, err.Error())
		return nil, err
	}
	log.Printf("service [%s] in project [%s] created", service.Name, project)
	return service, nil
}

/*
ListService list services
@param project
@param fieldSelector
@param labelSelector
@param limit
**/
func ListService(project, fieldSelector, labelSelector string, limit int64) (*apiv1.ServiceList, error) {
	log.Printf("listing service in project %q:\n", project)
	serviceClient := createServiceClient(project)
	listOptions := metav1.ListOptions{
		Limit:         limit,
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := serviceClient.List(listOptions)
	if err != nil {
		log.Printf("cannot get service list in project [%s], error: %s", project, err.Error())
		return nil, err
	}
	return list, nil
}

/*
GetService get service
@param project
@param serviceName
**/
func GetService(project, serviceName string) (*apiv1.Service, error) {
	log.Printf("getting service [%s] in project [%s]", serviceName, project)
	serviceClient := createServiceClient(project)
	result, err := serviceClient.Get(serviceName, metav1.GetOptions{})
	if err != nil {
		log.Printf("failed to get latest version of service [%s] in project [%s], error: [%s]", serviceName, project, err.Error())
		return nil, err
	}
	return result, nil
}

/*
UpdateService update service
@param service
@param proejct
**/
func UpdateService(project string, service *apiv1.Service) (*apiv1.Service, error) {
	log.Printf("updating service [%s] in project [%s]", service.Name, project)
	serviceClient := createServiceClient(project)
	var updatedService *apiv1.Service
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, err := GetService(project, service.Name)
		if err != nil {
			log.Printf("failed to update service [%s] in project [%s], failed to get lastest version of service [%s], error: [%s]", service.Name, project, service.Name, err.Error())
			return err
		}
		// TODO: compare and replace
		oldData, err := json.Marshal(result)
		if err != nil {
			log.Printf("failed to update service [%s] in project [%s], error: [%s]", service.Name, project, err.Error())
			return err
		}
		newData, err := json.Marshal(service)
		if err != nil {
			log.Printf("failed to update service [%s] in project [%s], error: [%s]", service.Name, project, err.Error())
			return err
		}
		patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, apiv1.Service{})
		if err != nil {
			log.Printf("failed to update service [%s] in project [%s], error: [%s]", service.Name, project, err.Error())
			return err
		}

		updatedService, err = serviceClient.Patch(service.Name, types.StrategicMergePatchType, patchBytes)
		return err
	})
	if err != nil {
		log.Printf("failed to update service [%s] in project [%s], error: [%s]", service.Name, project, err.Error())
		return nil, err
	}
	log.Printf("service [%s] in project [%s] updated", service.Name, project)
	return updatedService, nil
}

/*
DeleteService delete service
@param project
@param serviceName
**/
func DeleteService(project, serviceName string) error {
	log.Printf("deleting service [%s] in project [%s]", serviceName, project)
	deletePolicy := metav1.DeletePropagationForeground
	serviceClient := createServiceClient(project)
	if err := serviceClient.Delete(serviceName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Printf("failed to delete service [%s] in project [%s], error: [%s]", serviceName, project, err.Error())
		return err
	}
	log.Printf("service [%s] in project [%s] deleted", serviceName, project)
	return nil
}
