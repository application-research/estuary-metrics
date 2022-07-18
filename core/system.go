package core

import (
	"context"
	"github.com/application-research/estuary-metrics/core/generated/query"
)

func (m Metrics) GetTotalObjectsPinned() (int64, error) {
	u := query.Use(DB).Content
	count, err := u.WithContext(context.Background()).Count()
	//	return the result
	return count, err

}

func (m Metrics) GetTotalTibsUploaded() (int64, error) {
	//	get the database connection
	u := query.Use(DB).Object
	size, err := u.WithContext(context.Background()).Select(u.Size.Sum()).First()
	//	return the result
	return size.Size, err
}

func (m Metrics) GetTotalTibsSealedOnFilecoin() (int64, error) {
	//	get the database connection
	u := query.Use(DB).Object
	size, err := u.WithContext(context.Background()).Select(u.Size.Sum()).First()
	//	return the result
	return size.Size, err
}

func (m Metrics) GetAvailableFreeSpace() (int64, error) {
	//	get the database connection
	u := query.Use(DB).Object
	size, err := u.WithContext(context.Background()).Select(u.Size.Sum()).First()
	//	return the result
	return size.Size, err
}
