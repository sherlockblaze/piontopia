package client

import (
	"log"

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clientV1 "k8s.io/client-go/kubernetes/typed/batch/v1"
	"k8s.io/client-go/util/retry"
)

/*
createJobClient create job client by project
@param project
**/
func createJobClient(project string) clientV1.JobInterface {
	return k8sclient.BatchV1().Jobs(project)
}

/*
CreateJob create job
@param project
@param Job
**/
func CreateJob(project string, job *batchv1.Job) (*batchv1.Job, error) {
	log.Printf("creating job [%s] in project [%s]", job.Name, project)
	jobClient := createJobClient(project)
	result, err := jobClient.Create(job)
	if err != nil {
		log.Fatalf("failed to create job [%s] in project [%s], error: [%s]", job.Name, project, err.Error())
	}
	log.Printf("job [%s] created in project [%s]", result.GetObjectMeta().GetName(), project)
	return result, nil
}

/*
ListJob get job list
@param project
@param limit
@param fieldSelector
@param labelSelector
**/
func ListJob(project, fieldSelector, labelSelector string, limit int64) (*batchv1.JobList, error) {
	log.Printf("listing job in project [%s]", project)
	jobClient := createJobClient(project)
	listOptions := metav1.ListOptions{
		Limit:         limit,
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := jobClient.List(listOptions)
	if err != nil {
		log.Fatalf("cannot get job list in project [%s], error: %s", project, err.Error())
		return nil, err
	}
	return list, nil
}

/*
GetJob get job
@param project
@param jobName
**/
func GetJob(project, jobName string) (*batchv1.Job, error) {
	log.Printf("getting job [%s] in project [%s]", jobName, project)
	jobClient := createJobClient(project)
	result, err := jobClient.Get(jobName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("failed to get latest version of job [%s] in project [%s], error: [%s]", jobName, project, err.Error())
		return nil, err
	}
	return result, nil
}

/*
UpdateJob update job
@param project
@param job
**/
func UpdateJob(project string, job *batchv1.Job) (*batchv1.Job, error) {
	log.Printf("updating job [%s] in project [%s]", job.Name, project)
	var updatedJob *batchv1.Job
	jobClient := createJobClient(project)
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, err := GetJob(project, job.Name)
		if err != nil {
			log.Fatalf("failed to update job [%s] in project [%s], failed to get lastest version of job [%s], error: [%s]", job.Name, project, job.Name, err.Error())
			return err
		}
		// TODO: compare and replace

		updatedJob, err = jobClient.Update(result)
		return err
	})
	if err != nil {
		log.Fatalf("failed to update job [%s] in project [%s], error: [%s]", job, project, err.Error())
		return nil, err
	}
	log.Printf("job [%s] in project [%s] updated", updatedJob.Name, project)
	return updatedJob, nil
}

/*
DeleteJob delete job
@param project
@param jobName
**/
func DeleteJob(project, jobName string) error {
	log.Printf("deleting job [%s] in project [%s]", jobName, project)
	deletePolicy := metav1.DeletePropagationForeground
	jobClient := createJobClient(project)
	if err := jobClient.Delete(jobName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Fatalf("failed to delete job [%s] in project [%s], error: [%s]", jobName, project, err.Error())
		return err
	}
	log.Fatalf("job [%s] in project [%s] deleted", jobName, project)
	return nil
}
