package handler

import (
	"net/http"
	"piontopia/xtopia-web/model"

	"github.com/unrolled/render"
)

// CreateUser create user
func CreateUser(formatter *render.Render, userInfo *model.UserInfo) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusCreated, struct {
			Test string
		}{"This is a Test"})
	}
}

// DeleteUser delete a user
func DeleteUser(formatter *render.Render, userID string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{"This is a Test"})
	}
}

// CheckoutUser check user's info
func CheckoutUser(formatter *render.Render, userID string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct {
			Test string
		}{"This is a Test"})
	}
}
