package products

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"untitled9/internal/db"
	"untitled9/internal/http/model"
)

func GetProduct(c *gin.Context) {
	var product model.Product

	if err := db.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}
