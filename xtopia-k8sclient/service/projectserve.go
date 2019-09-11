package service

import (
	"k8s-client/handler"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initProjectRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/project/create", handler.CreateProject(formatter)).Methods("POST")
	mx.HandleFunc("/console/project/{projectName}", handler.GetProject(formatter)).Methods("GET")
	mx.HandleFunc("/console/project/{projectName}/delete", handler.DeleteProject(formatter)).Methods("DELETE")
	mx.HandleFunc("/console/projects", handler.ListProject(formatter)).Methods("GET")
	mx.HandleFunc("/console/project", handler.UpdateProject(formatter)).Methods("PATCH")
}
