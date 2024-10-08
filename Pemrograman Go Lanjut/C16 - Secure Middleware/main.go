package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/unrolled/secure"
)

func main() {
	e := echo.New()

	e.GET("/index", func(c echo.Context) error {
		c.Response().Header().Set("Access-Control-Allow-Origin", "*")

		return c.String(http.StatusOK, "Hello")
	})

	e.Logger.Fatal(e.StartTLS(":9000", "server.crt", "server.key"))

	secureMiddleware := secure.New(secure.Options{
		AllowedHosts:            []string{"localhost:9000", "www.google.com"},
		FrameDeny:               true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
	})

	e.Use(echo.WrapMiddleware(secureMiddleware.Handler))
}
