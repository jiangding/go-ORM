package session

import (
	"go-ORM/clause"
	"reflect"
)

// Insert 构造插入语句并执行
func (s *Session) Insert(values ...interface{}) (int64,error) {
	recordValues := make([]interface{}, 0) // 空切片

	for _, value := range values {
		table := s.Model(value).RefTable()
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value))
	}
	// 多次调用 clause.Set() 构造好每一个子句。
	s.clause.Set(clause.VALUES, recordValues...)
	// build
	// 调用一次 clause.Build() 按照传入的顺序构造出最终的 SQL 语句。
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		return 0, nil
	}
	return result.RowsAffected()
}


// Find 传入一个切片指针，查询的结果保存在切片中
// var users []Use s.Find(&users);
func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	// destSlice.Type().Elem() 获取切片的单个元素的类型 destType
	destType := destSlice.Type().Elem()
	table := s.Model(reflect.New(destType).Elem().Interface()).RefTable()

	s.clause.Set(clause.SELECT, table.Name, table.FieldNames)
	sql, vars := s.clause.Build(clause.SELECT, clause.WHERE, clause.ORDERBY, clause.LIMIT)
	rows, err := s.Raw(sql, vars...).QueryRows()
	if err != nil {
		return err
	}

	for rows.Next() {
		dest := reflect.New(destType).Elem()
		var values []interface{}
		for _, name := range table.FieldNames {
			values = append(values, dest.FieldByName(name).Addr().Interface())
		}
		if err := rows.Scan(values...); err != nil {
			return err
		}
		destSlice.Set(reflect.Append(destSlice, dest))
	}
	return rows.Close()

}