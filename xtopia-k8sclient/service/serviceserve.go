package service

import (
	"k8s-client/handler"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initServiceRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/service/create", handler.CreateService(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/service/{serviceName}", handler.GetService(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/services", handler.ListService(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/service", handler.UpdateService(formatter)).Methods("PATCH")
	mx.HandleFunc("/console/{project}/service/{serviceName}/delete", handler.DeleteService(formatter)).Methods("DELETE")
}
