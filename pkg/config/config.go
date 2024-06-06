package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Participants    []string          `yaml:"participants"`
	ExcludePaths    []string          `yaml:"excludePaths"`
	MessagePrefixes map[string]string `yaml:"messagePrefixes"`
}

func LoadConfig(filePath string) (Config, error) {
	var config Config
	configFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
