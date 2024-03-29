package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameMinerStorageAsk = "miner_storage_asks"

// MinerStorageAsk mapped from table <miner_storage_asks>
type MinerStorageAsk struct {
	ID            int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	Miner         string         `gorm:"column:miner" json:"miner"`
	Price         string         `gorm:"column:price" json:"price"`
	VerifiedPrice string         `gorm:"column:verified_price" json:"verified_price"`
	MinPieceSize  int64          `gorm:"column:min_piece_size" json:"min_piece_size"`
	MaxPieceSize  int64          `gorm:"column:max_piece_size" json:"max_piece_size"`
}

// TableName MinerStorageAsk's table name
func (*MinerStorageAsk) TableName() string {
	return TableNameMinerStorageAsk
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *MinerStorageAsk) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *MinerStorageAsk) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *MinerStorageAsk) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *MinerStorageAsk) TableInfo() *TableInfo {
	return nil
}
