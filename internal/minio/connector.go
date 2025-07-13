package minio

import (
	"context"
	"fmt"
	// "net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

var s3Client *minio.Client

func GetMinio(log *zap.Logger) *minio.Client {
	if s3Client == nil {
		InitMinio(log)
		if s3Client == nil {
			fmt.Println("Failed to initialize Minio client.")
			return nil
		}
	} else {
		fmt.Println("Using existing Minio client.")
	}
	return s3Client
}

func InitMinio(log *zap.Logger) error {
	// Requests are always secure (HTTPS) by default.
	// Set secure=false to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().
	MinioEndpt := os.Getenv("MINIO_ENDPOINT")
	MinioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey := os.Getenv("MINIO_SECRET_KEY")
	MinioBucket := os.Getenv("MINIO_BUCKET")

	fmt.Println(MinioEndpt, "+", MinioAccessKey,"+", MinioSecretKey, "+",MinioBucket)
	if MinioEndpt == "" || MinioAccessKey == "" || MinioSecretKey == "" || MinioBucket == "" {
		fmt.Println("Minio environment variables are not set properly.")
		return fmt.Errorf("minio environment variables are not set properly")
	}

	conn, err := minio.New(MinioEndpt, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
		Secure: false,
	})
	if err != nil {
		fmt.Println("Minio failed to connect:", err)
		return fmt.Errorf("minio failed to connect: %w", err)
	}
	found, err := conn.BucketExists(context.Background(), MinioBucket)
	if err != nil {
		return fmt.Errorf("minio bucket existence check failed: %w", err)
	}
	if found {
		fmt.Println("Connection to Minio successful.")
		fmt.Println("Endpoint: " + MinioEndpt)
		fmt.Println("Bucket  : " + MinioBucket)
	}

	s3Client = conn
	return nil
}
func HandlePanic() {
	r := recover()

	if r != nil {
		fmt.Println("RECOVER :", r)
	}
}
func GetPresignedURLFromMinio(log *zap.Logger, objectname string) (string, error) {
	defer HandlePanic()

	MinioBucket := os.Getenv("MINIO_BUCKET")
	if MinioBucket == "" {
		log.Error("MINIO_BUCKET environment variable is not set.")
		return "", fmt.Errorf("MINIO_BUCKET environment variable is not set")
	}
	// generate

	// Gernerate presigned get object url.
	presignedURL, err := GetMinio(log).PresignedPutObject(
		context.Background(),
		MinioBucket,
		objectname,
		time.Second*1*60*60,
	)
	if err != nil {
		log.Error("Error generating presigned URL", zap.Error(err))
		return "", err
	}
	return presignedURL.String(), nil
}

// minio/connector.go
func DownloadFileFromMinio(log *zap.Logger, objectname string, filePath string) error {
	// Uncomment this in case the code goes into panic at any point of time
	// defer HandlePanic()
	// Download and save the object as a file in the local filesystem.
	MinioBucket := os.Getenv("MINIO_BUCKET")
	if MinioBucket == "" {
		fmt.Println("MINIO_BUCKET environment variable is not set.")
		return fmt.Errorf("MINIO_BUCKET environment variable is not set")
	}

	err := GetMinio(log).FGetObject(context.Background(), MinioBucket, objectname, filePath, minio.GetObjectOptions{})
	if err != nil {
		log.Error("Error downloading file from MinIO", zap.Error(err))
		return err
	}
	return nil
}

func UploadFileInMinio(log *zap.Logger, objectname string, filePath string, contentType string)error{

	// Upload the test file with FPutObject

	MinioBucket := os.Getenv("MINIO_BUCKET")
	if MinioBucket == "" {
		fmt.Println("MINIO_BUCKET environment variable is not set.")
		return fmt.Errorf("MINIO_BUCKET environment variable is not set")
	}
	info, err := GetMinio(log).FPutObject(context.Background(), MinioBucket, objectname, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Error("Error uploading file to MinIO", zap.Error(err))
		return fmt.Errorf("Error uploading file to MinIO: %w", err)
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", objectname, info.Size)
	return nil
}
