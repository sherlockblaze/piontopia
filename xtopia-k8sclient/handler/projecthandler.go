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
		createdProject, _ := client.CreateProject(project)
		formatter.JSON(w, http.StatusCreated,
			struct {
				Project v1.Namespace
			}{*createdProject})
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
		project, _ := client.GetProject(projectName)
		formatter.JSON(w, http.StatusOK,
			struct {
				Project v1.Namespace
			}{*project})
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
		client.DeleteProject(projectName)
		formatter.JSON(w, http.StatusOK,
			struct{}{})
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
		list, _ := client.ListProject(fieldSelector, labelSelector, limit)
		var projectList []v1.Namespace
		if list != nil {
			projectList = list.Items
		}
		formatter.JSON(w, http.StatusOK,
			struct {
				ProjectList []v1.Namespace
			}{projectList})
	}
}
