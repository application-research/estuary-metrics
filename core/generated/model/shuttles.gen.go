package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameShuttle = "shuttles"

// Shuttle mapped from table <shuttles>
type Shuttle struct {
	ID             int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Handle         string         `gorm:"column:handle" json:"handle"`
	Token          string         `gorm:"column:token" json:"token"`
	LastConnection time.Time      `gorm:"column:last_connection" json:"last_connection"`
	Host           string         `gorm:"column:host" json:"host"`
	PeerID         string         `gorm:"column:peer_id" json:"peer_id"`
	Open           bool           `gorm:"column:open" json:"open"`
	Private        bool           `gorm:"column:private" json:"private"`
	Priority       int64          `gorm:"column:priority" json:"priority"`
}

// TableName Shuttle's table name
func (*Shuttle) TableName() string {
	return TableNameShuttle
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Shuttle) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Shuttle) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Shuttle) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Shuttle) TableInfo() *TableInfo {
	return nil
}
