package client

import (
	"fmt"
	"log"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientv1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/util/retry"
)

var deploymentsClient clientv1.DeploymentInterface

func init() {
	deploymentsClient = k8sclient.AppsV1().Deployments(apiv1.NamespaceDefault)
}

var deployment = &appsv1.Deployment{
	ObjectMeta: metav1.ObjectMeta{
		Name: "demo-deployment",
	},
	Spec: appsv1.DeploymentSpec{
		Replicas: int32Ptr(2),
		Selector: &metav1.LabelSelector{
			MatchLabels: map[string]string{
				"app": "demo",
			},
		},
		Template: apiv1.PodTemplateSpec{
			ObjectMeta: metav1.ObjectMeta{
				Labels: map[string]string{
					"app": "demo",
				},
			},
			Spec: apiv1.PodSpec{
				Containers: []apiv1.Container{
					{
						Name:  "web",
						Image: "nginx:1.12",
						Ports: []apiv1.ContainerPort{
							{
								Name:          "http",
								Protocol:      apiv1.ProtocolTCP,
								ContainerPort: 80,
							},
						},
					},
				},
			},
		},
	},
}

// CreateDeployment create deployment
func CreateDeployment() {
	log.Println("Creating deployment...")
	result, err := deploymentsClient.Create(deployment)
	if err != nil {
		panic(err)
	}
	log.Printf("Create deployment %q.\n", result.GetObjectMeta().GetName())
}

// ListDeployment list deployment
func ListDeployment() {
	log.Printf("Listing deployment in namespace %q:\n", apiv1.NamespaceDefault)
	list, err := deploymentsClient.List(metav1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, d := range list.Items {
		log.Printf(" * %s (%d replicas)\n", d.Name, *d.Spec.Replicas)
	}
}

// UpdateDeployment update deployment
func UpdateDeployment() {
	log.Println("Updating deployment")
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, getErr := deploymentsClient.Get("demo-deployment", metav1.GetOptions{})
		if getErr != nil {
			panic(fmt.Errorf("Failed to get latest version of Deployment: %v", getErr))
		}
		result.Spec.Replicas = int32Ptr(1)

		result.Spec.Template.Spec.Containers[0].Image = "nginx:1.13"
		_, updateErr := deploymentsClient.Update(result)
		return updateErr
	})
	if retryErr != nil {
		panic(fmt.Errorf("Update failed: %v", retryErr))
	}
	log.Println("Updated deployment...")
}

// DeleteDeployment delete deployment
func DeleteDeployment() {
	deletePolicy := metav1.DeletePropagationForeground
	if err := deploymentsClient.Delete("demo-deployment", &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		panic(err)
	}
	fmt.Println("Deleted deployment.")
}

func int32Ptr(i int32) *int32 {
	return &i
}
