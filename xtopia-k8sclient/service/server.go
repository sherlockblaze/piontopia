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
}

func initDeploymentRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/deployment/{user_id}/list", handler.ListUserDeployment(formatter)).Methods("GET")
}
