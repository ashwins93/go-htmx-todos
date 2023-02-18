package http

import (
	"html/template"
	"io"
	"log"

	"github.com/labstack/echo/v4"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data any, ctx echo.Context) error {
	log.Printf("Executing template %s\n", name)
	log.Printf("Available templates - %v\n", t.templates)
	return t.templates.ExecuteTemplate(w, name, data)
}

var _ echo.Renderer = (*Template)(nil)
