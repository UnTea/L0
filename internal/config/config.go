package config

import (
	"github.com/UnTea/L0/pkg/helpers"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
)

type Database struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Migrations string `yaml:"migrations"`
	Name       string `yaml:"name"`
	SslMode    string `yaml:"sslmode"`
}

type Nats struct {
	ClusterID string `yaml:"cluster_id"`
	ClientID  string `yaml:"client_id"`
	Channel   string `yaml:"channel"`
}

type HttpServer struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type Config struct {
	Database   Database   `yaml:"database"`
	Nats       Nats       `yaml:"nats"`
	HttpServer HttpServer `yaml:"http_server"`
}

func ReadConfigYML(filePath string) (config Config, err error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return config, err
	}

	defer helpers.Closer(file)

	decoder := yaml.NewDecoder(file)

	err = decoder.Decode(&config)
	if err != nil {
		return config, err
	}

	return config, nil
}
