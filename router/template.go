package router

import (
	"io"
	"text/template"

	"github.com/labstack/echo/v4"
)

type Template struct {
	tmpl *template.Template
}

func newTemplate() *Template {
	t := template.New("views")
	t.Funcs(template.FuncMap{
		"inc": increment,
		"arr": arr,
	})
	t.ParseGlob("views/*.html")
	return &Template{ tmpl: t }
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.tmpl.ExecuteTemplate(w, name, data)
}

func increment(n int) int {
	return n + 1
}

func arr(els ...any) []any {
	// {{ template "__question.html" (arr $quiz $question) }}
	return els
}
