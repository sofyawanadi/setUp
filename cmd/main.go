// @title			My API
// @version		1.0
// @description	This is a sample server.
// @host			localhost:3000
// @BasePath		/api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"fmt"
	"log"
	"os"

	"setUp/internal/logger"
	minioServ "setUp/internal/minio"
	noteServ "setUp/internal/services/notes"
	uploadServ "setUp/internal/services/upload"
	userServ "setUp/internal/services/users"
	"setUp/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "setUp/cmd/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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
	userRepo := userServ.NewUserRepository(db, logZap)
	userUC := userServ.NewUserUsecase(userRepo, logZap)
	userHandler := userServ.NewUserHandler(userUC, logZap)
	uploadHandler := uploadServ.NewUploadHandler(logZap)
	noteRepo := noteServ.NewNoteRepository(db, logZap)
	noteUC := noteServ.NewNoteUsecase(noteRepo, logZap)
	noteHandler := noteServ.NewNoteHandler(noteUC, logZap)

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
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiV1 := r.Group("/api/v1") // Grouping routes under /api/v1
	apiV1.Use(gin.Recovery())
	userServ.RouteUser(apiV1, userHandler)
	uploadServ.RouteUpload(apiV1, uploadHandler)
	noteServ.RouteNote(apiV1, noteHandler)
	fmt.Println("Server running at port", os.Getenv("PORT"))
	r.Run(":" + os.Getenv("PORT"))
}
