package client

import (
	batchv1 "k8s.io/api/batch/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientV1 "k8s.io/client-go/kubernetes/typed/batch/v1"
)

var jobClient clientV1.JobInterface

func init() {
	jobClient = k8sclient.BatchV1().Jobs(apiv1.NamespaceDefault)
}

var jobExample = &batchv1.Job{}

// CreateJob create job
func CreateJob() {
	jobClient.Create(jobExample)
}

// ListJob list job
func ListJob() {
	jobClient.List(metav1.ListOptions{})
}

// UpdateJob update job
func UpdateJob() {
	jobClient.Update(jobExample)
}

// DeleteJob delete job
func DeleteJob() {
	jobClient.Delete("", &metav1.DeleteOptions{})
}
