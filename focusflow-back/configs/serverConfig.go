package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	ServerAddress string `yaml:"server_address"`
	DataBaseURL   string `yaml:"data_base_url"`
}

func NewConfig() *Config {
	var cfg Config

	bytes, err := ioutil.ReadFile("configs/application.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	return &cfg
}
