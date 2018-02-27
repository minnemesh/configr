package fetcher

import (
	"github.com/minnemesh/configr/common/types"
	"github.com/minnemesh/configr/node/config"
)

type HTTPFetcher struct {
}

// Fetch fetches an AppConfig via HTTP.
func (f *HTTPFetcher) Fetch(afc *config.AppFetchConfig) types.EncryptedAppConfig {
	return types.EncryptedAppConfig{}
}
