package corsLibs

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ALLOWED_ORIGIN = []string{
		"http://localhost:8081",
	}
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, origin := range ALLOWED_ORIGIN {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		}

		c.Writer.Header().Set("Access-Allow-Control-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
