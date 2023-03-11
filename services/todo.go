package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
)

type NewTodo struct {
	Todo        string `json:"todo" binding:"required"`
	IsCompleted bool   `json:"isCompleted" binding:"required"`
	Description string `json:"description"`
}

type UpdatedTodo struct {
	Todo        string `json:"todo"`
	IsCompleted bool   `json:"isCompleted"`
	Description string `json:"description"`
}

func GetAllTodos(c *gin.Context) {
	var todos []model.Todo

	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	if error = db.Find(&todos).Error; error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}

func GetTodo(c *gin.Context) {
	var todo model.Todo

	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	if error = db.Where("id= ?", c.Param("id")).First(&todo).Error; error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found!"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func PostTodo(c *gin.Context) {
	var todo NewTodo

	if error := c.ShouldBindJSON(&todo); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}

	newTodo := model.Todo{Todo: todo.Todo, IsCompleted: todo.IsCompleted, Description: todo.Description}

	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	if error := db.Create(&newTodo).Error; error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	}

	c.JSON(http.StatusCreated, newTodo)
}

func UpdateTodo(c *gin.Context) {
	var todo model.Todo

	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	if error := db.Where("id= ?", c.Param("id")).First(&todo).Error; error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
		return
	}

	var updateTodo UpdatedTodo

	if error := c.ShouldBindJSON(&updateTodo); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	if error := db.Model(&todo).Updates(model.Todo{Todo: updateTodo.Todo, IsCompleted: updateTodo.IsCompleted, Description: updateTodo.Description}).Error; error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	c.JSON(http.StatusOK, todo.Todo)
}

func DeleteTodo(c *gin.Context) {
	var todo model.Todo

	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	if error := db.Where("id= ?", c.Param("id")).First(&todo).Error; error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
		return
	}

	if error := db.Delete(&todo).Error; error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
