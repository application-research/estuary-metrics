package core

import (
	"context"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/application-research/estuary-metrics/core/generated/query"
	"time"
)

func (m Metrics) GetNumberOfUsers() (int64, error) {
	userModel := query.Use(DB).User
	count, err := userModel.WithContext(context.Background()).Count()
	return count, err
}

func (m Metrics) GetNumberOfActiveUsers() (int64, error) {
	userModel := query.Use(DB).User
	count, err := userModel.WithContext(context.Background()).Count()
	return count, err
}

func (m Metrics) GetUsersWithinTimeRange(to time.Time, from time.Time) ([]*model.User, error) {
	userModel := query.Use(DB).User
	users, err := userModel.WithContext(context.Background()).Where(userModel.CreatedAt.Between(to, from)).Find()
	return users, err
}
