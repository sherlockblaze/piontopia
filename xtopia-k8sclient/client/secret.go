package client

import (
	"log"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
)

/*
createSecretClient create Secret client by project
@param project
**/
func createSecretClient(project string) corev1.SecretInterface {
	return k8sclient.CoreV1().Secrets(project)
}

/*
CreateSecret create a secret
@param project
@param secret
**/
func CreateSecret(project string, secret *apiv1.Secret) (*apiv1.Secret, error) {
	log.Printf("creating secret [%s] in project [%s]", secret.Name, project)
	secretClient := createSecretClient(project)
	secret, err := secretClient.Create(secret)
	if err != nil {
		log.Printf("failed to create secret [%s] in project [%s], error: [%s]", secret.Name, project, err.Error())
		return nil, err
	}
	log.Printf("secret [%s] in project [%s] created", secret.Name, project)
	return secret, nil
}

/*
GetSecret get secret
@param project
@param secretName
**/
func GetSecret(project, secretName string) (*apiv1.Secret, error) {
	log.Printf("getting secret [%s] in project [%s]", secretName, project)
	secretClient := createSecretClient(project)
	result, err := secretClient.Get(secretName, metav1.GetOptions{})
	if err != nil {
		log.Printf("failed to get latest version of secret [%s] in project [%s], error: [%s]", secretName, project, err.Error())
		return nil, err
	}
	return result, nil
}

/*
DeleteSecret delete secret
@param project
@param scretName
**/
func DeleteSecret(project, secretName string) error {
	log.Printf("deleting secret [%s] in project [%s]", secretName, project)
	deletePolicy := metav1.DeletePropagationForeground
	secretClient := createSecretClient(project)
	if err := secretClient.Delete(secretName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Printf("failed to delete secret [%s] in project [%s], error: [%s]", secretName, project, err.Error())
		return err
	}
	log.Printf("secret [%s] in project [%s] deleted", secretName, project)
	return nil
}
