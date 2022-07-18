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

// GetAllProposalRecords is a function to get a slice of record(s) from proposal_records table in the estuary database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - ErrNotFound, db Find error
func GetAllProposalRecords(ctx context.Context, page, pagesize int, order string) (results []*model.ProposalRecord, totalRows int64, err error) {

	resultOrm := DB.Model(&model.ProposalRecord{})
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

// GetProposalRecords is a function to get a single record from the proposal_records table in the estuary database
// error - ErrNotFound, db Find error
func GetProposalRecords(ctx context.Context, argPropCid string) (record *model.ProposalRecord, err error) {
	record = &model.ProposalRecord{}
	if err = DB.First(record, argPropCid).Error; err != nil {
		err = ErrNotFound
		return record, err
	}

	return record, nil
}

// AddProposalRecords is a function to add a single record to proposal_records table in the estuary database
// error - ErrInsertFailed, db save call failed
func AddProposalRecords(ctx context.Context, record *model.ProposalRecord) (result *model.ProposalRecord, RowsAffected int64, err error) {
	db := DB.Save(record)
	if err = db.Error; err != nil {
		return nil, -1, ErrInsertFailed
	}

	return record, db.RowsAffected, nil
}

// UpdateProposalRecords is a function to update a single record from proposal_records table in the estuary database
// error - ErrNotFound, db record for id not found
// error - ErrUpdateFailed, db meta data copy failed or db.Save call failed
func UpdateProposalRecords(ctx context.Context, argPropCid string, updated *model.ProposalRecord) (result *model.ProposalRecord, RowsAffected int64, err error) {

	result = &model.ProposalRecord{}
	db := DB.First(result, argPropCid)
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

// DeleteProposalRecords is a function to delete a single record from proposal_records table in the estuary database
// error - ErrNotFound, db Find error
// error - ErrDeleteFailed, db Delete failed error
func DeleteProposalRecords(ctx context.Context, argPropCid string) (rowsAffected int64, err error) {

	record := &model.ProposalRecord{}
	db := DB.First(record, argPropCid)
	if db.Error != nil {
		return -1, ErrNotFound
	}

	db = db.Delete(record)
	if err = db.Error; err != nil {
		return -1, ErrDeleteFailed
	}

	return db.RowsAffected, nil
}
