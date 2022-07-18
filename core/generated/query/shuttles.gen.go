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

func newShuttle(db *gorm.DB) shuttle {
	_shuttle := shuttle{}

	_shuttle.shuttleDo.UseDB(db)
	_shuttle.shuttleDo.UseModel(&model.Shuttle{})

	tableName := _shuttle.shuttleDo.TableName()
	_shuttle.ALL = field.NewField(tableName, "*")
	_shuttle.ID = field.NewInt64(tableName, "id")
	_shuttle.CreatedAt = field.NewTime(tableName, "created_at")
	_shuttle.UpdatedAt = field.NewTime(tableName, "updated_at")
	_shuttle.DeletedAt = field.NewField(tableName, "deleted_at")
	_shuttle.Handle = field.NewString(tableName, "handle")
	_shuttle.Token = field.NewString(tableName, "token")
	_shuttle.LastConnection = field.NewTime(tableName, "last_connection")
	_shuttle.Host = field.NewString(tableName, "host")
	_shuttle.PeerID = field.NewString(tableName, "peer_id")
	_shuttle.Open = field.NewBool(tableName, "open")
	_shuttle.Private = field.NewBool(tableName, "private")
	_shuttle.Priority = field.NewInt64(tableName, "priority")

	_shuttle.fillFieldMap()

	return _shuttle
}

type shuttle struct {
	shuttleDo

	ALL            field.Field
	ID             field.Int64
	CreatedAt      field.Time
	UpdatedAt      field.Time
	DeletedAt      field.Field
	Handle         field.String
	Token          field.String
	LastConnection field.Time
	Host           field.String
	PeerID         field.String
	Open           field.Bool
	Private        field.Bool
	Priority       field.Int64

	fieldMap map[string]field.Expr
}

func (s shuttle) Table(newTableName string) *shuttle {
	s.shuttleDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s shuttle) As(alias string) *shuttle {
	s.shuttleDo.DO = *(s.shuttleDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *shuttle) updateTableName(table string) *shuttle {
	s.ALL = field.NewField(table, "*")
	s.ID = field.NewInt64(table, "id")
	s.CreatedAt = field.NewTime(table, "created_at")
	s.UpdatedAt = field.NewTime(table, "updated_at")
	s.DeletedAt = field.NewField(table, "deleted_at")
	s.Handle = field.NewString(table, "handle")
	s.Token = field.NewString(table, "token")
	s.LastConnection = field.NewTime(table, "last_connection")
	s.Host = field.NewString(table, "host")
	s.PeerID = field.NewString(table, "peer_id")
	s.Open = field.NewBool(table, "open")
	s.Private = field.NewBool(table, "private")
	s.Priority = field.NewInt64(table, "priority")

	s.fillFieldMap()

	return s
}

func (s *shuttle) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *shuttle) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 12)
	s.fieldMap["id"] = s.ID
	s.fieldMap["created_at"] = s.CreatedAt
	s.fieldMap["updated_at"] = s.UpdatedAt
	s.fieldMap["deleted_at"] = s.DeletedAt
	s.fieldMap["handle"] = s.Handle
	s.fieldMap["token"] = s.Token
	s.fieldMap["last_connection"] = s.LastConnection
	s.fieldMap["host"] = s.Host
	s.fieldMap["peer_id"] = s.PeerID
	s.fieldMap["open"] = s.Open
	s.fieldMap["private"] = s.Private
	s.fieldMap["priority"] = s.Priority
}

func (s shuttle) clone(db *gorm.DB) shuttle {
	s.shuttleDo.ReplaceDB(db)
	return s
}

type shuttleDo struct{ gen.DO }

type IShuttleDo interface {
	Debug() IShuttleDo
	WithContext(ctx context.Context) IShuttleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IShuttleDo
	Not(conds ...gen.Condition) IShuttleDo
	Or(conds ...gen.Condition) IShuttleDo
	Select(conds ...field.Expr) IShuttleDo
	Where(conds ...gen.Condition) IShuttleDo
	Order(conds ...field.Expr) IShuttleDo
	Distinct(cols ...field.Expr) IShuttleDo
	Omit(cols ...field.Expr) IShuttleDo
	Join(table schema.Tabler, on ...field.Expr) IShuttleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IShuttleDo
	RightJoin(table schema.Tabler, on ...field.Expr) IShuttleDo
	Group(cols ...field.Expr) IShuttleDo
	Having(conds ...gen.Condition) IShuttleDo
	Limit(limit int) IShuttleDo
	Offset(offset int) IShuttleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IShuttleDo
	Unscoped() IShuttleDo
	Create(values ...*model.Shuttle) error
	CreateInBatches(values []*model.Shuttle, batchSize int) error
	Save(values ...*model.Shuttle) error
	First() (*model.Shuttle, error)
	Take() (*model.Shuttle, error)
	Last() (*model.Shuttle, error)
	Find() ([]*model.Shuttle, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Shuttle, err error)
	FindInBatches(result *[]*model.Shuttle, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete() (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IShuttleDo
	Assign(attrs ...field.AssignExpr) IShuttleDo
	Joins(fields ...field.RelationField) IShuttleDo
	Preload(fields ...field.RelationField) IShuttleDo
	FirstOrInit() (*model.Shuttle, error)
	FirstOrCreate() (*model.Shuttle, error)
	FindByPage(offset int, limit int) (result []*model.Shuttle, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IShuttleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s shuttleDo) Debug() IShuttleDo {
	return s.withDO(s.DO.Debug())
}

func (s shuttleDo) WithContext(ctx context.Context) IShuttleDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s shuttleDo) ReadDB() IShuttleDo {
	return s.Clauses(dbresolver.Read)
}

func (s shuttleDo) WriteDB() IShuttleDo {
	return s.Clauses(dbresolver.Write)
}

func (s shuttleDo) Clauses(conds ...clause.Expression) IShuttleDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s shuttleDo) Returning(value interface{}, columns ...string) IShuttleDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s shuttleDo) Not(conds ...gen.Condition) IShuttleDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s shuttleDo) Or(conds ...gen.Condition) IShuttleDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s shuttleDo) Select(conds ...field.Expr) IShuttleDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s shuttleDo) Where(conds ...gen.Condition) IShuttleDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s shuttleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) IShuttleDo {
	return s.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (s shuttleDo) Order(conds ...field.Expr) IShuttleDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s shuttleDo) Distinct(cols ...field.Expr) IShuttleDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s shuttleDo) Omit(cols ...field.Expr) IShuttleDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s shuttleDo) Join(table schema.Tabler, on ...field.Expr) IShuttleDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s shuttleDo) LeftJoin(table schema.Tabler, on ...field.Expr) IShuttleDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s shuttleDo) RightJoin(table schema.Tabler, on ...field.Expr) IShuttleDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s shuttleDo) Group(cols ...field.Expr) IShuttleDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s shuttleDo) Having(conds ...gen.Condition) IShuttleDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s shuttleDo) Limit(limit int) IShuttleDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s shuttleDo) Offset(offset int) IShuttleDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s shuttleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IShuttleDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s shuttleDo) Unscoped() IShuttleDo {
	return s.withDO(s.DO.Unscoped())
}

func (s shuttleDo) Create(values ...*model.Shuttle) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s shuttleDo) CreateInBatches(values []*model.Shuttle, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s shuttleDo) Save(values ...*model.Shuttle) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s shuttleDo) First() (*model.Shuttle, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Shuttle), nil
	}
}

func (s shuttleDo) Take() (*model.Shuttle, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Shuttle), nil
	}
}

func (s shuttleDo) Last() (*model.Shuttle, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Shuttle), nil
	}
}

func (s shuttleDo) Find() ([]*model.Shuttle, error) {
	result, err := s.DO.Find()
	return result.([]*model.Shuttle), err
}

func (s shuttleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Shuttle, err error) {
	buf := make([]*model.Shuttle, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s shuttleDo) FindInBatches(result *[]*model.Shuttle, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s shuttleDo) Attrs(attrs ...field.AssignExpr) IShuttleDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s shuttleDo) Assign(attrs ...field.AssignExpr) IShuttleDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s shuttleDo) Joins(fields ...field.RelationField) IShuttleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s shuttleDo) Preload(fields ...field.RelationField) IShuttleDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s shuttleDo) FirstOrInit() (*model.Shuttle, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Shuttle), nil
	}
}

func (s shuttleDo) FirstOrCreate() (*model.Shuttle, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Shuttle), nil
	}
}

func (s shuttleDo) FindByPage(offset int, limit int) (result []*model.Shuttle, count int64, err error) {
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

func (s shuttleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s shuttleDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s *shuttleDo) withDO(do gen.Dao) *shuttleDo {
	s.DO = *do.(*gen.DO)
	return s
}
