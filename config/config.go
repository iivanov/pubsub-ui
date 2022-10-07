package config

import (
	"github.com/jinzhu/configor"
	"pubsub-ui/src/pub_sub"
)

type Config struct {
	PubSub pub_sub.Config `yaml:"pubSub"`
}

func NewConfig() Config {
	config := Config{}
	if err := configor.New(&configor.Config{Verbose: true}).Load(&config, "config/config.yaml"); err != nil {
		panic(err)
	}
	return config
}
