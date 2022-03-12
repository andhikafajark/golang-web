package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("Content-Type")

	_, _ = fmt.Fprint(writer, contentType)
}

func TestRequestHeader(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	request.Header.Add("Content-Type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	bodyString := string(body)

	fmt.Println(bodyString)
}

func ResponseHeader(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Powered-By", "Master")

	_, _ = fmt.Fprint(writer, "Ok")
}

func TestResponseHeader(t *testing.T) {
	request, _ := http.NewRequest(http.MethodPost, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)

	poweredBy := recorder.Header().Get("X-Powered-By")

	fmt.Println(poweredBy)
}
