package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/*", func(c echo.Context) error {
		return c.HTML(http.StatusNotFound, "<p>No encontrado</p>")
	})

	e.GET("/gopher", func(c echo.Context) error {
		return c.HTML(http.StatusNotFound, "<p>Encontrado. Especifique el archivo en el hostname</p>")
	})

	e.GET("/gopher/:name", func(c echo.Context) error {
		p := c.Param("name")
		if p == "svg" {
			return c.Inline("img/gopher.svg", "gopher.svg")
		}
		if p == "jpg" {
			return c.Inline("img/gopher.jpg", "gopher.jpg")
		}

		if p == "svg-download" {
			return c.Attachment("img/gopher.svg", "gopher.svg")
		}

		if p == "jpg-download" {
			return c.Attachment("img/gopher.jpg", "gopher.jpg")
		}
		return c.HTML(http.StatusNotFound, "<p>No encontrado</p>")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
