package service

import (
	"k8s-client/handler"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initJobRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/job/create", handler.CreateJob(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/job/{jobName}", handler.GetJob(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/jobs", handler.ListJob(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/job", handler.UpdateJob(formatter)).Methods("PATCH")
	mx.HandleFunc("/console/{project}/job/{jobName}/delete", handler.DeleteJob(formatter)).Methods("DELETE")
}
