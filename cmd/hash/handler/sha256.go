package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HashRequest struct {
	Data string `json:"data"`
}

func (h *Handler) Sha256(c echo.Context) error {
	fmt.Println(c.QueryParam("data"))
	fmt.Println(c.FormValue("dd"))
	var u HashRequest
	if err := c.Bind(&u); err != nil {
		fmt.Println(err)
	}
	fmt.Println(u)
	return c.String(http.StatusOK, "Hello, From sha256!")
}
