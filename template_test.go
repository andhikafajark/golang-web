package golang_web

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func SimpleHTML(writer http.ResponseWriter, request *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	//t, err := template.New("SIMPLE").Parse(templateText)
	//if err != nil {
	//	panic(err)
	//}

	t := template.Must(template.New("SIMPLE").Parse(templateText))

	err := t.ExecuteTemplate(writer, "SIMPLE", "Hello HTML Template")
	if err != nil {
		panic(err)
	}
}

func TestSimpleHTML(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTML(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	err := t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
	if err != nil {
		panic(err)
	}
}

func TestSimpleHTMLFile(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	SimpleHTMLFile(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateDirectory(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	err := t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
	if err != nil {
		panic(err)
	}
}

func TestTemplateDirectory(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateDirectory(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templates embed.FS

func TemplateEmbed(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFS(templates, "templates/*.gohtml"))

	err := t.ExecuteTemplate(writer, "simple.gohtml", "Hello HTML Template")
	if err != nil {
		panic(err)
	}
}

func TestTemplateEmbed(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
