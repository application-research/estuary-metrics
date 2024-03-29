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

func newDealer(db *gorm.DB) dealer {
	_dealer := dealer{}

	_dealer.dealerDo.UseDB(db)
	_dealer.dealerDo.UseModel(&model.Dealer{})

	tableName := _dealer.dealerDo.TableName()
	_dealer.ALL = field.NewField(tableName, "*")
	_dealer.ID = field.NewInt64(tableName, "id")
	_dealer.CreatedAt = field.NewTime(tableName, "created_at")
	_dealer.UpdatedAt = field.NewTime(tableName, "updated_at")
	_dealer.DeletedAt = field.NewField(tableName, "deleted_at")
	_dealer.Handle = field.NewString(tableName, "handle")
	_dealer.Token = field.NewString(tableName, "token")
	_dealer.Host = field.NewString(tableName, "host")
	_dealer.PeerID = field.NewString(tableName, "peer_id")
	_dealer.Open = field.NewBool(tableName, "open")
	_dealer.LastConnection = field.NewTime(tableName, "last_connection")

	_dealer.fillFieldMap()

	return _dealer
}

type dealer struct {
	dealerDo

	ALL            field.Field
	ID             field.Int64
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	Handle         field.String
	Token          field.String
	Host           field.String
	PeerID         field.String
	Open           field.Bool
	LastConnection field.Time

	fieldMap map[string]field.Expr
}

func (d dealer) Table(newTableName string) *dealer {
	d.dealerDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d dealer) As(alias string) *dealer {
	d.dealerDo.DO = *(d.dealerDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *dealer) updateTableName(table string) *dealer {
	d.ALL = field.NewField(table, "*")
	d.ID = field.NewInt64(table, "id")
	d.CreatedAt = field.NewTime(table, "created_at")
	d.UpdatedAt = field.NewTime(table, "updated_at")
	d.DeletedAt = field.NewField(table, "deleted_at")
	d.Handle = field.NewString(table, "handle")
	d.Token = field.NewString(table, "token")
	d.Host = field.NewString(table, "host")
	d.PeerID = field.NewString(table, "peer_id")
	d.Open = field.NewBool(table, "open")
	d.LastConnection = field.NewTime(table, "last_connection")

	d.fillFieldMap()

	return d
}

func (d *dealer) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *dealer) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 10)
	d.fieldMap["id"] = d.ID
	d.fieldMap["created_at"] = d.CreatedAt
	d.fieldMap["updated_at"] = d.UpdatedAt
	d.fieldMap["deleted_at"] = d.DeletedAt
	d.fieldMap["handle"] = d.Handle
	d.fieldMap["token"] = d.Token
	d.fieldMap["host"] = d.Host
	d.fieldMap["peer_id"] = d.PeerID
	d.fieldMap["open"] = d.Open
	d.fieldMap["last_connection"] = d.LastConnection
}

func (d dealer) clone(db *gorm.DB) dealer {
	d.dealerDo.ReplaceDB(db)
	return d
}

type dealerDo struct{ gen.DO }

type IDealerDo interface {
	Debug() IDealerDo
	WithContext(ctx context.Context) IDealerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDealerDo
	Not(conds ...gen.Condition) IDealerDo
	Or(conds ...gen.Condition) IDealerDo
	Select(conds ...field.Expr) IDealerDo
	Where(conds ...gen.Condition) IDealerDo
	Order(conds ...field.Expr) IDealerDo
	Distinct(cols ...field.Expr) IDealerDo
	Omit(cols ...field.Expr) IDealerDo
	Join(table schema.Tabler, on ...field.Expr) IDealerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDealerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDealerDo
	Group(cols ...field.Expr) IDealerDo
	Having(conds ...gen.Condition) IDealerDo
	Limit(limit int) IDealerDo
	Offset(offset int) IDealerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDealerDo
	Unscoped() IDealerDo
	Create(values ...*model.Dealer) error
	CreateInBatches(values []*model.Dealer, batchSize int) error
	Save(values ...*model.Dealer) error
	First() (*model.Dealer, error)
	Take() (*model.Dealer, error)
	Last() (*model.Dealer, error)
	Find() ([]*model.Dealer, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Dealer, err error)
	FindInBatches(result *[]*model.Dealer, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDealerDo
	Assign(attrs ...field.AssignExpr) IDealerDo
	Joins(fields ...field.RelationField) IDealerDo
	Preload(fields ...field.RelationField) IDealerDo
	FirstOrInit() (*model.Dealer, error)
	FirstOrCreate() (*model.Dealer, error)
	FindByPage(offset int, limit int) (result []*model.Dealer, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDealerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d dealerDo) Debug() IDealerDo {
	return d.withDO(d.DO.Debug())
}

func (d dealerDo) WithContext(ctx context.Context) IDealerDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d dealerDo) ReadDB() IDealerDo {
	return d.Clauses(dbresolver.Read)
}

func (d dealerDo) WriteDB() IDealerDo {
	return d.Clauses(dbresolver.Write)
}

func (d dealerDo) Clauses(conds ...clause.Expression) IDealerDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d dealerDo) Returning(value interface{}, columns ...string) IDealerDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d dealerDo) Not(conds ...gen.Condition) IDealerDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d dealerDo) Or(conds ...gen.Condition) IDealerDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d dealerDo) Select(conds ...field.Expr) IDealerDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d dealerDo) Where(conds ...gen.Condition) IDealerDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d dealerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IDealerDo {
	return d.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (d dealerDo) Order(conds ...field.Expr) IDealerDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d dealerDo) Distinct(cols ...field.Expr) IDealerDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d dealerDo) Omit(cols ...field.Expr) IDealerDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d dealerDo) Join(table schema.Tabler, on ...field.Expr) IDealerDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d dealerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDealerDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d dealerDo) RightJoin(table schema.Tabler, on ...field.Expr) IDealerDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d dealerDo) Group(cols ...field.Expr) IDealerDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d dealerDo) Having(conds ...gen.Condition) IDealerDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d dealerDo) Limit(limit int) IDealerDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d dealerDo) Offset(offset int) IDealerDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d dealerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDealerDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d dealerDo) Unscoped() IDealerDo {
	return d.withDO(d.DO.Unscoped())
}

func (d dealerDo) Create(values ...*model.Dealer) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d dealerDo) CreateInBatches(values []*model.Dealer, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d dealerDo) Save(values ...*model.Dealer) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d dealerDo) First() (*model.Dealer, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealer), nil
	}
}

func (d dealerDo) Take() (*model.Dealer, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealer), nil
	}
}

func (d dealerDo) Last() (*model.Dealer, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealer), nil
	}
}

func (d dealerDo) Find() ([]*model.Dealer, error) {
	result, err := d.DO.Find()
	return result.([]*model.Dealer), err
}

func (d dealerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Dealer, err error) {
	buf := make([]*model.Dealer, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d dealerDo) FindInBatches(result *[]*model.Dealer, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d dealerDo) Attrs(attrs ...field.AssignExpr) IDealerDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d dealerDo) Assign(attrs ...field.AssignExpr) IDealerDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d dealerDo) Joins(fields ...field.RelationField) IDealerDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d dealerDo) Preload(fields ...field.RelationField) IDealerDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d dealerDo) FirstOrInit() (*model.Dealer, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealer), nil
	}
}

func (d dealerDo) FirstOrCreate() (*model.Dealer, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Dealer), nil
	}
}

func (d dealerDo) FindByPage(offset int, limit int) (result []*model.Dealer, count int64, err error) {
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

func (d dealerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d dealerDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d *dealerDo) withDO(do gen.Dao) *dealerDo {
	d.DO = *do.(*gen.DO)
	return d
}
