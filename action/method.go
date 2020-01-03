package action

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// Human ...
type Human struct {
	Name string `json:"name,omitempty"`
}

// GetByQuery ...
func GetByQuery(c echo.Context) error {
	name := c.QueryParam("name")

	if name != "" {
		return c.String(http.StatusOK, fmt.Sprintf("name is %s", name))
	}

	return c.String(http.StatusBadRequest, "not have name")
}

// GetHuman ...
func GetHuman(c echo.Context) error {
	var human Human
	qStr := c.QueryParam("q")
	if qStr == "" {
		return c.String(http.StatusBadRequest, "Invalid")
	}

	err := json.Unmarshal([]byte(qStr), &human)

	if err != nil {
		return c.String(http.StatusInternalServerError, "error")
	}

	return c.JSON(http.StatusOK, human)
}

// AddHuman ...
func AddHuman(c echo.Context) error {
	var human Human
	bodyBytes, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(bodyBytes, &human)

	if err != nil {
		return c.String(http.StatusInternalServerError, "")
	}

	return c.JSON(http.StatusOK, human)
}
