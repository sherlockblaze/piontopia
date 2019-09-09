package handler

import (
	"encoding/json"
	"io/ioutil"
	"k8s-client/client"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	batchbetav1 "k8s.io/api/batch/v1beta1"
)

/*
CreateCronJob create cronjob api
@param formatter
**/
func CreateCronJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var cronJob *batchbetav1.CronJob
		json.Unmarshal(body, &cronJob)
		createdCronJob, _ := client.CreateCronJob(project, cronJob)
		formatter.JSON(w, http.StatusCreated,
			struct {
				CronJob batchbetav1.CronJob
			}{*createdCronJob})
	}
}

/*
GetCronJob get cronjob api
@param formatter
**/
func GetCronJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		cronJobName := vars["cronJobName"]
		cronJob, _ := client.GetCronJob(project, cronJobName)
		formatter.JSON(w, http.StatusOK,
			struct {
				CronJob batchbetav1.CronJob
			}{*cronJob})
	}
}

/*
UpdateCronJob update cronjob api
@param formatter
**/
func UpdateCronJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
	}
}

/*
DeleteCronJob delete cronjob api
@param formatter
**/
func DeleteCronJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		cronJobName := vars["cronJobName"]
		client.DeleteJob(project, cronJobName)
		formatter.JSON(w, http.StatusOK, struct{}{})
	}
}

/*
ListCronJob list cronjob api
@param formatter
**/
func ListCronJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		var reqParam map[string]interface{}
		body, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &reqParam)
		var fieldSelector string
		if reqParam["fieldSelector"] != nil {
			fieldSelector = reqParam["fieldSelector"].(string)
		}
		var labelSelector string
		if reqParam["fieldSelector"] != nil {
			fieldSelector = reqParam["labelSelector"].(string)
		}
		var limit int64
		if reqParam["limit"] != nil {
			limit = reqParam["limit"].(int64)
		}
		list, _ := client.ListCronJob(project, fieldSelector, labelSelector, limit)
		var cronJobList []batchbetav1.CronJob
		if list != nil {
			cronJobList = list.Items
		}
		formatter.JSON(w, http.StatusOK,
			struct {
				CronJobList []batchbetav1.CronJob
			}{cronJobList})
	}
}
