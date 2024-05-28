package session

import (
	"bocabbage/geeorm-learn/clause"
	"bocabbage/geeorm-learn/log"
	"reflect"
)

// 接受 custom struct 对象，构建 sql 来进行 insert
func (s *Session) Insert(values ...interface{}) (int64, error) {
	// insert + value 子句构建
	recordValues := make([]interface{}, 0)
	for _, value := range values {
		table := s.Model(value).RefTable() // 获得对应的 table 映射对象（schema）
		s.clause.Set(clause.INSERT, table.Name, table.FieldNames)
		recordValues = append(recordValues, table.RecordValues(value)) // 读对象值
	}
	s.clause.Set(clause.VALUES, recordValues...)

	// 构建完整 sql 语句
	sql, vars := s.clause.Build(clause.INSERT, clause.VALUES)
	log.Infof("sql: %s\nvars: [%v]", sql, vars)
	// 执行 sql
	result, err := s.Raw(sql, vars...).Exec()
	if err != nil {
		log.Errorf("error happend for Raw: %v", err)
		return 0, err
	}
	return result.RowsAffected()
}

func (s *Session) Find(values interface{}) error {
	destSlice := reflect.Indirect(reflect.ValueOf(values))
	destType := destSlice.Type().Elem()

	// New(): 建立 destType 的新 reflect.Value 对象，作为 model 的入参来建立表信息
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
