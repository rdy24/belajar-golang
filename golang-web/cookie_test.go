package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "my-cookie"
	cookie.Value = "some value"
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprint(writer, "Cookie has been set")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("my-cookie")
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		fmt.Fprint(writer, "No cookie found")
	}
	fmt.Fprint(writer, "Your cookie: ", cookie)
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/set-cookie?name=test", nil)
	recoder := httptest.NewRecorder()

	SetCookie(recoder, request)

	cookies := recoder.Result().Cookies()
	fmt.Println(cookies)
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/get-cookie?name=test", nil)
	recoder := httptest.NewRecorder()

	cookie := new(http.Cookie)
	cookie.Name = "my-cookie"
	cookie.Value = "some value"
	request.AddCookie(cookie)

	recoder.Result().Cookies()
	GetCookie(recoder, request)

	body, _ := io.ReadAll(recoder.Result().Body)
	fmt.Println(string(body))
}