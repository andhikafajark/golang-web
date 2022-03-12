package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("Function").Parse(`{{.SayHello "One"}}`))

	_ = t.ExecuteTemplate(writer, "Function", MyPage{Name: "Master"})
}

func TestTemplateFunction(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("Function").Parse(`{{len .Name}}`))

	_ = t.ExecuteTemplate(writer, "Function", MyPage{Name: "Master"})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("Function")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{upper .Name}}`))

	_ = t.ExecuteTemplate(writer, "Function", MyPage{Name: "Master"})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionCreateGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("Function")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})
	t = template.Must(t.Parse(`{{sayHello .Name | upper}}`))

	_ = t.ExecuteTemplate(writer, "Function", MyPage{Name: "Master"})
}

func TestTemplateFunctionCreateGlobalPipeline(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionCreateGlobalPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
