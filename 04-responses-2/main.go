package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Person struct {
	Firstname string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	e := echo.New()
	p := Person{
		Firstname: "Nicolás",
		LastName:  "Lacave",
		Age:       22,
	}
	e.GET("/json", func(c echo.Context) error {
		return c.JSON(http.StatusOK, p)
	})

	e.GET("/xml", func(c echo.Context) error {
		return c.XML(http.StatusOK, p)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
