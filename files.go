package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	// "github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
//   err := godotenv.Load(".env")

//   if err != nil {
//     log.Fatalf("Error loading .env file")
//   }

  return os.Getenv(key)
}

var (
    R2_ENDPOINT    = goDotEnvVariable("R2_ENDPOINT")
    R2_BUCKET      = goDotEnvVariable("BUCKET_NAME")
    ACCESS_KEY_ID  = goDotEnvVariable("ACCESS_KEY_ID")
    SECRET_ACCESS_KEY = goDotEnvVariable("SECRET_ACCESS_KEY")
)

func handleGetSharedFile(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is a shared file"))
}

func handleUploadFile(w http.ResponseWriter, r *http.Request){
	r.ParseMultipartForm(10 << 20) // 10 MB max file size
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error retrieving file from form-data", http.StatusBadRequest)
		return
	}
	defer file.Close()

	sess, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(ACCESS_KEY_ID, SECRET_ACCESS_KEY, ""),
		Endpoint:    aws.String(R2_ENDPOINT),
		Region:      aws.String("us-east-1"),
	})
	if err != nil {
		http.Error(w, "Error creating AWS session", http.StatusInternalServerError)
		return
	}

	svc := s3.New(sess)
	objectKey := fmt.Sprintf("%d-%s", time.Now().UnixNano(), handler.Filename)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(R2_BUCKET),
		Key:    aws.String(objectKey),
		Body:   file,
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		http.Error(w, "Error uploading file", http.StatusInternalServerError)
		return
	}

	fileURL := fmt.Sprintf("%s/%s/%s", R2_ENDPOINT, R2_BUCKET, objectKey)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"url": fileURL})
}
