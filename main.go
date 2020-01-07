package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/hoangduc02011998/golang-echo/action"
)

func main() {

	e := echo.New()

	//g := e.Group("/admin")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}]  ${status}  ${method}${path} ${latency_human}" + "\n",
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.POST("/human", action.AddHuman)
	e.GET("/human", action.GetByQuery)
	e.GET("/human1", action.GetHuman)

	e.Start(":3000")
}
