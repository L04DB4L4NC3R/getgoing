package controller

import (
	"net/http"

	"github.com/L04DB4L4NC3R/getgoing/section4/todo-api/model"
)

func create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			// take some data
			// save it!
			if err := model.CreateTodo(); err != nil {
				w.Write([]byte("Some error"))
				return
			}
		}
	}
}
