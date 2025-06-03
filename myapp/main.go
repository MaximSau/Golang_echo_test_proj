package main

import (
	"html/template"
	"myapp/handlers"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Подготовка шаблонов для `TemplateRenderer`
	templatesGlobPattern := filepath.Join("templates", "*.html")
	templates, err := template.ParseGlob(templatesGlobPattern)
	if err != nil {
		e.Logger.Fatal(err)
	}
	// Создаем экземпляр `TemplateRenderer` и передаем ему загруженные шаблоны
	renderer := &handlers.TemplateRenderer{
		Templates: templates,
	}
	// Устанавливаем `renderer` для Echo
	e.Renderer = renderer

	e.GET("/", handlers.Home)
	e.GET("/secondary/", handlers.SecPage)
	e.Logger.Fatal(e.Start(":8000"))
}
