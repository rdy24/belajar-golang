package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", Page{
		Title: "Template Data Map",
		Name:  "John Doe",
	})
}

func TestTemplateActionIf(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"FinalValue":  90,
	})
}

func TestTemplateActionOperator(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Hobbies":  []string{"Reading", "Writing", "Coding"},
	})
}

func TestTemplateActionRange(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/address.gohtml"))
	t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name" : "John Doe",
		"Address" : map[string]interface{}{
			"Street" : "Jl. Jendral Sudirman",
			"City" : "Jakarta",
		},
	})
}

func TestTemplateActionWith(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}