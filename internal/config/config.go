package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

func NewConfig() ServiceConfiguration {
	return Load()
}

func Load() ServiceConfiguration {
	file, err := os.Open("config.yml")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	decoder := yaml.NewDecoder(file)
	var cfg ServiceConfiguration
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}

type ServiceConfiguration struct {
	APIConfig   API   `yaml:"api"`
	RedisConfig Redis `yaml:"redis_configuration"`
	AuthConfig  Auth  `yaml:"auth_configuration"`
}

type API struct {
	Port    string `yaml:"port"`
	Host    string `yaml:"host"`
	UseCORS bool   `yaml:"use_cors"`
}

func (api *API) GetAddr() string {
	return fmt.Sprintf("%s:%s", api.Host, api.Port)
}

type Redis struct {
	Addr string `yaml:"redis_addr"`
}

type Auth struct {
	TTL     time.Duration `yaml:"ttl"`
	CodeLen uint          `yaml:"code_len"`
}
