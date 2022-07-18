package model

import "gorm.io/gorm"

const TableNameProposalRecord = "proposal_records"

// ProposalRecord mapped from table <proposal_records>
type ProposalRecord struct {
	PropCid []uint8 `gorm:"column:prop_cid" json:"prop_cid"`
	Data    []uint8 `gorm:"column:data" json:"data"`
}

// TableName ProposalRecord's table name
func (*ProposalRecord) TableName() string {
	return TableNameProposalRecord
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ProposalRecord) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ProposalRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ProposalRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ProposalRecord) TableInfo() *TableInfo {
	return nil
}