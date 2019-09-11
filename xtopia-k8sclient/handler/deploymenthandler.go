package handler

import (
	"encoding/json"
	"io/ioutil"
	"k8s-client/client"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/unrolled/render"
	appsv1 "k8s.io/api/apps/v1"
)

/*
CreateDeployment create deployment api
@param formatter
**/
func CreateDeployment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var deployment *appsv1.Deployment
		json.Unmarshal(body, &deployment)
		createdDeployment, err := client.CreateDeployment(project, deployment)
		status := http.StatusCreated
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = createdDeployment
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
GetDeployment get deployment api
@param formatter
**/
func GetDeployment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		deploymentName := vars["deploymentName"]
		deployment, err := client.GetDeployment(project, deploymentName)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = deployment
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
UpdateDeployment update deployment api
@param formatter
**/
func UpdateDeployment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var deployment *appsv1.Deployment
		json.Unmarshal(body, &deployment)
		updatedDeployment, err := client.UpdateDeployment(project, deployment)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = updatedDeployment
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
DeleteDeployment delete deployment api
@param formatter
**/
func DeleteDeployment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		deploymentName := vars["deploymentName"]
		err := client.DeleteDeployment(project, deploymentName)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
ListDeployment list deployment api
@param formatter
**/
func ListDeployment(formatter *render.Render) http.HandlerFunc {
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
		list, err := client.ListDeployment(project, fieldSelector, labelSelector, limit)
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
			res["result"] = list.Items
		}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}
