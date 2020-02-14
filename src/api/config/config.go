package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Server ServerConfig
	Db     DbConfig
}

type ServerConfig struct {
	Host string `toml:"host"`
	Port string `toml:"port"`
}

type DbConfig struct {
	User string `toml:"user"`
	Pass string `toml:"pass"`
}

const configPath string = "config/default.toml"

func Load() Config {
	var config Config
	_, err := toml.DecodeFile(configPath, &config)
	if err != nil {
		panic(err)
	}
	return config
}
