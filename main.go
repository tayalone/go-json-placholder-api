package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type todo struct {
	UserID    uint   `json:"userId"`
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func (t *todo) set(title string, completed bool) {
	t.Title = title
	t.Completed = completed
}

type todos []todo

func (ts todos) findAll() []todo {
	return ts
}

func (ts todos) findByID(i int) (todo, error) {
	if i < 0 || len(ts)-1 < i {
		return todo{}, errors.New("id doest does not exist")
	}
	return ts[i], nil
}

func (ts *todos) setByID(i int, title string, completed bool) (todo, error) {
	if len(*ts)-1 < i {
		return todo{}, errors.New("id doest does not exist")
	}
	(*ts)[i].set(title, completed)

	return (*ts)[i], nil
}

func (ts *todos) removeBy(i int) error {
	if len(*ts)-1 < i {
		return errors.New("id doest does not exist")
	}
	s := *ts
	s = append(s[:i], s[i+1:]...)
	*ts = s
	return nil
}

var todoRepo todos = todos{
	todo{
		UserID:    1,
		ID:        1,
		Title:     "Lorem 1",
		Completed: true,
	},
	todo{
		UserID:    1,
		ID:        2,
		Title:     "Lorem 2",
		Completed: true,
	},
}

func main() {
	r := gin.Default()
	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
		})
	})
	rTodo := r.Group("/todos")
	{

		rTodo.GET("/", func(c *gin.Context) {
			todos := todoRepo.findAll()

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"todos":   todos,
			})
		})
		rTodo.GET("/:id", func(c *gin.Context) {
			idStr := c.Param("id")

			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Fail",
					"key":     "ID_MUST_BE_NUMBER",
				})
				return
			}

			todo, err := todoRepo.findByID(int(id))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"message": "Fail",
					"key":     "NOT_FOUND_TODO",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
				"todo":    todo,
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
