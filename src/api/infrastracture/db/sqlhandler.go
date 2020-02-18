package db

import (
	"api/infrastracture/env"
	"api/interfaces"
	"fmt"

	"github.com/jinzhu/gorm"
)

type SQLHandler struct {
	Conn *gorm.DB
}

func NewSQLHandler() (interfaces.SQLHandler, error) {
	var env = env.Load()
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.DbUser, env.DbPassword, env.DbHost, env.DbPort, env.DbDatabase)
	conn, err := gorm.Open(env.DbDriver, dataSourceName)
	if err != nil {
		panic(err.Error)
	}
	err = conn.DB().Ping()
	if err != nil {
		fmt.Printf("Can not connection DB : %s", err)
		fmt.Printf("dataSourceName : %s", dataSourceName)
		return nil, err
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.Conn = conn
	return sqlHandler, nil
}

func (handler *SQLHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SQLHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SQLHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SQLHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SQLHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SQLHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SQLHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SQLHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}
