package service

import (
	"github.com/sherlockblaze/piontopia/xtopia-gateway/handler"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server
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
	initUserRoutes(mx, formatter)
	initMissionRoutes(mx, formatter)
}

func initUserRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/user/create", handler.CreateUser(formatter)).Methods("POST")
	mx.HandleFunc("/user/{id}/delete", handler.DeleteUser(formatter)).Methods("PUT")
	mx.HandleFunc("/user/{id}", handler.CheckoutUser(formatter)).Methods("GET")
}

func initMissionRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/mission/{user_id}/list", handler.ShowAllMissions(formatter)).Methods("GET")
	mx.HandleFunc("/mission/{user_id}/received", handler.ShowReceviedMissions(formatter)).Methods("GET")
}
