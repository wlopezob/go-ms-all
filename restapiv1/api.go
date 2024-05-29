package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr  string
	store Store
}

func NewAPIServer(addr string, store Store) *APIServer {
	return &APIServer{
		addr:  addr,
		store: store,
	}
}

func (s *APIServer) Serve() {
	r := gin.Default()
	api := r.Group("/ux-api/v1")

	// Add a middleware to all routes in the group
	api.Use(func(ctx *gin.Context) {
		// Do something before each request
	})

	// Routes, registering our services
	api.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	taskService := NewTaskService(s.store)
	taskService.RegisterRoutes(api)
	log.Println("Starting the API server at", s.addr)
	r.Run(s.addr)
}
