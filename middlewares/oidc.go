package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
)

// OIDCMiddleware creates a Gin middleware for OIDC token verification
func OIDCMiddleware(ctx context.Context, oidcVerifier *oidc.IDTokenVerifier) gin.HandlerFunc {
	return func(c *gin.Context) {
		rawToken := extractToken(c.Request)
		if rawToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			return
		}

		// Verify the token
		_, err := oidcVerifier.Verify(ctx, rawToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			return
		}

		// Token is valid; proceed with the request
		c.Next()
	}
}

// extractToken extracts the bearer token from a request
func extractToken(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	// Check if the Authorization header is in the correct format
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return ""
	}

	return parts[1]
}
