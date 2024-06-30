package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	ServerInfo struct {
		Port string
		Info string
	}

	Paseto struct {
		Key string
	}

	GRPC struct {
		Domain string
		Port   string
	}
}

func NewConfig(path string) *Config {
	c := &Config{}

	if file, err := os.Open(path); err != nil {
		panic(err)
	} else {
		defer file.Close()

		if err = toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}
