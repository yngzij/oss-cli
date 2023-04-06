package db

import (
	"context"
	"fmt"
	"log"
	"sync"

	"oss-cli/configs"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var minioClient *Minio
var minioOnce sync.Once

type Minio struct {
	client *minio.Client
}

func MinioClient() *Minio {
	// Initialize minio client object.

	/*endpoint := "45.77.34.64:9000"
	accessKeyID := "FWLR1ceLxDjwQocY"
	secretAccessKey := "2GwwbwofklqMrbwqGpYKs4NpvVcV0KNf"
	useSSL := false*/
	fmt.Println(configs.Config.Minio.Endpoint)

	minioOnce.Do(func() {
		client, err := minio.New(configs.Config.Minio.Endpoint, &minio.Options{
			Creds: credentials.NewStaticV4(
				configs.Config.Minio.AccessKeyID,
				configs.Config.Minio.SecretAccessKey,
				""),
			Secure: configs.Config.Minio.UseSSL,
		})
		if err != nil {
			log.Fatalln(err)
		}
		minioClient = &Minio{
			client: client,
		}

	})
	return minioClient
}

func (mio Minio) PutObject(path string, bucketName string, objectName string, metadata map[string]string) error {
	n, err := mio.client.FPutObject(context.Background(), bucketName, objectName, path, minio.PutObjectOptions{UserMetadata: metadata})
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully uploaded bytes: ", n)
	return nil
}
