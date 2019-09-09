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
		createdService, _ := client.CreateService(project, service)
		formatter.JSON(w, http.StatusCreated,
			struct {
				Service apiv1.Service
			}{*createdService})
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
		service, _ := client.GetService(project, serviceName)
		formatter.JSON(w, http.StatusOK,
			struct {
				Service apiv1.Service
			}{*service})
	}
}

/*
UpdateService update service api
@param formatter
**/
func UpdateService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
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
		client.DeleteService(project, serviceName)
		formatter.JSON(w, http.StatusOK, struct{}{})
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
		list, _ := client.ListService(project, fieldSelector, labelSelector, limit)
		var serviceList []apiv1.Service
		if list != nil {
			serviceList = list.Items
		}
		formatter.JSON(w, http.StatusOK,
			struct {
				ServiceList []apiv1.Service
			}{serviceList})
	}
}
