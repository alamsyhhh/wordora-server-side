package middlewares

import (
	"strings"
	"wordora/app/utils/common"
	"wordora/app/utils/paseto"

	"net/http"

	"github.com/gin-gonic/gin"
)

func extractToken(authHeader string) string {
	if strings.HasPrefix(authHeader, "Bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer ")
	}
	return authHeader
}

func AuthMiddleware(tokenHelper *paseto.TokenHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			common.GenerateErrorResponse(c, http.StatusUnauthorized, "Missing token", nil)
			return
		}

		token := extractToken(authHeader)

		payload, err := tokenHelper.ValidateToken(token)
		if err != nil {
			common.GenerateErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token", nil)
			return
		}

		userID, ok := payload["sub"].(string)
		if !ok {
			common.GenerateErrorResponse(c, http.StatusUnauthorized, "Invalid token payload", nil)
			return
		}

		// log.Printf("Token Payload: %+v", payload)

		c.Set("sub", userID)
		c.Next()
	}
}

func AdminOnly(tokenHelper *paseto.TokenHelper) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			common.GenerateErrorResponse(c, http.StatusUnauthorized, "Missing token", nil)
			return
		}

		token := extractToken(authHeader)

		payload, err := tokenHelper.ValidateToken(token)
		if err != nil {
			common.GenerateErrorResponse(c, http.StatusUnauthorized, "Invalid or expired token", nil)
			return
		}

		role, ok := payload["role"].(string)
		if !ok || role != "admin" {
			common.GenerateErrorResponse(c, http.StatusForbidden, "Access denied", nil)
			return
		}

		c.Next()
	}
}
