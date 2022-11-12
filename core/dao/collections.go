package dao

import (
	"context"
	"time"

	"github.com/application-research/estuary-metrics/core/generated/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllCollections is a function to get a slice of record(s) from collections table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllCollections(ctx context.Context, page, pagesize int, order string) (results []*model.Collection, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Collection{})
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

// GetCollections is a function to get a single record from the collections table in the estuary database
// error - ErrNotFound, db Find error
func GetCollections(ctx context.Context, argID int64) (record *model.Collection, err error) {
	record = &model.Collection{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

type TopCollectionUsers struct {
	UserID   int64
	Username string
	Count    int64
}

func GetTopCollectionUsers(ctx context.Context, top int) (record []*TopCollectionUsers, err error) {
	//	select a.user_id, count(*) from collections a, users b where a.user_id = b.id group by a.user_id order by count(*) desc limit 10;
	if err = DB.Table("collections").Select("user_id, username, count(*)").Joins("join users on collections.user_id = users.id").Group("user_id, username").Order("count(*) desc").Limit(top).Scan(&record).Error; err != nil {
		err = ErrNotFound
		return record, err
	}
	return record, nil
}
