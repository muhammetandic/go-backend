package grocery

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/muhammetandic/go-backend/main/model"
)

type CreateGrocery struct {
    Name string `json:"name" binding:"required"`
    Quantity string `json:"quantity" binding:required""`
}

type UpdateGrocery struct {
    Name string `json:"name"`
    Quantity string `json:"qauntity"`
}

func GetAllGroceries(c *gin.Context) {
    var groceries []model.Grocery
    db, error := model.DbInit()
    if error != nil {
        log.Println(error)
    }
    
    if error = db.Find(&groceries).Error; error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": error.Error()})
        return
    }
    
    c.JSON(http.StatusOK, groceries)
}