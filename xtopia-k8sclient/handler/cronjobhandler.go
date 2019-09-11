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
		createdCronJob, err := client.CreateCronJob(project, cronJob)
		status := http.StatusCreated
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = createdCronJob
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
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
		cronJob, err := client.GetCronJob(project, cronJobName)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = cronJob
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
UpdateCronJob update cronjob api
@param formatter
**/
func UpdateCronJob(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var cronJob *batchbetav1.CronJob
		json.Unmarshal(body, &cronJob)
		updatedCronJob, err := client.UpdateCronJob(project, cronJob)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = updatedCronJob
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
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
		err := client.DeleteJob(project, cronJobName)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
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
		list, err := client.ListCronJob(project, fieldSelector, labelSelector, limit)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		res["result"] = struct{}{}
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		if list != nil {
			res["result"] = list
		}
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}
