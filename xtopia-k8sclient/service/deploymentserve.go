package service

import (
	"k8s-client/handler"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initDeploymentRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/deployment/create", handler.CreateDeployment(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/deployment/{deploymentName}", handler.GetDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/deployments", handler.ListDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/deployment", handler.UpdateDeployment(formatter)).Methods("PATCH")
	mx.HandleFunc("/console/{project}/deployment/{deploymentName}/delete", handler.DeleteDeployment(formatter)).Methods("DELETE")
}
