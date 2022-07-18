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

func newMinerStorageAsk(db *gorm.DB) minerStorageAsk {
	_minerStorageAsk := minerStorageAsk{}

	_minerStorageAsk.minerStorageAskDo.UseDB(db)
	_minerStorageAsk.minerStorageAskDo.UseModel(&model.MinerStorageAsk{})

	tableName := _minerStorageAsk.minerStorageAskDo.TableName()
	_minerStorageAsk.ALL = field.NewField(tableName, "*")
	_minerStorageAsk.ID = field.NewInt64(tableName, "id")
	_minerStorageAsk.CreatedAt = field.NewTime(tableName, "created_at")
	_minerStorageAsk.UpdatedAt = field.NewTime(tableName, "updated_at")
	_minerStorageAsk.DeletedAt = field.NewField(tableName, "deleted_at")
	_minerStorageAsk.Miner = field.NewString(tableName, "miner")
	_minerStorageAsk.Price = field.NewString(tableName, "price")
	_minerStorageAsk.VerifiedPrice = field.NewString(tableName, "verified_price")
	_minerStorageAsk.MinPieceSize = field.NewInt64(tableName, "min_piece_size")
	_minerStorageAsk.MaxPieceSize = field.NewInt64(tableName, "max_piece_size")

	_minerStorageAsk.fillFieldMap()

	return _minerStorageAsk
}

type minerStorageAsk struct {
	minerStorageAskDo

	ALL           field.Field
	ID            field.Int64
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	Miner         field.String
	Price         field.String
	VerifiedPrice field.String
	MinPieceSize  field.Int64
	MaxPieceSize  field.Int64

	fieldMap map[string]field.Expr
}

func (m minerStorageAsk) Table(newTableName string) *minerStorageAsk {
	m.minerStorageAskDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m minerStorageAsk) As(alias string) *minerStorageAsk {
	m.minerStorageAskDo.DO = *(m.minerStorageAskDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *minerStorageAsk) updateTableName(table string) *minerStorageAsk {
	m.ALL = field.NewField(table, "*")
	m.ID = field.NewInt64(table, "id")
	m.CreatedAt = field.NewTime(table, "created_at")
	m.UpdatedAt = field.NewTime(table, "updated_at")
	m.DeletedAt = field.NewField(table, "deleted_at")
	m.Miner = field.NewString(table, "miner")
	m.Price = field.NewString(table, "price")
	m.VerifiedPrice = field.NewString(table, "verified_price")
	m.MinPieceSize = field.NewInt64(table, "min_piece_size")
	m.MaxPieceSize = field.NewInt64(table, "max_piece_size")

	m.fillFieldMap()

	return m
}

func (m *minerStorageAsk) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *minerStorageAsk) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 9)
	m.fieldMap["id"] = m.ID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
	m.fieldMap["deleted_at"] = m.DeletedAt
	m.fieldMap["miner"] = m.Miner
	m.fieldMap["price"] = m.Price
	m.fieldMap["verified_price"] = m.VerifiedPrice
	m.fieldMap["min_piece_size"] = m.MinPieceSize
	m.fieldMap["max_piece_size"] = m.MaxPieceSize
}

func (m minerStorageAsk) clone(db *gorm.DB) minerStorageAsk {
	m.minerStorageAskDo.ReplaceDB(db)
	return m
}

type minerStorageAskDo struct{ gen.DO }

type IMinerStorageAskDo interface {
	Debug() IMinerStorageAskDo
	WithContext(ctx context.Context) IMinerStorageAskDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMinerStorageAskDo
	Not(conds ...gen.Condition) IMinerStorageAskDo
	Or(conds ...gen.Condition) IMinerStorageAskDo
	Select(conds ...field.Expr) IMinerStorageAskDo
	Where(conds ...gen.Condition) IMinerStorageAskDo
	Order(conds ...field.Expr) IMinerStorageAskDo
	Distinct(cols ...field.Expr) IMinerStorageAskDo
	Omit(cols ...field.Expr) IMinerStorageAskDo
	Join(table schema.Tabler, on ...field.Expr) IMinerStorageAskDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMinerStorageAskDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMinerStorageAskDo
	Group(cols ...field.Expr) IMinerStorageAskDo
	Having(conds ...gen.Condition) IMinerStorageAskDo
	Limit(limit int) IMinerStorageAskDo
	Offset(offset int) IMinerStorageAskDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMinerStorageAskDo
	Unscoped() IMinerStorageAskDo
	Create(values ...*model.MinerStorageAsk) error
	CreateInBatches(values []*model.MinerStorageAsk, batchSize int) error
	Save(values ...*model.MinerStorageAsk) error
	First() (*model.MinerStorageAsk, error)
	Take() (*model.MinerStorageAsk, error)
	Last() (*model.MinerStorageAsk, error)
	Find() ([]*model.MinerStorageAsk, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MinerStorageAsk, err error)
	FindInBatches(result *[]*model.MinerStorageAsk, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMinerStorageAskDo
	Assign(attrs ...field.AssignExpr) IMinerStorageAskDo
	Joins(fields ...field.RelationField) IMinerStorageAskDo
	Preload(fields ...field.RelationField) IMinerStorageAskDo
	FirstOrInit() (*model.MinerStorageAsk, error)
	FirstOrCreate() (*model.MinerStorageAsk, error)
	FindByPage(offset int, limit int) (result []*model.MinerStorageAsk, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMinerStorageAskDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m minerStorageAskDo) Debug() IMinerStorageAskDo {
	return m.withDO(m.DO.Debug())
}

func (m minerStorageAskDo) WithContext(ctx context.Context) IMinerStorageAskDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m minerStorageAskDo) ReadDB() IMinerStorageAskDo {
	return m.Clauses(dbresolver.Read)
}

func (m minerStorageAskDo) WriteDB() IMinerStorageAskDo {
	return m.Clauses(dbresolver.Write)
}

func (m minerStorageAskDo) Clauses(conds ...clause.Expression) IMinerStorageAskDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m minerStorageAskDo) Returning(value interface{}, columns ...string) IMinerStorageAskDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m minerStorageAskDo) Not(conds ...gen.Condition) IMinerStorageAskDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m minerStorageAskDo) Or(conds ...gen.Condition) IMinerStorageAskDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m minerStorageAskDo) Select(conds ...field.Expr) IMinerStorageAskDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m minerStorageAskDo) Where(conds ...gen.Condition) IMinerStorageAskDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m minerStorageAskDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IMinerStorageAskDo {
	return m.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (m minerStorageAskDo) Order(conds ...field.Expr) IMinerStorageAskDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m minerStorageAskDo) Distinct(cols ...field.Expr) IMinerStorageAskDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m minerStorageAskDo) Omit(cols ...field.Expr) IMinerStorageAskDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m minerStorageAskDo) Join(table schema.Tabler, on ...field.Expr) IMinerStorageAskDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m minerStorageAskDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMinerStorageAskDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m minerStorageAskDo) RightJoin(table schema.Tabler, on ...field.Expr) IMinerStorageAskDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m minerStorageAskDo) Group(cols ...field.Expr) IMinerStorageAskDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m minerStorageAskDo) Having(conds ...gen.Condition) IMinerStorageAskDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m minerStorageAskDo) Limit(limit int) IMinerStorageAskDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m minerStorageAskDo) Offset(offset int) IMinerStorageAskDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m minerStorageAskDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMinerStorageAskDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m minerStorageAskDo) Unscoped() IMinerStorageAskDo {
	return m.withDO(m.DO.Unscoped())
}

func (m minerStorageAskDo) Create(values ...*model.MinerStorageAsk) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m minerStorageAskDo) CreateInBatches(values []*model.MinerStorageAsk, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m minerStorageAskDo) Save(values ...*model.MinerStorageAsk) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m minerStorageAskDo) First() (*model.MinerStorageAsk, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinerStorageAsk), nil
	}
}

func (m minerStorageAskDo) Take() (*model.MinerStorageAsk, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinerStorageAsk), nil
	}
}

func (m minerStorageAskDo) Last() (*model.MinerStorageAsk, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinerStorageAsk), nil
	}
}

func (m minerStorageAskDo) Find() ([]*model.MinerStorageAsk, error) {
	result, err := m.DO.Find()
	return result.([]*model.MinerStorageAsk), err
}

func (m minerStorageAskDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MinerStorageAsk, err error) {
	buf := make([]*model.MinerStorageAsk, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m minerStorageAskDo) FindInBatches(result *[]*model.MinerStorageAsk, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m minerStorageAskDo) Attrs(attrs ...field.AssignExpr) IMinerStorageAskDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m minerStorageAskDo) Assign(attrs ...field.AssignExpr) IMinerStorageAskDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m minerStorageAskDo) Joins(fields ...field.RelationField) IMinerStorageAskDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m minerStorageAskDo) Preload(fields ...field.RelationField) IMinerStorageAskDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m minerStorageAskDo) FirstOrInit() (*model.MinerStorageAsk, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinerStorageAsk), nil
	}
}

func (m minerStorageAskDo) FirstOrCreate() (*model.MinerStorageAsk, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MinerStorageAsk), nil
	}
}

func (m minerStorageAskDo) FindByPage(offset int, limit int) (result []*model.MinerStorageAsk, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m minerStorageAskDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m minerStorageAskDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m *minerStorageAskDo) withDO(do gen.Dao) *minerStorageAskDo {
	m.DO = *do.(*gen.DO)
	return m
}
