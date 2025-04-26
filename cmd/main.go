// main.go
package main

import (
    "fmt"
    "log"
    "os"

    "project-name/config"
    "project-name/internal/delivery"
    "project-name/internal/repository"
    "project-name/internal/usecase"
    "project-name/pkg/database"

    "github.com/gin-gonic/gin"
)

func main() {
    cfg := config.Load()

    db, err := database.ConnectPostgres(cfg.DBUrl)
    if err != nil {
        log.Fatalf("failed to connect to DB: %v", err)
    }

    userRepo := repository.NewUserPostgres(db)
    userUC := usecase.NewUserUsecase(userRepo, cfg.JWTSecret)
    userHandler := delivery.NewUserHandler(userUC)

    r := gin.Default()
    r.POST("/login", userHandler.Login)

    fmt.Println("Server running at port", cfg.Port)
    r.Run(":" + cfg.Port)
}
