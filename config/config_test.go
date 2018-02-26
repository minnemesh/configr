package config

import (
	"os"
	"testing"
)

func TestReadConfigFromFile(t *testing.T) {
	const TestFile = "tests/config_basic.toml"
	f, err := os.Open(TestFile)
	if err != nil {
		t.Error("Could not open test file:", TestFile)
		t.FailNow()
	}

	_, err = ReadConfigFromFile(f)
	if err != nil {
		t.Error("Error parsing config:", err)
		t.FailNow()
	}

	// TODO: check contents of `config`
}
