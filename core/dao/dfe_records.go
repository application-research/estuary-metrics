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

// GetAllDfeRecords is a function to get a slice of record(s) from dfe_records table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllDfeRecords(ctx context.Context, page, pagesize int, order string) (results []*model.DfeRecord, totalRows int64, err error) {

	resultOrm := DB.Model(&model.DfeRecord{})
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

// GetDfeRecords is a function to get a single record from the dfe_records table in the estuary database
// error - ErrNotFound, db Find error
func GetDfeRecords(ctx context.Context, argID int64) (record *model.DfeRecord, err error) {
	record = &model.DfeRecord{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

func GetDfeStorageFailureRecords(ctx context.Context, limit int64, before string, u *model.User) ([]model.DfeRecord, error) {
	var defLimit = 2000
	if limit == 0 {
		limit = int64(defLimit)
	}

	q := DB.Model(model.DfeRecord{}).Limit(int(limit)).Order("created_at desc")
	if u != nil {
		q = q.Where("user_id=?", u.ID)
	}

	if bef := before; bef != "" {
		beftime, err := time.Parse(time.RFC3339, bef)
		if err != nil {
			return nil, err
		}
		q = q.Where("created_at <= ?", beftime)
	}

	var recs []model.DfeRecord
	if err := q.Scan(&recs).Error; err != nil {
		return nil, err
	}
	return recs, nil
}
