package core

import (
	"context"
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
	EquinixEndpoint = "https://objects-api.equinix.go.com/metal/v1/devices/"
)

var (
	//	should not be here. Environment variable instead.
	EquinixAuthToken = "JTB647zpDi2i8sThgvG9W4c98rkdMQAY"
)
