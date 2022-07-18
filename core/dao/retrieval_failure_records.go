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

// GetAllRetrievalFailureRecords is a function to get a slice of record(s) from retrieval_failure_records table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllRetrievalFailureRecords(ctx context.Context, page, pagesize int, order string) (results []*model.RetrievalFailureRecord, totalRows int64, err error) {

	resultOrm := DB.Model(&model.RetrievalFailureRecord{})
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

// GetRetrievalFailureRecords is a function to get a single record from the retrieval_failure_records table in the estuary database
// error - ErrNotFound, db Find error
func GetRetrievalFailureRecords(ctx context.Context, argID int64) (record *model.RetrievalFailureRecord, err error) {
	record = &model.RetrievalFailureRecord{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddRetrievalFailureRecords is a function to add a single record to retrieval_failure_records table in the estuary database
// error - ErrInsertFailed, db save call failed
func AddRetrievalFailureRecords(ctx context.Context, record *model.RetrievalFailureRecord) (result *model.RetrievalFailureRecord, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateRetrievalFailureRecords is a function to update a single record from retrieval_failure_records table in the estuary database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateRetrievalFailureRecords(ctx context.Context, argID int64, updated *model.RetrievalFailureRecord) (result *model.RetrievalFailureRecord, RowsAffected int64, err error) {

	result = &model.RetrievalFailureRecord{}
	db := DB.First(result, argID)
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// DeleteRetrievalFailureRecords is a function to delete a single record from retrieval_failure_records table in the estuary database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteRetrievalFailureRecords(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.RetrievalFailureRecord{}
	db := DB.First(record, argID)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
