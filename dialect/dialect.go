package dialect

import "reflect"

// 不同数据库支持的数据类型也是有差异的，
// 即使功能相同，在 SQL 语句的表达上也可能有差异。
// ORM 框架往往需要兼容多种数据库，因此我们需要将差异的这一部分提取出来
// 每一种数据库分别实现，实现最大程度的复用和解耦

type Dialect interface { // 接口
	DataTypeOf(tpy reflect.Value) string // 用于将 Go 语言的类型转换为该数据库的数据类型
	TableExistSQL(tableName string) (string, []interface{})
	// 返回某个表是否存在的 SQL 语句，参数是表名(table)。
}

var dialectsMap = map[string]Dialect{}

// RegisterDialect 注册
func RegisterDialect(name string, dialect Dialect){
	dialectsMap[name] = dialect
}
// GetDialect 获取实例
func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}