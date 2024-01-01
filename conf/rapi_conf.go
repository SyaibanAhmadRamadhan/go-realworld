package conf

import (
	"fmt"

	"github.com/SyaibanAhmadRamadhan/gocatch/genv"
)

// RapiConf is rest api configuration
type RapiConf struct {
	Hostname  string
	Port      string
	ServerKey string
}

func (w RapiConf) ListenerAddr() string {
	return fmt.Sprintf(":%s", w.Port)
}

func LoadEnvRapiConf() *RapiConf {
	return &RapiConf{
		Hostname:  genv.GetEnv("HOSTNAME_APP", "localhost"),
		Port:      genv.GetEnv("PORT_APP", "8080"),
		ServerKey: genv.GetEnv("SERVER_KEY_APP", "server key app"),
	}
}
