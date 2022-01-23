package main

import (
	"context"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go/v4"

	"google.golang.org/api/option"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Auth() echo.MiddlewareFunc {
	return auth
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// [START initialize_app_service_account_golang]
		opt := option.WithCredentialsFile("bonus-time-app-firebase-adminsdk-hydva-b8d0f9557b.json")
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
			return err
		}
		// [END initialize_app_service_account_golang]

		client, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
			return err
		}

		// JWT取得
		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT検証
		token, err := client.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			// JWTが無効ならエラーを返す
			log.Printf("error getting token: %v\n", err)
			message := struct {Message error `json:"message"`}{Message: err}
			return c.JSON(http.StatusUnauthorized, message)
		}

		c.Set("access_token", token)
		return next(c)
	}
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", "https://bonustime.vercel.app"},
	}))

	type Message struct {
		SendMessage string `json:"message"`
	}
	
	e.GET("/public", func(c echo.Context) error {
		message := &Message{
			SendMessage: "Public API.",
		}
		return c.JSON(http.StatusOK, message)
	})

	e.GET("/private", func(c echo.Context) error {
		message := &Message{
			SendMessage: "Private API.",
		}
		return c.JSON(http.StatusOK, message)
	}, Auth())
	
	e.Logger.Fatal(e.Start(":5000"))
}

// func initializeAppWithServiceAccount() *firebase.App {
// 	// [START initialize_app_service_account_golang]
// 	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
// 	app, err := firebase.NewApp(context.Background(), nil, opt)
// 	if err != nil {
// 		log.Fatalf("error initializing app: %v\n", err)
// 	}
// 	// [END initialize_app_service_account_golang]

// 	return app
// }
