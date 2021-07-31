package main

import (
	"hashing/cmd/hash/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	h := handler.NewHandler()
	h.InitRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}
