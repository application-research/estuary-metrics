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

func newObjRef(db *gorm.DB) objRef {
	_objRef := objRef{}

	_objRef.objRefDo.UseDB(db)
	_objRef.objRefDo.UseModel(&model.ObjRef{})

	tableName := _objRef.objRefDo.TableName()
	_objRef.ALL = field.NewField(tableName, "*")
	_objRef.ID = field.NewInt64(tableName, "id")
	_objRef.Content = field.NewInt64(tableName, "content")
	_objRef.Object = field.NewInt64(tableName, "object")
	_objRef.Offloaded = field.NewInt64(tableName, "offloaded")

	_objRef.fillFieldMap()

	return _objRef
}

type objRef struct {
	objRefDo

	ALL       field.Field
	ID        field.Int64
	Content   field.Int64
	Object    field.Int64
	Offloaded field.Int64

	fieldMap map[string]field.Expr
}

func (o objRef) Table(newTableName string) *objRef {
	o.objRefDo.UseTable(newTableName)
	return o.updateTableName(newTableName)
}

func (o objRef) As(alias string) *objRef {
	o.objRefDo.DO = *(o.objRefDo.As(alias).(*gen.DO))
	return o.updateTableName(alias)
}

func (o *objRef) updateTableName(table string) *objRef {
	o.ALL = field.NewField(table, "*")
	o.ID = field.NewInt64(table, "id")
	o.Content = field.NewInt64(table, "content")
	o.Object = field.NewInt64(table, "object")
	o.Offloaded = field.NewInt64(table, "offloaded")

	o.fillFieldMap()

	return o
}

func (o *objRef) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := o.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (o *objRef) fillFieldMap() {
	o.fieldMap = make(map[string]field.Expr, 4)
	o.fieldMap["id"] = o.ID
	o.fieldMap["content"] = o.Content
	o.fieldMap["object"] = o.Object
	o.fieldMap["offloaded"] = o.Offloaded
}

func (o objRef) clone(db *gorm.DB) objRef {
	o.objRefDo.ReplaceDB(db)
	return o
}

type objRefDo struct{ gen.DO }

type IObjRefDo interface {
	Debug() IObjRefDo
	WithContext(ctx context.Context) IObjRefDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IObjRefDo
	Not(conds ...gen.Condition) IObjRefDo
	Or(conds ...gen.Condition) IObjRefDo
	Select(conds ...field.Expr) IObjRefDo
	Where(conds ...gen.Condition) IObjRefDo
	Order(conds ...field.Expr) IObjRefDo
	Distinct(cols ...field.Expr) IObjRefDo
	Omit(cols ...field.Expr) IObjRefDo
	Join(table schema.Tabler, on ...field.Expr) IObjRefDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IObjRefDo
	RightJoin(table schema.Tabler, on ...field.Expr) IObjRefDo
	Group(cols ...field.Expr) IObjRefDo
	Having(conds ...gen.Condition) IObjRefDo
	Limit(limit int) IObjRefDo
	Offset(offset int) IObjRefDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IObjRefDo
	Unscoped() IObjRefDo
	Create(values ...*model.ObjRef) error
	CreateInBatches(values []*model.ObjRef, batchSize int) error
	Save(values ...*model.ObjRef) error
	First() (*model.ObjRef, error)
	Take() (*model.ObjRef, error)
	Last() (*model.ObjRef, error)
	Find() ([]*model.ObjRef, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ObjRef, err error)
	FindInBatches(result *[]*model.ObjRef, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IObjRefDo
	Assign(attrs ...field.AssignExpr) IObjRefDo
	Joins(fields ...field.RelationField) IObjRefDo
	Preload(fields ...field.RelationField) IObjRefDo
	FirstOrInit() (*model.ObjRef, error)
	FirstOrCreate() (*model.ObjRef, error)
	FindByPage(offset int, limit int) (result []*model.ObjRef, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IObjRefDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (o objRefDo) Debug() IObjRefDo {
	return o.withDO(o.DO.Debug())
}

func (o objRefDo) WithContext(ctx context.Context) IObjRefDo {
	return o.withDO(o.DO.WithContext(ctx))
}

func (o objRefDo) ReadDB() IObjRefDo {
	return o.Clauses(dbresolver.Read)
}

func (o objRefDo) WriteDB() IObjRefDo {
	return o.Clauses(dbresolver.Write)
}

func (o objRefDo) Clauses(conds ...clause.Expression) IObjRefDo {
	return o.withDO(o.DO.Clauses(conds...))
}

func (o objRefDo) Returning(value interface{}, columns ...string) IObjRefDo {
	return o.withDO(o.DO.Returning(value, columns...))
}

func (o objRefDo) Not(conds ...gen.Condition) IObjRefDo {
	return o.withDO(o.DO.Not(conds...))
}

func (o objRefDo) Or(conds ...gen.Condition) IObjRefDo {
	return o.withDO(o.DO.Or(conds...))
}

func (o objRefDo) Select(conds ...field.Expr) IObjRefDo {
	return o.withDO(o.DO.Select(conds...))
}

func (o objRefDo) Where(conds ...gen.Condition) IObjRefDo {
	return o.withDO(o.DO.Where(conds...))
}

func (o objRefDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IObjRefDo {
	return o.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (o objRefDo) Order(conds ...field.Expr) IObjRefDo {
	return o.withDO(o.DO.Order(conds...))
}

func (o objRefDo) Distinct(cols ...field.Expr) IObjRefDo {
	return o.withDO(o.DO.Distinct(cols...))
}

func (o objRefDo) Omit(cols ...field.Expr) IObjRefDo {
	return o.withDO(o.DO.Omit(cols...))
}

func (o objRefDo) Join(table schema.Tabler, on ...field.Expr) IObjRefDo {
	return o.withDO(o.DO.Join(table, on...))
}

func (o objRefDo) LeftJoin(table schema.Tabler, on ...field.Expr) IObjRefDo {
	return o.withDO(o.DO.LeftJoin(table, on...))
}

func (o objRefDo) RightJoin(table schema.Tabler, on ...field.Expr) IObjRefDo {
	return o.withDO(o.DO.RightJoin(table, on...))
}

func (o objRefDo) Group(cols ...field.Expr) IObjRefDo {
	return o.withDO(o.DO.Group(cols...))
}

func (o objRefDo) Having(conds ...gen.Condition) IObjRefDo {
	return o.withDO(o.DO.Having(conds...))
}

func (o objRefDo) Limit(limit int) IObjRefDo {
	return o.withDO(o.DO.Limit(limit))
}

func (o objRefDo) Offset(offset int) IObjRefDo {
	return o.withDO(o.DO.Offset(offset))
}

func (o objRefDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IObjRefDo {
	return o.withDO(o.DO.Scopes(funcs...))
}

func (o objRefDo) Unscoped() IObjRefDo {
	return o.withDO(o.DO.Unscoped())
}

func (o objRefDo) Create(values ...*model.ObjRef) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Create(values)
}

func (o objRefDo) CreateInBatches(values []*model.ObjRef, batchSize int) error {
	return o.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (o objRefDo) Save(values ...*model.ObjRef) error {
	if len(values) == 0 {
		return nil
	}
	return o.DO.Save(values)
}

func (o objRefDo) First() (*model.ObjRef, error) {
	if result, err := o.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ObjRef), nil
	}
}

func (o objRefDo) Take() (*model.ObjRef, error) {
	if result, err := o.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ObjRef), nil
	}
}

func (o objRefDo) Last() (*model.ObjRef, error) {
	if result, err := o.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ObjRef), nil
	}
}

func (o objRefDo) Find() ([]*model.ObjRef, error) {
	result, err := o.DO.Find()
	return result.([]*model.ObjRef), err
}

func (o objRefDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ObjRef, err error) {
	buf := make([]*model.ObjRef, 0, batchSize)
	err = o.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (o objRefDo) FindInBatches(result *[]*model.ObjRef, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return o.DO.FindInBatches(result, batchSize, fc)
}

func (o objRefDo) Attrs(attrs ...field.AssignExpr) IObjRefDo {
	return o.withDO(o.DO.Attrs(attrs...))
}

func (o objRefDo) Assign(attrs ...field.AssignExpr) IObjRefDo {
	return o.withDO(o.DO.Assign(attrs...))
}

func (o objRefDo) Joins(fields ...field.RelationField) IObjRefDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Joins(_f))
	}
	return &o
}

func (o objRefDo) Preload(fields ...field.RelationField) IObjRefDo {
	for _, _f := range fields {
		o = *o.withDO(o.DO.Preload(_f))
	}
	return &o
}

func (o objRefDo) FirstOrInit() (*model.ObjRef, error) {
	if result, err := o.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ObjRef), nil
	}
}

func (o objRefDo) FirstOrCreate() (*model.ObjRef, error) {
	if result, err := o.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ObjRef), nil
	}
}

func (o objRefDo) FindByPage(offset int, limit int) (result []*model.ObjRef, count int64, err error) {
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

func (o objRefDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = o.Count()
	if err != nil {
		return
	}

	err = o.Offset(offset).Limit(limit).Scan(result)
	return
}

func (o objRefDo) Scan(result interface{}) (err error) {
	return o.DO.Scan(result)
}

func (o *objRefDo) withDO(do gen.Dao) *objRefDo {
	o.DO = *do.(*gen.DO)
	return o
}
