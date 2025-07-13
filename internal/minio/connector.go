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
	ctx := context.Background()

	fmt.Println(MinioEndpt, "+", MinioAccessKey, "+", MinioSecretKey, "+", MinioBucket)
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
	found, err := conn.BucketExists(ctx, MinioBucket)
	if err != nil {
		return fmt.Errorf("minio bucket existence check failed: %w", err)
	}
	location := "us-east-1"
	if found {
		fmt.Println("Connection to Minio successful.")
		fmt.Println("Endpoint: " + MinioEndpt)
		fmt.Println("Bucket  : " + MinioBucket)
	}else{
		err = conn.MakeBucket(ctx, MinioBucket, minio.MakeBucketOptions{Region: location})
		if err != nil {
			return fmt.Errorf("minio failed to connect: %w", err)
		}
		fmt.Println("✅ Bucket dibuat:", MinioBucket)
	}

	// Set policy public read-only
	policy := fmt.Sprintf(`{
		"Version":"2012-10-17",
		"Statement":[{
			"Effect":"Allow",
			"Principal":{"AWS":["*"]},
			"Action":["s3:GetObject"],
			"Resource":["arn:aws:s3:::%s/*"]
		}]
	}`, MinioBucket)

	err = conn.SetBucketPolicy(ctx, MinioBucket, policy)
	if err != nil {
		return fmt.Errorf("minio failed to connect: %w", err)
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
	presignedURL, err := s3Client.PresignedPutObject(
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
func DownloadFileFromMinio(log *zap.Logger, objectname string, filePath string)(*minio.Object, error ){
	// Uncomment this in case the code goes into panic at any point of time
	// defer HandlePanic()
	// Download and save the object as a file in the local filesystem.
	MinioBucket := os.Getenv("MINIO_BUCKET")
	if MinioBucket == "" {
		fmt.Println("MINIO_BUCKET environment variable is not set.")
		return nil,fmt.Errorf("MINIO_BUCKET environment variable is not set")
	}

	// err := s3Client.FGetObject(ctx, MinioBucket, objectname, filePath, minio.GetObjectOptions{})
	// if err != nil {
	// 	log.Error("Error downloading file from MinIO", zap.Error(err))
	// 	return err
	// }
	obj,	err := s3Client.GetObject(context.Background(), MinioBucket, objectname,  minio.GetObjectOptions{})
	if err != nil {
		log.Error("Error downloading file from MinIO", zap.Error(err))
		return nil,err
	}

	return obj,nil
}

func UploadFileInMinio(log *zap.Logger, objectname string, filePath string, contentType string) error {

	// Upload the test file with FPutObject

	MinioBucket := os.Getenv("MINIO_BUCKET")
	if MinioBucket == "" {
		fmt.Println("MINIO_BUCKET environment variable is not set.")
		return fmt.Errorf("MINIO_BUCKET environment variable is not set")
	}
	info, err := s3Client.FPutObject(context.Background(), MinioBucket, objectname, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Error("Error uploading file to MinIO", zap.Error(err))
		return fmt.Errorf("error uploading file to MinIO: %w", err)
	}
	fmt.Printf("Successfully uploaded %s of size %d\n", objectname, info.Size)
	return nil
}
