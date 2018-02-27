package config

import (
	"io"

	"github.com/BurntSushi/toml"
)

type NodeConfig struct {
	Config []AppConfig `toml:"config"`
}

type AppConfig struct {
	Name  string           `toml:"name"`
	Fetch []AppFetchConfig `toml:"fetch"`
}

type AppFetchConfig struct {
	Method string `toml:"method"`
	URL    string `toml:"url"`
}

// ReadConfigFromFile reads a TOML-formatted node configuration into a
// config.NodeConfig struct and returns it.
func ReadConfigFromFile(reader io.Reader) (config NodeConfig, err error) {
	_, err = toml.DecodeReader(reader, &config)
	if err != nil {
		return
	}

	return
}
