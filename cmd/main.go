// main.go
package main

import (
	"fmt"
	"log"
	"os"

	// "os"
    userServ "setUp/internal/services/users"
    uploadServ "setUp/internal/services/upload"
    minioServ "setUp/internal/minio"
	"setUp/internal/logger"
	"setUp/pkg/database"
    "github.com/gin-contrib/cors"
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
    logZap := logger.NewLogger()
	defer logZap.Sync()
    // Initialize MinIO client
    err = minioServ.InitMinio(logZap)
    if err != nil {
        log.Fatal("failed to connect to MinIO: %v", err)
    }
    // Initialize repositories and usecases
    userRepo := userServ.NewUserRepository(db,logZap)
    userUC := userServ.NewUserUsecase(userRepo,logZap)
    userHandler := userServ.NewUserHandler(userUC,logZap)
    uploadHandler := uploadServ.NewUploadHandler(logZap)

    // Initialize Gin router
    r := gin.Default()

    // Gunakan CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // ganti dengan origin kamu
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// MaxAge:           12 * time.Hour,
	}))

    apiV1 := r.Group("/api/v1") // Grouping routes under /api/v1
    apiV1.Use(gin.Recovery())
    userServ.RouteUser(apiV1, userHandler)
    uploadServ.RouteUpload(apiV1, uploadHandler)

    fmt.Println("Server running at port", os.Getenv("PORT"))
    r.Run(":" + os.Getenv("PORT"))
}
