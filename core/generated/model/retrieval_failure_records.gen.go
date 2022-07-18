package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameRetrievalFailureRecord = "retrieval_failure_records"

// RetrievalFailureRecord mapped from table <retrieval_failure_records>
type RetrievalFailureRecord struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Miner     string         `gorm:"column:miner" json:"miner"`
	Phase     string         `gorm:"column:phase" json:"phase"`
	Message   string         `gorm:"column:message" json:"message"`
	Content   int64          `gorm:"column:content" json:"content"`
	Cid       []uint8        `gorm:"column:cid" json:"cid"`
}

// TableName RetrievalFailureRecord's table name
func (*RetrievalFailureRecord) TableName() string {
	return TableNameRetrievalFailureRecord
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *RetrievalFailureRecord) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *RetrievalFailureRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *RetrievalFailureRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *RetrievalFailureRecord) TableInfo() *TableInfo {
	return nil
}
