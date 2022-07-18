package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDealer = "dealers"

// Dealer mapped from table <dealers>
type Dealer struct {
	ID             int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt      time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Handle         string         `gorm:"column:handle" json:"handle"`
	Token          string         `gorm:"column:token" json:"token"`
	Host           string         `gorm:"column:host" json:"host"`
	PeerID         string         `gorm:"column:peer_id" json:"peer_id"`
	Open           bool           `gorm:"column:open" json:"open"`
	LastConnection time.Time      `gorm:"column:last_connection" json:"last_connection"`
}

// TableName Dealer's table name
func (*Dealer) TableName() string {
	return TableNameDealer
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *Dealer) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *Dealer) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *Dealer) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *Dealer) TableInfo() *TableInfo {
	return nil
}
