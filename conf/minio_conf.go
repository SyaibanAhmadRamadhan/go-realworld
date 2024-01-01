package conf

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/gcommon"
	"github.com/SyaibanAhmadRamadhan/gocatch/genv"
)

type MinioConf struct {
	AccessKeyId     string
	AccessSecretKey string
	Endpoint        string
	SSL             bool
	Bucket          string
}

func LoadMinioEnv() *MinioConf {
	ssl := genv.GetEnv("MINIO_SSL")
	return &MinioConf{
		AccessKeyId:     genv.GetEnv("MINIO_ACCESS_KEY_ID"),
		AccessSecretKey: genv.GetEnv("MINIO_ACCESS_SECRET_KEY"),
		Endpoint:        genv.GetEnv("MINIO_ENDPOINT"),
		SSL:             gcommon.Ternary(ssl == "true", true, false),
		Bucket:          genv.GetEnv("MINIO_BUCKET"),
	}
}
