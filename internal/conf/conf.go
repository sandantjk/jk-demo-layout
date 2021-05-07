package conf

import (
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	Server0Port int `yaml:"server0Port"`
	Server1Port int `yaml:"server1Port"`
}

func (c *Config) LoadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	bytes, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(bytes, c)
	if err != nil {
		return err
	}
	return nil
}
