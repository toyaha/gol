package gol

import (
	"errors"
	"fmt"
	"strings"
)

func NewQueryValue(config *Config) *QueryValue {
	if config == nil {
		config = NewConfig()
	}

	queryValue := &QueryValue{
		Config: config,
		Meta:   NewMeta(config),
	}

	return queryValue
}

type QueryValue struct {
	Config           *Config
	ErrorList        []error
	Meta             *Meta
	Table            *QueryTable
	FromList         []*QueryFrom
	JoinList         []*QueryJoin
	JoinWhereList    []*QueryJoinWhere
	ValuesColumnList []*QueryValuesColumn
	ValuesList       []*QueryValues
	ConfLictList     []*QueryConflict
	SetList          []*QuerySet
	SelectList       []*QuerySelect
	WhereList        []*QueryWhere
	GroupByList      []*QueryGroupBy
	HavingList       []*QueryHaving
	OrderByList      []*QueryOrderBy
	Limit            int
	Offset           int
}

func (rec *QueryValue) GetInsertQuery() (string, []interface{}, error) {
	var query = "INSERT"
	var valueList []interface{}

	err := func() error {
		{
			str, err := rec.BuildTable()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("table not exist")
			}

			query = fmt.Sprintf("%v INTO %v", query, str)
		}

		{
			str, err := rec.BuildValuesColumn()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, valList, err := rec.BuildValues()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("values not exist")
			}

			query = fmt.Sprintf("%v VALUES %v", query, str)
			valueList = append(valueList, valList...)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetInsertDoNothingQuery() (string, []interface{}, error) {
	var query = "INSERT"
	var valueList []interface{}

	err := func() error {
		{
			str, err := rec.BuildTable()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("table not exist")
			}

			query = fmt.Sprintf("%v INTO %v", query, str)
		}

		{
			str, err := rec.BuildValuesColumn()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, valList, err := rec.BuildValues()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("values not exist")
			}

			query = fmt.Sprintf("%v VALUES %v", query, str)
			valueList = append(valueList, valList...)
		}

		query = fmt.Sprintf("%v ON CONFLICT DO NOTHING", query)

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetInsertDoUpdateQuery() (string, []interface{}, error) {
	var query = "INSERT"
	var valueList []interface{}

	err := func() error {
		{
			str, err := rec.BuildTable()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("table not exist")
			}

			query = fmt.Sprintf("%v INTO %v", query, str)
		}

		{
			str, err := rec.BuildValuesColumn()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, valList, err := rec.BuildValues()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("values not exist")
			}

			query = fmt.Sprintf("%v VALUES %v", query, str)
			valueList = append(valueList, valList...)
		}

		{
			str, valList, err := rec.BuildConflict()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("conflict not exist")
			}

			query = fmt.Sprintf("%v ON CONFLICT (%v)", query, str)
			valueList = append(valueList, valList...)
		}

		{
			str, valList, err := rec.BuildSet()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("set not exist")
			}

			query = fmt.Sprintf("%v DO UPDATE SET %v", query, str)
			valueList = append(valueList, valList...)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetInsertIgnoreQuery() (string, []interface{}, error) {
	var query = "INSERT IGNORE"
	var valueList []interface{}

	err := func() error {
		{
			str, err := rec.BuildTable()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("table not exist")
			}

			query = fmt.Sprintf("%v INTO %v", query, str)
		}

		{
			str, err := rec.BuildValuesColumn()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, valList, err := rec.BuildValues()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("values not exist")
			}

			query = fmt.Sprintf("%v VALUES %v", query, str)
			valueList = append(valueList, valList...)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetInsertOnDuplicateKeyUpdateQuery() (string, []interface{}, error) {
	var query = "INSERT"
	var valueList []interface{}

	err := func() error {
		{
			str, err := rec.BuildTable()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("table not exist")
			}

			query = fmt.Sprintf("%v INTO %v", query, str)
		}

		{
			str, err := rec.BuildValuesColumn()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, valList, err := rec.BuildValues()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("values not exist")
			}

			query = fmt.Sprintf("%v VALUES %v", query, str)
			valueList = append(valueList, valList...)
		}

		{
			str, valList, err := rec.BuildSet()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("set not exist")
			}

			query = fmt.Sprintf("%v ON DUPLICATE KEY UPDATE %v", query, str)
			valueList = append(valueList, valList...)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetInsertSelectUnionQuery() (string, []interface{}, error) {
	var query = "INSERT"
	var valueList []interface{}

	err := func() error {
		{
			str, err := rec.BuildTable()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("table not exist")
			}

			query = fmt.Sprintf("%v INTO %v", query, str)
		}

		{
			str, err := rec.BuildValuesColumn()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, valList, err := rec.BuildValuesSelectUnion()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("values not exist")
			}

			query = fmt.Sprintf("%v %v", query, str)
			valueList = append(valueList, valList...)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetUpdateQuery() (string, []interface{}, error) {
	var query = "UPDATE"
	var valueList []interface{}

	err := func() error {
		{
			str, err := rec.BuildTable()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("table not exist")
			}

			query = fmt.Sprintf("%v %v", query, str)
		}

		{
			str, valList, err := rec.BuildSet()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("set not exist")
			}

			query = fmt.Sprintf("%v SET %v", query, str)
			valueList = append(valueList, valList...)
		}

		{
			str, valList, err := rec.BuildWhere()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
				valueList = append(valueList, valList...)
			}
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetDeleteQuery() (string, []interface{}, error) {
	var query = "DELETE"
	var valueList []interface{}

	err := func() error {
		{
			str, err := rec.BuildTable()
			if err != nil {
				return err
			}
			if str == "" {
				return errors.New("table not exist")
			}
			query = fmt.Sprintf("%v FROM %v", query, str)
		}

		{
			str, valList, err := rec.BuildWhere()
			if err != nil {
				return err
			}
			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
				valueList = append(valueList, valList...)
			}
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetTruncateQuery() (string, []interface{}, error) {
	var query = "TRUNCATE TABLE"
	var valueList []interface{}

	err := func() error {
		str, err := rec.BuildTable()
		if err != nil {
			return err
		}

		if str == "" {
			return errors.New("table not exist")
		}

		query = fmt.Sprintf("%v %v", query, str)

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetTruncateRestartIdentityQuery() (string, []interface{}, error) {
	query, valueList, err := rec.GetTruncateQuery()
	if err != nil {
		return query, valueList, err
	}

	query = fmt.Sprintf("%v RESTART IDENTITY", query)

	return query, valueList, nil
}

func (rec *QueryValue) GetSelectQuery() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		if rec.Table == nil && len(rec.FromList) < 1 {
			return errors.New("table not exist")
		}

		fnBuildSelect := rec.BuildSelect
		fnBuildTable := rec.BuildTable
		fnBuildWhere := rec.BuildWhere
		fnBuildGroupBy := rec.BuildGroupBy
		fnBuildHaving := rec.BuildHaving
		fnBuildOrderBy := rec.BuildOrderBy
		if len(rec.FromList) > 1 || len(rec.JoinList) > 0 {
			fnBuildSelect = rec.BuildSelectUseAs
			fnBuildTable = rec.BuildTableUseAs
			fnBuildWhere = rec.BuildWhereUseAs
			fnBuildGroupBy = rec.BuildGroupByUseAs
			fnBuildHaving = rec.BuildHavingUseAs
			fnBuildOrderBy = rec.BuildOrderByUseAs
		}

		{
			str, valList, err := fnBuildSelect()
			if err != nil {
				return err
			}

			if str == "" {
				return errors.New("select not exist")
			}

			query = str
			valueList = append(valueList, valList...)
		}

		{
			var strList []string
			{
				str, err := fnBuildTable()
				if err != nil {
					return err
				}

				if str != "" {
					strList = append(strList, str)
				}
			}

			{
				str, valList, err := rec.BuildFromUseAs()
				if err != nil {
					return err
				}

				if str != "" {
					strList = append(strList, str)
				}

				valueList = append(valueList, valList...)
			}

			if len(strList) < 1 {
				return errors.New("from not exist")
			}

			str := strings.Join(strList, ", ")
			query = fmt.Sprintf("%v FROM %v", query, str)
		}

		{
			str, valList, err := rec.BuildJoinUseAs()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
				valueList = append(valueList, valList...)
			}
		}

		{
			str, valList, err := fnBuildWhere()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
				valueList = append(valueList, valList...)
			}
		}

		{
			str, err := fnBuildGroupBy()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, valList, err := fnBuildHaving()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
				valueList = append(valueList, valList...)
			}
		}

		{
			str, err := fnBuildOrderBy()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, err := rec.BuildLimit()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, err := rec.BuildOffset()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetSelectCountQuery() (string, []interface{}, error) {
	var query = "SELECT count(*) as count"
	var valueList []interface{}

	err := func() error {
		if rec.Table == nil && len(rec.FromList) < 1 {
			return errors.New("table not exist")
		}

		fnBuildTable := rec.BuildTable
		fnBuildWhere := rec.BuildWhere
		fnBuildGroupBy := rec.BuildGroupBy
		fnBuildHaving := rec.BuildHaving
		fnBuildOrderBy := rec.BuildOrderBy
		if len(rec.FromList) > 1 || len(rec.JoinList) > 0 {
			fnBuildTable = rec.BuildTableUseAs
			fnBuildWhere = rec.BuildWhereUseAs
			fnBuildGroupBy = rec.BuildGroupByUseAs
			fnBuildHaving = rec.BuildHavingUseAs
			fnBuildOrderBy = rec.BuildOrderByUseAs
		}

		{
			var strList []string
			{
				str, err := fnBuildTable()
				if err != nil {
					return err
				}

				if str != "" {
					strList = append(strList, str)
				}
			}

			{
				str, valList, err := rec.BuildFromUseAs()
				if err != nil {
					return err
				}

				if str != "" {
					strList = append(strList, str)
				}

				valueList = append(valueList, valList...)
			}

			if len(strList) < 1 {
				return errors.New("from not exist")
			}

			str := strings.Join(strList, ", ")
			query = fmt.Sprintf("%v FROM %v", query, str)
		}

		{
			str, valList, err := rec.BuildJoinUseAs()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
				valueList = append(valueList, valList...)
			}
		}

		{
			str, valList, err := fnBuildWhere()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
				valueList = append(valueList, valList...)
			}
		}

		{
			str, err := fnBuildGroupBy()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, valList, err := fnBuildHaving()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
				valueList = append(valueList, valList...)
			}
		}

		{
			str, err := fnBuildOrderBy()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, err := rec.BuildLimit()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		{
			str, err := rec.BuildOffset()
			if err != nil {
				return err
			}

			if str != "" {
				query = fmt.Sprintf("%v %v", query, str)
			}
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) GetValuesCount() int {
	return len(rec.ValuesList)
}

func (rec *QueryValue) BuildTable() (string, error) {
	if rec.Table == nil {
		return "", errors.New("table not exist")
	}
	return rec.Table.Build(rec.Meta)
}

func (rec *QueryValue) BuildTableUseAs() (string, error) {
	if rec.Table == nil {
		return "", errors.New("table not exist")
	}
	return rec.Table.BuildUseAs(rec.Meta)
}

func (rec *QueryValue) BuildFromUseAs() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		for _, val := range rec.FromList {
			str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
			valueList = append(valueList, valList...)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildJoinUseAs() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		if len(rec.JoinList) < 1 {
			return nil
		}

		var strList []string
		for _, val := range rec.JoinList {
			str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return err
			}

			valueList = append(valueList, valList...)

			where, valList, err := rec.BuildJoinWhereUseAs(val.TablePtr)
			if err != nil {
				return err
			}

			str = fmt.Sprintf("%v %v", str, where)
			strList = append(strList, str)
			valueList = append(valueList, valList...)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, " ")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildJoinWhereUseAs(tablePtr interface{}) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		var flg = false
		for _, val := range rec.JoinWhereList {
			if tablePtr != val.TablePtr {
				continue
			}

			prefix, str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return err
			}

			if flg && val.Mode != QueryModeNestClose {
				strList = append(strList, prefix, str)
			} else {
				strList = append(strList, str)
				flg = true
			}

			valueList = append(valueList, valList...)

			if val.Mode == QueryModeNest {
				flg = false
			}
		}

		if len(strList) < 1 {
			return errors.New("joinWhere not exist")
		}

		query = strings.Join(strList, " ")

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildValuesColumn() (string, error) {
	var query string

	err := func() error {
		var strList []string
		for _, val := range rec.ValuesColumnList {
			str, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("(%v)", query)
		}

		return nil
	}()

	return query, err
}

func (rec *QueryValue) BuildValues() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		for _, val := range rec.ValuesList {
			str, valList, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			str = fmt.Sprintf("(%v)", str)
			strList = append(strList, str)
			valueList = append(valueList, valList...)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildValuesSelectUnion() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		for _, val := range rec.ValuesList {
			str, valList, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			str = fmt.Sprintf("SELECT %v", str)
			strList = append(strList, str)
			valueList = append(valueList, valList...)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, " UNION ")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildConflict() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		for _, val := range rec.ConfLictList {
			str, valueList, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
			valueList = append(valueList, valueList...)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildSet() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		for _, val := range rec.SetList {
			str, value, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
			valueList = append(valueList, value)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildSelect() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		for _, val := range rec.SelectList {
			str, valList, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
			valueList = append(valueList, valList...)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("SELECT %v", query)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildSelectUseAs() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		for _, val := range rec.SelectList {
			str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
			valueList = append(valueList, valList...)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("SELECT %v", query)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildWhere() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		var flg = false
		for _, val := range rec.WhereList {
			prefix, str, valList, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			if flg && val.Mode != QueryModeNestClose {
				strList = append(strList, prefix, str)
			} else {
				strList = append(strList, str)
				flg = true
			}

			valueList = append(valueList, valList...)

			if val.Mode == QueryModeNest {
				flg = false
			}
		}

		if len(strList) > 0 {
			query = strings.Join(strList, " ")
			query = fmt.Sprintf("WHERE %v", query)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildWhereUseAs() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		var flg = false
		for _, val := range rec.WhereList {
			prefix, str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return err
			}

			if flg && val.Mode != QueryModeNestClose {
				strList = append(strList, prefix, str)
			} else {
				strList = append(strList, str)
				flg = true
			}

			valueList = append(valueList, valList...)

			if val.Mode == QueryModeNest {
				flg = false
			}
		}

		if len(strList) > 0 {
			query = strings.Join(strList, " ")
			query = fmt.Sprintf("WHERE %v", query)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildGroupBy() (string, error) {
	var query string

	err := func() error {
		var strList []string
		for _, val := range rec.GroupByList {
			str, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("GROUP BY %v", query)
		}

		return nil
	}()

	return query, err
}

func (rec *QueryValue) BuildGroupByUseAs() (string, error) {
	var query string

	err := func() error {
		var strList []string
		for _, val := range rec.GroupByList {
			str, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("GROUP BY %v", query)
		}

		return nil
	}()

	return query, err
}

func (rec *QueryValue) BuildHaving() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		var flg = false
		for _, val := range rec.HavingList {
			prefix, str, valList, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			if flg && val.Mode != QueryModeNestClose {
				strList = append(strList, prefix, str)
			} else {
				strList = append(strList, str)
				flg = true
			}

			valueList = append(valueList, valList...)

			if val.Mode == QueryModeNest {
				flg = false
			}
		}

		if len(strList) > 0 {
			query = strings.Join(strList, " ")
			query = fmt.Sprintf("HAVING %v", query)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildHavingUseAs() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		var flg = false
		for _, val := range rec.HavingList {
			prefix, str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return err
			}

			if flg && val.Mode != QueryModeNestClose {
				strList = append(strList, prefix, str)
			} else {
				strList = append(strList, str)
				flg = true
			}

			valueList = append(valueList, valList...)

			if val.Mode == QueryModeNest {
				flg = false
			}
		}

		if len(strList) > 0 {
			query = strings.Join(strList, " ")
			query = fmt.Sprintf("HAVING %v", query)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValue) BuildOrderBy() (string, error) {
	var query string

	err := func() error {
		var strList []string
		for _, val := range rec.OrderByList {
			str, err := val.Build(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("ORDER BY %v", query)
		}

		return nil
	}()

	return query, err
}

func (rec *QueryValue) BuildOrderByUseAs() (string, error) {
	var query string

	err := func() error {
		var strList []string
		for _, val := range rec.OrderByList {
			str, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return err
			}

			strList = append(strList, str)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("ORDER BY %v", query)
		}

		return nil
	}()

	return query, err
}

func (rec *QueryValue) BuildLimit() (string, error) {
	query := ""
	if rec.Limit > 0 {
		query = fmt.Sprintf("LIMIT %v", rec.Limit)
	}
	return query, nil
}

func (rec *QueryValue) BuildOffset() (string, error) {
	query := ""
	if rec.Offset > 0 {
		query = fmt.Sprintf("OFFSET %v", rec.Offset)
	}
	return query, nil
}

func (rec *QueryValue) AddMeta(tablePtr interface{}, useAs bool) {
	err := rec.Meta.Add(tablePtr, useAs)
	if err != nil {
		rec.ErrorList = append(rec.ErrorList, err)
	}
}

func (rec *QueryValue) SetTable(mode int, tablePtr interface{}) {
	data := &QueryTable{}
	data.Set(mode, tablePtr)
	rec.Table = data
}

func (rec *QueryValue) AddFrom(mode int, tablePtr interface{}, valueList ...interface{}) {
	data := &QueryFrom{}
	data.Set(mode, tablePtr, valueList...)
	rec.FromList = append(rec.FromList, data)
}

func (rec *QueryValue) AddJoin(mode int, tablePtr interface{}, valueList ...interface{}) {
	data := &QueryJoin{}
	data.Set(mode, tablePtr, valueList...)
	rec.JoinList = append(rec.JoinList, data)
}

func (rec *QueryValue) AddJoinWhere(mode int, prefix string, tablePtr interface{}, valueList ...interface{}) {
	data := &QueryJoinWhere{}
	data.Set(mode, prefix, tablePtr, valueList...)
	rec.JoinWhereList = append(rec.JoinWhereList, data)
}

func (rec *QueryValue) AddValuesColumn(mode int, valueList ...interface{}) {
	data := &QueryValuesColumn{}
	data.Set(mode, valueList...)
	rec.ValuesColumnList = append(rec.ValuesColumnList, data)
}

func (rec *QueryValue) AddValues(mode int, valueList ...interface{}) {
	data := &QueryValues{}
	data.Set(mode, valueList...)
	rec.ValuesList = append(rec.ValuesList, data)
}

func (rec *QueryValue) ClearValues() {
	rec.ValuesList = make([]*QueryValues, 0)
}

func (rec *QueryValue) AddConflict(mode int, valueList ...interface{}) {
	data := &QueryConflict{}
	data.Set(mode, valueList...)
	rec.ConfLictList = append(rec.ConfLictList, data)
}

func (rec *QueryValue) AddSet(mode int, valueList ...interface{}) {
	data := &QuerySet{}
	data.Set(mode, valueList...)
	rec.SetList = append(rec.SetList, data)
}

func (rec *QueryValue) AddSelect(mode int, valueList ...interface{}) {
	data := &QuerySelect{}
	data.Set(mode, valueList...)
	rec.SelectList = append(rec.SelectList, data)
}

func (rec *QueryValue) AddWhere(mode int, prefix string, valueList ...interface{}) {
	data := &QueryWhere{}
	data.Set(mode, prefix, valueList...)
	rec.WhereList = append(rec.WhereList, data)
}

func (rec *QueryValue) AddGroupBy(mode int, valueList ...interface{}) {
	data := &QueryGroupBy{}
	data.Set(mode, valueList...)
	rec.GroupByList = append(rec.GroupByList, data)
}

func (rec *QueryValue) AddHaving(mode int, prefix string, valueList ...interface{}) {
	data := &QueryHaving{}
	data.Set(mode, prefix, valueList...)
	rec.HavingList = append(rec.HavingList, data)
}

func (rec *QueryValue) AddOrderBy(mode int, valueList ...interface{}) {
	data := &QueryOrderBy{}
	data.Set(mode, valueList...)
	rec.OrderByList = append(rec.OrderByList, data)
}

func (rec *QueryValue) SetLimit(value int) {
	rec.Limit = value
}

func (rec *QueryValue) SetOffset(value int) {
	rec.Offset = value
}

type QueryTable struct {
	Mode     int
	TablePtr interface{}
}

func (rec *QueryTable) Build(meta *Meta) (string, error) {
	var query string

	err := func() error {
		data := meta.Get(rec.TablePtr)
		if data == nil {
			return errors.New("table meta not exist")
		}

		query = data.SchemaTable

		return nil
	}()

	return query, err
}

func (rec *QueryTable) BuildUseAs(meta *Meta) (string, error) {
	var query string

	err := func() error {
		data := meta.Get(rec.TablePtr)
		if data == nil {
			return errors.New("table meta not exist")
		}

		query = data.SchemaTable
		if data.As != "" {
			query = fmt.Sprintf("%v as %v", query, data.As)
		}

		return nil
	}()

	return query, err
}

func (rec *QueryTable) Set(mode int, tablePtr interface{}) {
	rec.Mode = mode
	rec.TablePtr = tablePtr
}

type QueryFrom struct {
	Mode      int
	TablePtr  interface{}
	ValueList []interface{}
}

func (rec *QueryFrom) Build(meta *Meta) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		data := meta.Get(rec.TablePtr)
		if data == nil {
			return errors.New("from meta not exist")
		}

		query = data.SchemaTable
		if len(rec.ValueList) > 0 {
			var strList []string
			for _, val := range rec.ValueList {
				if fmt.Sprintf("%T", val) == "[]interface {}" {
					valueList = append(valueList, val.([]interface{})...)
				} else {
					strList = append(strList, fmt.Sprintf("%v", val))
				}
			}

			query = strings.Join(strList, "")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryFrom) BuildUseAs(meta *Meta) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		data := meta.Get(rec.TablePtr)
		if data == nil {
			return errors.New("from meta not exist")
		}

		query = data.SchemaTable
		if len(rec.ValueList) > 0 {
			var strList []string
			for _, val := range rec.ValueList {
				if fmt.Sprintf("%T", val) == "[]interface {}" {
					valueList = append(valueList, val.([]interface{})...)
				} else {
					strList = append(strList, fmt.Sprintf("%v", val))
				}
			}

			query = strings.Join(strList, "")
		}

		if data.As != "" {
			query = fmt.Sprintf("%v as %v", query, data.As)
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryFrom) Set(mode int, tablePtr interface{}, valueList ...interface{}) {
	rec.Mode = mode
	rec.TablePtr = tablePtr
	rec.ValueList = valueList
}

type QueryJoin struct {
	Mode      int
	TablePtr  interface{}
	ValueList []interface{}
}

func (rec *QueryJoin) BuildUseAs(meta *Meta) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var prefix string
		switch rec.Mode {
		case QueryJoinModeInner:
			prefix = "INNER"
		case QueryJoinModeLeft:
			prefix = "LEFT"
		case QueryJoinModeRight:
			prefix = "RIGHT"
		default:
			return errors.New("join mode not exist")
		}

		data := meta.Get(rec.TablePtr)
		if data == nil {
			return errors.New("join meta not exist")
		}

		var table string
		if len(rec.ValueList) > 0 {
			var strList []string
			for _, val := range rec.ValueList {
				if fmt.Sprintf("%T", val) == "[]interface {}" {
					valueList = append(valueList, val.([]interface{})...)
				} else {
					strList = append(strList, fmt.Sprintf("%v", val))
				}
			}

			table = strings.Join(strList, "")
		} else {
			table = data.SchemaTable
		}

		query = fmt.Sprintf("%v JOIN %v as %v ON", prefix, table, data.As)

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryJoin) Set(mode int, tablePtr interface{}, valueList ...interface{}) {
	rec.Mode = mode
	rec.TablePtr = tablePtr
	rec.ValueList = valueList
}

type QueryJoinWhere struct {
	Mode      int
	Prefix    string
	TablePtr  interface{}
	ValueList []interface{}
}

func (rec *QueryJoinWhere) BuildUseAs(meta *Meta) (string, string, []interface{}, error) {
	query, valueList, err := buildQueryWhere(meta, "joinWhere", true, rec.Mode, rec.ValueList...)
	return rec.Prefix, query, valueList, err
}

func (rec *QueryJoinWhere) Set(mode int, prefix string, tablePtr interface{}, valueList ...interface{}) {
	rec.Mode = mode
	rec.Prefix = prefix
	rec.TablePtr = tablePtr
	rec.ValueList = valueList
}

type QueryValuesColumn struct {
	Mode      int
	ValueList []interface{}
}

func (rec *QueryValuesColumn) Build(meta *Meta) (string, error) {
	var query string

	err := func() error {
		if len(rec.ValueList) != 1 {
			return errors.New(fmt.Sprintf("valuesColumn length must be 1"))
		}

		data := meta.Get(rec.ValueList[0])
		if data == nil {
			return errors.New("valuesColumn meta not exist")
		}

		query = data.Column

		return nil
	}()

	return query, err
}

func (rec *QueryValuesColumn) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}

type QueryValues struct {
	Mode      int
	ValueList []interface{}
}

func (rec *QueryValues) Build(_ *Meta) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		var strList []string
		for _, val := range rec.ValueList {
			strList = append(strList, PlaceHolder)
			valueList = append(valueList, val)
		}

		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryValues) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}

type QueryConflict struct {
	Mode      int
	ValueList []interface{}
}

func (rec *QueryConflict) Build(meta *Meta) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		if len(rec.ValueList) != 1 {
			return errors.New(fmt.Sprintf("conflict length must be 1"))
		}

		data := meta.Get(rec.ValueList[0])
		if data == nil {
			return errors.New("conflict meta not exist")
		}

		query = data.Column

		return nil
	}()

	return query, valueList, err
}

func (rec *QueryConflict) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}

type QuerySet struct {
	Mode      int
	ValueList []interface{}
}

func (rec *QuerySet) Build(meta *Meta) (string, interface{}, error) {
	var query string

	err := func() error {
		if len(rec.ValueList) != 2 {
			return errors.New(fmt.Sprintf("set length must be 2"))
		}

		data := meta.Get(rec.ValueList[0])
		if data == nil {
			return errors.New("set meta not exist")
		}

		query = fmt.Sprintf("%v = %v", data.Column, PlaceHolder)

		return nil
	}()

	return query, rec.ValueList[1], err
}

func (rec *QuerySet) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}

type QuerySelect struct {
	Mode      int
	ValueList []interface{}
}

func (rec *QuerySelect) Build(meta *Meta) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		switch rec.Mode {
		case QueryModeDefault:
			if len(rec.ValueList) < 1 {
				return errors.New("select length must be greater than 1")
			}

			var strList []string
			for _, val := range rec.ValueList {
				if fmt.Sprintf("%T", val) == "[]interface {}" {
					valueList = append(valueList, val.([]interface{})...)
				} else {
					if data := meta.Get(val); data != nil {
						strList = append(strList, data.SchemaTableColumn)
					} else {
						strList = append(strList, fmt.Sprintf("%v", val))
					}
				}
			}

			query = strings.Join(strList, "")
		case QueryModeAll:
			if len(rec.ValueList) != 1 {
				return errors.New("select length should be 1")
			}

			data := meta.Get(rec.ValueList[0])
			if data == nil {
				return errors.New("select not exist")
			}

			query = fmt.Sprintf("%s.*", data.SchemaTable)
		default:
			return errors.New("select mode not exist")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QuerySelect) BuildUseAs(meta *Meta) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	err := func() error {
		switch rec.Mode {
		case QueryModeDefault:
			if len(rec.ValueList) < 1 {
				return errors.New("select length must be greater than 1")
			}

			var strList []string
			for _, val := range rec.ValueList {
				if fmt.Sprintf("%T", val) == "[]interface {}" {
					valueList = append(valueList, val.([]interface{})...)
				} else {
					if data := meta.Get(val); data != nil {
						strList = append(strList, data.SchemaTableAsColumn)
					} else {
						strList = append(strList, fmt.Sprintf("%v", val))
					}
				}
			}

			query = strings.Join(strList, "")
		case QueryModeAll:
			if len(rec.ValueList) != 1 {
				return errors.New("select length should be 1")
			}

			data := meta.Get(rec.ValueList[0])
			if data == nil {
				return errors.New("select not exist")
			}

			query = fmt.Sprintf("%s.*", data.SchemaTableAs)
		default:
			return errors.New("select mode not exist")
		}

		return nil
	}()

	return query, valueList, err
}

func (rec *QuerySelect) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}

type QueryWhere struct {
	Mode      int
	Prefix    string
	ValueList []interface{}
}

func (rec *QueryWhere) Build(meta *Meta) (string, string, []interface{}, error) {
	query, valueList, err := buildQueryWhere(meta, "where", false, rec.Mode, rec.ValueList...)
	return rec.Prefix, query, valueList, err
}

func (rec *QueryWhere) BuildUseAs(meta *Meta) (string, string, []interface{}, error) {
	query, valueList, err := buildQueryWhere(meta, "where", true, rec.Mode, rec.ValueList...)
	return rec.Prefix, query, valueList, err
}

func (rec *QueryWhere) Set(mode int, prefix string, valueList ...interface{}) {
	rec.Mode = mode
	rec.Prefix = prefix
	rec.ValueList = valueList
}

type QueryGroupBy struct {
	Mode      int
	ValueList []interface{}
}

func (rec *QueryGroupBy) Build(meta *Meta) (string, error) {
	var query string

	err := func() error {
		if len(rec.ValueList) < 1 {
			return errors.New("select length must be greater than 1")
		}

		var strList []string
		for _, val := range rec.ValueList {
			data := meta.Get(val)
			if data != nil {
				strList = append(strList, data.SchemaTableColumn)
			} else {
				strList = append(strList, fmt.Sprintf("%v", val))
			}
		}

		query = strings.Join(strList, "")

		return nil
	}()

	return query, err
}

func (rec *QueryGroupBy) BuildUseAs(meta *Meta) (string, error) {
	var query string

	err := func() error {
		if len(rec.ValueList) < 1 {
			return errors.New("select length must be greater than 1")
		}

		var strList []string
		for _, val := range rec.ValueList {
			data := meta.Get(val)
			if data != nil {
				strList = append(strList, data.SchemaTableAsColumn)
			} else {
				strList = append(strList, fmt.Sprintf("%v", val))
			}
		}

		query = strings.Join(strList, "")

		return nil
	}()

	return query, err
}

func (rec *QueryGroupBy) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}

type QueryHaving struct {
	Mode      int
	Prefix    string
	ValueList []interface{}
}

func (rec *QueryHaving) Build(meta *Meta) (string, string, []interface{}, error) {
	query, valueList, err := buildQueryWhere(meta, "having", false, rec.Mode, rec.ValueList...)
	return rec.Prefix, query, valueList, err
}

func (rec *QueryHaving) BuildUseAs(meta *Meta) (string, string, []interface{}, error) {
	query, valueList, err := buildQueryWhere(meta, "having", true, rec.Mode, rec.ValueList...)
	return rec.Prefix, query, valueList, err
}

func (rec *QueryHaving) Set(mode int, prefix string, valueList ...interface{}) {
	rec.Mode = mode
	rec.Prefix = prefix
	rec.ValueList = valueList
}

type QueryOrderBy struct {
	Mode      int
	ValueList []interface{}
}

func (rec *QueryOrderBy) Build(meta *Meta) (string, error) {
	var query string

	err := func() error {
		if len(rec.ValueList) < 1 {
			return errors.New("orderBy length must be greater than 1")
		}

		{
			var strList []string
			for _, val := range rec.ValueList {
				data := meta.Get(val)
				if data != nil {
					strList = append(strList, data.SchemaTableColumn)
				} else {
					strList = append(strList, fmt.Sprintf("%v", val))
				}
			}

			query = strings.Join(strList, "")
		}

		switch rec.Mode {
		case QueryModeDefault:
		case QueryModeAsc:
			query = fmt.Sprintf("%s ASC", query)
		case QueryModeDesc:
			query = fmt.Sprintf("%s DESC", query)
		default:
			return errors.New("orderBy mode not exist")
		}

		return nil
	}()

	return query, err
}

func (rec *QueryOrderBy) BuildUseAs(meta *Meta) (string, error) {
	var query string

	err := func() error {
		if len(rec.ValueList) < 1 {
			return errors.New("orderBy length must be greater than 1")
		}

		{
			var strList []string
			for _, val := range rec.ValueList {
				data := meta.Get(val)
				if data != nil {
					strList = append(strList, data.SchemaTableAsColumn)
				} else {
					strList = append(strList, fmt.Sprintf("%v", val))
				}
			}

			query = strings.Join(strList, "")
		}

		switch rec.Mode {
		case QueryModeDefault:
		case QueryModeAsc:
			query = fmt.Sprintf("%s ASC", query)
		case QueryModeDesc:
			query = fmt.Sprintf("%s DESC", query)
		default:
			return errors.New("orderBy mode not exist")
		}

		return nil
	}()

	return query, err
}

func (rec *QueryOrderBy) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}
