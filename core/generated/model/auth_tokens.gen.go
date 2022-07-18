package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameAuthToken = "auth_tokens"

// AuthToken mapped from table <auth_tokens>
type AuthToken struct {
	ID         int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Token      string         `gorm:"column:token" json:"token"`
	User       int64          `gorm:"column:user" json:"user"`
	Expiry     time.Time      `gorm:"column:expiry" json:"expiry"`
	UploadOnly bool           `gorm:"column:upload_only" json:"upload_only"`
}

// TableName AuthToken's table name
func (*AuthToken) TableName() string {
	return TableNameAuthToken
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AuthToken) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AuthToken) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AuthToken) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AuthToken) TableInfo() *TableInfo {
	return nil
}
