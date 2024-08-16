package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Service Service `yaml:"service"`
	Storage Storage `yaml:"storage"`
	Broker  Broker  `yaml:"broker"`
}

type Service struct {
	Port string `yaml:"port"`
}

type Storage struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Broker struct {
	Port      string `yaml:"port"`
	Topic     string `yaml:"topic"`
	Partition string `yaml:"partition"`
}

func MustLoad() *Config {
	path := "./config/local.yaml"
	//if path == "" {
	//	panic("config path is empty")
	//}

	if _, err := os.Stat(path); err != nil {
		panic(fmt.Sprintf("no file by path: %s", path))
	}

	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("error reading config")
	}

	return &config
}
