package file_store

import (
	"admin_backend/global"
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"time"
)

func GenerateFileLink(ctx context.Context, objectName string) (string, error) {
	url, err := global.MinioConn.PresignedGetObject(
		ctx,
		global.MinioBucket,
		objectName,
		time.Hour*24*7,
		nil,
	)
	if err != nil {
		return "", err
	}
	return url.String(), nil
}

func StoreFile(ctx context.Context, objectName, filePath, contentType string) error {
	// contentType = "text/plain"
	// filePath    = "./text.txt
	// objectName  = "txt.txt
	info, err := global.MinioConn.FPutObject(
		ctx,
		global.MinioBucket,
		objectName,
		filePath,
		minio.PutObjectOptions{
			ContentType: contentType,
		})
	if err != nil {
		return err
	}
	log.Println("file upload success: ", objectName, info.Size)
	return nil
}
