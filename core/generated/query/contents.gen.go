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

func newContent(db *gorm.DB) content {
	_content := content{}

	_content.contentDo.UseDB(db)
	_content.contentDo.UseModel(&model.Content{})

	tableName := _content.contentDo.TableName()
	_content.ALL = field.NewField(tableName, "*")
	_content.ID = field.NewInt64(tableName, "id")
	_content.CreatedAt = field.NewTime(tableName, "created_at")
	_content.UpdatedAt = field.NewTime(tableName, "updated_at")
	_content.DeletedAt = field.NewField(tableName, "deleted_at")
	_content.Cid = field.NewBytes(tableName, "cid")
	_content.Name = field.NewString(tableName, "name")
	_content.User = field.NewString(tableName, "user")
	_content.Size = field.NewInt64(tableName, "size")
	_content.Active = field.NewBool(tableName, "active")
	_content.Description = field.NewString(tableName, "description")
	_content.Offloaded = field.NewBool(tableName, "offloaded")
	_content.UserID = field.NewInt64(tableName, "user_id")
	_content.Replication = field.NewInt64(tableName, "replication")
	_content.AggregatedIn = field.NewInt64(tableName, "aggregated_in")
	_content.Aggregate = field.NewBool(tableName, "aggregate")
	_content.Pinning = field.NewBool(tableName, "pinning")
	_content.PinMeta = field.NewString(tableName, "pin_meta")
	_content.Location = field.NewString(tableName, "location")
	_content.Failed = field.NewBool(tableName, "failed")
	_content.DagSplit = field.NewBool(tableName, "dag_split")
	_content.SplitFrom = field.NewInt64(tableName, "split_from")
	_content.Type = field.NewInt64(tableName, "type")
	_content.Replace = field.NewBool(tableName, "replace")
	_content.Origins = field.NewString(tableName, "origins")

	_content.fillFieldMap()

	return _content
}

type content struct {
	contentDo

	ALL          field.Field
	ID           field.Int64
	CreatedAt    field.Time
	UpdatedAt    field.Time
	DeletedAt    field.Field
	Cid          field.Bytes
	Name         field.String
	User         field.String
	Size         field.Int64
	Active       field.Bool
	Description  field.String
	Offloaded    field.Bool
	UserID       field.Int64
	Replication  field.Int64
	AggregatedIn field.Int64
	Aggregate    field.Bool
	Pinning      field.Bool
	PinMeta      field.String
	Location     field.String
	Failed       field.Bool
	DagSplit     field.Bool
	SplitFrom    field.Int64
	Type         field.Int64
	Replace      field.Bool
	Origins      field.String

	fieldMap map[string]field.Expr
}

func (c content) Table(newTableName string) *content {
	c.contentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c content) As(alias string) *content {
	c.contentDo.DO = *(c.contentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *content) updateTableName(table string) *content {
	c.ALL = field.NewField(table, "*")
	c.ID = field.NewInt64(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Cid = field.NewBytes(table, "cid")
	c.Name = field.NewString(table, "name")
	c.User = field.NewString(table, "user")
	c.Size = field.NewInt64(table, "size")
	c.Active = field.NewBool(table, "active")
	c.Description = field.NewString(table, "description")
	c.Offloaded = field.NewBool(table, "offloaded")
	c.UserID = field.NewInt64(table, "user_id")
	c.Replication = field.NewInt64(table, "replication")
	c.AggregatedIn = field.NewInt64(table, "aggregated_in")
	c.Aggregate = field.NewBool(table, "aggregate")
	c.Pinning = field.NewBool(table, "pinning")
	c.PinMeta = field.NewString(table, "pin_meta")
	c.Location = field.NewString(table, "location")
	c.Failed = field.NewBool(table, "failed")
	c.DagSplit = field.NewBool(table, "dag_split")
	c.SplitFrom = field.NewInt64(table, "split_from")
	c.Type = field.NewInt64(table, "type")
	c.Replace = field.NewBool(table, "replace")
	c.Origins = field.NewString(table, "origins")

	c.fillFieldMap()

	return c
}

func (c *content) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *content) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 24)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["cid"] = c.Cid
	c.fieldMap["name"] = c.Name
	c.fieldMap["user"] = c.User
	c.fieldMap["size"] = c.Size
	c.fieldMap["active"] = c.Active
	c.fieldMap["description"] = c.Description
	c.fieldMap["offloaded"] = c.Offloaded
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["replication"] = c.Replication
	c.fieldMap["aggregated_in"] = c.AggregatedIn
	c.fieldMap["aggregate"] = c.Aggregate
	c.fieldMap["pinning"] = c.Pinning
	c.fieldMap["pin_meta"] = c.PinMeta
	c.fieldMap["location"] = c.Location
	c.fieldMap["failed"] = c.Failed
	c.fieldMap["dag_split"] = c.DagSplit
	c.fieldMap["split_from"] = c.SplitFrom
	c.fieldMap["type"] = c.Type
	c.fieldMap["replace"] = c.Replace
	c.fieldMap["origins"] = c.Origins
}

func (c content) clone(db *gorm.DB) content {
	c.contentDo.ReplaceDB(db)
	return c
}

type contentDo struct{ gen.DO }

type IContentDo interface {
	Debug() IContentDo
	WithContext(ctx context.Context) IContentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IContentDo
	Not(conds ...gen.Condition) IContentDo
	Or(conds ...gen.Condition) IContentDo
	Select(conds ...field.Expr) IContentDo
	Where(conds ...gen.Condition) IContentDo
	Order(conds ...field.Expr) IContentDo
	Distinct(cols ...field.Expr) IContentDo
	Omit(cols ...field.Expr) IContentDo
	Join(table schema.Tabler, on ...field.Expr) IContentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IContentDo
	RightJoin(table schema.Tabler, on ...field.Expr) IContentDo
	Group(cols ...field.Expr) IContentDo
	Having(conds ...gen.Condition) IContentDo
	Limit(limit int) IContentDo
	Offset(offset int) IContentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IContentDo
	Unscoped() IContentDo
	Create(values ...*model.Content) error
	CreateInBatches(values []*model.Content, batchSize int) error
	Save(values ...*model.Content) error
	First() (*model.Content, error)
	Take() (*model.Content, error)
	Last() (*model.Content, error)
	Find() ([]*model.Content, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Content, err error)
	FindInBatches(result *[]*model.Content, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IContentDo
	Assign(attrs ...field.AssignExpr) IContentDo
	Joins(fields ...field.RelationField) IContentDo
	Preload(fields ...field.RelationField) IContentDo
	FirstOrInit() (*model.Content, error)
	FirstOrCreate() (*model.Content, error)
	FindByPage(offset int, limit int) (result []*model.Content, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IContentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c contentDo) Debug() IContentDo {
	return c.withDO(c.DO.Debug())
}

func (c contentDo) WithContext(ctx context.Context) IContentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c contentDo) ReadDB() IContentDo {
	return c.Clauses(dbresolver.Read)
}

func (c contentDo) WriteDB() IContentDo {
	return c.Clauses(dbresolver.Write)
}

func (c contentDo) Clauses(conds ...clause.Expression) IContentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c contentDo) Returning(value interface{}, columns ...string) IContentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c contentDo) Not(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c contentDo) Or(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c contentDo) Select(conds ...field.Expr) IContentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c contentDo) Where(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c contentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IContentDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c contentDo) Order(conds ...field.Expr) IContentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c contentDo) Distinct(cols ...field.Expr) IContentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c contentDo) Omit(cols ...field.Expr) IContentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c contentDo) Join(table schema.Tabler, on ...field.Expr) IContentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c contentDo) LeftJoin(table schema.Tabler, on ...field.Expr) IContentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c contentDo) RightJoin(table schema.Tabler, on ...field.Expr) IContentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c contentDo) Group(cols ...field.Expr) IContentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c contentDo) Having(conds ...gen.Condition) IContentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c contentDo) Limit(limit int) IContentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c contentDo) Offset(offset int) IContentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c contentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IContentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c contentDo) Unscoped() IContentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c contentDo) Create(values ...*model.Content) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c contentDo) CreateInBatches(values []*model.Content, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c contentDo) Save(values ...*model.Content) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c contentDo) First() (*model.Content, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) Take() (*model.Content, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) Last() (*model.Content, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) Find() ([]*model.Content, error) {
	result, err := c.DO.Find()
	return result.([]*model.Content), err
}

func (c contentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Content, err error) {
	buf := make([]*model.Content, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c contentDo) FindInBatches(result *[]*model.Content, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c contentDo) Attrs(attrs ...field.AssignExpr) IContentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c contentDo) Assign(attrs ...field.AssignExpr) IContentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c contentDo) Joins(fields ...field.RelationField) IContentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c contentDo) Preload(fields ...field.RelationField) IContentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c contentDo) FirstOrInit() (*model.Content, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) FirstOrCreate() (*model.Content, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Content), nil
	}
}

func (c contentDo) FindByPage(offset int, limit int) (result []*model.Content, count int64, err error) {
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

func (c contentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c contentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c *contentDo) withDO(do gen.Dao) *contentDo {
	c.DO = *do.(*gen.DO)
	return c
}
