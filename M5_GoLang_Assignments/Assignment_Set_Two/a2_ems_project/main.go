package main

import (
	"database/sql"
	db "ecommerce/config"
	"ecommerce/controller"
	"ecommerce/middleware"
	"ecommerce/repository"
	"ecommerce/service"
	"fmt"

	"github.com/gin-gonic/gin"
)

func InitializeAuthDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./ecommerce.db")
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to the database successfully.")

	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
	id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	username TEXT NOT NULL,
	password TEXT NOT NULL
	);`
	if _, err := db.Exec(createTableSQL); err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	}

	insertSQL := `INSERT INTO users (username, password) VALUES (?,?);`
	if _, err := db.Exec(insertSQL, "admin", "admin"); err != nil {
		return nil, err
	}

	return db, nil
}

func main() {

	auth_db, err := InitializeAuthDatabase()
	if err != nil {
		fmt.Println("Auth Database initialization failed:", err)
		return
	}
	defer auth_db.Close()

	db.InitializeDatabase()

	productRepo := repository.NewProductRepository(db.GetDB())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	r := gin.Default()

	r.Use(middleware.LoggingMiddlewareGin())
	authMiddleware := middleware.AuthMiddlewareGin(db.GetDB())

	authorized := r.Group("/")
	authorized.Use(authMiddleware)
	{
		authorized.POST("/product", productController.CreateProduct)
		authorized.GET("/product/:id", productController.GetProduct)
		authorized.GET("/products", productController.GetAllProducts)
		authorized.PUT("/product/:id", productController.UpdateProduct)
		authorized.DELETE("/product/:id", productController.DeleteProduct)
	}
a
	r.Run(":8080")
}
