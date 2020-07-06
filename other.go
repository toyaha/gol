package gol

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func ChangeQueryForMssql(query string) string {
	num := 1
	for strings.Contains(query, "?") {
		query = strings.Replace(query, "?", fmt.Sprintf("@v%v", num), 1)
		num++
	}
	return query
}

func ChangeValueListForMssql(valueList ...interface{}) []interface{} {
	var valList []interface{}
	count := 1
	for _, val := range valueList {
		if fmt.Sprintf("%T", val) != "sql.NamedArg" {
			valList = append(valList, sql.NamedArg{Name: fmt.Sprintf("v%v", count), Value: val})
			count++
		} else {
			valList = append(valList, val)
		}
	}
	return valList
}

func ChangeQueryForPostgresql(query string) string {
	num := 1
	for strings.Contains(query, "?") {
		query = strings.Replace(query, "?", fmt.Sprintf("$%v", num), 1)
		num++
	}
	return query
}

func buildQueryWhere(meta *Meta, errPrefix string, useAs bool, mode int, dataList ...interface{}) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	var fnGetName func(*MetaValue) string
	if useAs {
		fnGetName = func(meta *MetaValue) string {
			return meta.SchemaTableAsColumn
		}
	} else {
		fnGetName = func(meta *MetaValue) string {
			return meta.SchemaTableColumn
		}
	}

	switch mode {
	case QueryModeIs:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		indexLast := len(dataList) - 1

		for i := 0; i < indexLast; i++ {
			val := dataList[i]
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v%v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v%v", query, val)
			}
		}

		val := dataList[indexLast]
		if isValueNullStruct(val) {
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v = %v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v = %v", query, PlaceHolder)
				valueList = append(valueList, val)
			}
		} else {
			query = fmt.Sprintf("%v IS NULL", query)
		}
	case QueryModeIsNot:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		indexLast := len(dataList) - 1

		for i := 0; i < indexLast; i++ {
			val := dataList[i]
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v%v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v%v", query, val)
			}
		}

		val := dataList[indexLast]
		if isValueNullStruct(val) {
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v != %v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v != %v", query, PlaceHolder)
				valueList = append(valueList, val)
			}
		} else {
			query = fmt.Sprintf("%v IS NOT NULL", query)
		}
	case QueryModeLike:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		indexLast := len(dataList) - 1

		for i := 0; i < indexLast; i++ {
			val := dataList[i]
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v%v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v%v", query, val)
			}
		}
		query = fmt.Sprintf("%v LIKE %v", query, PlaceHolder)

		valueList = append(valueList, dataList[indexLast])
	case QueryModeLikeNot:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		indexLast := len(dataList) - 1

		for i := 0; i < indexLast; i++ {
			val := dataList[i]
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v%v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v%v", query, val)
			}
		}
		query = fmt.Sprintf("%v NOT LIKE %v", query, PlaceHolder)

		valueList = append(valueList, dataList[indexLast])
	case QueryModeIn:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		var strList []string
		for _, val := range dataList {
			if ty := fmt.Sprintf("%T", val); ty == "[]interface {}" {
				for _, value := range val.([]interface{}) {
					strList = append(strList, PlaceHolder)
					valueList = append(valueList, value)
				}
			} else {
				if data := meta.Get(val); data != nil {
					query = fmt.Sprintf("%v%v", query, fnGetName(data))
				} else {
					query = fmt.Sprintf("%v%v", query, val)
				}
			}
		}
		query = fmt.Sprintf("%v IN (%v)", query, strings.Join(strList, ", "))
	case QueryModeInNot:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		var strList []string
		for _, val := range dataList {
			if ty := fmt.Sprintf("%T", val); ty == "[]interface {}" {
				for _, value := range val.([]interface{}) {
					strList = append(strList, PlaceHolder)
					valueList = append(valueList, value)
				}
			} else {
				if data := meta.Get(val); data != nil {
					query = fmt.Sprintf("%v%v", query, fnGetName(data))
				} else {
					query = fmt.Sprintf("%v%v", query, val)
				}
			}
		}
		query = fmt.Sprintf("%v NOT IN (%v)", query, strings.Join(strList, ", "))
	case QueryModeGt:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		indexLast := len(dataList) - 1

		for i := 0; i < indexLast; i++ {
			val := dataList[i]
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v%v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v%v", query, val)
			}
		}

		val := dataList[indexLast]
		if data := meta.Get(val); data != nil {
			query = fmt.Sprintf("%v > %v", query, fnGetName(data))
		} else {
			query = fmt.Sprintf("%v > %v", query, PlaceHolder)
			valueList = append(valueList, val)
		}
	case QueryModeGte:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		indexLast := len(dataList) - 1

		for i := 0; i < indexLast; i++ {
			val := dataList[i]
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v%v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v%v", query, val)
			}
		}

		val := dataList[indexLast]
		if data := meta.Get(val); data != nil {
			query = fmt.Sprintf("%v >= %v", query, fnGetName(data))
		} else {
			query = fmt.Sprintf("%v >= %v", query, PlaceHolder)
			valueList = append(valueList, val)
		}
	case QueryModeLt:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		indexLast := len(dataList) - 1

		for i := 0; i < indexLast; i++ {
			val := dataList[i]
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v%v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v%v", query, val)
			}
		}

		val := dataList[indexLast]
		if data := meta.Get(val); data != nil {
			query = fmt.Sprintf("%v < %v", query, fnGetName(data))
		} else {
			query = fmt.Sprintf("%v < %v", query, PlaceHolder)
			valueList = append(valueList, val)
		}
	case QueryModeLte:
		if len(dataList) <= 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 1", errPrefix))
		}

		indexLast := len(dataList) - 1

		for i := 0; i < indexLast; i++ {
			val := dataList[i]
			if data := meta.Get(val); data != nil {
				query = fmt.Sprintf("%v%v", query, fnGetName(data))
			} else {
				query = fmt.Sprintf("%v%v", query, val)
			}
		}

		val := dataList[indexLast]
		if data := meta.Get(val); data != nil {
			query = fmt.Sprintf("%v <= %v", query, fnGetName(data))
		} else {
			query = fmt.Sprintf("%v <= %v", query, PlaceHolder)
			valueList = append(valueList, val)
		}
	case QueryModeNest:
		query = "("
	case QueryModeNestClose:
		query = ")"
	default:
		if len(dataList) < 1 {
			return "", nil, errors.New(fmt.Sprintf("%v length must be greater than 0", errPrefix))
		}
		for _, val := range dataList {
			if ty := fmt.Sprintf("%T", val); ty == "[]interface {}" {
				valueList = append(valueList, val.([]interface{})...)
			} else {
				if data := meta.Get(val); data != nil {
					query = fmt.Sprintf("%v%v", query, fnGetName(data))
				} else {
					query = fmt.Sprintf("%v%v", query, val)
				}
			}
		}
	}

	return query, valueList, nil
}

func isValueNullStruct(val interface{}) bool {
	if val == nil {
		return false
	}

	existFlag := true
	switch val.(type) {
	case NullBool:
		existFlag = val.(NullBool).Valid
	case *NullBool:
		existFlag = val.(*NullBool).Valid
	case NullFloat64:
		existFlag = val.(NullFloat64).Valid
	case *NullFloat64:
		existFlag = val.(*NullFloat64).Valid
	case NullInt32:
		existFlag = val.(NullInt32).Valid
	case *NullInt32:
		existFlag = val.(*NullInt32).Valid
	case NullInt64:
		existFlag = val.(NullInt64).Valid
	case *NullInt64:
		existFlag = val.(*NullInt64).Valid
	case NullString:
		existFlag = val.(NullString).Valid
	case *NullString:
		existFlag = val.(*NullString).Valid
	case NullTime:
		existFlag = val.(NullTime).Valid
	case *NullTime:
		existFlag = val.(*NullTime).Valid
	}
	return existFlag
}

func makeTagIndexMap(value reflect.Type, tagName string) (map[string][]int, error) {
	tagIndexMap := make(map[string][]int, 0)

	if value.Kind() != reflect.Struct {
		return tagIndexMap, errors.New("not struct")
	}

	tagIndexMap = makeTagIndexMapRe(tagIndexMap, []int{}, value, tagName)

	return tagIndexMap, nil
}

func makeTagIndexMapRe(tagIndexMap map[string][]int, indexList []int, value reflect.Type, tagName string) map[string][]int {
	if value.Kind() != reflect.Struct {
		return tagIndexMap
	}

	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		tag := field.Tag.Get(tagName)
		indexNextList := append(indexList, i)
		if tag == "" {
			tagIndexMap = makeTagIndexMapRe(tagIndexMap, indexNextList, field.Type, tagName)
		} else {
			tagIndexMap[tag] = indexNextList
		}
	}

	return tagIndexMap
}

func ChangeFromPascalCaseToConstantCase(value string) string {
	str := ""
	match := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	valueList := strings.Split(value, "")
	limit := len(valueList)
	for i := 0; i < limit; i++ {
		val := valueList[i]
		if strings.Contains(match, val) {
			val = strings.ToUpper(val)
			if i != 0 {
				val = fmt.Sprintf("%s%s", "_", val)
			}
		} else {
			val = strings.ToUpper(val)
		}
		str = fmt.Sprintf("%s%s", str, val)
	}

	return str
}

func ChangeFromPascalCaseToSnakeCase(value string) string {
	str := ""
	match := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	valueList := strings.Split(value, "")
	limit := len(valueList)
	for i := 0; i < limit; i++ {
		val := valueList[i]
		if strings.Contains(match, val) {
			val = strings.ToLower(val)
			if i != 0 {
				val = fmt.Sprintf("%s%s", "_", val)
			}
		}
		str = fmt.Sprintf("%s%s", str, val)
	}

	return str
}
