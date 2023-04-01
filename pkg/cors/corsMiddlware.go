package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func AllowAll() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	return cors.New(config)
}
