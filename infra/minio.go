package infra

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/rs/zerolog/log"

	"realworld-go/conf"
)

type Minio struct {
	Bucket string
	Client *minio.Client
}

func OpenConnMinio() *Minio {
	minioConf := conf.LoadMinioEnv()
	minioClient, err := minio.New(minioConf.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(minioConf.AccessKeyId, minioConf.AccessSecretKey, ""),
		Secure: minioConf.SSL,
	})
	if err != nil {
		log.Fatal().Err(err).Msgf("failed open minio client")
	}

	return &Minio{
		Bucket: minioConf.Bucket,
		Client: minioClient,
	}
}
