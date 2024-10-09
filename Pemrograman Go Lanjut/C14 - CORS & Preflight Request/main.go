package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

func main() {
	e := echo.New()

	e.GET("/index", func(c echo.Context) error {
		return c.String(http.StatusOK, "hello")
	})

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"https://novalagung.com", "https://www.google.com", "https://dasarpemrogramangolang.novalagung.com"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
		AllowedHeaders: []string{"Content-Type", "X-CSRF-Token"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.Logger.Fatal(e.Start(":9000"))
}
