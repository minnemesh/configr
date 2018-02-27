package fetcher

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/minnemesh/configr/common/types"
	"github.com/minnemesh/configr/node/config"
)

// HTTPFetcher implements the Fetcher interface by fetching app configs
// over the HTTP protocol.
type HTTPFetcher struct {
}

// Fetch fetches an AppConfig via HTTP.
func (f *HTTPFetcher) Fetch(afc *config.AppFetchConfig) (eac types.EncryptedAppConfig, err error) {
	resp, err := http.Get(afc.URL)
	if err != nil {
		err = fmt.Errorf("Error requesting config via HTTP: %v", err)
		return
	}

	eac.Data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("Error reading from HTTP response body: %v", err)
		return
	}

	return
}
