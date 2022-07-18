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

func newAuthToken(db *gorm.DB) authToken {
	_authToken := authToken{}

	_authToken.authTokenDo.UseDB(db)
	_authToken.authTokenDo.UseModel(&model.AuthToken{})

	tableName := _authToken.authTokenDo.TableName()
	_authToken.ALL = field.NewField(tableName, "*")
	_authToken.ID = field.NewInt64(tableName, "id")
	_authToken.CreatedAt = field.NewTime(tableName, "created_at")
	_authToken.UpdatedAt = field.NewTime(tableName, "updated_at")
	_authToken.DeletedAt = field.NewField(tableName, "deleted_at")
	_authToken.Token = field.NewString(tableName, "token")
	_authToken.User = field.NewInt64(tableName, "user")
	_authToken.Expiry = field.NewTime(tableName, "expiry")
	_authToken.UploadOnly = field.NewBool(tableName, "upload_only")

	_authToken.fillFieldMap()

	return _authToken
}

type authToken struct {
	authTokenDo

	ALL        field.Field
	ID         field.Int64
	CreatedAt  field.Time
	UpdatedAt  field.Time
	DeletedAt  field.Field
	Token      field.String
	User       field.Int64
	Expiry     field.Time
	UploadOnly field.Bool

	fieldMap map[string]field.Expr
}

func (a authToken) Table(newTableName string) *authToken {
	a.authTokenDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a authToken) As(alias string) *authToken {
	a.authTokenDo.DO = *(a.authTokenDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *authToken) updateTableName(table string) *authToken {
	a.ALL = field.NewField(table, "*")
	a.ID = field.NewInt64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Token = field.NewString(table, "token")
	a.User = field.NewInt64(table, "user")
	a.Expiry = field.NewTime(table, "expiry")
	a.UploadOnly = field.NewBool(table, "upload_only")

	a.fillFieldMap()

	return a
}

func (a *authToken) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *authToken) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 8)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["token"] = a.Token
	a.fieldMap["user"] = a.User
	a.fieldMap["expiry"] = a.Expiry
	a.fieldMap["upload_only"] = a.UploadOnly
}

func (a authToken) clone(db *gorm.DB) authToken {
	a.authTokenDo.ReplaceDB(db)
	return a
}

type authTokenDo struct{ gen.DO }

type IAuthTokenDo interface {
	Debug() IAuthTokenDo
	WithContext(ctx context.Context) IAuthTokenDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAuthTokenDo
	Not(conds ...gen.Condition) IAuthTokenDo
	Or(conds ...gen.Condition) IAuthTokenDo
	Select(conds ...field.Expr) IAuthTokenDo
	Where(conds ...gen.Condition) IAuthTokenDo
	Order(conds ...field.Expr) IAuthTokenDo
	Distinct(cols ...field.Expr) IAuthTokenDo
	Omit(cols ...field.Expr) IAuthTokenDo
	Join(table schema.Tabler, on ...field.Expr) IAuthTokenDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAuthTokenDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAuthTokenDo
	Group(cols ...field.Expr) IAuthTokenDo
	Having(conds ...gen.Condition) IAuthTokenDo
	Limit(limit int) IAuthTokenDo
	Offset(offset int) IAuthTokenDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthTokenDo
	Unscoped() IAuthTokenDo
	Create(values ...*model.AuthToken) error
	CreateInBatches(values []*model.AuthToken, batchSize int) error
	Save(values ...*model.AuthToken) error
	First() (*model.AuthToken, error)
	Take() (*model.AuthToken, error)
	Last() (*model.AuthToken, error)
	Find() ([]*model.AuthToken, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AuthToken, err error)
	FindInBatches(result *[]*model.AuthToken, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAuthTokenDo
	Assign(attrs ...field.AssignExpr) IAuthTokenDo
	Joins(fields ...field.RelationField) IAuthTokenDo
	Preload(fields ...field.RelationField) IAuthTokenDo
	FirstOrInit() (*model.AuthToken, error)
	FirstOrCreate() (*model.AuthToken, error)
	FindByPage(offset int, limit int) (result []*model.AuthToken, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAuthTokenDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a authTokenDo) Debug() IAuthTokenDo {
	return a.withDO(a.DO.Debug())
}

func (a authTokenDo) WithContext(ctx context.Context) IAuthTokenDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a authTokenDo) ReadDB() IAuthTokenDo {
	return a.Clauses(dbresolver.Read)
}

func (a authTokenDo) WriteDB() IAuthTokenDo {
	return a.Clauses(dbresolver.Write)
}

func (a authTokenDo) Clauses(conds ...clause.Expression) IAuthTokenDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a authTokenDo) Returning(value interface{}, columns ...string) IAuthTokenDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a authTokenDo) Not(conds ...gen.Condition) IAuthTokenDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a authTokenDo) Or(conds ...gen.Condition) IAuthTokenDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a authTokenDo) Select(conds ...field.Expr) IAuthTokenDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a authTokenDo) Where(conds ...gen.Condition) IAuthTokenDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a authTokenDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAuthTokenDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a authTokenDo) Order(conds ...field.Expr) IAuthTokenDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a authTokenDo) Distinct(cols ...field.Expr) IAuthTokenDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a authTokenDo) Omit(cols ...field.Expr) IAuthTokenDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a authTokenDo) Join(table schema.Tabler, on ...field.Expr) IAuthTokenDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a authTokenDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAuthTokenDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a authTokenDo) RightJoin(table schema.Tabler, on ...field.Expr) IAuthTokenDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a authTokenDo) Group(cols ...field.Expr) IAuthTokenDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a authTokenDo) Having(conds ...gen.Condition) IAuthTokenDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a authTokenDo) Limit(limit int) IAuthTokenDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a authTokenDo) Offset(offset int) IAuthTokenDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a authTokenDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAuthTokenDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a authTokenDo) Unscoped() IAuthTokenDo {
	return a.withDO(a.DO.Unscoped())
}

func (a authTokenDo) Create(values ...*model.AuthToken) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a authTokenDo) CreateInBatches(values []*model.AuthToken, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a authTokenDo) Save(values ...*model.AuthToken) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a authTokenDo) First() (*model.AuthToken, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthToken), nil
	}
}

func (a authTokenDo) Take() (*model.AuthToken, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthToken), nil
	}
}

func (a authTokenDo) Last() (*model.AuthToken, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthToken), nil
	}
}

func (a authTokenDo) Find() ([]*model.AuthToken, error) {
	result, err := a.DO.Find()
	return result.([]*model.AuthToken), err
}

func (a authTokenDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.AuthToken, err error) {
	buf := make([]*model.AuthToken, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a authTokenDo) FindInBatches(result *[]*model.AuthToken, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a authTokenDo) Attrs(attrs ...field.AssignExpr) IAuthTokenDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a authTokenDo) Assign(attrs ...field.AssignExpr) IAuthTokenDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a authTokenDo) Joins(fields ...field.RelationField) IAuthTokenDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a authTokenDo) Preload(fields ...field.RelationField) IAuthTokenDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a authTokenDo) FirstOrInit() (*model.AuthToken, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthToken), nil
	}
}

func (a authTokenDo) FirstOrCreate() (*model.AuthToken, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.AuthToken), nil
	}
}

func (a authTokenDo) FindByPage(offset int, limit int) (result []*model.AuthToken, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a authTokenDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a authTokenDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a *authTokenDo) withDO(do gen.Dao) *authTokenDo {
	a.DO = *do.(*gen.DO)
	return a
}
