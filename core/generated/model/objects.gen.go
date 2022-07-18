package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameObject = "objects"

// Object mapped from table <objects>
type Object struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Cid        []uint8   `gorm:"column:cid" json:"cid"`
	Size       int64     `gorm:"column:size" json:"size"`
	Reads      int64     `gorm:"column:reads" json:"reads"`
	LastAccess time.Time `gorm:"column:last_access" json:"last_access"`
}

// TableName Object's table name
func (*Object) TableName() string {
	return TableNameObject
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Object) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Object) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Object) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Object) TableInfo() *TableInfo {
	return nil
}