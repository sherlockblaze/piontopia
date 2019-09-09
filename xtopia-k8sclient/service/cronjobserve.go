package service

import (
	"k8s-client/handler"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func initCronJobRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/cronjob/create", handler.CreateCronJob(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/cronjob/{cronJobName}", handler.GetCronJob(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/cronjobs", handler.ListCronJob(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/cronjob", handler.UpdateCronJob(formatter)).Methods("PUT")
	mx.HandleFunc("/console/{project}/cronjob/{cronJobName}/delete", handler.DeleteCronJob(formatter)).Methods("DELETE")
}
