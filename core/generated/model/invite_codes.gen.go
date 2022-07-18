package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameInviteCode = "invite_codes"

// InviteCode mapped from table <invite_codes>
type InviteCode struct {
	ID        int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Code      string         `gorm:"column:code" json:"code"`
	CreatedBy int64          `gorm:"column:created_by" json:"created_by"`
	ClaimedBy int64          `gorm:"column:claimed_by" json:"claimed_by"`
}

// TableName InviteCode's table name
func (*InviteCode) TableName() string {
	return TableNameInviteCode
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *InviteCode) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *InviteCode) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *InviteCode) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *InviteCode) TableInfo() *TableInfo {
	return nil
}
