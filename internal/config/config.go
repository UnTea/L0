package config

import (
	"github.com/UnTea/L0/pkg/helpers"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Database - contains all parameters database configuration
type Database struct {
	Host       string `yaml:"host"`
	Port       string `yaml:"port"`
	User       string `yaml:"user"`
	Password   string `yaml:"password"`
	Migrations string `yaml:"migrations"`
	Name       string `yaml:"name"`
	SslMode    string `yaml:"sslmode"`
}

// Nats - contains all parameters nats configuration
type Nats struct {
	ClusterID string `yaml:"cluster_id"`
	ClientID  string `yaml:"client_id"`
	Channel   string `yaml:"channel"`
}

// HttpServer - contains ip and port for http server
type HttpServer struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

// Config - contains all configuration parameters in config package
type Config struct {
	Database   Database   `yaml:"database"`
	Nats       Nats       `yaml:"nats"`
	HttpServer HttpServer `yaml:"http_server"`
}

// ReadConfigYML - read configurations from file and init instance Config
func ReadConfigYML(filePath string) (cfg Config, err error) {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return cfg, err
	}
	defer helpers.Closer(file)

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
