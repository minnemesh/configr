package config

import (
	"io"

	"github.com/BurntSushi/toml"
)

type NodeConfig struct {
}

func ReadConfigFromFile(reader io.Reader) (config NodeConfig, err error) {
	_, err = toml.DecodeReader(reader, &config)
	if err != nil {
		return
	}

	return
}
