package core

import (
	"github.com/whyrusleeping/memo"
	"gorm.io/gorm"
)

type ConnectionManager struct {
	Connnection *Connection
}

type Connection struct {
	Db     *gorm.DB
	Cacher *memo.Cacher
}

func (c *ConnectionManager) Init() {
	c.Connnection = &Connection{}
}

func (c *ConnectionManager) SetConnection(db *gorm.DB, cacherm *memo.Cacher) (*Connection, error) {
	c.Connnection.Db = db
	c.Connnection.Cacher = cacherm
	return c.Connnection, nil
}

func (c *ConnectionManager) GetConnection() *Connection {
	return c.Connnection
}
