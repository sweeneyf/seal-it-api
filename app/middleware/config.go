package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sweeneyf/seal-it-api/pkg/config"
)

func AddConfig(config config.Configuration) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("config", config)
		c.Next()
	}
}
