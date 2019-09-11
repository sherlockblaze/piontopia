package client

import (
	"encoding/json"
	"log"

	betav1 "k8s.io/api/networking/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	clientBetaV1 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	"k8s.io/client-go/util/retry"
)

/*
createIngressClient create ingress client by project
@param project
**/
func createIngressClient(project string) clientBetaV1.IngressInterface {
	return k8sclient.NetworkingV1beta1().Ingresses(project)
}

/*
CreateIngress create ingress
@param project
@param ingress
**/
func CreateIngress(project string, ingress *betav1.Ingress) (*betav1.Ingress, error) {
	log.Printf("creating ingress [%s] in project [%s]", ingress.Name, project)
	ingressClient := createIngressClient(project)
	result, err := ingressClient.Create(ingress)
	if err != nil {
		log.Printf("failed to create ingress [%s] in project [%s], error: [%s]", ingress.Name, project, err.Error())
		return nil, err
	}
	log.Printf("ingress [%s] created in project [%s]", result.Name, project)
	return result, nil
}

/*
ListIngress get ingress list
@param project
@param limit
@param fieldSelector
@param labelSelector
**/
func ListIngress(project, fieldSelector, labelSelector string, limit int64) (*betav1.IngressList, error) {
	log.Printf("listing ingress in project [%s]", project)
	ingressClient := createIngressClient(project)
	listOptions := metav1.ListOptions{
		Limit:         limit,
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := ingressClient.List(listOptions)
	if err != nil {
		log.Printf("cannot get ingress list in project [%s], error: %s", project, err.Error())
		return nil, err
	}
	return list, nil
}

/*
GetIngress get ingress
@param project
@param ingressName
**/
func GetIngress(project, ingressName string) (*betav1.Ingress, error) {
	log.Printf("getting ingress [%s] in project [%s]", ingressName, project)
	ingressClient := createIngressClient(project)
	result, err := ingressClient.Get(ingressName, metav1.GetOptions{})
	if err != nil {
		log.Printf("failed to get latest version of ingress [%s] in project [%s], error: [%s]", ingressName, project, err.Error())
		return nil, err
	}
	return result, nil
}

/*
UpdateIngress update ingress
@param project
@param ingress
**/
func UpdateIngress(project string, ingress *betav1.Ingress) (*betav1.Ingress, error) {
	log.Printf("updating ingress [%s] in project [%s]", ingress.Name, project)
	var updatedIngress *betav1.Ingress
	ingressClient := createIngressClient(project)
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, err := GetIngress(project, ingress.Name)
		if err != nil {
			log.Printf("failed to update ingress [%s] in project [%s], failed to get lastest version of ingress [%s], error: [%s]", ingress.Name, project, ingress.Name, err.Error())
			return err
		}
		// TODO: compare and replace
		oldData, err := json.Marshal(result)
		if err != nil {
			log.Printf("failed to update ingress [%s] in project [%s], error: [%s]", ingress.Name, project, err.Error())
			return err
		}
		newData, err := json.Marshal(ingress)
		if err != nil {
			log.Printf("failed to update ingress [%s] in project [%s], error: [%s]", ingress.Name, project, err.Error())
			return err
		}
		patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, betav1.Ingress{})
		if err != nil {
			log.Printf("failed to update ingress [%s] in project [%s], error: [%s]", ingress.Name, project, err.Error())
			return err
		}

		updatedIngress, err = ingressClient.Patch(ingress.Name, types.StrategicMergePatchType, patchBytes)
		return err
	})
	if err != nil {
		log.Printf("failed to update ingress [%s] in project [%s], error: [%s]", ingress.Name, project, err.Error())
		return nil, err
	}
	log.Printf("ingress [%s] in project [%s] updated", updatedIngress.Name, project)
	return updatedIngress, nil
}

/*
DeleteIngress delete ingress
@param project
@param ingressName
**/
func DeleteIngress(project, ingressName string) error {
	log.Printf("deleting ingress [%s] in project [%s]", ingressName, project)
	deletePolicy := metav1.DeletePropagationForeground
	IngressClient := createIngressClient(project)
	if err := IngressClient.Delete(ingressName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Printf("failed to delete ingress [%s] in project [%s], error: [%s]", ingressName, project, err.Error())
		return err
	}
	log.Printf("ingress [%s] in project [%s] deleted", ingressName, project)
	return nil
}
