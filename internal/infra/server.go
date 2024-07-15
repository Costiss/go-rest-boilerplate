package infra

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

type Server struct {
	Server *http.Server
	Router *gin.Engine
}

func (s *Server) Serve() error {
	fmt.Println("Starting server on port: ", s.Server.Addr)
	return s.Server.ListenAndServe()
}

func (s *Server) Shutdown() error {
	fmt.Println("Shutting down server")
	return s.Server.Shutdown(context.Background())
}

func NewServer(lc fx.Lifecycle) *Server {
	router := gin.Default()

	router.Use(gin.Recovery())

	if os.Getenv("GIN_MODE") != "release" {
		router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	http := &http.Server{
		Addr:         getPort(),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	server := &Server{
		Server: http,
		Router: router,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go server.Serve()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Shutdown()
		},
	})

	return server
}

func getPort() string {
	port := 3000
	if os.Getenv("PORT") != "" {
		p, _ := strconv.Atoi(os.Getenv("PORT"))
		port = p
	}
	return fmt.Sprintf(":%d", port)
}
