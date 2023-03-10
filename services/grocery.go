package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/db"
	"github.com/muhammetandic/go-backend/main/db/model"
)

type NewGrocery struct {
	Name     string `json:"name" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
}

type UpdatedGrocery struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func GetAllGroceries(c *gin.Context) {
	var groceries []model.Grocery
	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	if error = db.Find(&groceries).Error; error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
		return
	}

	c.JSON(http.StatusOK, groceries)
}

func GetGrocery(c *gin.Context) {
	var grocery model.Grocery

	db, error := db.Connect()

	if error != nil {
		log.Println(error)
	}

	if error := db.Where("id= ?", c.Param("id")).First(&grocery).Error; error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grocery not found!"})
		return
	}

	c.JSON(http.StatusOK, grocery)
}

func PostGrocery(c *gin.Context) {
	var grocery NewGrocery

	if error := c.ShouldBindJSON(&grocery); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
	}

	newGrocery := model.Grocery{Name: grocery.Name, Quantity: grocery.Quantity}

	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	if error := db.Create(&newGrocery).Error; error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
	}

	c.JSON(http.StatusOK, newGrocery)
}

func UpdateGrocery(c *gin.Context) {

	var grocery model.Grocery

	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	if error := db.Where("id = ?", c.Param("id")).First(&grocery).Error; error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grocery not found!"})
		return
	}

	var updateGrocery UpdatedGrocery

	if error := c.ShouldBindJSON(&updateGrocery); error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": error.Error()})
		return
	}

	if error := db.Model(&grocery).Updates(model.Grocery{Name: updateGrocery.Name, Quantity: updateGrocery.Quantity}).Error; error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	c.JSON(http.StatusOK, grocery)
}

func DeleteGrocery(c *gin.Context) {

	var grocery model.Grocery

	db, error := db.Connect()
	if error != nil {
		log.Println(error)
	}

	if error := db.Where("id = ?", c.Param("id")).First(&grocery).Error; error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Grocery not found!"})
		return
	}

	if error := db.Delete(&grocery).Error; error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Grocery deleted"})
}
