package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameContentDeal = "content_deals"

// ContentDeal mapped from table <content_deals>
type ContentDeal struct {
	ID               int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt        time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Content          int64          `gorm:"column:content" json:"content"`
	PropCid          []uint8        `gorm:"column:prop_cid" json:"prop_cid"`
	Miner            string         `gorm:"column:miner" json:"miner"`
	DealID           int64          `gorm:"column:deal_id" json:"deal_id"`
	Failed           bool           `gorm:"column:failed" json:"failed"`
	FailedAt         time.Time      `gorm:"column:failed_at" json:"failed_at"`
	DtChan           string         `gorm:"column:dt_chan" json:"dt_chan"`
	Verified         bool           `gorm:"column:verified" json:"verified"`
	SealedAt         time.Time      `gorm:"column:sealed_at" json:"sealed_at"`
	OnChainAt        time.Time      `gorm:"column:on_chain_at" json:"on_chain_at"`
	TransferStarted  time.Time      `gorm:"column:transfer_started" json:"transfer_started"`
	TransferFinished time.Time      `gorm:"column:transfer_finished" json:"transfer_finished"`
	DealUUID         string         `gorm:"column:deal_uuid" json:"deal_uuid"`
	UserID           int64          `gorm:"column:user_id" json:"user_id"`
	Slashed          bool           `gorm:"column:slashed" json:"slashed"`
}

// TableName ContentDeal's table name
func (*ContentDeal) TableName() string {
	return TableNameContentDeal
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ContentDeal) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ContentDeal) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ContentDeal) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ContentDeal) TableInfo() *TableInfo {
	return nil
}
