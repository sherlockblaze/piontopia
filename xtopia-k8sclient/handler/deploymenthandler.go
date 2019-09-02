package handler

import (
	"net/http"

	"github.com/unrolled/render"
)

// ListUserDeployment list user's deployment
func ListUserDeployment(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK,
			struct {
				Test string
			}{"This is a Test"})
	}
}
