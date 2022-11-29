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

type MonthLookUp struct {
	MonthToLook []MonthPerMonth
}
type MonthPerMonth struct {
	Month            int
	MonthFirstDay    string
	MonthLastDay     string
	Year             int
	NumberOfContents int
}

type MonthResultResponse struct {
	MonthFrom        string
	MonthTo          string
	NumberOfContents int
}

func AllContentOverASpecificDates(ctx context.Context, fromDate string, toDate string) (MonthResultResponse, error) {

	var result MonthResultResponse
	result.MonthTo = toDate
	result.MonthFrom = fromDate

	//(select count(*) from contents where (created_at between '2000-01-01' and '2022-11-28') and pinning and not active and not failed) as older
	err := DB.Raw("select count(*) from contents where (created_at between ? and ?) and pinning and not active and not failed", fromDate, toDate).Scan(&result.NumberOfContents).Error

	if err != nil {
		return MonthResultResponse{}, err
	}
	return result, nil
}
func AllContentOverThePastMonths(ctx context.Context, month int64) (results MonthLookUp, err error) {

	// get current month
	var monthLookUp MonthLookUp
	currentTime := time.Now()
	currentLocation := currentTime.Location()
	timeLayout := "2006-01-02"

	for i := 0; i < int(month); i++ {
		fromWhenMonth := currentTime.AddDate(0, (int(month)-i)-12, 0)
		firstOfMonth := time.Date(fromWhenMonth.Year(), fromWhenMonth.Month(), 1, 0, 0, 0, 0, currentLocation)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		//
		stringFirstOfMonth := firstOfMonth.Format(timeLayout)
		stringLastOfMonth := lastOfMonth.Format(timeLayout)

		monthLookUp.MonthToLook = append(monthLookUp.MonthToLook, MonthPerMonth{Month: int(firstOfMonth.Month()), MonthFirstDay: stringFirstOfMonth, MonthLastDay: stringLastOfMonth, Year: firstOfMonth.Year()})

	}

	var result MonthLookUp
	for _, month := range monthLookUp.MonthToLook {
		var monthResult MonthPerMonth
		monthResult.Month = month.Month
		monthResult.MonthFirstDay = month.MonthFirstDay
		monthResult.MonthLastDay = month.MonthLastDay
		monthResult.Year = month.Year

		//(select count(*) from contents where (created_at between '2000-01-01' and '2022-11-28') and pinning and not active and not failed) as older
		DB.Raw("select count(*) from contents where (created_at between ? and ?) and pinning and not active and not failed", month.MonthFirstDay, month.MonthLastDay).Scan(&monthResult.NumberOfContents)
		result.MonthToLook = append(result.MonthToLook, monthResult)

	}

	return result, err

}

//
//func AllContentOverThePastHours(ctx context.Context) (results []*model.Content, err error) {
//	if err = DB.Model(&model.Content{}).Where("created_at > ?", time.Now().AddDate(0, -12, 0)).Find(&results).Error; err != nil {
//		err = ErrNotFound
//		return nil, err
//	}
//
//	return results, nil
//}
//
func AllContentOverThePastDays(ctx context.Context) (results []*model.Content, err error) {
	if err = DB.Model(&model.Content{}).Where("created_at > ?", time.Now().AddDate(0, -12, 0)).Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, err
	}

	return results, nil
}

//
//func AllContentOverThePastMin(ctx context.Context) (results []*model.Content, err error) {
//	if err = DB.Model(&model.Content{}).Where("created_at > ?", time.Now().AddDate(0, -12, 0)).Find(&results).Error; err != nil {
//		err = ErrNotFound
//		return nil, err
//	}
//
//	return results, nil
//}
