package minio

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var s3Client *minio.Client

func GetMinio() *minio.Client {
	if s3Client == nil {
		InitMinio()
		if s3Client == nil {
			fmt.Println("Failed to initialize Minio client.")
			return nil
		}
	} else {
		fmt.Println("Using existing Minio client.")
	}
	return s3Client
}

func InitMinio() {
	// Requests are always secure (HTTPS) by default.
	// Set secure=false to enable insecure (HTTP) access.
	// This boolean value is the last argument for New().
	MinioEndpt := os.Getenv("MINIO_ENDPOINT")
	MinioAccessKey := os.Getenv("MINIO_ACCESS_KEY")
	MinioSecretKey := os.Getenv("MINIO_SECRET_KEY")
	MinioBucket := os.Getenv("MINIO_BUCKET")

	if MinioEndpt == "" || MinioAccessKey == "" || MinioSecretKey == "" || MinioBucket == "" {
		fmt.Println("Minio environment variables are not set properly.")
		return
	}

	conn, err := minio.New(MinioEndpt, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKey, MinioSecretKey, ""),
		Secure: true,
	})
	if err != nil {
		fmt.Println(err)
	}
	found, err := conn.BucketExists(context.Background(), MinioBucket)
	if err != nil {
		fmt.Println(err)
	}
	if found {
		fmt.Println("Connection to Minio successful.")
		fmt.Println("Endpoint: " + MinioEndpt)
		fmt.Println("Bucket  : " + MinioBucket)
	}

	s3Client = conn
}
func HandlePanic() {
	r := recover()

	if r != nil {
		fmt.Println("RECOVER :", r)
	}
}
func GetPresignedURLFromMinio(objectname string) string {
	defer HandlePanic()

	MinioBucket := os.Getenv("MINIO_BUCKET")
	if MinioBucket == "" {
		fmt.Println("MINIO_BUCKET environment variable is not set.")
		return ""
	}
	reqParams := make(url.Values)

	// Gernerate presigned get object url.
	presignedURL, err := GetMinio().PresignedGetObject(
		context.Background(),
		MinioBucket,
		objectname,
		time.Second*24*60*60,
		reqParams,
	)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return presignedURL.String()
}

// minio/connector.go
func DownloadFileFromMinio(objectname string, filePath string) error {
	// Uncomment this in case the code goes into panic at any point of time
	// defer HandlePanic()
	// Download and save the object as a file in the local filesystem.
	MinioBucket := os.Getenv("MINIO_BUCKET")
	if MinioBucket == "" {
		fmt.Println("MINIO_BUCKET environment variable is not set.")
		return fmt.Errorf("MINIO_BUCKET environment variable is not set")
	}

	err := GetMinio().FGetObject(context.Background(), MinioBucket, objectname, filePath, minio.GetObjectOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}


func UploadFileInMinio(objectname string, filePath string, contentType string) string {

	// Upload the test file with FPutObject

	MinioBucket := os.Getenv("MINIO_BUCKET")
	if MinioBucket == "" {
		fmt.Println("MINIO_BUCKET environment variable is not set.")
		return ""
	}
	info, err := GetMinio().FPutObject(context.Background(), MinioBucket, objectname, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", objectname, info.Size)
	return info.ETag
}