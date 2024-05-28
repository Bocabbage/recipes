package schema

import (
	"bocabbage/geeorm-learn/dialect"
	"go/ast"
	"reflect"
)

// Field represents a column of database
type Field struct {
	Name string
	Type string // 数据库中类型的字段
	Tag  string
}

// Schema represents a table of database
type Schema struct {
	Model      interface{} // 被映射对象（用户自定义 struct）
	Name       string      // 表名
	Fields     []*Field    // 字段列表
	FieldNames []string
	fieldMap   map[string]*Field
}

func (schema *Schema) GetField(name string) *Field {
	return schema.fieldMap[name]
}

// [定义解析] 用于将任意 struct 解析到 Schema 对象
// dest: 对象的指针
func Parse(dest interface{}, d dialect.Dialect) *Schema {
	// Indirect：入参是对象指针，需要获取其指向的实例
	modelType := reflect.Indirect(reflect.ValueOf(dest)).Type()

	schema := &Schema{
		Model:    dest,
		Name:     modelType.Name(),
		fieldMap: make(map[string]*Field),
	}

	for i := 0; i < modelType.NumField(); i++ {
		p := modelType.Field(i)
		// 只有在非匿名对象 + Exported 的情况下创建field
		if !p.Anonymous && ast.IsExported(p.Name) {
			field := &Field{
				Name: p.Name,
				// dialect 实现类型转换
				Type: d.DataTypeOf(reflect.Indirect(reflect.New(p.Type))),
			}
			if v, ok := p.Tag.Lookup("geeorm"); ok {
				field.Tag = v
			}
			schema.Fields = append(schema.Fields, field)
			schema.FieldNames = append(schema.FieldNames, p.Name)
			schema.fieldMap[p.Name] = field
		}
	}
	return schema
}

// [数值解析] 用于将具体 custom 对象的值根据 Schema 解析为对应的 value slice
func (schema *Schema) RecordValues(dest interface{}) []interface{} {
	destValue := reflect.Indirect(reflect.ValueOf(dest))
	var fieldValues []interface{}

	for _, field := range schema.Fields {
		fieldValues = append(fieldValues, destValue.FieldByName(field.Name).Interface())
	}

	return fieldValues
}
