package service

import (
	"k8s-client/handler"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initIngressRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/ingress/create", handler.CreateIngress(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/ingress/{ingressName}", handler.GetIngress(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/ingresss", handler.ListIngress(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/ingress", handler.UpdateIngress(formatter)).Methods("PATCH")
	mx.HandleFunc("/console/{project}/ingress/{ingressName}/delete", handler.DeleteIngress(formatter)).Methods("DELETE")
}
