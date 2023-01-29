package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprint(writer, "Welcome!")
	})

	server := http.Server{
		Handler : router,
		Addr : "localhost:8080",
	}

	server.ListenAndServe()
}