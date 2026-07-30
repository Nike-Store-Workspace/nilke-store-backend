package main

import (
	"database/sql"
	"net/http"
	"nike_store_api/internal/data/repository"
	"nike_store_api/internal/handler"
	"nike_store_api/internal/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var jwtSecret = "my-super-secret-jwt-key"

func main() {
	// create db
	connStr := "host=127.0.0.1 port=5432 user=postgres password=secret dbname=nikeStoreDb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Manual Di
	userRepo := repository.NewPostgresUserRepository(db)
	productRepo := repository.NewPostgresProductRepository(db)

	authService := services.NewAuthService(userRepo, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)

	signupService := services.NewSignupService(userRepo, jwtSecret)
	signupHandler := handler.NewSignupHandler(signupService)

	productService := services.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productService)

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"message": "hi there",
			},
		)
	})

	v1 := r.Group("/api/v1")
	{
		v1.GET("/ping", handler.Ping)
		v1.POST("/login", authHandler.Login)
		v1.POST("/signup", signupHandler.Signup)
		v1.GET("/products", productHandler.GetProducts)
	}
	r.Run(":8090")
}
