package dao

import (
	"context"
	"time"

	model "github.com/application-research/estuary-metrics/core/generated/model"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

func GetAllActiveAuthTokenCount(ctx context.Context) (count int64, err error) {
	resultOrm := DB.Model(&model.AuthToken{})
	resultOrm.Where("deleted_at <> ?", nil).Count(&count)
	return count, nil
}

// GetAllAuthTokens is a function to get a slice of record(s) from auth_tokens table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllAuthTokens(ctx context.Context, page, pagesize int, order string) (results []*model.AuthToken, totalRows int64, err error) {
	resultOrm := DB.Model(&model.AuthToken{})
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

// GetAuthTokens is a function to get a single record from the auth_tokens table in the estuary database
// error - ErrNotFound, db Find error
func GetAuthTokens(ctx context.Context, argID int64) (record *model.AuthToken, err error) {
	record = &model.AuthToken{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddAuthTokens is a function to add a single record to auth_tokens table in the estuary database
// error - ErrInsertFailed, db save call failed
func AddAuthTokens(ctx context.Context, record *model.AuthToken) (result *model.AuthToken, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateAuthTokens is a function to update a single record from auth_tokens table in the estuary database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateAuthTokens(ctx context.Context, argID int64, updated *model.AuthToken) (result *model.AuthToken, RowsAffected int64, err error) {

	result = &model.AuthToken{}
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

// DeleteAuthTokens is a function to delete a single record from auth_tokens table in the estuary database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteAuthTokens(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.AuthToken{}
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
