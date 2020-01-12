package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hoangduc02011998/golang-echo/action"
)

type Account struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// jwt claims ...
type JwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

/// custome middleware
func serverHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("server", "Edar/Ha")

		return next(c)
	}
}

func login(c echo.Context) error {
	var account Account
	accBytes, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return c.String(http.StatusBadRequest, "something went wrong")
	}
	json.Unmarshal(accBytes, &account)

	if account.Username == "edar" && account.Password == "123" {
		token, err := createJwtToken()

		if err != nil {
			return c.String(http.StatusInternalServerError, "something went wrong")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "you were logged in",
			"token":   token,
		})
	}

	return c.String(http.StatusBadRequest, "username or password not match")
}

func mainJwt(c echo.Context) error {
	user := c.Get("user")
	token := user.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	fmt.Println("name", claims["name"])
	return c.String(http.StatusOK, "had token")
}

func createJwtToken() (string, error) {
	claims := JwtClaims{
		"edar",
		jwt.StandardClaims{
			Id:        "browser_id",
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte("mySecret"))

	if err != nil {
		return "", err
	}

	return token, nil
}

func main() {

	e := echo.New()

	e.Use(serverHeader)

	jwtGroup := e.Group("/jwt")
	jwtGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS512",
		SigningKey:    []byte("mySecret"),
	}))

	jwtGroup.GET("/main", mainJwt)

	e.POST("/login", login)

	//g := e.Group("/admin")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "[${time_rfc3339}]  ${status}  ${method}${path} ${latency_human}" + "\n",
	}))

	// e.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
	// 	if username == "edar" && password == "123" {
	// 		return true, nil
	// 	}

	// 	return false, nil
	// }))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello world")
	})

	e.POST("/human", action.AddHuman)
	e.GET("/human", action.GetByQuery)
	e.GET("/human1", action.GetHuman)

	e.Start(":3000")
}
