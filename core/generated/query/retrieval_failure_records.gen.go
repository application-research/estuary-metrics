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

func newRetrievalFailureRecord(db *gorm.DB) retrievalFailureRecord {
	_retrievalFailureRecord := retrievalFailureRecord{}

	_retrievalFailureRecord.retrievalFailureRecordDo.UseDB(db)
	_retrievalFailureRecord.retrievalFailureRecordDo.UseModel(&model.RetrievalFailureRecord{})

	tableName := _retrievalFailureRecord.retrievalFailureRecordDo.TableName()
	_retrievalFailureRecord.ALL = field.NewField(tableName, "*")
	_retrievalFailureRecord.ID = field.NewInt64(tableName, "id")
	_retrievalFailureRecord.CreatedAt = field.NewTime(tableName, "created_at")
	_retrievalFailureRecord.UpdatedAt = field.NewTime(tableName, "updated_at")
	_retrievalFailureRecord.DeletedAt = field.NewField(tableName, "deleted_at")
	_retrievalFailureRecord.Miner = field.NewString(tableName, "miner")
	_retrievalFailureRecord.Phase = field.NewString(tableName, "phase")
	_retrievalFailureRecord.Message = field.NewString(tableName, "message")
	_retrievalFailureRecord.Content = field.NewInt64(tableName, "content")
	_retrievalFailureRecord.Cid = field.NewBytes(tableName, "cid")

	_retrievalFailureRecord.fillFieldMap()

	return _retrievalFailureRecord
}

type retrievalFailureRecord struct {
	retrievalFailureRecordDo

	ALL       field.Field
	ID        field.Int64
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field
	Miner     field.String
	Phase     field.String
	Message   field.String
	Content   field.Int64
	Cid       field.Bytes

	fieldMap map[string]field.Expr
}

func (r retrievalFailureRecord) Table(newTableName string) *retrievalFailureRecord {
	r.retrievalFailureRecordDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r retrievalFailureRecord) As(alias string) *retrievalFailureRecord {
	r.retrievalFailureRecordDo.DO = *(r.retrievalFailureRecordDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *retrievalFailureRecord) updateTableName(table string) *retrievalFailureRecord {
	r.ALL = field.NewField(table, "*")
	r.ID = field.NewInt64(table, "id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")
	r.Miner = field.NewString(table, "miner")
	r.Phase = field.NewString(table, "phase")
	r.Message = field.NewString(table, "message")
	r.Content = field.NewInt64(table, "content")
	r.Cid = field.NewBytes(table, "cid")

	r.fillFieldMap()

	return r
}

func (r *retrievalFailureRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *retrievalFailureRecord) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 9)
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["miner"] = r.Miner
	r.fieldMap["phase"] = r.Phase
	r.fieldMap["message"] = r.Message
	r.fieldMap["content"] = r.Content
	r.fieldMap["cid"] = r.Cid
}

func (r retrievalFailureRecord) clone(db *gorm.DB) retrievalFailureRecord {
	r.retrievalFailureRecordDo.ReplaceDB(db)
	return r
}

type retrievalFailureRecordDo struct{ gen.DO }

type IRetrievalFailureRecordDo interface {
	Debug() IRetrievalFailureRecordDo
	WithContext(ctx context.Context) IRetrievalFailureRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRetrievalFailureRecordDo
	Not(conds ...gen.Condition) IRetrievalFailureRecordDo
	Or(conds ...gen.Condition) IRetrievalFailureRecordDo
	Select(conds ...field.Expr) IRetrievalFailureRecordDo
	Where(conds ...gen.Condition) IRetrievalFailureRecordDo
	Order(conds ...field.Expr) IRetrievalFailureRecordDo
	Distinct(cols ...field.Expr) IRetrievalFailureRecordDo
	Omit(cols ...field.Expr) IRetrievalFailureRecordDo
	Join(table schema.Tabler, on ...field.Expr) IRetrievalFailureRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRetrievalFailureRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRetrievalFailureRecordDo
	Group(cols ...field.Expr) IRetrievalFailureRecordDo
	Having(conds ...gen.Condition) IRetrievalFailureRecordDo
	Limit(limit int) IRetrievalFailureRecordDo
	Offset(offset int) IRetrievalFailureRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRetrievalFailureRecordDo
	Unscoped() IRetrievalFailureRecordDo
	Create(values ...*model.RetrievalFailureRecord) error
	CreateInBatches(values []*model.RetrievalFailureRecord, batchSize int) error
	Save(values ...*model.RetrievalFailureRecord) error
	First() (*model.RetrievalFailureRecord, error)
	Take() (*model.RetrievalFailureRecord, error)
	Last() (*model.RetrievalFailureRecord, error)
	Find() ([]*model.RetrievalFailureRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RetrievalFailureRecord, err error)
	FindInBatches(result *[]*model.RetrievalFailureRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRetrievalFailureRecordDo
	Assign(attrs ...field.AssignExpr) IRetrievalFailureRecordDo
	Joins(fields ...field.RelationField) IRetrievalFailureRecordDo
	Preload(fields ...field.RelationField) IRetrievalFailureRecordDo
	FirstOrInit() (*model.RetrievalFailureRecord, error)
	FirstOrCreate() (*model.RetrievalFailureRecord, error)
	FindByPage(offset int, limit int) (result []*model.RetrievalFailureRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRetrievalFailureRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r retrievalFailureRecordDo) Debug() IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Debug())
}

func (r retrievalFailureRecordDo) WithContext(ctx context.Context) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r retrievalFailureRecordDo) ReadDB() IRetrievalFailureRecordDo {
	return r.Clauses(dbresolver.Read)
}

func (r retrievalFailureRecordDo) WriteDB() IRetrievalFailureRecordDo {
	return r.Clauses(dbresolver.Write)
}

func (r retrievalFailureRecordDo) Clauses(conds ...clause.Expression) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r retrievalFailureRecordDo) Returning(value interface{}, columns ...string) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r retrievalFailureRecordDo) Not(conds ...gen.Condition) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r retrievalFailureRecordDo) Or(conds ...gen.Condition) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r retrievalFailureRecordDo) Select(conds ...field.Expr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r retrievalFailureRecordDo) Where(conds ...gen.Condition) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r retrievalFailureRecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IRetrievalFailureRecordDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r retrievalFailureRecordDo) Order(conds ...field.Expr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r retrievalFailureRecordDo) Distinct(cols ...field.Expr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r retrievalFailureRecordDo) Omit(cols ...field.Expr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r retrievalFailureRecordDo) Join(table schema.Tabler, on ...field.Expr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r retrievalFailureRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r retrievalFailureRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r retrievalFailureRecordDo) Group(cols ...field.Expr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r retrievalFailureRecordDo) Having(conds ...gen.Condition) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r retrievalFailureRecordDo) Limit(limit int) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r retrievalFailureRecordDo) Offset(offset int) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r retrievalFailureRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r retrievalFailureRecordDo) Unscoped() IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Unscoped())
}

func (r retrievalFailureRecordDo) Create(values ...*model.RetrievalFailureRecord) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r retrievalFailureRecordDo) CreateInBatches(values []*model.RetrievalFailureRecord, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r retrievalFailureRecordDo) Save(values ...*model.RetrievalFailureRecord) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r retrievalFailureRecordDo) First() (*model.RetrievalFailureRecord, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalFailureRecord), nil
	}
}

func (r retrievalFailureRecordDo) Take() (*model.RetrievalFailureRecord, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalFailureRecord), nil
	}
}

func (r retrievalFailureRecordDo) Last() (*model.RetrievalFailureRecord, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalFailureRecord), nil
	}
}

func (r retrievalFailureRecordDo) Find() ([]*model.RetrievalFailureRecord, error) {
	result, err := r.DO.Find()
	return result.([]*model.RetrievalFailureRecord), err
}

func (r retrievalFailureRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.RetrievalFailureRecord, err error) {
	buf := make([]*model.RetrievalFailureRecord, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r retrievalFailureRecordDo) FindInBatches(result *[]*model.RetrievalFailureRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r retrievalFailureRecordDo) Attrs(attrs ...field.AssignExpr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r retrievalFailureRecordDo) Assign(attrs ...field.AssignExpr) IRetrievalFailureRecordDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r retrievalFailureRecordDo) Joins(fields ...field.RelationField) IRetrievalFailureRecordDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r retrievalFailureRecordDo) Preload(fields ...field.RelationField) IRetrievalFailureRecordDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r retrievalFailureRecordDo) FirstOrInit() (*model.RetrievalFailureRecord, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalFailureRecord), nil
	}
}

func (r retrievalFailureRecordDo) FirstOrCreate() (*model.RetrievalFailureRecord, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.RetrievalFailureRecord), nil
	}
}

func (r retrievalFailureRecordDo) FindByPage(offset int, limit int) (result []*model.RetrievalFailureRecord, count int64, err error) {
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

func (r retrievalFailureRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r retrievalFailureRecordDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r *retrievalFailureRecordDo) withDO(do gen.Dao) *retrievalFailureRecordDo {
	r.DO = *do.(*gen.DO)
	return r
}
