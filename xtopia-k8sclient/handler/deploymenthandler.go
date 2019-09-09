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
		createdDeployment, _ := client.CreateDeployment(project, deployment)
		formatter.JSON(w, http.StatusCreated,
			struct {
				Deployment appsv1.Deployment
			}{*createdDeployment})
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
		deployment, _ := client.GetDeployment(project, deploymentName)
		formatter.JSON(w, http.StatusOK,
			struct {
				Deployment appsv1.Deployment
			}{*deployment})
	}
}

/*
UpdateDeployment update deployment api
@param formatter
**/
func UpdateDeployment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
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
		client.DeleteDeployment(project, deploymentName)
		formatter.JSON(w, http.StatusOK, struct{}{})
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
		list, _ := client.ListDeployment(project, fieldSelector, labelSelector, limit)
		var deploymentList []appsv1.Deployment
		if list != nil {
			deploymentList = list.Items
		}
		formatter.JSON(w, http.StatusOK,
			struct {
				DeploymentList []appsv1.Deployment
			}{deploymentList})
	}
}
