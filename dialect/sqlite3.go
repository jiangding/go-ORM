package dialect

import (
	"fmt"
	"reflect"
	"time"
)

// 添加sqlite3的 dialect支持

type sqlite3 struct {}

// 右边的要实现左边的接口
var _ Dialect = (*sqlite3)(nil)

func init()  {
	RegisterDialect("sqlite3", &sqlite3{})
}

// 接口方法

func (s *sqlite3) DataTypeOf(typ reflect.Value) string {
	// 使用反射 区分类型
	// 将 Go 语言的类型映射为 SQLite 的数据类型
	switch typ.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64,reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	case reflect.Struct:
		// 结构体类型转换为时间格式
		if _, ok := typ.Interface().(time.Time); ok {
			return "datetime"
		}
	}
	// 最后报告异常
	panic(fmt.Sprintf("invalid sql type %s (%s)", typ.Type().Name(), typ.Kind()))
}

// TableExistSQL 返回了在 SQLite 中判断表 tableName 是否存在的 SQL 语句
func (s *sqlite3) TableExistSQL(tableName string) (string, []interface{}) {
	args := []interface{}{tableName}
	return "SELECT name FROM sqlite_master WHERE type='table' and name = ?", args
}