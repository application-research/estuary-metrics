package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDfeRecord = "dfe_records"

// DfeRecord mapped from table <dfe_records>
type DfeRecord struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Miner        string         `gorm:"column:miner" json:"miner"`
	Phase        string         `gorm:"column:phase" json:"phase"`
	Message      string         `gorm:"column:message" json:"message"`
	Content      int64          `gorm:"column:content" json:"content"`
	MinerVersion string         `gorm:"column:miner_version" json:"miner_version"`
	UserID       int64          `gorm:"column:user_id" json:"user_id"`
}

// TableName DfeRecord's table name
func (*DfeRecord) TableName() string {
	return TableNameDfeRecord
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *DfeRecord) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *DfeRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *DfeRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *DfeRecord) TableInfo() *TableInfo {
	return nil
}
