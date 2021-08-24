package middleware

import (
	"github.com/gin-gonic/gin"
)

// Mnemonic is a middleware for share seed.
func Mnemonic(seed string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("Seed", seed)
		c.Next()
	}
}