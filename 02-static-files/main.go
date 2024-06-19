package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	//e.File("/", "public/index.html")
	//e.File("/styles.css", "public/styles.css")
	//e.File("/script.js", "public/script.js")
	e.Static("/", "public/pagina-principal")
	e.Logger.Fatal(e.Start(":8080"))
}
