package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAutoretrieve = "autoretrieves"

// Autoretriefe mapped from table <autoretrieves>
type Autoretrieve struct {
	ID             int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Handle         string         `gorm:"column:handle" json:"handle"`
	Token          string         `gorm:"column:token" json:"token"`
	LastConnection time.Time      `gorm:"column:last_connection" json:"last_connection"`
	PeerID         string         `gorm:"column:peer_id" json:"peer_id"`
	Addresses      string         `gorm:"column:addresses" json:"addresses"`
}

// TableName Autoretriefe's table name
func (*Autoretrieve) TableName() string {
	return TableNameAutoretrieve
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Autoretrieve) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Autoretrieve) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Autoretrieve) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Autoretrieve) TableInfo() *TableInfo {
	return nil
}
