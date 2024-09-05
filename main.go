package main

import (
	"klinik-hewan/config"
	"klinik-hewan/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/klinik_hewan_go?charset=utf8mb4&parseTime=True&loc=Local"
	db := config.InitDB(dsn)
	handlers.InitDB(db)

	r := gin.Default()

	// Konfigurasi CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Type"},
		AllowCredentials: true,
	}))

	// Rute untuk User
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)
	
	// Rute untuk hewan
	r.GET("/hewan", handlers.GetHewan)
	r.POST("/hewan", handlers.CreateHewan)
	r.PUT("/hewan/:id", handlers.UpdateHewan)
	r.DELETE("/hewan/:id", handlers.DeleteHewan)

	// Rute untuk produk
	r.GET("/produk", handlers.GetProduk)
	r.POST("/produk", handlers.CreateProduk)
	r.PUT("/produk/:id", handlers.UpdateProduk)
	r.DELETE("/produk/:id", handlers.DeleteProduk)

	r.Run(":8080")
}
