package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret []byte

func InitJWT() {
	if secret := os.Getenv("JWT_SECRET"); secret != "" {
		JWTSecret = []byte(secret)
		fmt.Println("[Payment Service] JWT_SECRET initialized:", secret)
	} else {
		JWTSecret = []byte("your-dev-secret-key")
		fmt.Println("[Payment Service] Using default JWT_SECRET")
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			fmt.Println("[Auth] No authorization header")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		fmt.Println("[Auth] Token received:", tokenString[:20]+"...")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return JWTSecret, nil
		})

		if err != nil {
			fmt.Println("[Auth] Token parse error:", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if !token.Valid {
			fmt.Println("[Auth] Token is not valid")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userId, exists := claims["userId"]; exists {
				c.Set("userID", uint(userId.(float64)))
				fmt.Println("[Auth] User authenticated:", uint(userId.(float64)))
			} else if userId, exists := claims["user_id"]; exists {
				c.Set("userID", uint(userId.(float64)))
				fmt.Println("[Auth] User authenticated:", uint(userId.(float64)))
			}
			if email, exists := claims["email"]; exists {
				c.Set("email", email.(string))
			}
		}

		c.Next()
	}
}
