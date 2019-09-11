package client

import (
	"encoding/json"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	clientv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/util/retry"
)

/*
createDeploymentClient create deployment client by project
@param project
**/
func createDeploymentClient(namespace string) clientv1.DeploymentInterface {
	return k8sclient.AppsV1().Deployments(namespace)
}

/*
CreateDeployment create deployment
@param project
@param deployment
**/
func CreateDeployment(project string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	log.Printf("creating deployment [%s] in project [%s]...", deployment.Name, project)
	deploymentsClient := createDeploymentClient(project)
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		log.Printf("failed to create deployment [%s] in project [%s], error: [%s]", deployment.Name, project, err.Error())
		return nil, err
	}
	log.Printf("deployment [%s] created in project [%s].", result.GetObjectMeta().GetName(), project)
	return result, nil
}

/*
ListDeployment list deployment
@param project
@param limit
@param fieldSelector
@param labelSelector
**/
func ListDeployment(project, fieldSelector, labelSelector string, limit int64) (*appsv1.DeploymentList, error) {
	log.Printf("listing deployment in project [%s]", project)
	deploymentsClient := createDeploymentClient(project)
	listOptions := metav1.ListOptions{}
	if fieldSelector != "" {
		listOptions.FieldSelector = fieldSelector
	}
	if labelSelector != "" {
		listOptions.LabelSelector = labelSelector
	}
	if limit != 0 {
		listOptions.Limit = limit
	}

	list, err := deploymentsClient.List(listOptions)
	if err != nil {
		log.Printf("cannot get deployment list in project [%s], error: %s", project, err.Error())
		return nil, err
	}
	return list, nil
}

/*
GetDeployment get target deployment by name
@param project
@param deploymentName
**/
func GetDeployment(project, deploymentName string) (*appsv1.Deployment, error) {
	log.Printf("getting deployment [%s] in project [%s]", deploymentName, project)
	deploymentsClient := createDeploymentClient(project)
	result, err := deploymentsClient.Get(deploymentName, metav1.GetOptions{})
	if err != nil {
		log.Printf("failed to get latest version of deployment [%s] in project [%s], error: [%s]", deploymentName, project, err.Error())
		return nil, err
	}
	return result, nil
}

/*
UpdateDeployment update a target deployment
@param project
@param deployment
**/
func UpdateDeployment(project string, deployment *appsv1.Deployment) (*appsv1.Deployment, error) {
	log.Printf("updating deployment [%s] in project [%s]", deployment.Name, project)
	var updatedDeployment *appsv1.Deployment
	deploymentsClient := createDeploymentClient(project)
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, err := deploymentsClient.Get(deployment.Name, metav1.GetOptions{})
		if err != nil {
			log.Printf("failed to update deployment [%s] in project [%s], failed to get lastest version of deployment [%s], error: [%s]]", deployment.Name, project, deployment.Name, err.Error())
			return err
		}
		// TODO: compare and replace
		oldData, err := json.Marshal(result)
		if err != nil {
			log.Printf("failed to update deployment [%s] in project [%s], error: [%s]", deployment.Name, project, err.Error())
			return err
		}
		newData, err := json.Marshal(deployment)
		if err != nil {
			log.Printf("failed to update deployment [%s] in project [%s], error: [%s]", deployment.Name, project, err.Error())
			return err
		}
		patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, appsv1.Deployment{})
		if err != nil {
			log.Printf("failed to update deployment [%s] in project [%s], error: [%s]", deployment.Name, project, err.Error())
			return err
		}
		updatedDeployment, err = deploymentsClient.Patch(deployment.Name, types.StrategicMergePatchType, patchBytes)
		return err
	})
	if err != nil {
		log.Printf("failed to update deployment [%s] in project [%s], error: [%s]", deployment.Name, project, err.Error())
		return nil, err
	}
	log.Printf("deployment [%s] in project [%s] updated", updatedDeployment.Name, project)
	return updatedDeployment, nil
}

/*
DeleteDeployment delete target deployment
@param project
@param deploymentName
**/
func DeleteDeployment(project, deploymentName string) error {
	log.Printf("deleting deployment [%s] in project [%s]", deploymentName, project)
	deletePolicy := metav1.DeletePropagationForeground
	deploymentsClient := createDeploymentClient(project)
	if err := deploymentsClient.Delete(deploymentName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Printf("failed to delete deployment [%s] in project [%s], error: [%s]", deploymentName, project, err.Error())
		return err
	}
	log.Printf("deployment [%s] in project [%s] deleted", deploymentName, project)
	return nil
}

func int32Ptr(i int32) *int32 {
	return &i
}
