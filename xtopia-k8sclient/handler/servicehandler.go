package handler

import (
	"encoding/json"
	"io/ioutil"
	"k8s-client/client"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	apiv1 "k8s.io/api/core/v1"
)

/*
CreateService create service api
@param formatter
**/
func CreateService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var service *apiv1.Service
		json.Unmarshal(body, &service)
		createdService, err := client.CreateService(project, service)
		status := http.StatusCreated
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = createdService
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
GetService get service api
@param formatter
**/
func GetService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		serviceName := vars["serviceName"]
		service, err := client.GetService(project, serviceName)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = service
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
UpdateService update service api
@param formatter
**/
func UpdateService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var service *apiv1.Service
		json.Unmarshal(body, &service)
		updatedService, err := client.UpdateService(project, service)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = updatedService
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
DeleteService delete service api
@param formatter
**/
func DeleteService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		serviceName := vars["serviceName"]
		err := client.DeleteService(project, serviceName)
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
		formatter.JSON(w, status, res)
	}
}

/*
ListService list service api
@param formatter
**/
func ListService(formatter *render.Render) http.HandlerFunc {
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
		list, err := client.ListService(project, fieldSelector, labelSelector, limit)
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
		formatter.JSON(w, status, res)
	}
}
