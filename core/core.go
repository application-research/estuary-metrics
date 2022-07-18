package core

import (
	"context"
	"github.com/whyrusleeping/memo"
	"gorm.io/gorm"
	"log"
)

func Init(db *gorm.DB, cacherm *memo.Cacher) (*Metrics, error) {
	//	initialize all connection with DB and cacher.
	if db == nil {
		log.Fatal("Database Connection (gorm.db) cannot be null.")
	}

	if cacherm == nil {
		Cacher = memo.NewCacher() // just set to a new one.
	}

	DB = db

	return &Metrics{
		Context: context.Background(),
	}, nil
}
