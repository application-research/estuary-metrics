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

// GetAllRetrievalSuccessRecords is a function to get a slice of record(s) from retrieval_success_records table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRetrievalSuccessRecords(ctx context.Context, page, pagesize int, order string) (results []*model.RetrievalSuccessRecord, totalRows int64, err error) {

	resultOrm := DB.Model(&model.RetrievalSuccessRecord{})
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

// GetRetrievalSuccessRecords is a function to get a single record from the retrieval_success_records table in the estuary database
// error - ErrNotFound, db Find error
func GetRetrievalSuccessRecords(ctx context.Context, argPropCid string) (record *model.RetrievalSuccessRecord, err error) {
	record = &model.RetrievalSuccessRecord{}
	if err = DB.First(record, argPropCid).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}
