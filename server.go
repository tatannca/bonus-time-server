package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/wd30gsrc/bonus-time-server/customMiddleware"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("環境変数読み込み失敗: %v", err)
	}
}

func main() {
	loadEnv()

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
	}, customMiddleware.Auth())
	
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
