package customMiddleware

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	// "github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch host := r.Header.Get("Origin"); host {
		case "http://localhost:3000":
			w.Header().Set("Access-Control-Allow-Origin", host)
		case "https://bonustime.vercel.app":
			w.Header().Set("Access-Control-Allow-Origin", host)
		}
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		next.ServeHTTP(w, r)
	})
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// [START initialize_app_service_account_golang]
		opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			log.Fatalf("error initializing app: %v\n", err)
			return
		}
		// [END initialize_app_service_account_golang]
		client, err := app.Auth(context.Background())
		if err != nil {
			log.Fatalf("error getting Auth client: %v\n", err)
			return
		}
		// JWT取得
		authHeader := r.Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

		// JWT検証
		token, err := client.VerifyIDToken(context.Background(), idToken)
		log.Printf("token: %v\n", token)
		// JWTが無効ならエラーを返す
		if err != nil {
			log.Printf("error getting token: %v\n", err)
			message := struct {Message error `json:"message"`}{Message: err}
			data, err := json.Marshal(message)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// func Auth() echo.MiddlewareFunc {
// 	return auth
// }

// func auth(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		// [START initialize_app_service_account_golang]
// 		opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
// 		app, err := firebase.NewApp(context.Background(), nil, opt)
// 		if err != nil {
// 			log.Fatalf("error initializing app: %v\n", err)
// 			return err
// 		}
// 		// [END initialize_app_service_account_golang]

// 		client, err := app.Auth(context.Background())
// 		if err != nil {
// 			log.Fatalf("error getting Auth client: %v\n", err)
// 			return err
// 		}

// 		// JWT取得
// 		authHeader := c.Request().Header.Get("Authorization")
// 		idToken := strings.Replace(authHeader, "Bearer ", "", 1)

// 		// JWT検証
// 		token, err := client.VerifyIDToken(context.Background(), idToken)
// 		if err != nil {
// 			// JWTが無効ならエラーを返す
// 			log.Printf("error getting token: %v\n", err)
// 			message := struct {Message error `json:"message"`}{Message: err}
// 			return c.JSON(http.StatusUnauthorized, message)
// 		}

// 		c.Set("access_token", token)
// 		return next(c)
// 	}
// }