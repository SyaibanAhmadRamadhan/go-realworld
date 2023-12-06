package conf

import (
	"github.com/SyaibanAhmadRamadhan/gocatch/genv"
)

type MongodbConf struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
}

func (m *MongodbConf) URI() string {
	return "mongodb://" + m.Host + ":" + m.Port
}

func EnvMongodb() *MongodbConf {
	return &MongodbConf{
		Host:     genv.GetEnv("MONGO_HOST", "localhost"),
		Port:     genv.GetEnv("MONGO_PORT", "27017"),
		Username: genv.GetEnv("MONGO_USER", "root"),
		Password: genv.GetEnv("MONGO_PORT", "root"),
		Database: genv.GetEnv("MONGO_DB", "mongodb"),
	}
}
