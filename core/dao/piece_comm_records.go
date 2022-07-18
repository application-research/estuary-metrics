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

// GetAllPieceCommRecords is a function to get a slice of record(s) from piece_comm_records table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllPieceCommRecords(ctx context.Context, page, pagesize int, order string) (results []*model.PieceCommRecord, totalRows int64, err error) {

	resultOrm := DB.Model(&model.PieceCommRecord{})
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

// GetPieceCommRecords is a function to get a single record from the piece_comm_records table in the estuary database
// error - ErrNotFound, db Find error
func GetPieceCommRecords(ctx context.Context, argData string) (record *model.PieceCommRecord, err error) {
	record = &model.PieceCommRecord{}
	if err = DB.First(record, argData).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddPieceCommRecords is a function to add a single record to piece_comm_records table in the estuary database
// error - ErrInsertFailed, db save call failed
func AddPieceCommRecords(ctx context.Context, record *model.PieceCommRecord) (result *model.PieceCommRecord, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdatePieceCommRecords is a function to update a single record from piece_comm_records table in the estuary database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdatePieceCommRecords(ctx context.Context, argData string, updated *model.PieceCommRecord) (result *model.PieceCommRecord, RowsAffected int64, err error) {

	result = &model.PieceCommRecord{}
	db := DB.First(result, argData)
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

// DeletePieceCommRecords is a function to delete a single record from piece_comm_records table in the estuary database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeletePieceCommRecords(ctx context.Context, argData string) (rowsAffected int64, err error) {

	record := &model.PieceCommRecord{}
	db := DB.First(record, argData)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
