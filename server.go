package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":5000"))
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 			c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
// 			c.Header("Access-Control-Allow-Credentials", "true")
// 			c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 			c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE")
// 			c.Next()
// 	}
// }

// func public(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "hello public!",
// 	})
// }
