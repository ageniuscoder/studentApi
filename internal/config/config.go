package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}
type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"` //env-default:"production"
	StoragePath string `yaml:"storage_path" env_required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config", "", "path to configuration file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file doesn,t exist %s", configPath)
	}

	var cfg Config

	e := cleanenv.ReadConfig(configPath, &cfg)
	if e != nil {
		log.Fatalf("can,t read config file: %s", e.Error())
	}

	return &cfg
}
