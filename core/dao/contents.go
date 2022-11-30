package dao

import (
	"context"
	"github.com/gofrs/uuid"
	"time"

	"github.com/application-research/estuary-metrics/core/generated/model"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAllContents is a function to get a slice of record(s) from contents table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllContents(ctx context.Context, page, pagesize int, order string) (results []*model.Content, totalRows int64, err error) {

	resultOrm := DB.Model(&model.Content{})
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

// GetContents is a function to get a single record from the contents table in the estuary database
// error - ErrNotFound, db Find error
func GetContents(ctx context.Context, argID int64) (record *model.Content, err error) {
	record = &model.Content{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// GetSizeOverPastMonths is a function get uploaded data sizes over the past month from the contents table
func GetSizeOverPastMonths(ctx context.Context, months int64) (results MonthLookUp, err error) {
	monthLookUp := ConstructMonthLookUpList(months)

	var result MonthLookUp
	for _, month := range monthLookUp.MonthToLook {
		var monthResult MonthPerMonth
		monthResult.Month = month.Month
		monthResult.MonthFirstDay = month.MonthFirstDay
		monthResult.MonthLastDay = month.MonthLastDay
		monthResult.Year = month.Year
		DB.Model(model.Content{}).Where("created_at between ? and ?", month.MonthFirstDay, month.MonthLastDay).Select("sum(size) as size").Scan(&monthResult.AggregatedResult)
		result.MonthToLook = append(result.MonthToLook, monthResult)
	}

	return result, err

}

// AddContents is a function to add a single record to contents table in the estuary database
// error - ErrInsertFailed, db save call failed
func AddContents(ctx context.Context, record *model.Content) (result *model.Content, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateContents is a function to update a single record from contents table in the estuary database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateContents(ctx context.Context, argID int64, updated *model.Content) (result *model.Content, RowsAffected int64, err error) {

	result = &model.Content{}
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

// DeleteContents is a function to delete a single record from contents table in the estuary database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteContents(ctx context.Context, argID int64) (rowsAffected int64, err error) {

	record := &model.Content{}
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
