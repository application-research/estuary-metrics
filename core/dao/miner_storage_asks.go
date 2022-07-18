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

// GetAllMinerStorageAsks is a function to get a slice of record(s) from miner_storage_asks table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllMinerStorageAsks(ctx context.Context, page, pagesize int, order string) (results []*model.MinerStorageAsk, totalRows int64, err error) {

	resultOrm := DB.Model(&model.MinerStorageAsk{})
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

// GetMinerStorageAsks is a function to get a single record from the miner_storage_asks table in the estuary database
// error - ErrNotFound, db Find error
func GetMinerStorageAsks(ctx context.Context, argID int64) (record *model.MinerStorageAsk, err error) {
	record = &model.MinerStorageAsk{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddMinerStorageAsks is a function to add a single record to miner_storage_asks table in the estuary database
// error - ErrInsertFailed, db save call failed
func AddMinerStorageAsks(ctx context.Context, record *model.MinerStorageAsk) (result *model.MinerStorageAsk, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateMinerStorageAsks is a function to update a single record from miner_storage_asks table in the estuary database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateMinerStorageAsks(ctx context.Context, argID int64, updated *model.MinerStorageAsk) (result *model.MinerStorageAsk, RowsAffected int64, err error) {

	result = &model.MinerStorageAsk{}
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

// DeleteMinerStorageAsks is a function to delete a single record from miner_storage_asks table in the estuary database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteMinerStorageAsks(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.MinerStorageAsk{}
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
