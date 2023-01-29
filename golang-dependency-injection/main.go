package main

import (
	"net/http"
	"rdy24/golang-dependency-injection/helper"
	"rdy24/golang-dependency-injection/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: authMiddleware,
	}
}

func main() {

	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
