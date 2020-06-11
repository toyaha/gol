package gol

import (
	"errors"
	"fmt"
	"reflect"
)

func NewMeta(config *Config) *Meta {
	if config == nil {
		config = NewConfig()
	}

	Meta := &Meta{
		DatabaseType:              config.DatabaseType,
		NamingConventionForTable:  config.NamingConvention,
		NamingConventionForColumn: config.NamingConvention,
		Value:                     make(map[string]*MetaValue),
	}

	return Meta
}

type Meta struct {
	DatabaseType              string
	NamingConventionForTable  string
	NamingConventionForColumn string
	Value                     map[string]*MetaValue
}

func (rec *Meta) Add(schema string, tablePtr interface{}, as string) error {
	tableType := reflect.TypeOf(tablePtr).Elem()
	tableVal := reflect.ValueOf(tablePtr).Elem()

	table := func() string {
		val := tableType.Name()
		switch rec.NamingConventionForTable {
		case NamingConventionConstantCase:
			return ChangeFromPascalCaseToConstantCase(val)
		}
		return ChangeFromPascalCaseToSnakeCase(val)
	}()

	numField := tableType.NumField()
	if numField < 1 {
		return errors.New("table none field")
	}

	for i := 0; i < tableType.NumField(); i++ {
		fieldType := tableType.Field(i)
		fieldVal := tableVal.FieldByName(fieldType.Name)

		baseSchema := func() string {
			if schema != "" {
				return schema
			}
			return fieldType.Tag.Get(StructFieldTagNameSchema)
		}()

		baseTable := func() string {
			val := fieldType.Tag.Get(StructFieldTagNameTable)
			if val != "" {
				return val
			}
			return table
		}()

		baseColumn := func() string {
			val := fieldType.Tag.Get(StructFieldTagNameColumn)
			if val != "" {
				return val
			}
			val = fieldType.Name
			switch rec.NamingConventionForColumn {
			case NamingConventionConstantCase:
				return ChangeFromPascalCaseToConstantCase(val)
			}
			return ChangeFromPascalCaseToSnakeCase(val)
		}()

		addr := func() string {
			v := reflect.Indirect(fieldVal).Addr().Interface()
			return fmt.Sprintf("%p", v)
		}()

		rec.Value[addr] = NewMetaValue(baseSchema, baseTable, as, baseColumn, rec.DatabaseType)
	}

	return nil
}

func (rec *Meta) Get(any interface{}) *MetaValue {
	addr := fmt.Sprintf("%p", any)
	val, ok := rec.Value[addr]
	if !ok {
		return nil
	}
	return val
}

func (rec *Meta) GetBaseSchema(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.BaseSchema
}

func (rec *Meta) GetBaseTable(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.BaseTable
}

func (rec *Meta) GetBaseAs(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.BaseAs
}

func (rec *Meta) GetBaseColumn(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.BaseColumn
}

func (rec *Meta) GetSchema(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.Schema
}

func (rec *Meta) GetTable(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.Table
}

func (rec *Meta) GetAs(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.As
}

func (rec *Meta) GetColumn(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.Column
}

func (rec *Meta) GetTableAs(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.TableAs
}

func (rec *Meta) GetSchemaTable(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.SchemaTable
}

func (rec *Meta) GetSchemaTableColumn(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.SchemaTableColumn
}

func (rec *Meta) GetSchemaTableAs(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.SchemaTableAs
}

func (rec *Meta) GetSchemaTableAsColumn(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.SchemaTableAsColumn
}

func (rec *Meta) GetTableColumn(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.TableColumn
}

func (rec *Meta) GetTableAsColumn(any interface{}) string {
	val := rec.Get(any)
	if val == nil {
		return "_"
	}
	return val.TableAsColumn
}

func NewMetaValue(schema string, table string, as string, column string, dbType string) *MetaValue {
	data := &MetaValue{
		BaseSchema: schema,
		BaseTable:  table,
		BaseAs:     as,
		BaseColumn: column,
	}

	format := func() string {
		switch dbType {
		case DatabaseTypeMssql:
			return "\"%v\""
		case DatabaseTypeMysql:
			return "%v"
		case DatabaseTypePostgresql:
			return "\"%v\""
		}
		return "\"%v\""
	}()

	if data.BaseSchema != "" {
		data.Schema = fmt.Sprintf(format, data.BaseSchema)
	}
	data.Table = fmt.Sprintf(format, data.BaseTable)
	if data.BaseAs != "" {
		data.As = fmt.Sprintf(format, data.BaseAs)
	}
	data.Column = fmt.Sprintf(format, data.BaseColumn)

	data.SchemaTable = data.Table
	if data.BaseSchema != "" {
		data.SchemaTable = fmt.Sprintf("%v.%v", data.Schema, data.Table)
	}

	data.TableAs = data.Table
	data.SchemaTableAs = data.SchemaTable
	if data.BaseAs != "" {
		data.TableAs = data.As
		data.SchemaTableAs = data.As
	}

	data.TableColumn = fmt.Sprintf("%v.%v", data.Table, data.Column)
	data.TableAsColumn = fmt.Sprintf("%v.%v", data.TableAs, data.Column)
	data.SchemaTableColumn = fmt.Sprintf("%v.%v", data.SchemaTable, data.Column)
	data.SchemaTableAsColumn = fmt.Sprintf("%v.%v", data.SchemaTableAs, data.Column)

	return data
}

type MetaValue struct {
	BaseSchema          string
	BaseTable           string
	BaseAs              string
	BaseColumn          string
	Schema              string
	Table               string
	As                  string
	Column              string
	TableAs             string
	SchemaTable         string
	SchemaTableColumn   string
	SchemaTableAs       string
	SchemaTableAsColumn string
	TableColumn         string
	TableAsColumn       string
}

func (rec *MetaValue) IsAs() bool {
	return rec.As != ""
}
