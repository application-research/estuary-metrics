package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/application-research/estuary-metrics/core"
	"github.com/application-research/estuary-metrics/core/generated/model"
	"github.com/whyrusleeping/memo"
	"gorm.io/gorm"
	"reflect"
	"time"
)

// BuildInfo is used to define the application build info, and inject values into via the build process.
type BuildInfo struct {

	// BuildDate date string of when build was performed filled in by -X compile flag
	BuildDate string

	// LatestCommit date string of when build was performed filled in by -X compile flag
	LatestCommit string

	// BuildNumber date string of when build was performed filled in by -X compile flag
	BuildNumber string

	// BuiltOnIP date string of when build was performed filled in by -X compile flag
	BuiltOnIP string

	// BuiltOnOs date string of when build was performed filled in by -X compile flag
	BuiltOnOs string

	// RuntimeVer date string of when build was performed filled in by -X compile flag
	RuntimeVer string
}

type LogSql func(ctx context.Context, sql string)

var (
	// ErrNotFound error when record not found
	ErrNotFound = fmt.Errorf("record Not Found")

	// ErrUnableToMarshalJSON error when json payload corrupt
	ErrUnableToMarshalJSON = fmt.Errorf("json payload corrupt")

	// ErrUpdateFailed error when update fails
	ErrUpdateFailed = fmt.Errorf("db update error")

	// ErrInsertFailed error when insert fails
	ErrInsertFailed = fmt.Errorf("db insert error")

	// ErrDeleteFailed error when delete fails
	ErrDeleteFailed = fmt.Errorf("db delete error")

	// ErrBadParams error when bad params passed in
	ErrBadParams = fmt.Errorf("bad params error")

	// DB reference to database
	DB *gorm.DB

	//	Caching
	Cacher *memo.Cacher

	// AppBuildInfo reference to build info
	AppBuildInfo *BuildInfo

	// Logger function that will be invoked before executing sql
	Logger LogSql

	Metrics *core.Metrics
)

// Copy a src struct into a destination struct
func Copy(dst interface{}, src interface{}) error {
	dstV := reflect.Indirect(reflect.ValueOf(dst))
	srcV := reflect.Indirect(reflect.ValueOf(src))

	if !dstV.CanAddr() {
		return errors.New("copy to value is unaddressable")
	}

	if srcV.Type() != dstV.Type() {
		return errors.New("different types can be copied")
	}

	for i := 0; i < dstV.NumField(); i++ {
		f := srcV.Field(i)
		if !isZeroOfUnderlyingType(f.Interface()) {
			dstV.Field(i).Set(f)
		}
	}

	return nil
}

func isZeroOfUnderlyingType(x interface{}) bool {
	return x == nil || reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface())
}

func RunDynamicQuery(ctx context.Context, modelForQuery interface{}, query map[string]interface{}, page, pagesize int, order string) (results []*model.AuthToken, totalRows int64, err error) {

	resultOrm := DB.Model(modelForQuery)
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

	if err = resultOrm.Where(query).Find(&results).Error; err != nil {
		err = ErrNotFound
		return nil, -1, err
	}

	return results, totalRows, nil
}

type MonthLookUp struct {
	MonthToLook []MonthPerMonth
}
type MonthPerMonth struct {
	Month            int
	MonthFirstDay    string
	MonthLastDay     string
	Year             int
	NumberOfContents int64
}

type MonthResultResponse struct {
	MonthFrom        string
	MonthTo          string
	NumberOfContents int64
}

func AllDataOverThePastMonth(ctx context.Context, model interface{}, month int64) (results MonthLookUp, err error) {

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
		DB.Model(model).Where("created_at between ? and ?", month.MonthFirstDay, month.MonthLastDay).Count(&monthResult.NumberOfContents)
		result.MonthToLook = append(result.MonthToLook, monthResult)
	}

	return result, err

}

func AllDataOverASpecificDates(ctx context.Context, model interface{}, fromDate string, toDate string) (MonthResultResponse, error) {

	var result MonthResultResponse
	result.MonthTo = toDate
	result.MonthFrom = fromDate

	err := DB.Model(model).Where("created_at between ? and ?", fromDate, toDate).Count(&result.NumberOfContents).Error
	if err != nil {
		return MonthResultResponse{}, err
	}
	return result, nil
}
