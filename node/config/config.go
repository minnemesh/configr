package config

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/BurntSushi/toml"
)

type NodeConfig struct {
	Config map[string]AppConfig `toml:"config"`
}

type AppConfig struct {
	Fetch []AppFetchConfig `toml:"fetch"`
}

type AppFetchConfig struct {
	Method string `toml:"method"`
	URL    string `toml:"url"`
}

// ReadConfig finds the correct configuration path for this system and then
// passes a Reader to `ReadConfigFromFile` to read the file.
func ReadConfig() (NodeConfig, error) {
	homePath := path.Join(os.Getenv("HOME"), ".config", "configr", "config.toml")
	etcPath := "/etc/configr/config.toml"
	file, err := os.Open(homePath)
	if !os.IsNotExist(err) {
		return ReadConfigFromFile(file)
	}

	file, err = os.Open(etcPath)
	if !os.IsNotExist(err) {
		return ReadConfigFromFile(file)
	}

	err = fmt.Errorf("Could not find config file.")
	return NodeConfig{}, err
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
