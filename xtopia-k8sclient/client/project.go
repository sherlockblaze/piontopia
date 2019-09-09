package client

import (
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

/*
createNsClient create namespace client
**/
func createNsClient() corev1.NamespaceInterface {
	return k8sclient.CoreV1().Namespaces()
}

/*
CreateProject create project
@param project
**/
func CreateProject(project *v1.Namespace) (*v1.Namespace, error) {
	log.Printf("creating project [%s]", project.GetObjectMeta().GetName())
	nsClient := createNsClient()
	createdNs, err := nsClient.Create(project)
	if err != nil {
		log.Printf("failed to create project [%s], error: [%s]", project.GetObjectMeta().GetName(), err.Error())
		return nil, err
	}
	log.Printf("project [%s] created", createdNs.GetObjectMeta().GetName())
	return createdNs, nil
}

/*
DeleteProject delete project
@param projectName
**/
func DeleteProject(projectName string) error {
	log.Printf("deleting project [%s]", projectName)
	deletePolicy := metav1.DeletePropagationForeground
	nsClient := createNsClient()
	err := nsClient.Delete(projectName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	})
	if err != nil {
		log.Printf("failed to delete project [%s]", projectName)
		return err
	}
	log.Printf("project [%s] deleted", projectName)
	return nil
}

/*
GetProject get project
@param projectName
**/
func GetProject(projectName string) (*v1.Namespace, error) {
	log.Printf("getting project [%s]", projectName)
	nsClient := createNsClient()
	result, err := nsClient.Get(projectName, metav1.GetOptions{})
	if err != nil {
		log.Printf("failed to get latest version of project [%s], error: [%s]", projectName, err.Error())
		return nil, err
	}
	return result, nil
}

/*
ListProject list project
@param fieldSelector
@param labelSelector
@param limit
**/
func ListProject(fieldSelector, labelSelector string, limit int64) (*v1.NamespaceList, error) {
	log.Printf("listing projects...")
	nsClient := createNsClient()
	listOptions := metav1.ListOptions{
		Limit:         limit,
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := nsClient.List(listOptions)
	if err != nil {
		log.Printf("cannot get projects, error: %s", err.Error())
		return nil, err
	}
	return list, nil
}
