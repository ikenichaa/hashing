package handler

import "github.com/labstack/echo/v4"

func (h *Handler) InitRoutes(e *echo.Echo) {
	e.GET("/", h.Sha256)
}
