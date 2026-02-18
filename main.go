package main

import (
	"github.com/gin-gonic/gin"
	"untitled9/internal/db"
	"untitled9/internal/http/controler"
	"untitled9/internal/http/controler/auth"
	"untitled9/internal/http/controler/products"
	"untitled9/internal/http/model"
)

func main() {
	db.InitDB()
	db.DB.AutoMigrate(&model.User{}, &model.Product{})

	route := gin.Default()

	route.GET("/health", controler.HealthCheck)

	authGroup := route.Group("/auth")
	authGroup.POST("/register", auth.Register)
	authGroup.POST("/login", auth.Login)

	route.POST("/products", products.CreateProduct)
	route.GET("/products", products.GetProducts)
	route.GET("/products/:id", products.GetProduct)
	route.PUT("/products/:id", products.UpdateProduct)
	route.DELETE("/products/:id", products.DeleteProduct)

	route.Run(":8080")
}
