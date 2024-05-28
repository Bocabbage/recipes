package session

import (
	"bocabbage/geeorm-learn/clause"
	"bocabbage/geeorm-learn/dialect"
	"bocabbage/geeorm-learn/log"
	"bocabbage/geeorm-learn/schema"
	"database/sql"
	"reflect"
	"strings"
)

type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect // dialect is by-session
	refTable *schema.Schema
	clause   clause.Clause
	sql      strings.Builder
	sqlVars  []interface{}
}

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
	}
}

// 获取 db handler
func (s *Session) DB() *sql.DB {
	return s.db
}

// 装载 sql 和 参数
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

// 清空 sql 和参数
func (s *Session) Clear() {
	s.clause = clause.Clause{}
	s.sql.Reset()
	s.sqlVars = nil
}

// 执行
func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.DB().QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

// ---------- Model 相关接口 ------------

// [加载Model] 接收用户自定义 struct 对象，解析为当前 session 的reftable
func (s *Session) Model(value interface{}) *Session {
	// nil or different model, update refTable
	if s.refTable == nil || reflect.TypeOf(value) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(value, s.dialect)
	}
	return s
}

func (s *Session) RefTable() *schema.Schema {
	if s.refTable == nil {
		log.Error("Model is not set")
	}
	return s.refTable
}
