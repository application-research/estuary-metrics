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

// GetAllStorageMiners is a function to get a slice of record(s) from storage_miners table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllStorageMiners(ctx context.Context, page, pagesize int, order string) (results []*model.StorageMiner, totalRows int64, err error) {

	resultOrm := DB.Model(&model.StorageMiner{})
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

// GetStorageMiners is a function to get a single record from the storage_miners table in the estuary database
// error - ErrNotFound, db Find error
func GetStorageMiners(ctx context.Context, argID int64) (record *model.StorageMiner, err error) {
	record = &model.StorageMiner{}
	if err = DB.First(record, argID).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

type TopMiner struct {
	Miner string
	Name  string
	Count int64
}

func GetTopStorageMiners(ctx context.Context, top int) (results []*TopMiner, err error) {
	//select a.miner, b.name, count(*) from content_deals a, storage_miners b where a.miner = b.address group by a.miner, b.name order by count(*) desc limit 10;
	resultOrm := DB.Raw("select a.miner, b.name, count(*) from content_deals a, storage_miners b where a.miner = b.address group by a.miner, b.name order by count(*) desc limit ?", top)

	if err = resultOrm.Scan(&results).Error; err != nil {
		err = ErrNotFound
		return nil, err
	}

	return results, nil
}
