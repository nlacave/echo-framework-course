package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8989"},
	}))

	e.GET("/data", LoadData)
	e.Start(":8888")
}

func LoadData(c echo.Context) error {
	a := []struct {
		Name string `json:"name"`
	}{
		{"Leidy"},
		{"Carol"},
		{"Daniel"},
	}

	return c.JSON(http.StatusOK, a)
}
