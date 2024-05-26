package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filepath string) *Config {
	c := new(Config)

	if file, err := os.Open(filepath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}
}
