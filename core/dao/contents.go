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

type MonthDataSizeLookUp struct {
	MonthToLook []MonthDataSizeResult
}

type MonthDataSizeResult struct {
	Month                int
	MonthFirstDay        string
	MonthLastDay         string
	Year                 int
	SizeTotalForTheMonth int64
}

func AllContentDataSizeOverThePastMonth(ctx context.Context, model interface{}, month int64) (MonthDataSizeLookUp, error) {

	// get current month
	var monthLookUp MonthLookUp
	currentTime := time.Now()
	currentLocation := currentTime.Location()
	timeLayout := "2006-01-02"

	for i := 0; i < int(month); i++ {
		fromWhenMonth := currentTime.AddDate(0, (int(month)-i)-12, 0)
		firstOfMonth := time.Date(fromWhenMonth.Year(), fromWhenMonth.Month(), 1, 0, 0, 0, 0, currentLocation)
		lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

		stringFirstOfMonth := firstOfMonth.Format(timeLayout)
		stringLastOfMonth := lastOfMonth.Format(timeLayout)

		monthLookUp.MonthToLook = append(monthLookUp.MonthToLook, MonthPerMonth{Month: int(firstOfMonth.Month()), MonthFirstDay: stringFirstOfMonth, MonthLastDay: stringLastOfMonth, Year: firstOfMonth.Year()})

	}

	var result MonthDataSizeLookUp
	var err error
	for _, month := range monthLookUp.MonthToLook {
		var monthResult MonthDataSizeResult
		monthResult.Month = month.Month
		monthResult.MonthFirstDay = month.MonthFirstDay
		monthResult.MonthLastDay = month.MonthLastDay
		monthResult.Year = month.Year

		//select sum(size) from contents where (created_at between '2000-01-01' and '2022-08-28')
		if err = DB.Model(model).Where("created_at between ? and ?", month.MonthFirstDay, month.MonthLastDay).Select("sum(size) as size").Scan(&monthResult.SizeTotalForTheMonth).Error; err != nil {
			result.MonthToLook = append(result.MonthToLook, monthResult)
		}
	}

	return result, err

}
