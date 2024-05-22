package dialect

/*
	用于实现 go <-> database 之间的类型映射的接口定义
*/

import "reflect"

// 所有支持的 dialect 集合
var dialectsMap = map[string]Dialect{}

// 方言实现接口
type Dialect interface {
	// go-type -> sql-type-name
	DataTypeOf(typ reflect.Value) string
	// 检查表是否存在
	TableExistSQL(tableName string) (string, []interface{})
}

// 注册一种 dialect
func RegisterDialect(name string, dialect Dialect) {
	dialectsMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = dialectsMap[name]
	return
}
