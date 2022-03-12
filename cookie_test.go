package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-Master-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	_, _ = fmt.Fprint(writer, "Success Create Cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-Master-Name")
	if err != nil {
		_, _ = fmt.Fprint(writer, "No Cookie")
	} else {
		name := cookie.Value
		_, _ = fmt.Fprint(writer, name)
	}
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
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080?name=Master", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	response := recorder.Result().Cookies()

	for _, cookie := range response {
		fmt.Printf("Cookie %s:%s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-Master-Name"
	cookie.Value = "Eko"
	request.AddCookie(cookie)
	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
