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

func newContentDeal(db *gorm.DB) contentDeal {
	_contentDeal := contentDeal{}

	_contentDeal.contentDealDo.UseDB(db)
	_contentDeal.contentDealDo.UseModel(&model.ContentDeal{})

	tableName := _contentDeal.contentDealDo.TableName()
	_contentDeal.ALL = field.NewField(tableName, "*")
	_contentDeal.ID = field.NewInt64(tableName, "id")
	_contentDeal.CreatedAt = field.NewTime(tableName, "created_at")
	_contentDeal.UpdatedAt = field.NewTime(tableName, "updated_at")
	_contentDeal.DeletedAt = field.NewField(tableName, "deleted_at")
	_contentDeal.Content = field.NewInt64(tableName, "content")
	_contentDeal.PropCid = field.NewBytes(tableName, "prop_cid")
	_contentDeal.Miner = field.NewString(tableName, "miner")
	_contentDeal.DealID = field.NewInt64(tableName, "deal_id")
	_contentDeal.Failed = field.NewBool(tableName, "failed")
	_contentDeal.FailedAt = field.NewTime(tableName, "failed_at")
	_contentDeal.DtChan = field.NewString(tableName, "dt_chan")
	_contentDeal.Verified = field.NewBool(tableName, "verified")
	_contentDeal.SealedAt = field.NewTime(tableName, "sealed_at")
	_contentDeal.OnChainAt = field.NewTime(tableName, "on_chain_at")
	_contentDeal.TransferStarted = field.NewTime(tableName, "transfer_started")
	_contentDeal.TransferFinished = field.NewTime(tableName, "transfer_finished")
	_contentDeal.DealUUID = field.NewString(tableName, "deal_uuid")
	_contentDeal.UserID = field.NewInt64(tableName, "user_id")
	_contentDeal.Slashed = field.NewBool(tableName, "slashed")

	_contentDeal.fillFieldMap()

	return _contentDeal
}

type contentDeal struct {
	contentDealDo

	ALL              field.Field
	ID               field.Int64
	CreatedAt        field.Time
	UpdatedAt        field.Time
	DeletedAt        field.Field
	Content          field.Int64
	PropCid          field.Bytes
	Miner            field.String
	DealID           field.Int64
	Failed           field.Bool
	FailedAt         field.Time
	DtChan           field.String
	Verified         field.Bool
	SealedAt         field.Time
	OnChainAt        field.Time
	TransferStarted  field.Time
	TransferFinished field.Time
	DealUUID         field.String
	UserID           field.Int64
	Slashed          field.Bool

	fieldMap map[string]field.Expr
}

func (c contentDeal) Table(newTableName string) *contentDeal {
	c.contentDealDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c contentDeal) As(alias string) *contentDeal {
	c.contentDealDo.DO = *(c.contentDealDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *contentDeal) updateTableName(table string) *contentDeal {
	c.ALL = field.NewField(table, "*")
	c.ID = field.NewInt64(table, "id")
	c.CreatedAt = field.NewTime(table, "created_at")
	c.UpdatedAt = field.NewTime(table, "updated_at")
	c.DeletedAt = field.NewField(table, "deleted_at")
	c.Content = field.NewInt64(table, "content")
	c.PropCid = field.NewBytes(table, "prop_cid")
	c.Miner = field.NewString(table, "miner")
	c.DealID = field.NewInt64(table, "deal_id")
	c.Failed = field.NewBool(table, "failed")
	c.FailedAt = field.NewTime(table, "failed_at")
	c.DtChan = field.NewString(table, "dt_chan")
	c.Verified = field.NewBool(table, "verified")
	c.SealedAt = field.NewTime(table, "sealed_at")
	c.OnChainAt = field.NewTime(table, "on_chain_at")
	c.TransferStarted = field.NewTime(table, "transfer_started")
	c.TransferFinished = field.NewTime(table, "transfer_finished")
	c.DealUUID = field.NewString(table, "deal_uuid")
	c.UserID = field.NewInt64(table, "user_id")
	c.Slashed = field.NewBool(table, "slashed")

	c.fillFieldMap()

	return c
}

func (c *contentDeal) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *contentDeal) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 19)
	c.fieldMap["id"] = c.ID
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
	c.fieldMap["deleted_at"] = c.DeletedAt
	c.fieldMap["content"] = c.Content
	c.fieldMap["prop_cid"] = c.PropCid
	c.fieldMap["miner"] = c.Miner
	c.fieldMap["deal_id"] = c.DealID
	c.fieldMap["failed"] = c.Failed
	c.fieldMap["failed_at"] = c.FailedAt
	c.fieldMap["dt_chan"] = c.DtChan
	c.fieldMap["verified"] = c.Verified
	c.fieldMap["sealed_at"] = c.SealedAt
	c.fieldMap["on_chain_at"] = c.OnChainAt
	c.fieldMap["transfer_started"] = c.TransferStarted
	c.fieldMap["transfer_finished"] = c.TransferFinished
	c.fieldMap["deal_uuid"] = c.DealUUID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["slashed"] = c.Slashed
}

func (c contentDeal) clone(db *gorm.DB) contentDeal {
	c.contentDealDo.ReplaceDB(db)
	return c
}

type contentDealDo struct{ gen.DO }

type IContentDealDo interface {
	Debug() IContentDealDo
	WithContext(ctx context.Context) IContentDealDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IContentDealDo
	Not(conds ...gen.Condition) IContentDealDo
	Or(conds ...gen.Condition) IContentDealDo
	Select(conds ...field.Expr) IContentDealDo
	Where(conds ...gen.Condition) IContentDealDo
	Order(conds ...field.Expr) IContentDealDo
	Distinct(cols ...field.Expr) IContentDealDo
	Omit(cols ...field.Expr) IContentDealDo
	Join(table schema.Tabler, on ...field.Expr) IContentDealDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IContentDealDo
	RightJoin(table schema.Tabler, on ...field.Expr) IContentDealDo
	Group(cols ...field.Expr) IContentDealDo
	Having(conds ...gen.Condition) IContentDealDo
	Limit(limit int) IContentDealDo
	Offset(offset int) IContentDealDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IContentDealDo
	Unscoped() IContentDealDo
	Create(values ...*model.ContentDeal) error
	CreateInBatches(values []*model.ContentDeal, batchSize int) error
	Save(values ...*model.ContentDeal) error
	First() (*model.ContentDeal, error)
	Take() (*model.ContentDeal, error)
	Last() (*model.ContentDeal, error)
	Find() ([]*model.ContentDeal, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ContentDeal, err error)
	FindInBatches(result *[]*model.ContentDeal, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IContentDealDo
	Assign(attrs ...field.AssignExpr) IContentDealDo
	Joins(fields ...field.RelationField) IContentDealDo
	Preload(fields ...field.RelationField) IContentDealDo
	FirstOrInit() (*model.ContentDeal, error)
	FirstOrCreate() (*model.ContentDeal, error)
	FindByPage(offset int, limit int) (result []*model.ContentDeal, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IContentDealDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c contentDealDo) Debug() IContentDealDo {
	return c.withDO(c.DO.Debug())
}

func (c contentDealDo) WithContext(ctx context.Context) IContentDealDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c contentDealDo) ReadDB() IContentDealDo {
	return c.Clauses(dbresolver.Read)
}

func (c contentDealDo) WriteDB() IContentDealDo {
	return c.Clauses(dbresolver.Write)
}

func (c contentDealDo) Clauses(conds ...clause.Expression) IContentDealDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c contentDealDo) Returning(value interface{}, columns ...string) IContentDealDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c contentDealDo) Not(conds ...gen.Condition) IContentDealDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c contentDealDo) Or(conds ...gen.Condition) IContentDealDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c contentDealDo) Select(conds ...field.Expr) IContentDealDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c contentDealDo) Where(conds ...gen.Condition) IContentDealDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c contentDealDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IContentDealDo {
	return c.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (c contentDealDo) Order(conds ...field.Expr) IContentDealDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c contentDealDo) Distinct(cols ...field.Expr) IContentDealDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c contentDealDo) Omit(cols ...field.Expr) IContentDealDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c contentDealDo) Join(table schema.Tabler, on ...field.Expr) IContentDealDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c contentDealDo) LeftJoin(table schema.Tabler, on ...field.Expr) IContentDealDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c contentDealDo) RightJoin(table schema.Tabler, on ...field.Expr) IContentDealDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c contentDealDo) Group(cols ...field.Expr) IContentDealDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c contentDealDo) Having(conds ...gen.Condition) IContentDealDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c contentDealDo) Limit(limit int) IContentDealDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c contentDealDo) Offset(offset int) IContentDealDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c contentDealDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IContentDealDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c contentDealDo) Unscoped() IContentDealDo {
	return c.withDO(c.DO.Unscoped())
}

func (c contentDealDo) Create(values ...*model.ContentDeal) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c contentDealDo) CreateInBatches(values []*model.ContentDeal, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c contentDealDo) Save(values ...*model.ContentDeal) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c contentDealDo) First() (*model.ContentDeal, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentDeal), nil
	}
}

func (c contentDealDo) Take() (*model.ContentDeal, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentDeal), nil
	}
}

func (c contentDealDo) Last() (*model.ContentDeal, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentDeal), nil
	}
}

func (c contentDealDo) Find() ([]*model.ContentDeal, error) {
	result, err := c.DO.Find()
	return result.([]*model.ContentDeal), err
}

func (c contentDealDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ContentDeal, err error) {
	buf := make([]*model.ContentDeal, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c contentDealDo) FindInBatches(result *[]*model.ContentDeal, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c contentDealDo) Attrs(attrs ...field.AssignExpr) IContentDealDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c contentDealDo) Assign(attrs ...field.AssignExpr) IContentDealDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c contentDealDo) Joins(fields ...field.RelationField) IContentDealDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c contentDealDo) Preload(fields ...field.RelationField) IContentDealDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c contentDealDo) FirstOrInit() (*model.ContentDeal, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentDeal), nil
	}
}

func (c contentDealDo) FirstOrCreate() (*model.ContentDeal, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ContentDeal), nil
	}
}

func (c contentDealDo) FindByPage(offset int, limit int) (result []*model.ContentDeal, count int64, err error) {
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

func (c contentDealDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c contentDealDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c *contentDealDo) withDO(do gen.Dao) *contentDealDo {
	c.DO = *do.(*gen.DO)
	return c
}
