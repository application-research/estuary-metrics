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

func newInviteCode(db *gorm.DB) inviteCode {
	_inviteCode := inviteCode{}

	_inviteCode.inviteCodeDo.UseDB(db)
	_inviteCode.inviteCodeDo.UseModel(&model.InviteCode{})

	tableName := _inviteCode.inviteCodeDo.TableName()
	_inviteCode.ALL = field.NewField(tableName, "*")
	_inviteCode.ID = field.NewInt64(tableName, "id")
	_inviteCode.CreatedAt = field.NewTime(tableName, "created_at")
	_inviteCode.UpdatedAt = field.NewTime(tableName, "updated_at")
	_inviteCode.DeletedAt = field.NewField(tableName, "deleted_at")
	_inviteCode.Code = field.NewString(tableName, "code")
	_inviteCode.CreatedBy = field.NewInt64(tableName, "created_by")
	_inviteCode.ClaimedBy = field.NewInt64(tableName, "claimed_by")

	_inviteCode.fillFieldMap()

	return _inviteCode
}

type inviteCode struct {
	inviteCodeDo

	ALL       field.Field
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Code      field.String
	CreatedBy field.Int64
	ClaimedBy field.Int64

	fieldMap map[string]field.Expr
}

func (i inviteCode) Table(newTableName string) *inviteCode {
	i.inviteCodeDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i inviteCode) As(alias string) *inviteCode {
	i.inviteCodeDo.DO = *(i.inviteCodeDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *inviteCode) updateTableName(table string) *inviteCode {
	i.ALL = field.NewField(table, "*")
	i.ID = field.NewInt64(table, "id")
	i.CreatedAt = field.NewTime(table, "created_at")
	i.UpdatedAt = field.NewTime(table, "updated_at")
	i.DeletedAt = field.NewField(table, "deleted_at")
	i.Code = field.NewString(table, "code")
	i.CreatedBy = field.NewInt64(table, "created_by")
	i.ClaimedBy = field.NewInt64(table, "claimed_by")

	i.fillFieldMap()

	return i
}

func (i *inviteCode) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *inviteCode) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 7)
	i.fieldMap["id"] = i.ID
	i.fieldMap["created_at"] = i.CreatedAt
	i.fieldMap["updated_at"] = i.UpdatedAt
	i.fieldMap["deleted_at"] = i.DeletedAt
	i.fieldMap["code"] = i.Code
	i.fieldMap["created_by"] = i.CreatedBy
	i.fieldMap["claimed_by"] = i.ClaimedBy
}

func (i inviteCode) clone(db *gorm.DB) inviteCode {
	i.inviteCodeDo.ReplaceDB(db)
	return i
}

type inviteCodeDo struct{ gen.DO }

type IInviteCodeDo interface {
	Debug() IInviteCodeDo
	WithContext(ctx context.Context) IInviteCodeDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IInviteCodeDo
	Not(conds ...gen.Condition) IInviteCodeDo
	Or(conds ...gen.Condition) IInviteCodeDo
	Select(conds ...field.Expr) IInviteCodeDo
	Where(conds ...gen.Condition) IInviteCodeDo
	Order(conds ...field.Expr) IInviteCodeDo
	Distinct(cols ...field.Expr) IInviteCodeDo
	Omit(cols ...field.Expr) IInviteCodeDo
	Join(table schema.Tabler, on ...field.Expr) IInviteCodeDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IInviteCodeDo
	RightJoin(table schema.Tabler, on ...field.Expr) IInviteCodeDo
	Group(cols ...field.Expr) IInviteCodeDo
	Having(conds ...gen.Condition) IInviteCodeDo
	Limit(limit int) IInviteCodeDo
	Offset(offset int) IInviteCodeDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IInviteCodeDo
	Unscoped() IInviteCodeDo
	Create(values ...*model.InviteCode) error
	CreateInBatches(values []*model.InviteCode, batchSize int) error
	Save(values ...*model.InviteCode) error
	First() (*model.InviteCode, error)
	Take() (*model.InviteCode, error)
	Last() (*model.InviteCode, error)
	Find() ([]*model.InviteCode, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InviteCode, err error)
	FindInBatches(result *[]*model.InviteCode, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IInviteCodeDo
	Assign(attrs ...field.AssignExpr) IInviteCodeDo
	Joins(fields ...field.RelationField) IInviteCodeDo
	Preload(fields ...field.RelationField) IInviteCodeDo
	FirstOrInit() (*model.InviteCode, error)
	FirstOrCreate() (*model.InviteCode, error)
	FindByPage(offset int, limit int) (result []*model.InviteCode, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IInviteCodeDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i inviteCodeDo) Debug() IInviteCodeDo {
	return i.withDO(i.DO.Debug())
}

func (i inviteCodeDo) WithContext(ctx context.Context) IInviteCodeDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i inviteCodeDo) ReadDB() IInviteCodeDo {
	return i.Clauses(dbresolver.Read)
}

func (i inviteCodeDo) WriteDB() IInviteCodeDo {
	return i.Clauses(dbresolver.Write)
}

func (i inviteCodeDo) Clauses(conds ...clause.Expression) IInviteCodeDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i inviteCodeDo) Returning(value interface{}, columns ...string) IInviteCodeDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i inviteCodeDo) Not(conds ...gen.Condition) IInviteCodeDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i inviteCodeDo) Or(conds ...gen.Condition) IInviteCodeDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i inviteCodeDo) Select(conds ...field.Expr) IInviteCodeDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i inviteCodeDo) Where(conds ...gen.Condition) IInviteCodeDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i inviteCodeDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IInviteCodeDo {
	return i.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (i inviteCodeDo) Order(conds ...field.Expr) IInviteCodeDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i inviteCodeDo) Distinct(cols ...field.Expr) IInviteCodeDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i inviteCodeDo) Omit(cols ...field.Expr) IInviteCodeDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i inviteCodeDo) Join(table schema.Tabler, on ...field.Expr) IInviteCodeDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i inviteCodeDo) LeftJoin(table schema.Tabler, on ...field.Expr) IInviteCodeDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i inviteCodeDo) RightJoin(table schema.Tabler, on ...field.Expr) IInviteCodeDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i inviteCodeDo) Group(cols ...field.Expr) IInviteCodeDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i inviteCodeDo) Having(conds ...gen.Condition) IInviteCodeDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i inviteCodeDo) Limit(limit int) IInviteCodeDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i inviteCodeDo) Offset(offset int) IInviteCodeDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i inviteCodeDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IInviteCodeDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i inviteCodeDo) Unscoped() IInviteCodeDo {
	return i.withDO(i.DO.Unscoped())
}

func (i inviteCodeDo) Create(values ...*model.InviteCode) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i inviteCodeDo) CreateInBatches(values []*model.InviteCode, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i inviteCodeDo) Save(values ...*model.InviteCode) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i inviteCodeDo) First() (*model.InviteCode, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteCode), nil
	}
}

func (i inviteCodeDo) Take() (*model.InviteCode, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteCode), nil
	}
}

func (i inviteCodeDo) Last() (*model.InviteCode, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteCode), nil
	}
}

func (i inviteCodeDo) Find() ([]*model.InviteCode, error) {
	result, err := i.DO.Find()
	return result.([]*model.InviteCode), err
}

func (i inviteCodeDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.InviteCode, err error) {
	buf := make([]*model.InviteCode, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i inviteCodeDo) FindInBatches(result *[]*model.InviteCode, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i inviteCodeDo) Attrs(attrs ...field.AssignExpr) IInviteCodeDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i inviteCodeDo) Assign(attrs ...field.AssignExpr) IInviteCodeDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i inviteCodeDo) Joins(fields ...field.RelationField) IInviteCodeDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i inviteCodeDo) Preload(fields ...field.RelationField) IInviteCodeDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i inviteCodeDo) FirstOrInit() (*model.InviteCode, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteCode), nil
	}
}

func (i inviteCodeDo) FirstOrCreate() (*model.InviteCode, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.InviteCode), nil
	}
}

func (i inviteCodeDo) FindByPage(offset int, limit int) (result []*model.InviteCode, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i inviteCodeDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i inviteCodeDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i *inviteCodeDo) withDO(do gen.Dao) *inviteCodeDo {
	i.DO = *do.(*gen.DO)
	return i
}
