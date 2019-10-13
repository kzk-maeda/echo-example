package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// Get Req at /
	e.GET("/", func(c echo.Context) error {
		fmt.Println("request GET")
		name := c.QueryParam("name")
		return c.String(http.StatusOK, "Hello "+name)
	})

	// Post Req at /
	e.POST("/", func(c echo.Context) error {
		fmt.Println("Request POST")
		name := c.FormValue("name")
		return c.String(http.StatusOK, "Hello "+name)
	})

	// Post Req at /:name
	e.POST("/:name", func(c echo.Context) error {
		fmt.Println("Request POST")
		name := c.Param("name")
		return c.String(http.StatusOK, "Hello "+name)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
