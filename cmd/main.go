// main.go
package main

import (
    "fmt"
    "log"
    // "os"
    "strconv"
    "setUp/config"
    "setUp/internal/delivery"
    "setUp/internal/logger"
    "setUp/internal/repository"
    "setUp/internal/usecase"
    "setUp/pkg/database"

    "github.com/gin-gonic/gin"
)

func main() {
    cfg := config.Load()

    db, err := database.ConnectPostgres()
    if err != nil {
        log.Fatalf("failed to connect to DB: %v", err)
    }
    log := logger.NewLogger()
	defer log.Sync()
    userRepo := repository.NewUserRepository(db,log)
    userUC := usecase.NewUserUsecase(userRepo,log)
    userHandler := delivery.NewUserHandler(userUC,log)

    r := gin.Default()
    r.POST("/login", userHandler.Login)

    fmt.Println("Server running at port", cfg.Port)
    r.Run(":" + strconv.Itoa(cfg.Port))
}
