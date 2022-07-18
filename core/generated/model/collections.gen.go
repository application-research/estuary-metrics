package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameCollection = "collections"

// Collection mapped from table <collections>
type Collection struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UUID        string    `gorm:"column:uuid" json:"uuid"`
	Name        string    `gorm:"column:name" json:"name"`
	Description string    `gorm:"column:description" json:"description"`
	UserID      int64     `gorm:"column:user_id" json:"user_id"`
	CID         string    `gorm:"column:c_id" json:"c_id"`
}

// TableName Collection's table name
func (*Collection) TableName() string {
	return TableNameCollection
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Collection) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Collection) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Collection) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Collection) TableInfo() *TableInfo {
	return nil
}