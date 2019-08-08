package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/unrolled/render"
)

// CreateUser create user
func CreateUser(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusCreated, struct {
			Test string
		}{"This is a Test"})
	}
}

// DeleteUser delete a user
func DeleteUser(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{Test: vars["id"]})
	}
}

// CheckoutUser check user's info
func CheckoutUser(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{Test: vars["id"]})
	}
}
