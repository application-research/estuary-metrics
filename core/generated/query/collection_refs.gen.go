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

func newCollectionRef(db *gorm.DB) collectionRef {
	_collectionRef := collectionRef{}

	_collectionRef.collectionRefDo.UseDB(db)
	_collectionRef.collectionRefDo.UseModel(&model.CollectionRef{})

	tableName := _collectionRef.collectionRefDo.TableName()
	_collectionRef.ALL = field.NewField(tableName, "*")
	_collectionRef.ID = field.NewInt64(tableName, "id")
	_collectionRef.CreatedAt = field.NewTime(tableName, "created_at")
	_collectionRef.Collection = field.NewInt64(tableName, "collection")
	_collectionRef.Content = field.NewInt64(tableName, "content")
	_collectionRef.Path = field.NewString(tableName, "path")

	_collectionRef.fillFieldMap()

	return _collectionRef
}

type collectionRef struct {
	collectionRefDo

	ALL        field.Field
	ID         field.Int64
	CreatedAt  field.Time
	Collection field.Int64
	Content    field.Int64
	Path       field.String

	fieldMap map[string]field.Expr
}

func (c collectionRef) Table(newTableName string) *collectionRef {
	c.collectionRefDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c collectionRef) As(alias string) *collectionRef {
	c.collectionRefDo.DO = *(c.collectionRefDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *collectionRef) updateTableName(table string) *collectionRef {
	c.ALL = field.NewField(table, "*")
	c.ID = field.NewInt64(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.Collection = field.NewInt64(table, "collection")
	c.Content = field.NewInt64(table, "content")
	c.Path = field.NewString(table, "path")

	c.fillFieldMap()

	return c
}

func (c *collectionRef) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *collectionRef) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["collection"] = c.Collection
	c.fieldMap["content"] = c.Content
	c.fieldMap["path"] = c.Path
}

func (c collectionRef) clone(db *gorm.DB) collectionRef {
	c.collectionRefDo.ReplaceDB(db)
	return c
}

type collectionRefDo struct{ gen.DO }

type ICollectionRefDo interface {
	Debug() ICollectionRefDo
	WithContext(ctx context.Context) ICollectionRefDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICollectionRefDo
	Not(conds ...gen.Condition) ICollectionRefDo
	Or(conds ...gen.Condition) ICollectionRefDo
	Select(conds ...field.Expr) ICollectionRefDo
	Where(conds ...gen.Condition) ICollectionRefDo
	Order(conds ...field.Expr) ICollectionRefDo
	Distinct(cols ...field.Expr) ICollectionRefDo
	Omit(cols ...field.Expr) ICollectionRefDo
	Join(table schema.Tabler, on ...field.Expr) ICollectionRefDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICollectionRefDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICollectionRefDo
	Group(cols ...field.Expr) ICollectionRefDo
	Having(conds ...gen.Condition) ICollectionRefDo
	Limit(limit int) ICollectionRefDo
	Offset(offset int) ICollectionRefDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectionRefDo
	Unscoped() ICollectionRefDo
	Create(values ...*model.CollectionRef) error
	CreateInBatches(values []*model.CollectionRef, batchSize int) error
	Save(values ...*model.CollectionRef) error
	First() (*model.CollectionRef, error)
	Take() (*model.CollectionRef, error)
	Last() (*model.CollectionRef, error)
	Find() ([]*model.CollectionRef, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CollectionRef, err error)
	FindInBatches(result *[]*model.CollectionRef, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICollectionRefDo
	Assign(attrs ...field.AssignExpr) ICollectionRefDo
	Joins(fields ...field.RelationField) ICollectionRefDo
	Preload(fields ...field.RelationField) ICollectionRefDo
	FirstOrInit() (*model.CollectionRef, error)
	FirstOrCreate() (*model.CollectionRef, error)
	FindByPage(offset int, limit int) (result []*model.CollectionRef, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICollectionRefDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c collectionRefDo) Debug() ICollectionRefDo {
	return c.withDO(c.DO.Debug())
}

func (c collectionRefDo) WithContext(ctx context.Context) ICollectionRefDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c collectionRefDo) ReadDB() ICollectionRefDo {
	return c.Clauses(dbresolver.Read)
}

func (c collectionRefDo) WriteDB() ICollectionRefDo {
	return c.Clauses(dbresolver.Write)
}

func (c collectionRefDo) Clauses(conds ...clause.Expression) ICollectionRefDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c collectionRefDo) Returning(value interface{}, columns ...string) ICollectionRefDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c collectionRefDo) Not(conds ...gen.Condition) ICollectionRefDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c collectionRefDo) Or(conds ...gen.Condition) ICollectionRefDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c collectionRefDo) Select(conds ...field.Expr) ICollectionRefDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c collectionRefDo) Where(conds ...gen.Condition) ICollectionRefDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c collectionRefDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ICollectionRefDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c collectionRefDo) Order(conds ...field.Expr) ICollectionRefDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c collectionRefDo) Distinct(cols ...field.Expr) ICollectionRefDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c collectionRefDo) Omit(cols ...field.Expr) ICollectionRefDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c collectionRefDo) Join(table schema.Tabler, on ...field.Expr) ICollectionRefDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c collectionRefDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICollectionRefDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c collectionRefDo) RightJoin(table schema.Tabler, on ...field.Expr) ICollectionRefDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c collectionRefDo) Group(cols ...field.Expr) ICollectionRefDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c collectionRefDo) Having(conds ...gen.Condition) ICollectionRefDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c collectionRefDo) Limit(limit int) ICollectionRefDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c collectionRefDo) Offset(offset int) ICollectionRefDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c collectionRefDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICollectionRefDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c collectionRefDo) Unscoped() ICollectionRefDo {
	return c.withDO(c.DO.Unscoped())
}

func (c collectionRefDo) Create(values ...*model.CollectionRef) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c collectionRefDo) CreateInBatches(values []*model.CollectionRef, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c collectionRefDo) Save(values ...*model.CollectionRef) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c collectionRefDo) First() (*model.CollectionRef, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionRef), nil
	}
}

func (c collectionRefDo) Take() (*model.CollectionRef, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionRef), nil
	}
}

func (c collectionRefDo) Last() (*model.CollectionRef, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionRef), nil
	}
}

func (c collectionRefDo) Find() ([]*model.CollectionRef, error) {
	result, err := c.DO.Find()
	return result.([]*model.CollectionRef), err
}

func (c collectionRefDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CollectionRef, err error) {
	buf := make([]*model.CollectionRef, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c collectionRefDo) FindInBatches(result *[]*model.CollectionRef, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c collectionRefDo) Attrs(attrs ...field.AssignExpr) ICollectionRefDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c collectionRefDo) Assign(attrs ...field.AssignExpr) ICollectionRefDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c collectionRefDo) Joins(fields ...field.RelationField) ICollectionRefDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c collectionRefDo) Preload(fields ...field.RelationField) ICollectionRefDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c collectionRefDo) FirstOrInit() (*model.CollectionRef, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionRef), nil
	}
}

func (c collectionRefDo) FirstOrCreate() (*model.CollectionRef, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CollectionRef), nil
	}
}

func (c collectionRefDo) FindByPage(offset int, limit int) (result []*model.CollectionRef, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c collectionRefDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c collectionRefDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c *collectionRefDo) withDO(do gen.Dao) *collectionRefDo {
	c.DO = *do.(*gen.DO)
	return c
}
