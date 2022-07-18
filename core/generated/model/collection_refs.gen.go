package model

import (
	"gorm.io/gorm"
	"time"
)

const TableNameCollectionRef = "collection_refs"

// CollectionRef mapped from table <collection_refs>
type CollectionRef struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"created_at"`
	Collection int64     `gorm:"column:collection" json:"collection"`
	Content    int64     `gorm:"column:content" json:"content"`
	Path       string    `gorm:"column:path" json:"path"`
}

// TableName CollectionRef's table name
func (*CollectionRef) TableName() string {
	return TableNameCollectionRef
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *CollectionRef) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *CollectionRef) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *CollectionRef) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *CollectionRef) TableInfo() *TableInfo {
	return nil
}
