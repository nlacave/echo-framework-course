package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hola Mundo con string")
	})

	e.GET("/html", func(c echo.Context) error {
		return c.HTML(http.StatusOK,
			`<h1>Hola Mundo con HTML</h1>
			<p>Este es un párrafo</p>`)
	})

	e.GET("/no-content", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	e.Logger.Fatal(e.Start(":8080"))
}
