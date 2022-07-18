package model

import "gorm.io/gorm"

const TableNameObjRef = "obj_refs"

// ObjRef mapped from table <obj_refs>
type ObjRef struct {
	ID        int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Content   int64 `gorm:"column:content" json:"content"`
	Object    int64 `gorm:"column:object" json:"object"`
	Offloaded int64 `gorm:"column:offloaded" json:"offloaded"`
}

// TableName ObjRef's table name
func (*ObjRef) TableName() string {
	return TableNameObjRef
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ObjRef) BeforeSave(db *gorm.DB) error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ObjRef) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ObjRef) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ObjRef) TableInfo() *TableInfo {
	return nil
}
