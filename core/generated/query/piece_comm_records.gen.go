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

func newPieceCommRecord(db *gorm.DB) pieceCommRecord {
	_pieceCommRecord := pieceCommRecord{}

	_pieceCommRecord.pieceCommRecordDo.UseDB(db)
	_pieceCommRecord.pieceCommRecordDo.UseModel(&model.PieceCommRecord{})

	tableName := _pieceCommRecord.pieceCommRecordDo.TableName()
	_pieceCommRecord.ALL = field.NewField(tableName, "*")
	_pieceCommRecord.Data = field.NewBytes(tableName, "data")
	_pieceCommRecord.Piece = field.NewBytes(tableName, "piece")
	_pieceCommRecord.Size = field.NewInt64(tableName, "size")
	_pieceCommRecord.CarSize = field.NewInt64(tableName, "car_size")

	_pieceCommRecord.fillFieldMap()

	return _pieceCommRecord
}

type pieceCommRecord struct {
	pieceCommRecordDo

	ALL     field.Field
	Data    field.Bytes
	Piece   field.Bytes
	Size    field.Int64
	CarSize field.Int64

	fieldMap map[string]field.Expr
}

func (p pieceCommRecord) Table(newTableName string) *pieceCommRecord {
	p.pieceCommRecordDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p pieceCommRecord) As(alias string) *pieceCommRecord {
	p.pieceCommRecordDo.DO = *(p.pieceCommRecordDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *pieceCommRecord) updateTableName(table string) *pieceCommRecord {
	p.ALL = field.NewField(table, "*")
	p.Data = field.NewBytes(table, "data")
	p.Piece = field.NewBytes(table, "piece")
	p.Size = field.NewInt64(table, "size")
	p.CarSize = field.NewInt64(table, "car_size")

	p.fillFieldMap()

	return p
}

func (p *pieceCommRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *pieceCommRecord) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 4)
	p.fieldMap["data"] = p.Data
	p.fieldMap["piece"] = p.Piece
	p.fieldMap["size"] = p.Size
	p.fieldMap["car_size"] = p.CarSize
}

func (p pieceCommRecord) clone(db *gorm.DB) pieceCommRecord {
	p.pieceCommRecordDo.ReplaceDB(db)
	return p
}

type pieceCommRecordDo struct{ gen.DO }

type IPieceCommRecordDo interface {
	Debug() IPieceCommRecordDo
	WithContext(ctx context.Context) IPieceCommRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPieceCommRecordDo
	Not(conds ...gen.Condition) IPieceCommRecordDo
	Or(conds ...gen.Condition) IPieceCommRecordDo
	Select(conds ...field.Expr) IPieceCommRecordDo
	Where(conds ...gen.Condition) IPieceCommRecordDo
	Order(conds ...field.Expr) IPieceCommRecordDo
	Distinct(cols ...field.Expr) IPieceCommRecordDo
	Omit(cols ...field.Expr) IPieceCommRecordDo
	Join(table schema.Tabler, on ...field.Expr) IPieceCommRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPieceCommRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPieceCommRecordDo
	Group(cols ...field.Expr) IPieceCommRecordDo
	Having(conds ...gen.Condition) IPieceCommRecordDo
	Limit(limit int) IPieceCommRecordDo
	Offset(offset int) IPieceCommRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPieceCommRecordDo
	Unscoped() IPieceCommRecordDo
	Create(values ...*model.PieceCommRecord) error
	CreateInBatches(values []*model.PieceCommRecord, batchSize int) error
	Save(values ...*model.PieceCommRecord) error
	First() (*model.PieceCommRecord, error)
	Take() (*model.PieceCommRecord, error)
	Last() (*model.PieceCommRecord, error)
	Find() ([]*model.PieceCommRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PieceCommRecord, err error)
	FindInBatches(result *[]*model.PieceCommRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPieceCommRecordDo
	Assign(attrs ...field.AssignExpr) IPieceCommRecordDo
	Joins(fields ...field.RelationField) IPieceCommRecordDo
	Preload(fields ...field.RelationField) IPieceCommRecordDo
	FirstOrInit() (*model.PieceCommRecord, error)
	FirstOrCreate() (*model.PieceCommRecord, error)
	FindByPage(offset int, limit int) (result []*model.PieceCommRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPieceCommRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pieceCommRecordDo) Debug() IPieceCommRecordDo {
	return p.withDO(p.DO.Debug())
}

func (p pieceCommRecordDo) WithContext(ctx context.Context) IPieceCommRecordDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pieceCommRecordDo) ReadDB() IPieceCommRecordDo {
	return p.Clauses(dbresolver.Read)
}

func (p pieceCommRecordDo) WriteDB() IPieceCommRecordDo {
	return p.Clauses(dbresolver.Write)
}

func (p pieceCommRecordDo) Clauses(conds ...clause.Expression) IPieceCommRecordDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pieceCommRecordDo) Returning(value interface{}, columns ...string) IPieceCommRecordDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pieceCommRecordDo) Not(conds ...gen.Condition) IPieceCommRecordDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pieceCommRecordDo) Or(conds ...gen.Condition) IPieceCommRecordDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pieceCommRecordDo) Select(conds ...field.Expr) IPieceCommRecordDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pieceCommRecordDo) Where(conds ...gen.Condition) IPieceCommRecordDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pieceCommRecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IPieceCommRecordDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p pieceCommRecordDo) Order(conds ...field.Expr) IPieceCommRecordDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pieceCommRecordDo) Distinct(cols ...field.Expr) IPieceCommRecordDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pieceCommRecordDo) Omit(cols ...field.Expr) IPieceCommRecordDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pieceCommRecordDo) Join(table schema.Tabler, on ...field.Expr) IPieceCommRecordDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pieceCommRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPieceCommRecordDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pieceCommRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IPieceCommRecordDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pieceCommRecordDo) Group(cols ...field.Expr) IPieceCommRecordDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pieceCommRecordDo) Having(conds ...gen.Condition) IPieceCommRecordDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pieceCommRecordDo) Limit(limit int) IPieceCommRecordDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pieceCommRecordDo) Offset(offset int) IPieceCommRecordDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pieceCommRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPieceCommRecordDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pieceCommRecordDo) Unscoped() IPieceCommRecordDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pieceCommRecordDo) Create(values ...*model.PieceCommRecord) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pieceCommRecordDo) CreateInBatches(values []*model.PieceCommRecord, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pieceCommRecordDo) Save(values ...*model.PieceCommRecord) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pieceCommRecordDo) First() (*model.PieceCommRecord, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PieceCommRecord), nil
	}
}

func (p pieceCommRecordDo) Take() (*model.PieceCommRecord, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PieceCommRecord), nil
	}
}

func (p pieceCommRecordDo) Last() (*model.PieceCommRecord, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PieceCommRecord), nil
	}
}

func (p pieceCommRecordDo) Find() ([]*model.PieceCommRecord, error) {
	result, err := p.DO.Find()
	return result.([]*model.PieceCommRecord), err
}

func (p pieceCommRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PieceCommRecord, err error) {
	buf := make([]*model.PieceCommRecord, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pieceCommRecordDo) FindInBatches(result *[]*model.PieceCommRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pieceCommRecordDo) Attrs(attrs ...field.AssignExpr) IPieceCommRecordDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pieceCommRecordDo) Assign(attrs ...field.AssignExpr) IPieceCommRecordDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pieceCommRecordDo) Joins(fields ...field.RelationField) IPieceCommRecordDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pieceCommRecordDo) Preload(fields ...field.RelationField) IPieceCommRecordDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pieceCommRecordDo) FirstOrInit() (*model.PieceCommRecord, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PieceCommRecord), nil
	}
}

func (p pieceCommRecordDo) FirstOrCreate() (*model.PieceCommRecord, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PieceCommRecord), nil
	}
}

func (p pieceCommRecordDo) FindByPage(offset int, limit int) (result []*model.PieceCommRecord, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p pieceCommRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pieceCommRecordDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p *pieceCommRecordDo) withDO(do gen.Dao) *pieceCommRecordDo {
	p.DO = *do.(*gen.DO)
	return p
}
