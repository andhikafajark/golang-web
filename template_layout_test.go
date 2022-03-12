package golang_web

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionLayout(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/layout.gohtml", "./templates/footer.gohtml"))

	err := t.ExecuteTemplate(writer, "layout", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Master",
	})
	if err != nil {
		panic(err)
	}
}

func TestTemplateActionLayout(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionLayout(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
