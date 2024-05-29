package main

import "github.com/gin-gonic/gin"

type TaskService struct {
	store Store
}

func NewTaskService(store Store) *TaskService {
	return &TaskService{store: store}
}

func (s *TaskService) RegisterRoutes(api *gin.RouterGroup) {
	api.POST("/tasks", s.createTask)
	api.GET("/taks/:id", s.getTask)
}

func (s *TaskService) getTask(c *gin.Context) {
	c.JSON(200, "get task")
}

func (s *TaskService) createTask(c *gin.Context) {

}