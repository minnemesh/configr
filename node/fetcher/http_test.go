package fetcher

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/minnemesh/configr/node/config"
)

func TestImplementsFetcherInterface(t *testing.T) {
	var fetcher Fetcher
	fetcher = &HTTPFetcher{}

	if fetcher == nil {
		t.Error("Pointer to Fetcher is nil")
		t.Fail()
	}
}

func TestFetchErrorConnecting(t *testing.T) {
	fetcher := HTTPFetcher{}
	_, err := fetcher.Fetch(&config.AppFetchConfig{
		Method: "http",
		URL:    "https://999.999.999.999:23413/sdfljaksjflasj",
	})

	if err == nil {
		t.Error("Expected an error fetching from an invalid URL")
		t.Fail()
	}
}

func TestFetchURL(t *testing.T) {
	encryptedAppConfig := []byte("test config data")
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Write(encryptedAppConfig)
		},
	))
	defer ts.Close()

	fetcher := HTTPFetcher{}
	encconfig, err := fetcher.Fetch(&config.AppFetchConfig{
		Method: "http",
		URL:    ts.URL,
	})

	if err != nil {
		t.Error("Error fetching encrypted app config:", err)
		t.FailNow()
	}

	if len(encconfig.Data) != len(encryptedAppConfig) {
		fmt.Println(len(encconfig.Data))
		fmt.Println(encconfig.Data)
		t.Error("Returned encrypted config data has the wrong length")
		t.FailNow()
	}

	for i := range encconfig.Data {
		if encconfig.Data[i] != encryptedAppConfig[i] {
			t.Error("Returned config data does not match served config data")
			t.FailNow()
		}
	}
}
