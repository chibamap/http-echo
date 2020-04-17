package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func root(ctx echo.Context) error {
	return ctx.String(http.StatusOK, fmt.Sprintln(ctx.RealIP()))
}

func main() {
	e := echo.New()

	e.GET("/", root)
	e.Logger.Fatal(e.Start(":80"))
}
