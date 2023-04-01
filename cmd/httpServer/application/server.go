package application

import (
	"github.com/gin-gonic/gin"
	"github.com/lara-store/damo/pkg/config"
	"log"
	"net/http"
	"time"
)

type Engine struct {
	*gin.Engine
}

func New() *Engine {
	return &Engine{gin.New()}
}

func (route Engine) Listen() error {
	var port = config.Get[string]("app.port", "8000")

	route.boot()
	route.load()

	server := &http.Server{
		Addr:           ":" + port,
		Handler:        route,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server is running on port " + port)
	return server.ListenAndServe()
}
