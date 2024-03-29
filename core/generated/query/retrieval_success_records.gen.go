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

func newRetrievalSuccessRecord(db *gorm.DB) retrievalSuccessRecord {
	_retrievalSuccessRecord := retrievalSuccessRecord{}

	_retrievalSuccessRecord.retrievalSuccessRecordDo.UseDB(db)
	_retrievalSuccessRecord.retrievalSuccessRecordDo.UseModel(&model.RetrievalSuccessRecord{})

	tableName := _retrievalSuccessRecord.retrievalSuccessRecordDo.TableName()
	_retrievalSuccessRecord.ALL = field.NewField(tableName, "*")
	_retrievalSuccessRecord.PropCid = field.NewBytes(tableName, "prop_cid")
	_retrievalSuccessRecord.Miner = field.NewString(tableName, "miner")
	_retrievalSuccessRecord.Peer = field.NewString(tableName, "peer")
	_retrievalSuccessRecord.Size = field.NewInt64(tableName, "size")
	_retrievalSuccessRecord.DurationMs = field.NewInt64(tableName, "duration_ms")
	_retrievalSuccessRecord.AverageSpeed = field.NewInt64(tableName, "average_speed")
	_retrievalSuccessRecord.TotalPayment = field.NewString(tableName, "total_payment")
	_retrievalSuccessRecord.NumPayments = field.NewInt64(tableName, "num_payments")
	_retrievalSuccessRecord.AskPrice = field.NewString(tableName, "ask_price")
	_retrievalSuccessRecord.ID = field.NewInt64(tableName, "id")
	_retrievalSuccessRecord.CreatedAt = field.NewTime(tableName, "created_at")
	_retrievalSuccessRecord.Cid = field.NewBytes(tableName, "cid")

	_retrievalSuccessRecord.fillFieldMap()

	return _retrievalSuccessRecord
}

type retrievalSuccessRecord struct {
	retrievalSuccessRecordDo

	ALL          field.Field
	PropCid      field.Bytes
	Miner        field.String
	Peer         field.String
	Size         field.Int64
	DurationMs   field.Int64
	AverageSpeed field.Int64
	TotalPayment field.String
	NumPayments  field.Int64
	AskPrice     field.String
	ID           field.Int64
	CreatedAt    field.Time
	Cid          field.Bytes

	fieldMap map[string]field.Expr
}

func (r retrievalSuccessRecord) Table(newTableName string) *retrievalSuccessRecord {
	r.retrievalSuccessRecordDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r retrievalSuccessRecord) As(alias string) *retrievalSuccessRecord {
	r.retrievalSuccessRecordDo.DO = *(r.retrievalSuccessRecordDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *retrievalSuccessRecord) updateTableName(table string) *retrievalSuccessRecord {
	r.ALL = field.NewField(table, "*")
	r.PropCid = field.NewBytes(table, "prop_cid")
	r.Miner = field.NewString(table, "miner")
	r.Peer = field.NewString(table, "peer")
	r.Size = field.NewInt64(table, "size")
	r.DurationMs = field.NewInt64(table, "duration_ms")
	r.AverageSpeed = field.NewInt64(table, "average_speed")
	r.TotalPayment = field.NewString(table, "total_payment")
	r.NumPayments = field.NewInt64(table, "num_payments")
	r.AskPrice = field.NewString(table, "ask_price")
	r.ID = field.NewInt64(table, "id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.Cid = field.NewBytes(table, "cid")

	r.fillFieldMap()

	return r
}

func (r *retrievalSuccessRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *retrievalSuccessRecord) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 12)
	r.fieldMap["prop_cid"] = r.PropCid
	r.fieldMap["miner"] = r.Miner
	r.fieldMap["peer"] = r.Peer
	r.fieldMap["size"] = r.Size
	r.fieldMap["duration_ms"] = r.DurationMs
	r.fieldMap["average_speed"] = r.AverageSpeed
	r.fieldMap["total_payment"] = r.TotalPayment
	r.fieldMap["num_payments"] = r.NumPayments
	r.fieldMap["ask_price"] = r.AskPrice
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["cid"] = r.Cid
}

func (r retrievalSuccessRecord) clone(db *gorm.DB) retrievalSuccessRecord {
	r.retrievalSuccessRecordDo.ReplaceDB(db)
	return r
}

type retrievalSuccessRecordDo struct{ gen.DO }

type IRetrievalSuccessRecordDo interface {
	Debug() IRetrievalSuccessRecordDo
	WithContext(ctx context.Context) IRetrievalSuccessRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRetrievalSuccessRecordDo
	Not(conds ...gen.Condition) IRetrievalSuccessRecordDo
	Or(conds ...gen.Condition) IRetrievalSuccessRecordDo
	Select(conds ...field.Expr) IRetrievalSuccessRecordDo
	Where(conds ...gen.Condition) IRetrievalSuccessRecordDo
	Order(conds ...field.Expr) IRetrievalSuccessRecordDo
	Distinct(cols ...field.Expr) IRetrievalSuccessRecordDo
	Omit(cols ...field.Expr) IRetrievalSuccessRecordDo
	Join(table schema.Tabler, on ...field.Expr) IRetrievalSuccessRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRetrievalSuccessRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRetrievalSuccessRecordDo
	Group(cols ...field.Expr) IRetrievalSuccessRecordDo
	Having(conds ...gen.Condition) IRetrievalSuccessRecordDo
	Limit(limit int) IRetrievalSuccessRecordDo
	Offset(offset int) IRetrievalSuccessRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRetrievalSuccessRecordDo
	Unscoped() IRetrievalSuccessRecordDo
	Create(values ...*model.RetrievalSuccessRecord) error
	CreateInBatches(values []*model.RetrievalSuccessRecord, batchSize int) error
	Save(values ...*model.RetrievalSuccessRecord) error
	First() (*model.RetrievalSuccessRecord, error)
	Take() (*model.RetrievalSuccessRecord, error)
	Last() (*model.RetrievalSuccessRecord, error)
	Find() ([]*model.RetrievalSuccessRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RetrievalSuccessRecord, err error)
	FindInBatches(result *[]*model.RetrievalSuccessRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRetrievalSuccessRecordDo
	Assign(attrs ...field.AssignExpr) IRetrievalSuccessRecordDo
	Joins(fields ...field.RelationField) IRetrievalSuccessRecordDo
	Preload(fields ...field.RelationField) IRetrievalSuccessRecordDo
	FirstOrInit() (*model.RetrievalSuccessRecord, error)
	FirstOrCreate() (*model.RetrievalSuccessRecord, error)
	FindByPage(offset int, limit int) (result []*model.RetrievalSuccessRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRetrievalSuccessRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r retrievalSuccessRecordDo) Debug() IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Debug())
}

func (r retrievalSuccessRecordDo) WithContext(ctx context.Context) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r retrievalSuccessRecordDo) ReadDB() IRetrievalSuccessRecordDo {
	return r.Clauses(dbresolver.Read)
}

func (r retrievalSuccessRecordDo) WriteDB() IRetrievalSuccessRecordDo {
	return r.Clauses(dbresolver.Write)
}

func (r retrievalSuccessRecordDo) Clauses(conds ...clause.Expression) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r retrievalSuccessRecordDo) Returning(value interface{}, columns ...string) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r retrievalSuccessRecordDo) Not(conds ...gen.Condition) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r retrievalSuccessRecordDo) Or(conds ...gen.Condition) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r retrievalSuccessRecordDo) Select(conds ...field.Expr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r retrievalSuccessRecordDo) Where(conds ...gen.Condition) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r retrievalSuccessRecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IRetrievalSuccessRecordDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r retrievalSuccessRecordDo) Order(conds ...field.Expr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r retrievalSuccessRecordDo) Distinct(cols ...field.Expr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r retrievalSuccessRecordDo) Omit(cols ...field.Expr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r retrievalSuccessRecordDo) Join(table schema.Tabler, on ...field.Expr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r retrievalSuccessRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r retrievalSuccessRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r retrievalSuccessRecordDo) Group(cols ...field.Expr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r retrievalSuccessRecordDo) Having(conds ...gen.Condition) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r retrievalSuccessRecordDo) Limit(limit int) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r retrievalSuccessRecordDo) Offset(offset int) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r retrievalSuccessRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r retrievalSuccessRecordDo) Unscoped() IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Unscoped())
}

func (r retrievalSuccessRecordDo) Create(values ...*model.RetrievalSuccessRecord) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r retrievalSuccessRecordDo) CreateInBatches(values []*model.RetrievalSuccessRecord, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r retrievalSuccessRecordDo) Save(values ...*model.RetrievalSuccessRecord) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r retrievalSuccessRecordDo) First() (*model.RetrievalSuccessRecord, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalSuccessRecord), nil
	}
}

func (r retrievalSuccessRecordDo) Take() (*model.RetrievalSuccessRecord, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalSuccessRecord), nil
	}
}

func (r retrievalSuccessRecordDo) Last() (*model.RetrievalSuccessRecord, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalSuccessRecord), nil
	}
}

func (r retrievalSuccessRecordDo) Find() ([]*model.RetrievalSuccessRecord, error) {
	result, err := r.DO.Find()
	return result.([]*model.RetrievalSuccessRecord), err
}

func (r retrievalSuccessRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RetrievalSuccessRecord, err error) {
	buf := make([]*model.RetrievalSuccessRecord, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r retrievalSuccessRecordDo) FindInBatches(result *[]*model.RetrievalSuccessRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r retrievalSuccessRecordDo) Attrs(attrs ...field.AssignExpr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r retrievalSuccessRecordDo) Assign(attrs ...field.AssignExpr) IRetrievalSuccessRecordDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r retrievalSuccessRecordDo) Joins(fields ...field.RelationField) IRetrievalSuccessRecordDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r retrievalSuccessRecordDo) Preload(fields ...field.RelationField) IRetrievalSuccessRecordDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r retrievalSuccessRecordDo) FirstOrInit() (*model.RetrievalSuccessRecord, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalSuccessRecord), nil
	}
}

func (r retrievalSuccessRecordDo) FirstOrCreate() (*model.RetrievalSuccessRecord, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalSuccessRecord), nil
	}
}

func (r retrievalSuccessRecordDo) FindByPage(offset int, limit int) (result []*model.RetrievalSuccessRecord, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r retrievalSuccessRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r retrievalSuccessRecordDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r *retrievalSuccessRecordDo) withDO(do gen.Dao) *retrievalSuccessRecordDo {
	r.DO = *do.(*gen.DO)
	return r
}
