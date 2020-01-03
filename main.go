package main

import (
	"net/http"

	"github.com/labstack/echo"

	"github.com/hoangduc02011998/golang-echo/action"
)

func main() {

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.POST("/human", action.AddHuman)
	e.GET("/human", action.GetByQuery)
	e.GET("/human1", action.GetHuman)

	e.Start(":3000")
}
