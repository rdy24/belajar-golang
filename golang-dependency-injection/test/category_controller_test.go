package test

import (
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"rdy24/golang-dependency-injection/app"
	"rdy24/golang-dependency-injection/controller"
	"rdy24/golang-dependency-injection/helper"
	"rdy24/golang-dependency-injection/middleware"
	"rdy24/golang-dependency-injection/model/domain"
	"rdy24/golang-dependency-injection/repository"
	"rdy24/golang-dependency-injection/service"
	"strconv"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/test_golang_restful_api")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler {
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func truncateTable(db *sql.DB) {
	db.Exec("TRUNCATE TABLE categories")
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "test"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["name"])
}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/api/categories", requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusBadRequest), responseBody["status"])
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "test"})
	tx.Commit()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": "test"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, "test", responseBody["data"].(map[string]interface{})["name"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: ""})
	tx.Commit()
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), requestBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 400, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusBadRequest), responseBody["status"])
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "test"})
	tx.Commit()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:8080/api/categories/404", nil)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusNotFound), responseBody["status"])
}

func TestGetCategorySuccess(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "test"})
	tx.Commit()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories/"+strconv.Itoa(category.Id), nil)
	request.Header.Set("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, category.Id, int(responseBody["data"].(map[string]interface{})["id"].(float64)))
	assert.Equal(t, category.Name, responseBody["data"].(map[string]interface{})["name"])
}

func TestGetCategoryFailed(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories/404", nil)
	request.Header.Set("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 404, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusNotFound), responseBody["status"])
}

func TestListCategoriesSuccess(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{Name: "test"})
	tx.Commit()
	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories", nil)
	request.Header.Set("X-API-KEY", "SECRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])

	var categories = responseBody["data"].([]interface{})
	categoryResponse := categories[0].(map[string]interface{})
	assert.Equal(t, category.Id, int(categoryResponse["id"].(float64)))
	assert.Equal(t, category.Name, categoryResponse["name"])
}

func TestUnathorized(t *testing.T) {
	db := setupTestDB()
	truncateTable(db)

	router := setupRouter(db)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/api/categories", nil)
	request.Header.Set("X-API-KEY", "SECasasdRET")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)
	assert.Equal(t, 401, int(responseBody["code"].(float64)))
	assert.Equal(t, http.StatusText(http.StatusUnauthorized), responseBody["status"])
}
