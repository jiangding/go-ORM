package orm

import (
	"database/sql"
	"go-ORM/session"
	"log"
)

// 连接/测试数据库
// 交互后的收尾工作（关闭连接）
// 等就交给 Engine 来负责了。
// Engine是与用户交互的入口。代码位于根目录的

type Engine struct {
	db *sql.DB
}
// NewDB 开启一个数据库连接
func NewDB(driver, source string) ( *Engine, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Send a ping to make sure the database connection is alive.
	if err := db.Ping(); err != nil {
		log.Println(err)
		return nil, err
	}
	// 返回
	e := &Engine{db: db}
	log.Println("Connect database success")
	return e, nil

}

// Close 关闭数据库连接
func (engine *Engine) Close()  {
	if err := engine.db.Close(); err != nil {
		log.Println("Failed to close database")
	}
	log.Println("closed database success")
}

// NewSession 开启一个关联
func (engine *Engine) NewSession() *session.Session  {
	// 返回个session
	return session.New(engine.db)
}

