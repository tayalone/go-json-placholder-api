package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("x")
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	rTodo := r.Group("/todos")
	{
		rTodo.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"from":    "Get /todos",
				"message": "OK",
			})
		})
		rTodo.GET("/:id", func(c *gin.Context) {
			id := c.Param("id")

			c.JSON(http.StatusOK, gin.H{
				"from":    fmt.Sprintf("Get /todos/%s", id),
				"message": "OK",
			})
		})
		rTodo.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"from":    "POST /todos",
				"message": "OK",
			})
		})
		rTodo.PATCH("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"from":    "PATH /todos",
				"message": "OK",
			})
		})
		rTodo.DELETE("/:id", func(c *gin.Context) {
			id := c.Param("id")

			c.JSON(http.StatusOK, gin.H{
				"from":    fmt.Sprintf("DELETE /todos/%s", id),
				"message": "OK",
			})
		})
	}
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
