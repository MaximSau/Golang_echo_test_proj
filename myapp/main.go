package main

import (
	"html/template"
	"myapp/db"
	"myapp/handlers"
	"myapp/models"
	"path/filepath"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if db.Err != nil {
		panic("failed to connect database")
	}
	// Migrate the schema
	db.Db.AutoMigrate(&models.User{})

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

	//url`s
	e.GET("/", handlers.Home)
	e.GET("/secondary/", handlers.SecPage)

	e.Logger.Fatal(e.Start(":8000"))
}
