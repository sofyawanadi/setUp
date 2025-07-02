// main.go
package main

import (
	"fmt"
	"log"
	"os"

	// "os"
	"setUp/internal/delivery"
	"setUp/internal/logger"
	"setUp/internal/repository"
	"setUp/internal/usecase"
	"setUp/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load("../.env")
    if err != nil {
      log.Fatal("Error loading .env file")
    }

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

    fmt.Println("Server running at port", os.Getenv("PORT"))
    r.Run(":" + os.Getenv("PORT"))
}
