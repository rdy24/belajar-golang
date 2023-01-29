package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestPanic(t *testing.T) {
	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		fmt.Fprint(writer, "Panic : ",err)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Something went wrong!")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	assert.Equal(t, "Panic : Something went wrong!", string(body))
}