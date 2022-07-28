package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameUser = "users"

// User mapped from table <users>
type User struct {
	ID              int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt       time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	UUID            string         `gorm:"column:uuid" json:"uuid"`
	Username        string         `gorm:"column:username" json:"username"`
	Salt            string         `gorm:"column:salt" json:"salt"`
	PassHash        string         `gorm:"column:pass_hash" json:"pass_hash"`
	UserEmail       string         `gorm:"column:user_email" json:"user_email"`
	Perm            int64          `gorm:"column:perm" json:"perm"`
	Flags           int64          `gorm:"column:flags" json:"flags"`
	Address         string         `gorm:"column:address" json:"address"`
	StorageDisabled bool           `gorm:"column:storage_disabled" json:"storage_disabled"`
	DID             string         `gorm:"column:d_id" json:"d_id"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *User) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *User) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *User) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *User) TableInfo() *TableInfo {
	return nil
}
