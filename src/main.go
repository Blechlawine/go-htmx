package main

import (
	"context"

	"github.com/blechlawine/go-htmx-template/src/views"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	app.Use(middleware.Logger())

	app.GET("/", func(c echo.Context) error {
		component := views.Index("Hi there from Go")
		return component.Render(context.Background(), c.Response().Writer)
	})

	app.Static("/", "./static")

	app.Logger.Fatal(app.Start(":3000"))
}
