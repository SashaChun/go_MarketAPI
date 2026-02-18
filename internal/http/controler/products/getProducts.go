package products

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"untitled9/internal/db"
	"untitled9/internal/http/model"
)

func GetProducts(c *gin.Context) {
	var items []model.Product

	db.DB.Find(&items)

	c.JSON(http.StatusOK, items)
}
