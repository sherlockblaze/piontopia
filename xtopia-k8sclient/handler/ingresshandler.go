package handler

import (
	"encoding/json"
	"io/ioutil"
	"k8s-client/client"
	"net/http"

	"github.com/gorilla/mux"
	betav1 "k8s.io/api/networking/v1beta1"

	"github.com/unrolled/render"
)

/*
CreateIngress create ingress api
@param formatter
**/
func CreateIngress(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var ingress *betav1.Ingress
		json.Unmarshal(body, &ingress)
		createdIngress, err := client.CreateIngress(project, ingress)
		status := http.StatusCreated
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = createdIngress
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
GetIngress get ingress api
@param formatter
**/
func GetIngress(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		ingressName := vars["ingressName"]
		ingress, err := client.GetIngress(project, ingressName)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = ingress
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
UpdateIngress update ingress api
@param formatter
**/
func UpdateIngress(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var ingress *betav1.Ingress
		json.Unmarshal(body, &ingress)
		updatedIngress, err := client.UpdateIngress(project, ingress)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["error"] = err.Error()
			res["code"] = -1
		}
		res["result"] = updatedIngress
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
DeleteIngress delete ingress api
@param formatter
**/
func DeleteIngress(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		ingressName := vars["ingressName"]
		err := client.DeleteIngress(project, ingressName)
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
ListIngress list ingress api
@param formatter
**/
func ListIngress(formatter *render.Render) http.HandlerFunc {
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
		list, err := client.ListIngress(project, fieldSelector, labelSelector, limit)
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
