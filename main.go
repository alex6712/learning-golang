package main

import (
	"log"

	"github.com/alex6712/learning-golang/config"
	"github.com/alex6712/learning-golang/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	db := config.InitDB(cfg)

	router := gin.Default()

	routes.RegisterRoutes(router, db)

	log.Printf("Сервер запущен на порту %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Не удалось запустить сервер:", err)
	}
}
