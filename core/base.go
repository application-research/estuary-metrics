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
)
