package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/gofiber/fiber/v2/middleware/compress"
)

type Config struct {
	Webserver struct {
		Header  string
		AppName string `toml:"app_name"`
		Port    string
	}
	Logger struct {
		Enabled    bool
		Timezone   string
		Timeformat string
		Format     string
	}
	Limiter struct {
		Enabled bool
		Max     int
		ExpSecs int `toml:"expiration_seconds"`
	}
	Session struct {
		Enabled bool
		ExpHrs  int `toml:"expiration_hours"`
	}
	Compress struct {
		Enabled bool
		Level   compress.Level
	}
}

func ParseConfig(file string) (*Config, error) {
	var config *Config

	contents, err := ioutil.ReadFile("./config/" + file + ".toml")
	if err != nil {
		return &Config{}, err
	}

	toml.Decode(string(contents), &config)
	if err != nil {
		return &Config{}, err
	}

	return config, nil
}
