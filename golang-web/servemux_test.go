package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World!")
	})
	mux.HandleFunc("/hai", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Goodbye!")
	})
	mux.HandleFunc("/hello/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello!")
	})

	mux.HandleFunc("/hello/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello!")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}