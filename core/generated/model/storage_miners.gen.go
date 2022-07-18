package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameStorageMiner = "storage_miners"

// StorageMiner mapped from table <storage_miners>
type StorageMiner struct {
	ID              int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Address         string         `gorm:"column:address" json:"address"`
	Suspended       bool           `gorm:"column:suspended" json:"suspended"`
	SuspendedReason string         `gorm:"column:suspended_reason" json:"suspended_reason"`
	Name            string         `gorm:"column:name" json:"name"`
	Version         string         `gorm:"column:version" json:"version"`
	Location        string         `gorm:"column:location" json:"location"`
	Owner           int64          `gorm:"column:owner" json:"owner"`
}

// TableName StorageMiner's table name
func (*StorageMiner) TableName() string {
	return TableNameStorageMiner
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *StorageMiner) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *StorageMiner) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *StorageMiner) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *StorageMiner) TableInfo() *TableInfo {
	return nil
}