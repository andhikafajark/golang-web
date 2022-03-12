package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprint(writer, "Name is empty")
	} else {
		_, _ = fmt.Fprint(writer, name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(bodyString)
}

func TestResponseCodeValid(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "http://localhost:8080?name=Master", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(bodyString)
}
