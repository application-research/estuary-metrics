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

// GetAllCollectionRefs is a function to get a slice of record(s) from collection_refs table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllCollectionRefs(ctx context.Context, page, pagesize int, order string) (results []*model.CollectionRef, totalRows int64, err error) {

	resultOrm := DB.Model(&model.CollectionRef{})
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

// GetCollectionRefs is a function to get a single record from the collection_refs table in the estuary database
// error - ErrNotFound, db Find error
func GetCollectionRefs(ctx context.Context, argID int64) (record *model.CollectionRef, err error) {
	record = &model.CollectionRef{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddCollectionRefs is a function to add a single record to collection_refs table in the estuary database
// error - ErrInsertFailed, db save call failed
func AddCollectionRefs(ctx context.Context, record *model.CollectionRef) (result *model.CollectionRef, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateCollectionRefs is a function to update a single record from collection_refs table in the estuary database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateCollectionRefs(ctx context.Context, argID int64, updated *model.CollectionRef) (result *model.CollectionRef, RowsAffected int64, err error) {

	result = &model.CollectionRef{}
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

// DeleteCollectionRefs is a function to delete a single record from collection_refs table in the estuary database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteCollectionRefs(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.CollectionRef{}
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
