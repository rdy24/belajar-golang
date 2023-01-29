package main

import (
	"net/http"
	"rdy24/golang-restful-api/app"
	"rdy24/golang-restful-api/controller"
	"rdy24/golang-restful-api/helper"
	"rdy24/golang-restful-api/middleware"
	"rdy24/golang-restful-api/repository"
	"rdy24/golang-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {

	validate := validator.New()
	db := app.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
