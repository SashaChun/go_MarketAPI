package products

import (
	"net/http"
	"untitled9/internal/db"
	"untitled9/internal/http/model"

	"github.com/gin-gonic/gin"
)

func DeleteProduct(c *gin.Context) {
	var product model.Product

	if err := db.DB.Where("id = ?", c.Param("id")).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if err := db.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
