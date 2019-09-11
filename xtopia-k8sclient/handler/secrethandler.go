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
CreateSecret create secret api
@param formatter
**/
func CreateSecret(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		body, _ := ioutil.ReadAll(req.Body)
		var secret *apiv1.Secret
		json.Unmarshal(body, &secret)
		createdSecret, err := client.CreateSecret(project, secret)
		status := http.StatusCreated
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = createdSecret
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
GetSecret get secret api
@param formatter
**/
func GetSecret(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		secretName := vars["secretName"]
		secret, err := client.GetSecret(project, secretName)
		status := http.StatusOK
		res := make(map[string]interface{})
		res["code"] = 0
		if err != nil {
			status = http.StatusBadRequest
			res["code"] = -1
			res["error"] = err.Error()
		}
		res["result"] = secret
		w.WriteHeader(status)
		w.Header().Add("Content-Type", "application/json")
		formatter.JSON(w, status, res)
	}
}

/*
DeleteSecret delete secret api
@param formatter
**/
func DeleteSecret(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		project := vars["project"]
		secretName := vars["secretName"]
		err := client.DeleteSecret(project, secretName)
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
