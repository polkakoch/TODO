package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo"
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	e.Static("/public", "public")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = renderer

	// Маршруты
	e.GET("/home", homeHandler)
	e.GET("/login", loginHandler)
	e.GET("/register", registerHandler)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}

// Обработчик для домашней страницы
func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

// Обработчик для страницы входа в аккаунт
func loginHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

// Обработчик для страницы регистрации
func registerHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}
