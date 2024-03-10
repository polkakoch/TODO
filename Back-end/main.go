package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// Определяем маршруты и их обработчики
	e.GET("/home", homeHandler)
	e.GET("/login", loginHandler)
	e.GET("/register", registerHandler)

	// Запускаем сервер на порту 8080
	e.Logger.Fatal(e.Start(":8080"))
}

// Обработчик для домашней страницы
func homeHandler(c echo.Context) error {
	html := `
    <!DOCTYPE html>
    <html lang="en">
    <head>
        <meta charset="UTF-8">
        <title>Home Page</title>	
    </head>
    <body>
        <h1>Welcome to the Home Page!</h1>
        <a href="/login">Login</a>
    </body>
    </html>
    `
	return c.HTML(http.StatusOK, html)
}

// Обработчик для страницы входа в аккаунт
func loginHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Страница для входа в аккаунт")
}

// Обработчик для страницы регистрации
func registerHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Страница для регистрации")
}
