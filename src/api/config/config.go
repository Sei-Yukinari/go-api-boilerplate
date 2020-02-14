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
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		panic(err)
	}
	return config
}
