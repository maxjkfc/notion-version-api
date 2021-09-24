package config

import (
	"errors"

	"github.com/spf13/viper"
)

// Config -
type Config struct {
	Notion Notion `yaml:"notion"`
}

// Notion -
type Notion struct {
	Token    string         `yaml:"token"`
	Host     string         `yaml:"host"`
	Version  string         `yaml:"version"`
	Database NotionDatabase `yaml:"database"`
}

// NotionDatabase -
type NotionDatabase struct {
	ID string `yaml:"id"`
}

// New -
func New(path string) (c Config, err error) {
	c = Config{}

	if path == "" {
		return c, errors.New("error: empty path for config file")
	}

	viper.SetConfigFile(path)
	if err = viper.ReadInConfig(); err != nil {
		return c, errors.New("read file failed: " + err.Error())
	}

	if err := viper.Unmarshal(&c); err != nil {
		return c, errors.New("Unmarshal failed : " + err.Error())
	}

	return
}
