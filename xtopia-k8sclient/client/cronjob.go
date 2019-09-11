package client

import (
	"encoding/json"
	"log"

	batchbetav1 "k8s.io/api/batch/v1beta1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	clientbetaV1 "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	"k8s.io/client-go/util/retry"
)

/*
createCronClient create cron job client by project
@param project
**/
func createCronClient(project string) clientbetaV1.CronJobInterface {
	return k8sclient.BatchV1beta1().CronJobs(project)
}

/*
CreateCronJob create a cronjob
@param project
@param cronjob
**/
func CreateCronJob(project string, cronJob *batchbetav1.CronJob) (*batchbetav1.CronJob, error) {
	log.Printf("creating cronjob [%s] in project [%s]", cronJob.Name, project)
	cronClient := createCronClient(project)
	result, err := cronClient.Create(cronJob)
	if err != nil {
		log.Printf("failed to create cronjob [%s] in project [%s], error: [%s]", cronJob.Name, project, err.Error())
		return nil, err
	}
	log.Printf("cronjob [%s] in project [%s] created", result.GetObjectMeta().GetName(), project)
	return result, nil
}

/*
ListCronJob list cron job
@param project
@param fieldSelector
@param labelSelector
@param limit
**/
func ListCronJob(project, fieldSelector, labelSelector string, limit int64) (*batchbetav1.CronJobList, error) {
	log.Printf("listing cronjob in project [%s]", project)
	cronClient := createCronClient(project)
	listOptions := metav1.ListOptions{
		Limit:         limit,
		FieldSelector: fieldSelector,
		LabelSelector: labelSelector,
	}
	list, err := cronClient.List(listOptions)
	if err != nil {
		log.Printf("cannot get cronjob list of [%s], error: [%s]", project, err.Error())
		return nil, err
	}

	return list, nil
}

/*
GetCronJob get cron job
@param project
@param cronName
**/
func GetCronJob(project, cronName string) (*batchbetav1.CronJob, error) {
	log.Printf("getting cronjob [%s] in project [%s]", cronName, project)
	cronClient := createCronClient(project)
	result, err := cronClient.Get(cronName, metav1.GetOptions{})
	if err != nil {
		log.Printf("failed to get latest version of cronjob [%s] in project [%s], error: [%s]", cronName, project, err.Error())
		return nil, err
	}
	return result, nil
}

/*
UpdateCronJob update cron job
@param project
@param cronjob
**/
func UpdateCronJob(project string, cronJob *batchbetav1.CronJob) (*batchbetav1.CronJob, error) {
	log.Printf("updating cronjob [%s] in project [%s]", cronJob.Name, project)
	var updatedCron *batchbetav1.CronJob
	cronClient := createCronClient(project)
	err := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		result, err := GetCronJob(project, cronJob.Name)
		if err != nil {
			log.Printf("failed to update cronjob [%s] in project [%s], failed to get lastest version of CronJob [%s], error: [%s]", cronJob.Name, project, cronJob.Name, err.Error())
			return err
		}
		// TODO: compare and replace
		oldData, err := json.Marshal(result)
		if err != nil {
			log.Printf("failed to update cronjob [%s] in project [%s], error: [%s]", cronJob.Name, project, err.Error())
			return err
		}
		newData, err := json.Marshal(cronJob)
		if err != nil {
			log.Printf("failed to update cronjob [%s] in project [%s], error: [%s]", cronJob.Name, project, err.Error())
			return err
		}
		patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, batchbetav1.CronJob{})
		if err != nil {
			log.Printf("failed to update cronjob [%s] in project [%s], error: [%s]", cronJob.Name, project, err.Error())
			return err
		}
		updatedCron, err = cronClient.Patch(cronJob.Name, types.StrategicMergePatchType, patchBytes)
		return err
	})
	if err != nil {
		log.Printf("failed to update cronjob [%s] in project [%s], error: [%s]", cronJob.Name, project, err.Error())
		return nil, err
	}
	log.Printf("cronjob [%s] updated in project [%s]...", cronJob.Name, project)
	return updatedCron, nil
}

/*
DeleteCronJob delete cron job
@param project
@param cronName
**/
func DeleteCronJob(project, cronName string) error {
	log.Printf("deleting cronJob [%s] in project [%s]", cronName, project)
	deletePolicy := metav1.DeletePropagationForeground
	cronClient := createCronClient(project)
	if err := cronClient.Delete(cronName, &metav1.DeleteOptions{
		PropagationPolicy: &deletePolicy,
	}); err != nil {
		log.Printf("failed to delete cronjob [%s] in project [%s], error: [%s]", cronName, project, err.Error())
		return err
	}
	log.Printf("cronjob [%s] in project [%s] deleted", cronName, project)
	return nil
}
