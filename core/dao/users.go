package dao

import (
	"context"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/guregu/null"
	"github.com/satori/go.uuid"
	"time"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

func GetUsersDynamicQuery(ctx context.Context, query map[string]interface{}, page, pagesize int, order string) (results []*model.User, err error) {
	resultOrm := DB.Model(&model.User{})

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Where(query).Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, err
	}

	return results, nil

}

func GetNumberOfUsersWithinRange(ctx context.Context, start, end time.Time) ([]*model.User, error) {
	var users []*model.User
	err := DB.Model(&model.User{}).Select("id", "username", "user_email", "perm").Where("created_at >= ? AND created_at <= ?", start, end).Find(&users).Error
	return users, err
}

func GetNumberOfUsers(ctx context.Context) (int64, error) {
	return Metrics.GetNumberOfUsers()
}

// GetAllUsers is a function to get a slice of record(s) from users table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllUsers(ctx context.Context, page, pagesize int, order string) (results []*model.User, totalRows int64, err error) {

	resultOrm := DB.Model(&model.User{})
	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	if err = resultOrm.Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

// GetUsers is a function to get a single record from the users table in the estuary database
// error - ErrNotFound, db Find error
func GetUsers(ctx context.Context, argID int64) (record *model.User, err error) {
	record = &model.User{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

func GetTopUsers(ctx context.Context, top int) (record []*model.User, err error) {
	record = []*model.User{}
	if err = DB.Order("total_bytes DESC").Limit(top).Find(&record).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

func GetTopUsersWithCollections(ctx context.Context, top int) (record []*model.User, err error) {
	record = []*model.User{}
	if err = DB.Order("total_bytes DESC").Limit(top).Find(&record).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}
