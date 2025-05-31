package lib

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Serve struct {
		Port      int    `toml:"port"`
		Mantela   string `toml:"mantela"`
		Static    bool   `toml:"static"`
		StaticDir string `toml:"static_dir"`
	} `toml:"serve"`
	Mantela struct {
		Source string `toml:"source"`
		Diff   string `toml:"diff"`
	} `toml:"mantela"`
}

func GetConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	c, err := ParseConfig(data)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func ParseConfig(data []byte) (*Config, error) {
	var config Config
	if err := toml.Unmarshal(data, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
