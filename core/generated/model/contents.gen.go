package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameContent = "contents"

// Content mapped from table <contents>
type Content struct {
	ID           int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt    time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Cid          []uint8        `gorm:"column:cid" json:"cid"`
	Name         string         `gorm:"column:name" json:"filename"`
	User         string         `gorm:"column:user" json:"user"`
	Size         int64          `gorm:"column:size" json:"size"`
	Active       bool           `gorm:"column:active" json:"active"`
	Description  string         `gorm:"column:description" json:"description"`
	Offloaded    bool           `gorm:"column:offloaded" json:"offloaded"`
	UserID       int64          `gorm:"column:user_id" json:"user_id"`
	Replication  int64          `gorm:"column:replication" json:"replication"`
	AggregatedIn int64          `gorm:"column:aggregated_in" json:"aggregated_in"`
	Aggregate    bool           `gorm:"column:aggregate" json:"aggregate"`
	Pinning      bool           `gorm:"column:pinning" json:"pinning"`
	PinMeta      string         `gorm:"column:pin_meta" json:"pin_meta"`
	Location     string         `gorm:"column:location" json:"location"`
	Failed       bool           `gorm:"column:failed" json:"failed"`
	DagSplit     bool           `gorm:"column:dag_split" json:"dag_split"`
	SplitFrom    int64          `gorm:"column:split_from" json:"split_from"`
	Type         int64          `gorm:"column:type" json:"type"`
	Replace      bool           `gorm:"column:replace" json:"replace"`
	Origins      string         `gorm:"column:origins" json:"origins"`
}

// TableName Content's table name
func (*Content) TableName() string {
	return TableNameContent
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Content) BeforeSave(DB *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Content) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Content) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Content) TableInfo() *TableInfo {
	return nil
}
