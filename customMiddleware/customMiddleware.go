package customMiddleware

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

func Auth() echo.MiddlewareFunc {
	return auth
}

func auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// [START initialize_app_service_account_golang]
		opt := option.WithCredentialsFile(os.Getenv("CREDENTIALS"))
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