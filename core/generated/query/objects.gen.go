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

func newObject(db *gorm.DB) object {
	_object := object{}

	_object.objectDo.UseDB(db)
	_object.objectDo.UseModel(&model.Object{})

	tableName := _object.objectDo.TableName()
	_object.ALL = field.NewField(tableName, "*")
	_object.ID = field.NewInt64(tableName, "id")
	_object.Cid = field.NewBytes(tableName, "cid")
	_object.Size = field.NewInt64(tableName, "size")
	_object.Reads = field.NewInt64(tableName, "reads")
	_object.LastAccess = field.NewTime(tableName, "last_access")

	_object.fillFieldMap()

	return _object
}

type object struct {
	objectDo

	ALL        field.Field
	ID         field.Int64
	Cid        field.Bytes
	Size       field.Int64
	Reads      field.Int64
	LastAccess field.Time

	fieldMap map[string]field.Expr
}

func (o object) Table(newTableName string) *object {
	o.objectDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o object) As(alias string) *object {
	o.objectDo.DO = *(o.objectDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *object) updateTableName(table string) *object {
	o.ALL = field.NewField(table, "*")
	o.ID = field.NewInt64(table, "id")
	o.Cid = field.NewBytes(table, "cid")
	o.Size = field.NewInt64(table, "size")
	o.Reads = field.NewInt64(table, "reads")
	o.LastAccess = field.NewTime(table, "last_access")

	o.fillFieldMap()

	return o
}

func (o *object) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *object) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 5)
	o.fieldMap["id"] = o.ID
	o.fieldMap["cid"] = o.Cid
	o.fieldMap["size"] = o.Size
	o.fieldMap["reads"] = o.Reads
	o.fieldMap["last_access"] = o.LastAccess
}

func (o object) clone(db *gorm.DB) object {
	o.objectDo.ReplaceDB(db)
	return o
}

type objectDo struct{ gen.DO }

type IObjectDo interface {
	Debug() IObjectDo
	WithContext(ctx context.Context) IObjectDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IObjectDo
	Not(conds ...gen.Condition) IObjectDo
	Or(conds ...gen.Condition) IObjectDo
	Select(conds ...field.Expr) IObjectDo
	Where(conds ...gen.Condition) IObjectDo
	Order(conds ...field.Expr) IObjectDo
	Distinct(cols ...field.Expr) IObjectDo
	Omit(cols ...field.Expr) IObjectDo
	Join(table schema.Tabler, on ...field.Expr) IObjectDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IObjectDo
	RightJoin(table schema.Tabler, on ...field.Expr) IObjectDo
	Group(cols ...field.Expr) IObjectDo
	Having(conds ...gen.Condition) IObjectDo
	Limit(limit int) IObjectDo
	Offset(offset int) IObjectDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IObjectDo
	Unscoped() IObjectDo
	Create(values ...*model.Object) error
	CreateInBatches(values []*model.Object, batchSize int) error
	Save(values ...*model.Object) error
	First() (*model.Object, error)
	Take() (*model.Object, error)
	Last() (*model.Object, error)
	Find() ([]*model.Object, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Object, err error)
	FindInBatches(result *[]*model.Object, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IObjectDo
	Assign(attrs ...field.AssignExpr) IObjectDo
	Joins(fields ...field.RelationField) IObjectDo
	Preload(fields ...field.RelationField) IObjectDo
	FirstOrInit() (*model.Object, error)
	FirstOrCreate() (*model.Object, error)
	FindByPage(offset int, limit int) (result []*model.Object, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IObjectDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o objectDo) Debug() IObjectDo {
	return o.withDO(o.DO.Debug())
}

func (o objectDo) WithContext(ctx context.Context) IObjectDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o objectDo) ReadDB() IObjectDo {
	return o.Clauses(dbresolver.Read)
}

func (o objectDo) WriteDB() IObjectDo {
	return o.Clauses(dbresolver.Write)
}

func (o objectDo) Clauses(conds ...clause.Expression) IObjectDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o objectDo) Returning(value interface{}, columns ...string) IObjectDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o objectDo) Not(conds ...gen.Condition) IObjectDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o objectDo) Or(conds ...gen.Condition) IObjectDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o objectDo) Select(conds ...field.Expr) IObjectDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o objectDo) Where(conds ...gen.Condition) IObjectDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o objectDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IObjectDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o objectDo) Order(conds ...field.Expr) IObjectDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o objectDo) Distinct(cols ...field.Expr) IObjectDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o objectDo) Omit(cols ...field.Expr) IObjectDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o objectDo) Join(table schema.Tabler, on ...field.Expr) IObjectDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o objectDo) LeftJoin(table schema.Tabler, on ...field.Expr) IObjectDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o objectDo) RightJoin(table schema.Tabler, on ...field.Expr) IObjectDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o objectDo) Group(cols ...field.Expr) IObjectDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o objectDo) Having(conds ...gen.Condition) IObjectDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o objectDo) Limit(limit int) IObjectDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o objectDo) Offset(offset int) IObjectDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o objectDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IObjectDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o objectDo) Unscoped() IObjectDo {
	return o.withDO(o.DO.Unscoped())
}

func (o objectDo) Create(values ...*model.Object) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o objectDo) CreateInBatches(values []*model.Object, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o objectDo) Save(values ...*model.Object) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o objectDo) First() (*model.Object, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Object), nil
	}
}

func (o objectDo) Take() (*model.Object, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Object), nil
	}
}

func (o objectDo) Last() (*model.Object, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Object), nil
	}
}

func (o objectDo) Find() ([]*model.Object, error) {
	result, err := o.DO.Find()
	return result.([]*model.Object), err
}

func (o objectDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Object, err error) {
	buf := make([]*model.Object, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o objectDo) FindInBatches(result *[]*model.Object, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o objectDo) Attrs(attrs ...field.AssignExpr) IObjectDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o objectDo) Assign(attrs ...field.AssignExpr) IObjectDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o objectDo) Joins(fields ...field.RelationField) IObjectDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o objectDo) Preload(fields ...field.RelationField) IObjectDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o objectDo) FirstOrInit() (*model.Object, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Object), nil
	}
}

func (o objectDo) FirstOrCreate() (*model.Object, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Object), nil
	}
}

func (o objectDo) FindByPage(offset int, limit int) (result []*model.Object, count int64, err error) {
	result, err = o.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = o.Offset(-1).Limit(-1).Count()
	return
}

func (o objectDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o objectDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o *objectDo) withDO(do gen.Dao) *objectDo {
	o.DO = *do.(*gen.DO)
	return o
}
