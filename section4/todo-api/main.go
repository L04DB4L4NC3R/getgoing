package main

import (
	"net/http"

	"github.com/L04DB4L4NC3R/getgoing/section4/todo-api/controller"
)

func main() {
	mux := controller.Register()
	http.ListenAndServe(":3000", mux)
}
