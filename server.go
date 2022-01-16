package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
	}))

	type Message struct {
		Public string `json:"message"`
	}
	
	e.GET("/", func(c echo.Context) error {
		m := &Message{
			Public: "Hello Public World.",
		}
		return c.JSON(http.StatusOK, m)
	})
	e.Logger.Fatal(e.Start(":5000"))
}
