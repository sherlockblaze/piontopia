package handler

import (
	"net/http"

	"github.com/unrolled/render"
)

// CreateService create service
func CreateService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
	}
}

// GetService get target service
func GetService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
	}
}

// UpdateService update service
func UpdateService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
	}
}

// DeleteService delete service
func DeleteService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
	}
}

// ListService list service
func ListService(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
	}
}
