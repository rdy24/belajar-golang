package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHtml(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`

	t := template.Must(template.New("simple").Parse(templateText))
	t.ExecuteTemplate(writer, "simple", "Hello World")
}

func TestSimpleHtml(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtml(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHtmlFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("template/simple.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Html Template")
}

func TestSimpleHtmlFile(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHtmlFile(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("template/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Html Template")
}

func TestTemplateDirectory(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}



func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "template/*.gohtml"))
	t.ExecuteTemplate(writer, "simple.gohtml", "Hello Html Template")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}