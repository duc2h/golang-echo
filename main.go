package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/hoangduc02011998/golang-echo/action"
)

/// custome middleware
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("server", "Edar/Ha")

		return next(c)
	}
}

func main() {

	e := echo.New()

	e.Use(serverHeader)

	//g := e.Group("/admin")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}]  ${status}  ${method}${path} ${latency_human}" + "\n",
	}))

	e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "edar" && password == "123" {
			return true, nil
		}

		return false, nil
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.POST("/human", action.AddHuman)
	e.GET("/human", action.GetByQuery)
	e.GET("/human1", action.GetHuman)

	e.Start(":3000")
}
