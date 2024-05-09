package config

import (
	_ "embed"
	"gopkg.in/yaml.v3"
)

var Config Data

//go:embed default_config.yaml
var defaultData string

type Data struct {
	Valkey Valkey `yaml:"valkey"`
	Mqtt   Mqtt   `yaml:"mqtt"`
	Minio  Minio  `yaml:"minio"`
}

type Valkey struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Mqtt struct {
	Host      string         `yaml:"host"`
	Listeners map[string]int `yaml:"listeners"`
}

type Minio struct {
	Endpoint string `yaml:"endpoint"`
	Secure   bool   `yaml:"secure"`
}

func init() {
	err := yaml.Unmarshal([]byte(defaultData), &Config)
	if err != nil {
		panic(err)
	}
}
