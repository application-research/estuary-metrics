// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameDfeRecord = "dfe_records"

// DfeRecord mapped from table <dfe_records>
type DfeRecord struct {
	ID                  int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt           time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt           time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Miner               string         `gorm:"column:miner" json:"miner"`
	Phase               string         `gorm:"column:phase" json:"phase"`
	Message             string         `gorm:"column:message" json:"message"`
	Content             int64          `gorm:"column:content" json:"content"`
	MinerVersion        string         `gorm:"column:miner_version" json:"miner_version"`
	UserID              int64          `gorm:"column:user_id" json:"user_id"`
	DealProtocolVersion string         `gorm:"column:deal_protocol_version" json:"deal_protocol_version"`
	DealUUID            string         `gorm:"column:deal_uuid" json:"deal_uuid"`
}

// TableName DfeRecord's table name
func (*DfeRecord) TableName() string {
	return TableNameDfeRecord
}
