package handlers

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

// TemplateRenderer - это структура, которая будет реализовывать интерфейс echo.Renderer
type TemplateRenderer struct {
	Templates *template.Template // Изменено на Templates
}

// Render реализует интерфейс echo.Renderer.
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data) // Изменено на t.Templates
}
