package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()

	e.Static("/", "public")
	e.GET("/gopher", func(c echo.Context) error {
		return c.Inline("img/gopher.svg", "gopher.svg")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
