package orm

import (
	"database/sql"
	"errors"
	"go-ORM/dialect"
	"go-ORM/log"
	"go-ORM/session"
)

// 连接/测试数据库
// 交互后的收尾工作（关闭连接）
// 等就交给 Engine 来负责了。
// Engine是与用户交互的入口。代码位于根目录的

type Engine struct {
	db *sql.DB

	// +
	dialect dialect.Dialect


}
// NewDB 开启一个数据库连接
func NewDB(driver, source string) ( *Engine, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	// Send a ping to make sure the database connection is alive.
	if err := db.Ping(); err != nil {
		log.Error(err)
		return nil, err
	}

	// make sure the specific dialect exists
	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Errorf("dialect %s Not Found", driver)
		return nil, errors.New("dialect not found")
	}

	// 返回
	e := &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return e, nil

}

// Close 关闭数据库连接
func (engine *Engine) Close()  {
	if err := engine.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("closed database success")
}

// NewSession 开启一个关联
func (engine *Engine) NewSession() *session.Session  {
	// 返回个session
	return session.New(engine.db, engine.dialect)
}

