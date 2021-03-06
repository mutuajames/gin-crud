package main

import (
	"gin-CRUD/controllers"
	"gin-CRUD/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	db := models.SetupDB()
	db.AutoMigrate(&models.Task{})
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("tasks/:id", controllers.DeleteTask)
	r.Run()
}
