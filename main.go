package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"untitled9/internal/config"
	"untitled9/internal/db"
	"untitled9/internal/http/controler"
	"untitled9/internal/http/controler/auth"
	"untitled9/internal/http/controler/products"
	"untitled9/internal/http/model"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}
	
	log.Printf("Loaded env: APP_PORT=%s, DB_HOST=%s, DB_USER=%s", 
		config.GetAppPort(), config.GetDBHost(), config.GetDBUser())

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

	port := ":" + config.GetAppPort()
	if config.GetAppPort() == "" {
		port = ":8080"
	}
	route.Run(port)
}
