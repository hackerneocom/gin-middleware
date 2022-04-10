package middleware

import "github.com/gin-gonic/gin"

type (
	JWTConfig struct {
	}
)

var (
	DefaultJWTConfig = JWTConfig{}
)

func JWT() gin.HandlerFunc {
	config := DefaultJWTConfig
	return JWTWithConfig(config)
}

func JWTWithConfig(config JWTConfig) gin.HandlerFunc {
	// Defaults
	return func(c *gin.Context) {

	}
}
