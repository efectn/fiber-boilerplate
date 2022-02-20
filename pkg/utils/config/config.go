package config

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog"
)

type Config struct {
	Webserver struct {
		Header     string
		AppName    string `toml:"app_name"`
		Port       string
		Prefork    bool
		Production bool
	}
	Logger struct {
		TimeFormat string        `toml:"time-format"`
		Level      zerolog.Level `toml:"level"`
		Prettier   bool          `toml:"prettier"`
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
	Recover struct {
		Enabled bool
	}
	Monitor struct {
		Enabled bool
	}
	Filesystem struct {
		Enabled bool
		Browse  bool
		MaxAge  int `toml:"max_age"`
		Index   string
		Root    string
	}
}

func ParseConfig(name string, debug ...bool) (*Config, error) {
	var contents *Config
	var file []byte
	var err error

	if len(debug) > 0 {
		file, err = os.ReadFile(name)
	} else {
		file, err = os.ReadFile("./config/" + name + ".toml")
	}

	if err != nil {
		return &Config{}, err
	}

	err = toml.Unmarshal(file, &contents)

	return contents, err
}

// ParseAddr From https://github.com/gofiber/fiber/blob/master/helpers.go#L305.
func ParseAddr(raw string) (host, port string) {
	if i := strings.LastIndex(raw, ":"); i != -1 {
		return raw[:i], raw[i+1:]
	}
	return raw, ""
}
