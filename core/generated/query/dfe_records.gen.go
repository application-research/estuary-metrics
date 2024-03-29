package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/application-research/estuary-metrics/core/generated/model"
)

func newDfeRecord(db *gorm.DB) dfeRecord {
	_dfeRecord := dfeRecord{}

	_dfeRecord.dfeRecordDo.UseDB(db)
	_dfeRecord.dfeRecordDo.UseModel(&model.DfeRecord{})

	tableName := _dfeRecord.dfeRecordDo.TableName()
	_dfeRecord.ALL = field.NewField(tableName, "*")
	_dfeRecord.ID = field.NewInt64(tableName, "id")
	_dfeRecord.CreatedAt = field.NewTime(tableName, "created_at")
	_dfeRecord.UpdatedAt = field.NewTime(tableName, "updated_at")
	_dfeRecord.DeletedAt = field.NewField(tableName, "deleted_at")
	_dfeRecord.Miner = field.NewString(tableName, "miner")
	_dfeRecord.Phase = field.NewString(tableName, "phase")
	_dfeRecord.Message = field.NewString(tableName, "message")
	_dfeRecord.Content = field.NewInt64(tableName, "content")
	_dfeRecord.MinerVersion = field.NewString(tableName, "miner_version")
	_dfeRecord.UserID = field.NewInt64(tableName, "user_id")

	_dfeRecord.fillFieldMap()

	return _dfeRecord
}

type dfeRecord struct {
	dfeRecordDo

	ALL          field.Field
	ID           field.Int64
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	Miner        field.String
	Phase        field.String
	Message      field.String
	Content      field.Int64
	MinerVersion field.String
	UserID       field.Int64

	fieldMap map[string]field.Expr
}

func (d dfeRecord) Table(newTableName string) *dfeRecord {
	d.dfeRecordDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dfeRecord) As(alias string) *dfeRecord {
	d.dfeRecordDo.DO = *(d.dfeRecordDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dfeRecord) updateTableName(table string) *dfeRecord {
	d.ALL = field.NewField(table, "*")
	d.ID = field.NewInt64(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.Miner = field.NewString(table, "miner")
	d.Phase = field.NewString(table, "phase")
	d.Message = field.NewString(table, "message")
	d.Content = field.NewInt64(table, "content")
	d.MinerVersion = field.NewString(table, "miner_version")
	d.UserID = field.NewInt64(table, "user_id")

	d.fillFieldMap()

	return d
}

func (d *dfeRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dfeRecord) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 10)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["miner"] = d.Miner
	d.fieldMap["phase"] = d.Phase
	d.fieldMap["message"] = d.Message
	d.fieldMap["content"] = d.Content
	d.fieldMap["miner_version"] = d.MinerVersion
	d.fieldMap["user_id"] = d.UserID
}

func (d dfeRecord) clone(db *gorm.DB) dfeRecord {
	d.dfeRecordDo.ReplaceDB(db)
	return d
}

type dfeRecordDo struct{ gen.DO }

type IDfeRecordDo interface {
	Debug() IDfeRecordDo
	WithContext(ctx context.Context) IDfeRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDfeRecordDo
	Not(conds ...gen.Condition) IDfeRecordDo
	Or(conds ...gen.Condition) IDfeRecordDo
	Select(conds ...field.Expr) IDfeRecordDo
	Where(conds ...gen.Condition) IDfeRecordDo
	Order(conds ...field.Expr) IDfeRecordDo
	Distinct(cols ...field.Expr) IDfeRecordDo
	Omit(cols ...field.Expr) IDfeRecordDo
	Join(table schema.Tabler, on ...field.Expr) IDfeRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDfeRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDfeRecordDo
	Group(cols ...field.Expr) IDfeRecordDo
	Having(conds ...gen.Condition) IDfeRecordDo
	Limit(limit int) IDfeRecordDo
	Offset(offset int) IDfeRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDfeRecordDo
	Unscoped() IDfeRecordDo
	Create(values ...*model.DfeRecord) error
	CreateInBatches(values []*model.DfeRecord, batchSize int) error
	Save(values ...*model.DfeRecord) error
	First() (*model.DfeRecord, error)
	Take() (*model.DfeRecord, error)
	Last() (*model.DfeRecord, error)
	Find() ([]*model.DfeRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DfeRecord, err error)
	FindInBatches(result *[]*model.DfeRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDfeRecordDo
	Assign(attrs ...field.AssignExpr) IDfeRecordDo
	Joins(fields ...field.RelationField) IDfeRecordDo
	Preload(fields ...field.RelationField) IDfeRecordDo
	FirstOrInit() (*model.DfeRecord, error)
	FirstOrCreate() (*model.DfeRecord, error)
	FindByPage(offset int, limit int) (result []*model.DfeRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDfeRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dfeRecordDo) Debug() IDfeRecordDo {
	return d.withDO(d.DO.Debug())
}

func (d dfeRecordDo) WithContext(ctx context.Context) IDfeRecordDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dfeRecordDo) ReadDB() IDfeRecordDo {
	return d.Clauses(dbresolver.Read)
}

func (d dfeRecordDo) WriteDB() IDfeRecordDo {
	return d.Clauses(dbresolver.Write)
}

func (d dfeRecordDo) Clauses(conds ...clause.Expression) IDfeRecordDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dfeRecordDo) Returning(value interface{}, columns ...string) IDfeRecordDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dfeRecordDo) Not(conds ...gen.Condition) IDfeRecordDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dfeRecordDo) Or(conds ...gen.Condition) IDfeRecordDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dfeRecordDo) Select(conds ...field.Expr) IDfeRecordDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dfeRecordDo) Where(conds ...gen.Condition) IDfeRecordDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dfeRecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDfeRecordDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dfeRecordDo) Order(conds ...field.Expr) IDfeRecordDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dfeRecordDo) Distinct(cols ...field.Expr) IDfeRecordDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dfeRecordDo) Omit(cols ...field.Expr) IDfeRecordDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dfeRecordDo) Join(table schema.Tabler, on ...field.Expr) IDfeRecordDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dfeRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDfeRecordDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dfeRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IDfeRecordDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dfeRecordDo) Group(cols ...field.Expr) IDfeRecordDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dfeRecordDo) Having(conds ...gen.Condition) IDfeRecordDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dfeRecordDo) Limit(limit int) IDfeRecordDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dfeRecordDo) Offset(offset int) IDfeRecordDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dfeRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDfeRecordDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dfeRecordDo) Unscoped() IDfeRecordDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dfeRecordDo) Create(values ...*model.DfeRecord) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dfeRecordDo) CreateInBatches(values []*model.DfeRecord, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dfeRecordDo) Save(values ...*model.DfeRecord) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dfeRecordDo) First() (*model.DfeRecord, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.DfeRecord), nil
	}
}

func (d dfeRecordDo) Take() (*model.DfeRecord, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.DfeRecord), nil
	}
}

func (d dfeRecordDo) Last() (*model.DfeRecord, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.DfeRecord), nil
	}
}

func (d dfeRecordDo) Find() ([]*model.DfeRecord, error) {
	result, err := d.DO.Find()
	return result.([]*model.DfeRecord), err
}

func (d dfeRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.DfeRecord, err error) {
	buf := make([]*model.DfeRecord, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dfeRecordDo) FindInBatches(result *[]*model.DfeRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dfeRecordDo) Attrs(attrs ...field.AssignExpr) IDfeRecordDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dfeRecordDo) Assign(attrs ...field.AssignExpr) IDfeRecordDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dfeRecordDo) Joins(fields ...field.RelationField) IDfeRecordDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dfeRecordDo) Preload(fields ...field.RelationField) IDfeRecordDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dfeRecordDo) FirstOrInit() (*model.DfeRecord, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.DfeRecord), nil
	}
}

func (d dfeRecordDo) FirstOrCreate() (*model.DfeRecord, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.DfeRecord), nil
	}
}

func (d dfeRecordDo) FindByPage(offset int, limit int) (result []*model.DfeRecord, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d dfeRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dfeRecordDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d *dfeRecordDo) withDO(do gen.Dao) *dfeRecordDo {
	d.DO = *do.(*gen.DO)
	return d
}
