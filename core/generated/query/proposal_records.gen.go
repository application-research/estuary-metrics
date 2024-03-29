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

func newProposalRecord(db *gorm.DB) proposalRecord {
	_proposalRecord := proposalRecord{}

	_proposalRecord.proposalRecordDo.UseDB(db)
	_proposalRecord.proposalRecordDo.UseModel(&model.ProposalRecord{})

	tableName := _proposalRecord.proposalRecordDo.TableName()
	_proposalRecord.ALL = field.NewField(tableName, "*")
	_proposalRecord.PropCid = field.NewBytes(tableName, "prop_cid")
	_proposalRecord.Data = field.NewBytes(tableName, "data")

	_proposalRecord.fillFieldMap()

	return _proposalRecord
}

type proposalRecord struct {
	proposalRecordDo

	ALL     field.Field
	PropCid field.Bytes
	Data    field.Bytes

	fieldMap map[string]field.Expr
}

func (p proposalRecord) Table(newTableName string) *proposalRecord {
	p.proposalRecordDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p proposalRecord) As(alias string) *proposalRecord {
	p.proposalRecordDo.DO = *(p.proposalRecordDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *proposalRecord) updateTableName(table string) *proposalRecord {
	p.ALL = field.NewField(table, "*")
	p.PropCid = field.NewBytes(table, "prop_cid")
	p.Data = field.NewBytes(table, "data")

	p.fillFieldMap()

	return p
}

func (p *proposalRecord) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *proposalRecord) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 2)
	p.fieldMap["prop_cid"] = p.PropCid
	p.fieldMap["data"] = p.Data
}

func (p proposalRecord) clone(db *gorm.DB) proposalRecord {
	p.proposalRecordDo.ReplaceDB(db)
	return p
}

type proposalRecordDo struct{ gen.DO }

type IProposalRecordDo interface {
	Debug() IProposalRecordDo
	WithContext(ctx context.Context) IProposalRecordDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProposalRecordDo
	Not(conds ...gen.Condition) IProposalRecordDo
	Or(conds ...gen.Condition) IProposalRecordDo
	Select(conds ...field.Expr) IProposalRecordDo
	Where(conds ...gen.Condition) IProposalRecordDo
	Order(conds ...field.Expr) IProposalRecordDo
	Distinct(cols ...field.Expr) IProposalRecordDo
	Omit(cols ...field.Expr) IProposalRecordDo
	Join(table schema.Tabler, on ...field.Expr) IProposalRecordDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProposalRecordDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProposalRecordDo
	Group(cols ...field.Expr) IProposalRecordDo
	Having(conds ...gen.Condition) IProposalRecordDo
	Limit(limit int) IProposalRecordDo
	Offset(offset int) IProposalRecordDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProposalRecordDo
	Unscoped() IProposalRecordDo
	Create(values ...*model.ProposalRecord) error
	CreateInBatches(values []*model.ProposalRecord, batchSize int) error
	Save(values ...*model.ProposalRecord) error
	First() (*model.ProposalRecord, error)
	Take() (*model.ProposalRecord, error)
	Last() (*model.ProposalRecord, error)
	Find() ([]*model.ProposalRecord, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProposalRecord, err error)
	FindInBatches(result *[]*model.ProposalRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProposalRecordDo
	Assign(attrs ...field.AssignExpr) IProposalRecordDo
	Joins(fields ...field.RelationField) IProposalRecordDo
	Preload(fields ...field.RelationField) IProposalRecordDo
	FirstOrInit() (*model.ProposalRecord, error)
	FirstOrCreate() (*model.ProposalRecord, error)
	FindByPage(offset int, limit int) (result []*model.ProposalRecord, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProposalRecordDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p proposalRecordDo) Debug() IProposalRecordDo {
	return p.withDO(p.DO.Debug())
}

func (p proposalRecordDo) WithContext(ctx context.Context) IProposalRecordDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p proposalRecordDo) ReadDB() IProposalRecordDo {
	return p.Clauses(dbresolver.Read)
}

func (p proposalRecordDo) WriteDB() IProposalRecordDo {
	return p.Clauses(dbresolver.Write)
}

func (p proposalRecordDo) Clauses(conds ...clause.Expression) IProposalRecordDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p proposalRecordDo) Returning(value interface{}, columns ...string) IProposalRecordDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p proposalRecordDo) Not(conds ...gen.Condition) IProposalRecordDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p proposalRecordDo) Or(conds ...gen.Condition) IProposalRecordDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p proposalRecordDo) Select(conds ...field.Expr) IProposalRecordDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p proposalRecordDo) Where(conds ...gen.Condition) IProposalRecordDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p proposalRecordDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IProposalRecordDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p proposalRecordDo) Order(conds ...field.Expr) IProposalRecordDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p proposalRecordDo) Distinct(cols ...field.Expr) IProposalRecordDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p proposalRecordDo) Omit(cols ...field.Expr) IProposalRecordDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p proposalRecordDo) Join(table schema.Tabler, on ...field.Expr) IProposalRecordDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p proposalRecordDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProposalRecordDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p proposalRecordDo) RightJoin(table schema.Tabler, on ...field.Expr) IProposalRecordDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p proposalRecordDo) Group(cols ...field.Expr) IProposalRecordDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p proposalRecordDo) Having(conds ...gen.Condition) IProposalRecordDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p proposalRecordDo) Limit(limit int) IProposalRecordDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p proposalRecordDo) Offset(offset int) IProposalRecordDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p proposalRecordDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProposalRecordDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p proposalRecordDo) Unscoped() IProposalRecordDo {
	return p.withDO(p.DO.Unscoped())
}

func (p proposalRecordDo) Create(values ...*model.ProposalRecord) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p proposalRecordDo) CreateInBatches(values []*model.ProposalRecord, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p proposalRecordDo) Save(values ...*model.ProposalRecord) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p proposalRecordDo) First() (*model.ProposalRecord, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalRecord), nil
	}
}

func (p proposalRecordDo) Take() (*model.ProposalRecord, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalRecord), nil
	}
}

func (p proposalRecordDo) Last() (*model.ProposalRecord, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalRecord), nil
	}
}

func (p proposalRecordDo) Find() ([]*model.ProposalRecord, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProposalRecord), err
}

func (p proposalRecordDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProposalRecord, err error) {
	buf := make([]*model.ProposalRecord, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p proposalRecordDo) FindInBatches(result *[]*model.ProposalRecord, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p proposalRecordDo) Attrs(attrs ...field.AssignExpr) IProposalRecordDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p proposalRecordDo) Assign(attrs ...field.AssignExpr) IProposalRecordDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p proposalRecordDo) Joins(fields ...field.RelationField) IProposalRecordDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p proposalRecordDo) Preload(fields ...field.RelationField) IProposalRecordDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p proposalRecordDo) FirstOrInit() (*model.ProposalRecord, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalRecord), nil
	}
}

func (p proposalRecordDo) FirstOrCreate() (*model.ProposalRecord, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProposalRecord), nil
	}
}

func (p proposalRecordDo) FindByPage(offset int, limit int) (result []*model.ProposalRecord, count int64, err error) {
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

func (p proposalRecordDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p proposalRecordDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p *proposalRecordDo) withDO(do gen.Dao) *proposalRecordDo {
	p.DO = *do.(*gen.DO)
	return p
}
