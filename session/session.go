package session

import (
	"database/sql"
	"go-ORM/clause"
	"go-ORM/dialect"
	"go-ORM/log"
	"go-ORM/schema"
	"strings"
)

// session 用于实现与数据库的交互

type Session struct {
	db *sql.DB
	sql strings.Builder // strings.Builder 用来高效拼接字符串
	sqlVals []interface{}

	// +
	dialect  dialect.Dialect
	refTable *schema.Schema


	//+
	clause clause.Clause


}

// New 开启一个session
func New(db *sql.DB, d dialect.Dialect) *Session {
	return &Session{ db: db, dialect: d}
}

// DB returns *sql.DB
func (s *Session) DB() *sql.DB {
	return s.db
}

// Raw 拼接sql 和 参数
func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ") // 以空字符串隔开

	s.sqlVals = append(s.sqlVals, values...) // 添加多个参数
	return s
}

func (s *Session) Clear() {
	s.sql.Reset() // 清空Builder
	s.sqlVals = nil // 清空slice

	// + 清空clause
	s.clause = clause.Clause{}

}


// Exec raw sql with sqlVars
func (s *Session) Exec() (result sql.Result, err error){
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVals)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVals...); err != nil {
		log.Error(err)
	}
	return
}

// QueryRow gets a record from db
func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVals)
	return s.DB().QueryRow(s.sql.String(), s.sqlVals...)
}

// QueryRows gets a list of records from db
func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVals)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVals...); err != nil {
		log.Error(err)
	}
	return
}