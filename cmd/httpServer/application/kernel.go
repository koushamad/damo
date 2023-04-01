package application

import (
	"github.com/gin-gonic/gin"
	"github.com/lara-store/damo/pkg/cors"
)

func (route Engine) boot() {
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	route.Use(gin.ErrorLogger())
	route.Use(cors.AllowAll())
}
