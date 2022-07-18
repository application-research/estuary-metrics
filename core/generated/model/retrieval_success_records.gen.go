package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameRetrievalSuccessRecord = "retrieval_success_records"

// RetrievalSuccessRecord mapped from table <retrieval_success_records>
type RetrievalSuccessRecord struct {
	PropCid      []uint8   `gorm:"column:prop_cid;type:bytea" json:"prop_cid"`
	Miner        string    `gorm:"column:miner;type:text" json:"miner"`
	Peer         string    `gorm:"column:peer;type:text" json:"peer"`
	Size         int64     `gorm:"column:size;type:int8" json:"size"`
	DurationMs   int64     `gorm:"column:duration_ms;type:int8" json:"duration_ms"`
	AverageSpeed int64     `gorm:"column:average_speed;type:int8" json:"average_speed"`
	TotalPayment string    `gorm:"column:total_payment;type:text" json:"total_payment"`
	NumPayments  int64     `gorm:"column:num_payments;type:int8" json:"num_payments"`
	AskPrice     string    `gorm:"column:ask_price;type:text" json:"ask_price"`
	ID           int64     `gorm:"column:id;type:int8;not null" json:"id"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamptz" json:"created_at"`
	Cid          []uint8   `gorm:"column:cid;type:bytea" json:"cid"`
}

// TableName RetrievalSuccessRecord's table name
func (*RetrievalSuccessRecord) TableName() string {
	return TableNameRetrievalSuccessRecord
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *RetrievalSuccessRecord) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *RetrievalSuccessRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *RetrievalSuccessRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *RetrievalSuccessRecord) TableInfo() *TableInfo {
	return nil
}