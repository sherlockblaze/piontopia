package service

import (
	"k8s-client/handler"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initSecretRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/secret/create", handler.CreateSecret(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/secret/{secretName}", handler.GetSecret(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/secret/{secretName}/delete", handler.DeleteSecret(formatter)).Methods("DELETE")
}
