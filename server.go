package main

import (
	"context"
	"log"
	"net/http"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://bonustime.vercel.app"},
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

func initializeAppWithServiceAccount() *firebase.App {
	// [START initialize_app_service_account_golang]
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	// [END initialize_app_service_account_golang]

	return app
}
