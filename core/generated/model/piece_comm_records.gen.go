package model

import "gorm.io/gorm"

const TableNamePieceCommRecord = "piece_comm_records"

// PieceCommRecord mapped from table <piece_comm_records>
type PieceCommRecord struct {
	Data    []uint8 `gorm:"column:data" json:"data"`
	Piece   []uint8 `gorm:"column:piece" json:"piece"`
	Size    int64   `gorm:"column:size" json:"size"`
	CarSize int64   `gorm:"column:car_size" json:"car_size"`
}

// TableName PieceCommRecord's table name
func (*PieceCommRecord) TableName() string {
	return TableNamePieceCommRecord
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *PieceCommRecord) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *PieceCommRecord) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *PieceCommRecord) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *PieceCommRecord) TableInfo() *TableInfo {
	return nil
}