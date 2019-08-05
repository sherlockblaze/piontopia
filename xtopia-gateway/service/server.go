package service

import (
	"github.com/sherlockblaze/piontopia/xtopia-gateway/handler"
	"github.com/sherlockblaze/piontopia/xtopia-gateway/model"

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
	mx.HandleFunc("/user/create", handler.CreateUser(formatter, &model.UserInfo{})).Methods("POST")
	mx.HandleFunc("/user/delete", handler.DeleteUser(formatter, "abc")).Methods("PUT")
	mx.HandleFunc("/user/info", handler.CheckoutUser(formatter, "abc")).Methods("GET")
}
