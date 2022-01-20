package schema

import (
	"go-ORM/dialect"
	"go/ast"
	"reflect"
)

// 对象(object)和表(table)的转换。给定一个任意的对象，转换为关系型数据库中的表结构。

// Schema represents a table of database
type Schema struct {
	Model interface{} // 被映射的对象 Model
	Name string // 表名
	Fields []*Field // 字段 Fields
	FieldNames []string  // 包含所有的字段名(列名)
	fieldMap map[string]*Field // fieldMap 记录字段名和 Field 的映射关系，方便之后直接使用，无需遍历 Fields
}

// Field represents a column of database
type Field struct {
	Name string // 字段名
	Type string // 类型
	Tag string  // 约束条件 Tag
}

// GetField 获取Field
func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}


// Parse 函数，将任意的对象解析为 Schema 实例
func Parse(dest interface{} , d dialect.Dialect) *Schema {
	// 获取model的类型
	// ValueOf() 获取类型的值,  reflect.Indirect() 获取指针指向的实例
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(), //  获取到结构体的名称作为表名
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ { // NumField() 获取实例的字段的个数
		p := modelType.Field(i)
		// p.Name 即字段名，p.Type 即字段类型
		// 通过 (Dialect).DataTypeOf() 转换为数据库的字段类型
		// p.Tag 即额外的约束条件。
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				// 返回字段类型
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			// Lookup 查找与标签字符串中的键相关联的值
			if v, ok := p.Tag.Lookup("orm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}


// RecordValues Values return the values of dest's member variables
// 即 u1、u2 转换为 ("Tom", 18), ("Same", 25) 这样的格式。
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}
	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}
	return fieldValues
}

type ITableName interface {
	TableName() string
}