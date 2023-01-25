package dao

import (
	"context"
	"time"

	model "github.com/application-research/estuary-metrics/core/generated/query/model"
	"github.com/gofrs/uuid"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = null.Bool{}
	_ = uuid.UUID{}
)

// GetAuthTokensDynamicQuery is a function to get a slice of record(s) from auth_tokens table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// results - slice of record(s)
// totalRows - total number of records in the auth_tokens table
// error - ErrNotFound, db Find error
func GetAuthTokensDynamicQuery(ctx context.Context, query map[string]interface{}, page, pagesize int, order string) (results []*model.AuthToken, totalRows int64, err error) {
	return RunDynamicQuery(ctx, &model.AuthToken{}, query, page, pagesize, order)
}

// GetAllActiveAuthTokenCount is a function to get the count of all active records in auth_tokens table in the estuary database
// count - count of active records
// error - ErrNotFound, db Find error
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

type AuthTokensStats struct {
	Issued  int64 `json:"issued"`
	Active  int64 `json:"active"`
	Expired int64 `json:"expired"`
	Deleted int64 `json:"deleted"`
}

func GetAuthTokensIssued(ctx context.Context) (record *AuthTokensStats, err error) {

	err = DB.Raw("select (select count(*) from auth_tokens) \"Issued\", (select count(*) from auth_tokens a where a.expiry > now()) \"Active\", (select count(*) from auth_tokens a where a.expiry < now()) \"Expired\", (select count(*) from auth_tokens a where a.deleted_at is not null) \"Deleted\"").Scan(&record).Error
	if err != nil {
		return nil, err
	}
	return record, nil
}
