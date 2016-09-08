package util

import (
	"fmt"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Title  string
	DB     database `toml:"database"`
	Server server   `toml:"server"`
}

type database struct {
	Name     string
	Host     string
	User     string
	Password string
}

type server struct {
	Port       string
	GinMode    string
	HmacSecret string
}

func ReadConfig() (Config, error) {
	var conf Config
	if _, err := toml.DecodeFile("config.toml", &conf); err != nil {
		return Config{}, fmt.Errorf("Unable to parse TOML config file")
	}

	return conf, nil
}
