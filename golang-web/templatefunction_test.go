package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + " from " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello "John"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "John Doe",
	})
}

func TestTemplateFunction(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{len ".Name"}}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "John Doe",
	})
}

func TestTemplateFunctionGlobal(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper" : func (value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{upper .Name}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "John Doe",
	})
}

func TestTemplateCreateGlobal(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateCreateGlobal(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionPipelines(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello" : func (name string) string {
			return "Hello " + name
		},
		"upper" : func (value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))

	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name: "John Doe",
	})
}

func TestTemplateFunctionPipelines(t *testing.T){
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionPipelines(recorder, request)

	body, _ :=  io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
