package main

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
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
	e.POST("/login", postLoginHandler)
	e.GET("/register", registerHandler)

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}

// Обработчик для домашней страницы, до создания полноценной страницы home ее не трогать
func homeHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", nil)
}

// Обработчик для страницы входа в аккаунт
func loginHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "login.html", nil)
}

// Обработчик для обработки POST-запроса на страницу входа
func postLoginHandler(c echo.Context) error {
	// Получение данных из формы
	login := c.FormValue("Email")
	password := c.FormValue("Password")

	// Проверка данных пользователя 
	if login == "login-test" && password == "password123" {
		// Данные верны, перенаправление на страницу home
		return c.Redirect(http.StatusSeeOther, "/home")
	}

	// Если данные неверны
	return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
}

// Обработчик для страницы регистрации
func registerHandler(c echo.Context) error {
	return c.Render(http.StatusOK, "register.html", nil)
}

