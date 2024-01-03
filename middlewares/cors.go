package middlewares

import "github.com/gin-gonic/gin"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// Define your allowed domains
		var allowedOrigins = []string{"https://web.xairline.org", "https://xairline.org", "https://www.xairline.org", "http://localhost:3000"}

		// Check if the request's origin is in the list of allowed origins
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
