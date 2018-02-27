package fetcher

import (
	"github.com/minnemesh/configr/common/types"
	"github.com/minnemesh/configr/node/config"
)

type Fetcher interface {
	Fetch(afc *config.AppFetchConfig) (types.EncryptedAppConfig, error)
}
