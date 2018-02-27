package config

import (
	"os"
	"testing"
)

func TestReadEmptyConfig(t *testing.T) {
	const TestFile = "tests/config_empty.toml"
	f, err := os.Open(TestFile)
	if err != nil {
		t.Error("Could not open test file:", TestFile)
		t.FailNow()
	}

	result, err := ReadConfigFromFile(f)
	if err != nil {
		t.Error("Error parsing config:", err)
		t.FailNow()
	}

	if len(result.Config) != 0 {
		t.Error("Returned a non-empty config from an empty file")
		t.Fail()
	}
}

func TestReadConfigFromFile(t *testing.T) {
	const TestFile = "tests/config_basic.toml"
	f, err := os.Open(TestFile)
	if err != nil {
		t.Error("Could not open test file:", TestFile)
		t.FailNow()
	}

	result, err := ReadConfigFromFile(f)
	if err != nil {
		t.Error("Error parsing config:", err)
		t.FailNow()
	}

	if result.Config[0].Name != "webserver" {
		t.Error("Expected first app name to be 'webserver'")
		t.Fail()
	}
	if len(result.Config[0].Fetch) != 2 {
		t.Error("Expected one AppFetchConfig")
		t.Fail()
	}
}

func TestReadConfigBadInput(t *testing.T) {
	const TestFile = "tests/config_bad.toml"
	f, err := os.Open(TestFile)
	if err != nil {
		t.Error("Could not open test file:", TestFile)
		t.FailNow()
	}

	_, err = ReadConfigFromFile(f)
	if err == nil {
		t.Error("Expected error while parsing invalid config but got none")
		t.FailNow()
	}
}
