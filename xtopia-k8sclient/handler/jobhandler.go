package handler

import (
	"encoding/json"
	"io/ioutil"
	"k8s-client/client"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	batchv1 "k8s.io/api/batch/v1"
)

/*
CreateJob create job api
@param formatter
**/
func CreateJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var job *batchv1.Job
		json.Unmarshal(body, &job)
		createdJob, _ := client.CreateJob(project, job)
		formatter.JSON(w, http.StatusCreated,
			struct {
				Job batchv1.Job
			}{*createdJob})
	}
}

/*
GetJob get job api
@param formatter
**/
func GetJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		jobName := vars["jobName"]
		job, _ := client.GetJob(project, jobName)
		formatter.JSON(w, http.StatusOK,
			struct {
				Job batchv1.Job
			}{*job})
	}
}

/*
UpdateJob update job api
@param formatter
**/
func UpdateJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
	}
}

/*
DeleteJob delete job api
@param formatter
**/
func DeleteJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		jobName := vars["jobName"]
		client.DeleteJob(project, jobName)
		formatter.JSON(w, http.StatusOK, struct{}{})
	}
}

/*
ListJob list job api
@param formatter
**/
func ListJob(formatter *render.Render) http.HandlerFunc {
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
		list, _ := client.ListJob(project, fieldSelector, labelSelector, limit)
		var jobList []batchv1.Job
		if list != nil {
			jobList = list.Items
		}
		formatter.JSON(w, http.StatusOK,
			struct {
				JobList []batchv1.Job
			}{jobList})
	}
}
