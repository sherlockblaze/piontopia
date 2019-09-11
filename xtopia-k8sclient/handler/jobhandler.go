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
		createdJob, err := client.CreateJob(project, job)
		status := http.StatusCreated
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = createdJob
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
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
		job, err := client.GetJob(project, jobName)
		status := http.StatusCreated
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = job
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
UpdateJob update job api
@param formatter
**/
func UpdateJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var job *batchv1.Job
		json.Unmarshal(body, &job)
		updatedJob, err := client.UpdateJob(project, job)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = updatedJob
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
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
		err := client.DeleteJob(project, jobName)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, http.StatusOK, res)
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
		list, err := client.ListJob(project, fieldSelector, labelSelector, limit)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		res["result"] = struct{}{}
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		if list != nil {
			res["result"] = list.Items
		}
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}
