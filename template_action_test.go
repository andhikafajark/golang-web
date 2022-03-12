package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	err := t.ExecuteTemplate(writer, "if.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Master",
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionIf(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionOperator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	err := t.ExecuteTemplate(writer, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Data Map",
		"FinalValue": 50,
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionOperator(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	err := t.ExecuteTemplate(writer, "range.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Hobbies": []string{
			"Game", "Read", "Code",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionRange(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

	err := t.ExecuteTemplate(writer, "address.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Master",
		"Address": map[string]interface{}{
			"Street": "Unknown Street",
			"City":   "Unknown City",
		},
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionWith(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
