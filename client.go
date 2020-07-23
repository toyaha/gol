package gol

import (
	"database/sql"
	"errors"
	"fmt"
	"reflect"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func NewClient(databaseType string, host string, port string, user string, pass string, database string, optionMap map[string]string) (*Client, error) {
	var err error

	data := &Client{
		Config: NewConfig(),
	}

	err = data.Init(databaseType, host, port, user, pass, database, optionMap)
	if err != nil {
		return nil, err
	}

	data.Meta = NewMeta(data.Config)

	return data, nil
}

type Client struct {
	Config *Config
	DB     *sql.DB
	TX     *sql.Tx
	Meta   *Meta
}

func (rec *Client) Init(databaseType string, host string, port string, user string, pass string, database string, optionMap map[string]string) error {
	var err error

	switch databaseType {
	case DatabaseTypeMssql:
		source := fmt.Sprintf(
			"server=%s;user id=%s;password=%s;port=%s;database=%s;",
			host,
			user,
			pass,
			port,
			database,
		)

		rec.DB, err = sql.Open(DatabaseTypeMssql, source)
		if err != nil {
			return err
		}

		rec.Config.InitMssql()
	case DatabaseTypeMysql:
		// [username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
		source := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s",
			user,
			pass,
			host,
			port,
			database,
		)

		rec.DB, err = sql.Open(DatabaseTypeMysql, source)
		if err != nil {
			return err
		}

		rec.Config.InitMysql()
	case DatabaseTypePostgresql:

		sslMode := PostgresqlSslModeDisable
		if v, ok := optionMap["sslMode"]; ok {
			switch v {
			case PostgresqlSslModeDisable:
				sslMode = PostgresqlSslModeDisable
			case PostgresqlSslModeRequire:
				sslMode = PostgresqlSslModeRequire
			case PostgresqlSslModeVerifyCa:
				sslMode = PostgresqlSslModeVerifyCa
			case PostgresqlSslModeVerifyFull:
				sslMode = PostgresqlSslModeVerifyFull
			}
		}

		source := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			host,
			port,
			user,
			pass,
			database,
			sslMode,
		)

		rec.DB, err = sql.Open(DatabaseTypePostgresql, source)
		if err != nil {
			return err
		}

		rec.Config.InitPostgresql()
	default:
		return errors.New("unknown databaseType")
	}

	err = rec.DB.Ping()
	if err != nil {
		return err
	}

	return nil
}

func (rec *Client) AddMeta(tablePtrList ...interface{}) {
	for _, val := range tablePtrList {
		rec.Meta.Add(val, false)
	}
}

func (rec *Client) GetMeta(any interface{}) *MetaValue {
	return rec.Meta.Get(any)
}

func (rec *Client) GetBaseSchema(any interface{}) string {
	return rec.Meta.GetBaseSchema(any)
}

func (rec *Client) GetBaseTable(any interface{}) string {
	return rec.Meta.GetBaseTable(any)
}

func (rec *Client) GetBaseAs(any interface{}) string {
	return rec.Meta.GetBaseAs(any)
}

func (rec *Client) GetBaseColumn(any interface{}) string {
	return rec.Meta.GetBaseColumn(any)
}

func (rec *Client) GetSchema(any interface{}) string {
	return rec.Meta.GetSchema(any)
}

func (rec *Client) GetTable(any interface{}) string {
	return rec.Meta.GetTable(any)
}

func (rec *Client) GetAs(any interface{}) string {
	return rec.Meta.GetAs(any)
}

func (rec *Client) GetColumn(any interface{}) string {
	return rec.Meta.GetColumn(any)
}

func (rec *Client) GetTableAs(any interface{}) string {
	return rec.Meta.GetTableAs(any)
}

func (rec *Client) GetSchemaTable(any interface{}) string {
	return rec.Meta.GetSchemaTable(any)
}

func (rec *Client) GetSchemaTableColumn(any interface{}) string {
	return rec.Meta.GetSchemaTableColumn(any)
}

func (rec *Client) GetSchemaTableAs(any interface{}) string {
	return rec.Meta.GetSchemaTableAs(any)
}

func (rec *Client) GetSchemaTableAsColumn(any interface{}) string {
	return rec.Meta.GetSchemaTableAsColumn(any)
}

func (rec *Client) GetTableColumn(any interface{}) string {
	return rec.Meta.GetTableColumn(any)
}

func (rec *Client) GetTableAsColumn(any interface{}) string {
	return rec.Meta.GetTableAsColumn(any)
}

func (rec *Client) NewQuery(tablePtrList ...interface{}) *Query {
	query := NewQuery(rec.Config)
	query.SetClient(rec)
	if len(tablePtrList) > 0 {
		for _, val := range tablePtrList {
			query.AddMeta(val)
		}
		query.SetTable(tablePtrList[0])
	}
	return query
}

func (rec *Client) Exec(query string, valueList ...interface{}) (sql.Result, error) {
	var err error
	var result sql.Result

	if rec.DB == nil && rec.TX == nil {
		return nil, errors.New("database does not exist")
	}

	switch rec.Config.DatabaseType {
	case DatabaseTypeMssql:
		query = ChangeQueryForMssql(query)
		valueList = ChangeValueListForMssql(valueList...)
	case DatabaseTypePostgresql:
		query = ChangeQueryForPostgresql(query)
	}

	if rec.Config.Log {
		fmt.Printf("query: %v\n", query)
		fmt.Printf("value: %v\n", valueList)
	}

	if rec.TX != nil {
		result, err = rec.TX.Exec(query, valueList...)
	} else {
		result, err = rec.DB.Exec(query, valueList...)
	}
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rec *Client) Query(query string, valueList ...interface{}) (*sql.Rows, error) {
	var err error
	var rows *sql.Rows

	if rec.DB == nil && rec.TX == nil {
		return nil, errors.New("database is null")
	}

	switch rec.Config.DatabaseType {
	case DatabaseTypeMssql:
		query = ChangeQueryForMssql(query)
		valueList = ChangeValueListForMssql(valueList...)
	case DatabaseTypePostgresql:
		query = ChangeQueryForPostgresql(query)
	}

	if rec.Config.Log {
		fmt.Printf("query: %v\n", query)
		fmt.Printf("value: %v\n", valueList)
	}

	if rec.TX != nil {
		rows, err = rec.TX.Query(query, valueList...)
	} else {
		rows, err = rec.DB.Query(query, valueList...)
	}
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (rec *Client) QueryRow(query string, valueList ...interface{}) *Row {
	var row = NewRow()

	if rec.DB == nil && rec.TX == nil {
		return row
	}

	switch rec.Config.DatabaseType {
	case DatabaseTypeMssql:
		query = ChangeQueryForMssql(query)
		valueList = ChangeValueListForMssql(valueList...)
	case DatabaseTypePostgresql:
		query = ChangeQueryForPostgresql(query)
	}

	if rec.Config.Log {
		fmt.Printf("query: %v\n", query)
		fmt.Printf("value: %v\n", valueList)
	}

	if rec.TX != nil {
		row.Row = rec.TX.QueryRow(query, valueList...)
	} else {
		row.Row = rec.DB.QueryRow(query, valueList...)
	}

	return row
}

func (rec *Client) Find(dest interface{}, query string, valueList ...interface{}) error {
	rows, err := rec.Query(query, valueList...)
	if err != nil {
		return err
	}

	err = rec.ExtractRows(dest, rows)
	if err != nil {
		return err
	}

	return nil
}

func (rec *Client) FindRow(dest interface{}, query string, valueList ...interface{}) error {
	rows, err := rec.Query(query, valueList...)
	if err != nil {
		return err
	}

	err = rec.ExtractRow(dest, rows)
	if err != nil {
		return err
	}

	return nil
}

func (rec *Client) ExtractRows(dest interface{}, rows *sql.Rows) error {
	if rows == nil {
		return errors.New("rows is nil")
	}
	defer func() {
		_ = rows.Close()
	}()

	columnList, err := rows.Columns()
	if err != nil {
		return err
	}

	destValue := reflect.ValueOf(dest)
	var destValueType string
	destDirect := reflect.Indirect(destValue)
	var base reflect.Type
	{
		ty := destValue.Type()
		if ty.Kind() != reflect.Ptr {
			return errors.New("dest is not pointer. Should be type *[]struct or *[]*struct or *[]map[string]interface {}")
		}
		ty = ty.Elem()
		base = ty.Elem()
		switch ty.Kind() {
		case reflect.Slice:
			if base.Kind() == reflect.Ptr {
				base = base.Elem()
				if base.Kind() != reflect.Struct {
					return errors.New("dest is not pointer. Should be type *[]struct or *[]*struct or *[]map[string]interface {}")
				}
				destValueType = "*[]*struct"
			} else {
				if base.Kind() == reflect.Struct {
					destValueType = "*[]struct"
				} else {
					destValueType = destValue.Type().String()
					switch destValueType {
					case "*[]map[string]interface {}":
					default:
						return errors.New("dest is not pointer. Should be type *[]struct or *[]*struct or *[]map[string]interface {}")
					}
				}
			}
		case reflect.Map:
			destValueType = destValue.Type().String()
		default:
			return errors.New("dest is not pointer. Should be type *[]struct or *[]*struct or *[]map[string]interface {}")
		}
	}

	switch destValueType {
	case "*[]*struct", "*[]struct":
		var tagIndexMap map[string][]int
		{
			tagIndexMap, err = makeTagIndexMap(base, StructFieldTagNameColumn)
			if err != nil {
				return err
			}

			if len(tagIndexMap) != len(columnList) {
				tagIndexMapFlag := true
				for _, column := range columnList {
					_, ok := tagIndexMap[column]
					if !ok {
						tagIndexMapFlag = true
						break
					}
					tagIndexMapFlag = false
				}
				if tagIndexMapFlag {
					baseValue := reflect.New(base)
					val := reflect.Indirect(baseValue)
					if val.NumField() != len(columnList) {
						return errors.New("length does not match")
					}
					for key, column := range columnList {
						tagIndexMap[column] = []int{key}
					}
				}
			}
		}

		scanList := make([]interface{}, len(columnList))
		if destValueType == "*[]*struct" {
			for rows.Next() {
				baseValue := reflect.New(base)
				val := reflect.Indirect(baseValue)

				for key, column := range columnList {
					indexList, ok := tagIndexMap[column]
					if !ok {
						return errors.New("column not exist")
					}

					field := val
					for _, index := range indexList {
						field = field.Field(index)
					}

					scanList[key] = field.Addr().Interface()
				}

				err = rows.Scan(scanList...)
				if err != nil {
					return err
				}

				destDirect.Set(reflect.Append(destDirect, val.Addr()))
			}
		} else {
			for rows.Next() {
				baseValue := reflect.New(base)
				val := reflect.Indirect(baseValue)

				for key, column := range columnList {
					indexList, ok := tagIndexMap[column]
					if !ok {
						return errors.New("column not exist")
					}

					field := val
					for _, index := range indexList {
						field = field.Field(index)
					}

					scanList[key] = field.Addr().Interface()
				}

				err = rows.Scan(scanList...)
				if err != nil {
					return err
				}

				destDirect.Set(reflect.Append(destDirect, val))
			}
		}
	case "*[]map[string]interface {}":
		columnChangeMap := make(map[string]string)
		for _, val := range columnList {
			columnChangeMap[val] = val
		}

		var valList = make([]interface{}, len(columnList))
		var scanList = make([]interface{}, len(columnList))
		for key := range columnList {
			scanList[key] = &valList[key]
		}

		for rows.Next() {
			err = rows.Scan(scanList...)
			if err != nil {
				return err
			}

			scanMap := make(map[string]interface{})
			for key, column := range columnList {
				name := columnChangeMap[column]
				scanMap[name] = valList[key]
			}

			scanMapType := reflect.ValueOf(scanMap)
			destDirect.Set(reflect.Append(destDirect, scanMapType))
		}
	default:
		return errors.New("type *[]struct or *[]*struct or *[]map[string]interface{}")
	}

	return nil
}

func (rec *Client) ExtractRow(dest interface{}, rows *sql.Rows) error {
	if rows == nil {
		return errors.New("rows not exist")
	}
	defer func() {
		_ = rows.Close()
	}()

	for !rows.Next() {
		return errors.New(fmt.Sprintf("number of lines is not 1"))
	}

	columnList, err := rows.Columns()
	if err != nil {
		return err
	}

	var destType = reflect.TypeOf(dest)
	var destValue = reflect.ValueOf(dest)
	if destType.Kind() != reflect.Ptr {
		return errors.New("type must be type *struct or *map[string]interface {}")
	}

	var originType = destType.Elem()
	var originValue = destValue.Elem()

	var destTypeName string
	if originType.Kind() == reflect.Struct {
		destTypeName = "*struct"
	} else {
		return errors.New("type must be type *struct or *map[string]interface {}")
	}

	switch destTypeName {
	case "*struct":
		var columnIndexList [][]int
		{
			tagIndexMap, err := makeTagIndexMap(originType, StructFieldTagNameColumn)
			if err != nil {
				return err
			}

			for _, val := range columnList {
				indexList, ok := tagIndexMap[val]
				if !ok {
					return errors.New(fmt.Sprintf("not found column %v", val))
				}

				columnIndexList = append(columnIndexList, indexList)
			}
		}

		scanList := make([]interface{}, len(columnList))

		direct := reflect.Indirect(originValue)
		for key := range columnList {
			field := direct
			for _, v := range columnIndexList[key] {
				field = field.Field(v)
			}

			scanList[key] = field.Addr().Interface()
		}

		err = rows.Scan(scanList...)
		if err != nil {
			return err
		}
	default:
		return errors.New("type must be type *struct or *map[string]interface {}")
	}

	if rows.Next() {
		return errors.New(fmt.Sprintf("Number of lines is not 1"))
	}

	if err := rows.Err(); err != nil {
		return err
	}

	if err := rows.Close(); err != nil {
		return err
	}

	return nil
}

func (rec *Client) Close() error {
	if rec.Config.Test {
		return nil
	}

	if rec.DB == nil {
		return nil
	}

	err := rec.DB.Close()
	if err != nil {
		return err
	}

	rec.DB = nil
	rec.TX = nil

	return nil
}

func (rec *Client) Begin() (*Client, error) {
	if rec.Config.Test {
		return rec, nil
	}

	if rec.TX != nil {
		return nil, errors.New("exist transaction")
	}

	tx, err := rec.DB.Begin()
	if err != nil {
		return nil, err
	}

	queryData := *rec
	queryData.TX = tx

	return &queryData, nil
}

func (rec *Client) Commit() error {
	var err error

	if rec.Config.Test {
		return nil
	}

	if rec.TX == nil {
		return nil
	}

	err = rec.TX.Commit()
	if err != nil {
		return err
	}

	rec.TX = nil

	return nil
}

func (rec *Client) Rollback() error {
	var err error

	if rec.Config.Test {
		return nil
	}

	if rec.TX == nil {
		return nil
	}

	err = rec.TX.Rollback()
	if err != nil {
		return err
	}

	rec.TX = nil

	return nil
}

func (rec *Client) TestBegin() {
	if rec.DB == nil {
		return
	}

	tx, _ := rec.DB.Begin()
	rec.Config.Test = true
	rec.TX = tx
}

func (rec *Client) TestEnd() {
	var err error

	err = rec.TX.Rollback()
	if err != nil {
		panic(err)
	}

	rec.TX = nil

	_ = rec.DB.Close()
	rec.DB = nil
}
