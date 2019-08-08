package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// ShowAllMissions show all mission of one target User
func ShowAllMissions(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{Test: vars["user_id"]})
	}
}

// ShowReceviedMissions show all recevied mission of one target User
func ShowReceviedMissions(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{Test: vars["user_id"]})
	}
}
