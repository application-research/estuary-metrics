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

// GetAllPublishedBatches is a function to get a slice of record(s) from published_batches table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllPublishedBatches(ctx context.Context, page, pagesize int, order string) (results []*model.PublishedBatch, totalRows int64, err error) {

	resultOrm := DB.Model(&model.PublishedBatch{})
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

// GetPublishedBatches is a function to get a single record from the published_batches table in the estuary database
// params - argID - primary key value
// error - ErrNotFound, db Find error
func GetPublishedBatches(ctx context.Context, argID int64) (record *model.PublishedBatch, err error) {
	record = &model.PublishedBatch{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

type PublishedBatchesCount struct {
	Count int64
}

func GetTotalPublishedBatches(ctx context.Context) (record PublishedBatchesCount, err error) {
	if err = DB.Raw("SELECT sum(count)FROM published_batches").Scan(&record.Count).Error; err != nil {
		err = ErrNotFound
		return record, err
	}
	return record, nil
}
