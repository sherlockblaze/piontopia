package handler

import (
	"encoding/json"
	"io/ioutil"
	"k8s-client/client"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	v1 "k8s.io/api/core/v1"
)

/*
CreateProject create project api
@param formatter
**/
func CreateProject(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		var project *v1.Namespace
		json.Unmarshal(body, &project)
		createdProject, err := client.CreateProject(project)
		status := http.StatusCreated
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = createdProject
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
GetProject get project api
@param fomatter
**/
func GetProject(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		projectName := vars["projectName"]
		project, err := client.GetProject(projectName)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = project
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
UpdateProject update project api
@param formatter
**/
func UpdateProject(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		body, _ := ioutil.ReadAll(req.Body)
		var project *v1.Namespace
		json.Unmarshal(body, &project)
		updatedProject, err := client.UpdateProject(project)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = updatedProject
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
DeleteProject delete project api
@param formatter
**/
func DeleteProject(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		projectName := vars["projectName"]
		err := client.DeleteProject(projectName)
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
ListProject list project api
@param formatter
**/
func ListProject(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
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
		list, err := client.ListProject(fieldSelector, labelSelector, limit)
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
