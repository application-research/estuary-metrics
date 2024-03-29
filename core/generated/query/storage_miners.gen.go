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

func newStorageMiner(db *gorm.DB) storageMiner {
	_storageMiner := storageMiner{}

	_storageMiner.storageMinerDo.UseDB(db)
	_storageMiner.storageMinerDo.UseModel(&model.StorageMiner{})

	tableName := _storageMiner.storageMinerDo.TableName()
	_storageMiner.ALL = field.NewField(tableName, "*")
	_storageMiner.ID = field.NewInt64(tableName, "id")
	_storageMiner.CreatedAt = field.NewTime(tableName, "created_at")
	_storageMiner.UpdatedAt = field.NewTime(tableName, "updated_at")
	_storageMiner.DeletedAt = field.NewField(tableName, "deleted_at")
	_storageMiner.Address = field.NewString(tableName, "address")
	_storageMiner.Suspended = field.NewBool(tableName, "suspended")
	_storageMiner.SuspendedReason = field.NewString(tableName, "suspended_reason")
	_storageMiner.Name = field.NewString(tableName, "name")
	_storageMiner.Version = field.NewString(tableName, "version")
	_storageMiner.Location = field.NewString(tableName, "location")
	_storageMiner.Owner = field.NewInt64(tableName, "owner")

	_storageMiner.fillFieldMap()

	return _storageMiner
}

type storageMiner struct {
	storageMinerDo

	ALL             field.Field
	ID              field.Int64
	CreatedAt       field.Time
	UpdatedAt       field.Time
	DeletedAt       field.Field
	Address         field.String
	Suspended       field.Bool
	SuspendedReason field.String
	Name            field.String
	Version         field.String
	Location        field.String
	Owner           field.Int64

	fieldMap map[string]field.Expr
}

func (s storageMiner) Table(newTableName string) *storageMiner {
	s.storageMinerDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s storageMiner) As(alias string) *storageMiner {
	s.storageMinerDo.DO = *(s.storageMinerDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *storageMiner) updateTableName(table string) *storageMiner {
	s.ALL = field.NewField(table, "*")
	s.ID = field.NewInt64(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.Address = field.NewString(table, "address")
	s.Suspended = field.NewBool(table, "suspended")
	s.SuspendedReason = field.NewString(table, "suspended_reason")
	s.Name = field.NewString(table, "name")
	s.Version = field.NewString(table, "version")
	s.Location = field.NewString(table, "location")
	s.Owner = field.NewInt64(table, "owner")

	s.fillFieldMap()

	return s
}

func (s *storageMiner) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *storageMiner) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 11)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["address"] = s.Address
	s.fieldMap["suspended"] = s.Suspended
	s.fieldMap["suspended_reason"] = s.SuspendedReason
	s.fieldMap["name"] = s.Name
	s.fieldMap["version"] = s.Version
	s.fieldMap["location"] = s.Location
	s.fieldMap["owner"] = s.Owner
}

func (s storageMiner) clone(db *gorm.DB) storageMiner {
	s.storageMinerDo.ReplaceDB(db)
	return s
}

type storageMinerDo struct{ gen.DO }

type IStorageMinerDo interface {
	Debug() IStorageMinerDo
	WithContext(ctx context.Context) IStorageMinerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IStorageMinerDo
	Not(conds ...gen.Condition) IStorageMinerDo
	Or(conds ...gen.Condition) IStorageMinerDo
	Select(conds ...field.Expr) IStorageMinerDo
	Where(conds ...gen.Condition) IStorageMinerDo
	Order(conds ...field.Expr) IStorageMinerDo
	Distinct(cols ...field.Expr) IStorageMinerDo
	Omit(cols ...field.Expr) IStorageMinerDo
	Join(table schema.Tabler, on ...field.Expr) IStorageMinerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IStorageMinerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IStorageMinerDo
	Group(cols ...field.Expr) IStorageMinerDo
	Having(conds ...gen.Condition) IStorageMinerDo
	Limit(limit int) IStorageMinerDo
	Offset(offset int) IStorageMinerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IStorageMinerDo
	Unscoped() IStorageMinerDo
	Create(values ...*model.StorageMiner) error
	CreateInBatches(values []*model.StorageMiner, batchSize int) error
	Save(values ...*model.StorageMiner) error
	First() (*model.StorageMiner, error)
	Take() (*model.StorageMiner, error)
	Last() (*model.StorageMiner, error)
	Find() ([]*model.StorageMiner, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StorageMiner, err error)
	FindInBatches(result *[]*model.StorageMiner, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IStorageMinerDo
	Assign(attrs ...field.AssignExpr) IStorageMinerDo
	Joins(fields ...field.RelationField) IStorageMinerDo
	Preload(fields ...field.RelationField) IStorageMinerDo
	FirstOrInit() (*model.StorageMiner, error)
	FirstOrCreate() (*model.StorageMiner, error)
	FindByPage(offset int, limit int) (result []*model.StorageMiner, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IStorageMinerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s storageMinerDo) Debug() IStorageMinerDo {
	return s.withDO(s.DO.Debug())
}

func (s storageMinerDo) WithContext(ctx context.Context) IStorageMinerDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s storageMinerDo) ReadDB() IStorageMinerDo {
	return s.Clauses(dbresolver.Read)
}

func (s storageMinerDo) WriteDB() IStorageMinerDo {
	return s.Clauses(dbresolver.Write)
}

func (s storageMinerDo) Clauses(conds ...clause.Expression) IStorageMinerDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s storageMinerDo) Returning(value interface{}, columns ...string) IStorageMinerDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s storageMinerDo) Not(conds ...gen.Condition) IStorageMinerDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s storageMinerDo) Or(conds ...gen.Condition) IStorageMinerDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s storageMinerDo) Select(conds ...field.Expr) IStorageMinerDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s storageMinerDo) Where(conds ...gen.Condition) IStorageMinerDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s storageMinerDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IStorageMinerDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s storageMinerDo) Order(conds ...field.Expr) IStorageMinerDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s storageMinerDo) Distinct(cols ...field.Expr) IStorageMinerDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s storageMinerDo) Omit(cols ...field.Expr) IStorageMinerDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s storageMinerDo) Join(table schema.Tabler, on ...field.Expr) IStorageMinerDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s storageMinerDo) LeftJoin(table schema.Tabler, on ...field.Expr) IStorageMinerDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s storageMinerDo) RightJoin(table schema.Tabler, on ...field.Expr) IStorageMinerDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s storageMinerDo) Group(cols ...field.Expr) IStorageMinerDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s storageMinerDo) Having(conds ...gen.Condition) IStorageMinerDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s storageMinerDo) Limit(limit int) IStorageMinerDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s storageMinerDo) Offset(offset int) IStorageMinerDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s storageMinerDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IStorageMinerDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s storageMinerDo) Unscoped() IStorageMinerDo {
	return s.withDO(s.DO.Unscoped())
}

func (s storageMinerDo) Create(values ...*model.StorageMiner) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s storageMinerDo) CreateInBatches(values []*model.StorageMiner, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s storageMinerDo) Save(values ...*model.StorageMiner) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s storageMinerDo) First() (*model.StorageMiner, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorageMiner), nil
	}
}

func (s storageMinerDo) Take() (*model.StorageMiner, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorageMiner), nil
	}
}

func (s storageMinerDo) Last() (*model.StorageMiner, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorageMiner), nil
	}
}

func (s storageMinerDo) Find() ([]*model.StorageMiner, error) {
	result, err := s.DO.Find()
	return result.([]*model.StorageMiner), err
}

func (s storageMinerDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.StorageMiner, err error) {
	buf := make([]*model.StorageMiner, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s storageMinerDo) FindInBatches(result *[]*model.StorageMiner, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s storageMinerDo) Attrs(attrs ...field.AssignExpr) IStorageMinerDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s storageMinerDo) Assign(attrs ...field.AssignExpr) IStorageMinerDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s storageMinerDo) Joins(fields ...field.RelationField) IStorageMinerDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s storageMinerDo) Preload(fields ...field.RelationField) IStorageMinerDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s storageMinerDo) FirstOrInit() (*model.StorageMiner, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorageMiner), nil
	}
}

func (s storageMinerDo) FirstOrCreate() (*model.StorageMiner, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.StorageMiner), nil
	}
}

func (s storageMinerDo) FindByPage(offset int, limit int) (result []*model.StorageMiner, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s storageMinerDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s storageMinerDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s *storageMinerDo) withDO(do gen.Dao) *storageMinerDo {
	s.DO = *do.(*gen.DO)
	return s
}
