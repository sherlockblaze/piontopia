package client

import (
	batchbetav1 "k8s.io/api/batch/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientbetaV1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
)

var cronClient clientbetaV1.CronJobInterface

func init() {
	cronClient = k8sclient.BatchV1beta1().CronJobs(apiv1.NamespaceDefault)
}

var cornExample = &batchbetav1.CronJob{}

// CreateCronJob create cronjob
func CreateCronJob() {
	cronClient.Create(cornExample)
}

// ListCronJob list cronjob
func ListCronJob() {
	cronClient.List(metav1.ListOptions{})
}

// UpdateCronJob update cronjob
func UpdateCronJob() {
	cronClient.Update(cornExample)
}

// DeleteCronJob delete cronjob
func DeleteCronJob() {
	cronClient.Delete("", &metav1.DeleteOptions{})
}
