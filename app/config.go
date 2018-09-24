package app

import ()

var config = NewDefaultConfiguration()

type Config struct {
}

func NewDefaultConfiguration() *Config {

	return &Config{}
}
