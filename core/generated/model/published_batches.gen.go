// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNamePublishedBatch = "published_batches"

// PublishedBatch mapped from table <published_batches>
type PublishedBatch struct {
	ID                 int64          `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CreatedAt          time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt          time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt          gorm.DeletedAt `gorm:"column:deleted_at" json:"deleted_at"`
	FirstContentID     int64          `gorm:"column:first_content_id" json:"first_content_id"`
	Count              int64          `gorm:"column:count" json:"count"`
	AutoretrieveHandle string         `gorm:"column:autoretrieve_handle" json:"autoretrieve_handle"`
}

// TableName PublishedBatch's table name
func (*PublishedBatch) TableName() string {
	return TableNamePublishedBatch
}
