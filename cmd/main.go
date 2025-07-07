// main.go
package main

import (
	"fmt"
	"log"
	"os"

	// "os"
    userServ "setUp/internal/services/users"
	"setUp/internal/logger"
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

    // Initialize repositories and usecases
    userRepo := userServ.NewUserRepository(db,log)
    userUC := userServ.NewUserUsecase(userRepo,log)
    userHandler := userServ.NewUserHandler(userUC,log)

    // Initialize Gin router
    r := gin.Default()
    apiV1 := r.Group("/api/v1") // Grouping routes under /api/v1
    apiV1.Use(gin.Recovery())
    userServ.RouteUser(apiV1, userHandler)

    fmt.Println("Server running at port", os.Getenv("PORT"))
    r.Run(":" + os.Getenv("PORT"))
}
