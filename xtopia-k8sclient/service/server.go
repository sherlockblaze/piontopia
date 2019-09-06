package service

import (
	"k8s-client/handler"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer new client server
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)

	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	initDeploymentRoutes(mx, formatter)
	initServiceRoutes(mx, formatter)
	initJobRoutes(mx, formatter)
	initCronJobRoutes(mx, formatter)
}

func initDeploymentRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/deployment/create", handler.CreateDeployment(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/deployment/{deploymentName}", handler.GetDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/deployment/list", handler.ListDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/deployment/update", handler.UpdateDeployment(formatter)).Methods("PUT")
	mx.HandleFunc("/console/{project}/deployment/{deploymentName}/delete", handler.DeleteDeployment(formatter)).Methods("DELETE")
}

func initServiceRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/service/create", handler.CreateDeployment(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/service/{serviceName}", handler.GetDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/service/list", handler.ListDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/service/update", handler.UpdateDeployment(formatter)).Methods("PUT")
	mx.HandleFunc("/console/{project}/service/{serviceName}/delete", handler.DeleteDeployment(formatter)).Methods("DELETE")
}

func initJobRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/job/create", handler.CreateDeployment(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/job/{jobName}", handler.GetDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/job/list", handler.ListDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/job/update", handler.UpdateDeployment(formatter)).Methods("PUT")
	mx.HandleFunc("/console/{project}/job/{jobName}/delete", handler.DeleteDeployment(formatter)).Methods("DELETE")
}

func initCronJobRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/console/{project}/cronjob/create", handler.CreateDeployment(formatter)).Methods("POST")
	mx.HandleFunc("/console/{project}/cronjob/{cronJobName}", handler.GetDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/cronjob/list", handler.ListDeployment(formatter)).Methods("GET")
	mx.HandleFunc("/console/{project}/cronjob/update", handler.UpdateDeployment(formatter)).Methods("PUT")
	mx.HandleFunc("/console/{project}/cronjob/{cronJobName}/delete", handler.DeleteDeployment(formatter)).Methods("DELETE")
}
