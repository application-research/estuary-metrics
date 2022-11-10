package core

import (
	"context"
	"github.com/spf13/viper"
	"github.com/whyrusleeping/memo"
	"gorm.io/gorm"
)

type Metrics struct {
	Context context.Context
}

var (
	DB              *gorm.DB
	Cacher          *memo.Cacher
	DefaultCacheTTL = 10
	EquinixEndpoint = "https://api.equinix.com/metal/v1/devices/"
)

var (
	//	should not be here. Environment variable instead.
	EquinixAuthToken = viper.Get("EQUINIX_AUTH_TOKEN")
)
