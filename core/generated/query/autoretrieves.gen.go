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

func newAutoretrieve(db *gorm.DB) autoretrieve {
	_autoretrieve := autoretrieve{}

	_autoretrieve.autoretrieveDo.UseDB(db)
	_autoretrieve.autoretrieveDo.UseModel(&model.Autoretrieve{})

	tableName := _autoretrieve.autoretrieveDo.TableName()
	_autoretrieve.ALL = field.NewField(tableName, "*")
	_autoretrieve.ID = field.NewInt64(tableName, "id")
	_autoretrieve.CreatedAt = field.NewTime(tableName, "created_at")
	_autoretrieve.UpdatedAt = field.NewTime(tableName, "updated_at")
	_autoretrieve.DeletedAt = field.NewField(tableName, "deleted_at")
	_autoretrieve.Handle = field.NewString(tableName, "handle")
	_autoretrieve.Token = field.NewString(tableName, "token")
	_autoretrieve.LastConnection = field.NewTime(tableName, "last_connection")
	_autoretrieve.PeerID = field.NewString(tableName, "peer_id")
	_autoretrieve.Addresses = field.NewString(tableName, "addresses")

	_autoretrieve.fillFieldMap()

	return _autoretrieve
}

type autoretrieve struct {
	autoretrieveDo

	ALL            field.Field
	ID             field.Int64
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	Handle         field.String
	Token          field.String
	LastConnection field.Time
	PeerID         field.String
	Addresses      field.String

	fieldMap map[string]field.Expr
}

func (a autoretrieve) Table(newTableName string) *autoretrieve {
	a.autoretrieveDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a autoretrieve) As(alias string) *autoretrieve {
	a.autoretrieveDo.DO = *(a.autoretrieveDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *autoretrieve) updateTableName(table string) *autoretrieve {
	a.ALL = field.NewField(table, "*")
	a.ID = field.NewInt64(table, "id")
	a.CreatedAt = field.NewTime(table, "created_at")
	a.UpdatedAt = field.NewTime(table, "updated_at")
	a.DeletedAt = field.NewField(table, "deleted_at")
	a.Handle = field.NewString(table, "handle")
	a.Token = field.NewString(table, "token")
	a.LastConnection = field.NewTime(table, "last_connection")
	a.PeerID = field.NewString(table, "peer_id")
	a.Addresses = field.NewString(table, "addresses")

	a.fillFieldMap()

	return a
}

func (a *autoretrieve) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *autoretrieve) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 9)
	a.fieldMap["id"] = a.ID
	a.fieldMap["created_at"] = a.CreatedAt
	a.fieldMap["updated_at"] = a.UpdatedAt
	a.fieldMap["deleted_at"] = a.DeletedAt
	a.fieldMap["handle"] = a.Handle
	a.fieldMap["token"] = a.Token
	a.fieldMap["last_connection"] = a.LastConnection
	a.fieldMap["peer_id"] = a.PeerID
	a.fieldMap["addresses"] = a.Addresses
}

func (a autoretrieve) clone(db *gorm.DB) autoretrieve {
	a.autoretrieveDo.ReplaceDB(db)
	return a
}

type autoretrieveDo struct{ gen.DO }

type IAutoretrieveDo interface {
	Debug() IAutoretrieveDo
	WithContext(ctx context.Context) IAutoretrieveDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAutoretrieveDo
	Not(conds ...gen.Condition) IAutoretrieveDo
	Or(conds ...gen.Condition) IAutoretrieveDo
	Select(conds ...field.Expr) IAutoretrieveDo
	Where(conds ...gen.Condition) IAutoretrieveDo
	Order(conds ...field.Expr) IAutoretrieveDo
	Distinct(cols ...field.Expr) IAutoretrieveDo
	Omit(cols ...field.Expr) IAutoretrieveDo
	Join(table schema.Tabler, on ...field.Expr) IAutoretrieveDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAutoretrieveDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAutoretrieveDo
	Group(cols ...field.Expr) IAutoretrieveDo
	Having(conds ...gen.Condition) IAutoretrieveDo
	Limit(limit int) IAutoretrieveDo
	Offset(offset int) IAutoretrieveDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAutoretrieveDo
	Unscoped() IAutoretrieveDo
	Create(values ...*model.Autoretrieve) error
	CreateInBatches(values []*model.Autoretrieve, batchSize int) error
	Save(values ...*model.Autoretrieve) error
	First() (*model.Autoretrieve, error)
	Take() (*model.Autoretrieve, error)
	Last() (*model.Autoretrieve, error)
	Find() ([]*model.Autoretrieve, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Autoretrieve, err error)
	FindInBatches(result *[]*model.Autoretrieve, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAutoretrieveDo
	Assign(attrs ...field.AssignExpr) IAutoretrieveDo
	Joins(fields ...field.RelationField) IAutoretrieveDo
	Preload(fields ...field.RelationField) IAutoretrieveDo
	FirstOrInit() (*model.Autoretrieve, error)
	FirstOrCreate() (*model.Autoretrieve, error)
	FindByPage(offset int, limit int) (result []*model.Autoretrieve, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAutoretrieveDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a autoretrieveDo) Debug() IAutoretrieveDo {
	return a.withDO(a.DO.Debug())
}

func (a autoretrieveDo) WithContext(ctx context.Context) IAutoretrieveDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a autoretrieveDo) ReadDB() IAutoretrieveDo {
	return a.Clauses(dbresolver.Read)
}

func (a autoretrieveDo) WriteDB() IAutoretrieveDo {
	return a.Clauses(dbresolver.Write)
}

func (a autoretrieveDo) Clauses(conds ...clause.Expression) IAutoretrieveDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a autoretrieveDo) Returning(value interface{}, columns ...string) IAutoretrieveDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a autoretrieveDo) Not(conds ...gen.Condition) IAutoretrieveDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a autoretrieveDo) Or(conds ...gen.Condition) IAutoretrieveDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a autoretrieveDo) Select(conds ...field.Expr) IAutoretrieveDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a autoretrieveDo) Where(conds ...gen.Condition) IAutoretrieveDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a autoretrieveDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IAutoretrieveDo {
	return a.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (a autoretrieveDo) Order(conds ...field.Expr) IAutoretrieveDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a autoretrieveDo) Distinct(cols ...field.Expr) IAutoretrieveDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a autoretrieveDo) Omit(cols ...field.Expr) IAutoretrieveDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a autoretrieveDo) Join(table schema.Tabler, on ...field.Expr) IAutoretrieveDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a autoretrieveDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAutoretrieveDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a autoretrieveDo) RightJoin(table schema.Tabler, on ...field.Expr) IAutoretrieveDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a autoretrieveDo) Group(cols ...field.Expr) IAutoretrieveDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a autoretrieveDo) Having(conds ...gen.Condition) IAutoretrieveDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a autoretrieveDo) Limit(limit int) IAutoretrieveDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a autoretrieveDo) Offset(offset int) IAutoretrieveDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a autoretrieveDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAutoretrieveDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a autoretrieveDo) Unscoped() IAutoretrieveDo {
	return a.withDO(a.DO.Unscoped())
}

func (a autoretrieveDo) Create(values ...*model.Autoretrieve) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a autoretrieveDo) CreateInBatches(values []*model.Autoretrieve, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a autoretrieveDo) Save(values ...*model.Autoretrieve) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a autoretrieveDo) First() (*model.Autoretrieve, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Autoretrieve), nil
	}
}

func (a autoretrieveDo) Take() (*model.Autoretrieve, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Autoretrieve), nil
	}
}

func (a autoretrieveDo) Last() (*model.Autoretrieve, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Autoretrieve), nil
	}
}

func (a autoretrieveDo) Find() ([]*model.Autoretrieve, error) {
	result, err := a.DO.Find()
	return result.([]*model.Autoretrieve), err
}

func (a autoretrieveDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Autoretrieve, err error) {
	buf := make([]*model.Autoretrieve, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a autoretrieveDo) FindInBatches(result *[]*model.Autoretrieve, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a autoretrieveDo) Attrs(attrs ...field.AssignExpr) IAutoretrieveDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a autoretrieveDo) Assign(attrs ...field.AssignExpr) IAutoretrieveDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a autoretrieveDo) Joins(fields ...field.RelationField) IAutoretrieveDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a autoretrieveDo) Preload(fields ...field.RelationField) IAutoretrieveDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a autoretrieveDo) FirstOrInit() (*model.Autoretrieve, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Autoretrieve), nil
	}
}

func (a autoretrieveDo) FirstOrCreate() (*model.Autoretrieve, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Autoretrieve), nil
	}
}

func (a autoretrieveDo) FindByPage(offset int, limit int) (result []*model.Autoretrieve, count int64, err error) {
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

func (a autoretrieveDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a autoretrieveDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a *autoretrieveDo) withDO(do gen.Dao) *autoretrieveDo {
	a.DO = *do.(*gen.DO)
	return a
}
