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
	JoinList         []*QueryJoin
	JoinWhereList    []*QueryJoinWhere
	ValuesColumnList []*QueryValuesColumn
	ValuesList       []*QueryValues
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
	query := "INSERT"
	var valueList []interface{}

	{
		str, err := rec.BuildTable()
		if err != nil {
			return "", nil, err
		}
		if str == "" {
			return "", nil, errors.New("table not exist")
		}
		query = fmt.Sprintf("%v INTO %v", query, str)
	}

	{
		str, err := rec.BuildValuesColumn()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
		}
	}

	{
		str, valList, err := rec.BuildValues()
		if err != nil {
			return "", nil, err
		}
		if str == "" {
			return "", nil, errors.New("values not exist")
		}
		query = fmt.Sprintf("%v VALUES %v", query, str)
		valueList = append(valueList, valList...)
	}

	return query, valueList, nil
}

func (rec *QueryValue) GetUpdateQuery() (string, []interface{}, error) {
	query := "UPDATE"
	var valueList []interface{}

	{
		str, err := rec.BuildTable()
		if err != nil {
			return "", nil, err
		}
		if str == "" {
			return "", nil, errors.New("table not exist")
		}
		query = fmt.Sprintf("%v %v", query, str)
	}

	{
		str, valList, err := rec.BuildSet()
		if err != nil {
			return "", nil, err
		}
		if str == "" {
			return "", nil, errors.New("set not exist")
		}
		query = fmt.Sprintf("%v %v", query, str)
		valueList = append(valueList, valList...)
	}

	{
		str, valList, err := rec.BuildWhere()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
			valueList = append(valueList, valList...)
		}
	}

	return query, valueList, nil
}

func (rec *QueryValue) GetDeleteQuery() (string, []interface{}, error) {
	query := "DELETE"
	var valueList []interface{}

	{
		str, err := rec.BuildTable()
		if err != nil {
			return "", nil, err
		}
		if str == "" {
			return "", nil, errors.New("table not exist")
		}
		query = fmt.Sprintf("%v FROM %v", query, str)
	}

	{
		str, valList, err := rec.BuildWhere()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
			valueList = append(valueList, valList...)
		}
	}

	return query, valueList, nil
}

func (rec *QueryValue) GetTruncateQuery() (string, []interface{}, error) {
	query := "TRUNCATE TABLE"
	var valueList []interface{}

	{
		str, err := rec.BuildTable()
		if err != nil {
			return "", nil, err
		}
		if str == "" {
			return "", nil, errors.New("table not exist")
		}
		query = fmt.Sprintf("%v %v", query, str)
	}

	return query, valueList, nil
}

func (rec *QueryValue) GetTruncateRestartIdentityQuery() (string, []interface{}, error) {
	query, valueList, err := rec.GetTruncateQuery()
	if err != nil {
		return "", nil, err
	}

	query = fmt.Sprintf("%v RESTART IDENTITY", query)

	return query, valueList, nil
}

func (rec *QueryValue) GetSelectQuery() (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	{
		str, valList, err := rec.BuildSelectUseAs()
		if err != nil {
			return "", nil, err
		}
		if str == "" {
			return "", nil, errors.New("select not exist")
		}
		query = str
		valueList = append(valueList, valList...)
	}

	{
		str, err := rec.BuildTableUseAs()
		if err != nil {
			return "", nil, err
		}
		if str == "" {
			return "", nil, errors.New("table not exist")
		}
		query = fmt.Sprintf("%v FROM %v", query, str)
	}

	{
		str, valList, err := rec.BuildJoinUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
			valueList = append(valueList, valList...)
		}
	}

	{
		str, valList, err := rec.BuildWhereUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
			valueList = append(valueList, valList...)
		}
	}

	{
		str, err := rec.BuildGroupByUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
		}
	}

	{
		str, valList, err := rec.BuildHavingUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
			valueList = append(valueList, valList...)
		}
	}

	{
		str, err := rec.BuildOrderByUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
		}
	}

	{
		str, err := rec.BuildLimit()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
		}
	}

	{
		str, err := rec.BuildOffset()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
		}
	}

	return query, valueList, nil
}

func (rec *QueryValue) GetSelectCountQuery() (string, []interface{}, error) {
	query := "SELECT count(*) as count"
	var valueList []interface{}

	{
		str, err := rec.BuildTableUseAs()
		if err != nil {
			return "", nil, err
		}
		if str == "" {
			return "", nil, errors.New("table not exist")
		}
		query = fmt.Sprintf("%v FROM %v", query, str)
	}

	{
		str, valList, err := rec.BuildJoinUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
			valueList = append(valueList, valList...)
		}
	}

	{
		str, valList, err := rec.BuildWhereUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
			valueList = append(valueList, valList...)
		}
	}

	{
		str, err := rec.BuildGroupByUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
		}
	}

	{
		str, valList, err := rec.BuildHavingUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
			valueList = append(valueList, valList...)
		}
	}

	{
		str, err := rec.BuildOrderByUseAs()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
		}
	}

	{
		str, err := rec.BuildLimit()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
		}
	}

	{
		str, err := rec.BuildOffset()
		if err != nil {
			return "", nil, err
		}
		if str != "" {
			query = fmt.Sprintf("%v %v", query, str)
		}
	}

	return query, valueList, nil
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

func (rec *QueryValue) BuildJoinUseAs() (string, []interface{}, error) {
	var query string
	var valueList []interface{}
	{
		var strList []string
		for _, val := range rec.JoinList {
			str, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return "", nil, err
			}

			where, valList, err := rec.BuildJoinWhereUseAs(val.TablePtr)
			if err != nil {
				return "", nil, err
			}

			str = fmt.Sprintf("%v %v", str, where)

			strList = append(strList, str)
			valueList = append(valueList, valList...)
		}
		if len(strList) > 0 {
			query = strings.Join(strList, " ")
		}
	}
	return query, valueList, nil
}

func (rec *QueryValue) BuildJoinWhereUseAs(tablePtr interface{}) (string, []interface{}, error) {
	var query string
	var valueList []interface{}
	{
		var strList []string
		flg := false
		for _, val := range rec.JoinWhereList {
			if tablePtr != val.TablePtr {
				continue
			}
			prefix, str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return "", nil, err
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
			return "", nil, errors.New("joinWhere not exist")
		}
		query = strings.Join(strList, " ")
	}
	return query, valueList, nil
}

func (rec *QueryValue) BuildValuesColumn() (string, error) {
	var query string
	{
		var strList []string
		for _, val := range rec.ValuesColumnList {
			str, err := val.Build(rec.Meta)
			if err != nil {
				return "", err
			}
			strList = append(strList, str)
		}
		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("(%v)", query)
		}
	}
	return query, nil
}

func (rec *QueryValue) BuildValues() (string, []interface{}, error) {
	var query string
	var valueList []interface{}
	{
		var strList []string
		for _, val := range rec.ValuesList {
			str, valList, err := val.Build(rec.Meta)
			if err != nil {
				return "", nil, err
			}
			strList = append(strList, str)
			valueList = append(valueList, valList...)
		}
		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
		}
	}
	return query, valueList, nil
}

func (rec *QueryValue) BuildSet() (string, []interface{}, error) {
	var query string
	var valueList []interface{}
	{
		var strList []string
		for _, val := range rec.SetList {
			str, value, err := val.Build(rec.Meta)
			if err != nil {
				return "", nil, err
			}
			strList = append(strList, str)
			valueList = append(valueList, value)
		}
		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("SET %v", query)
		}
	}
	return query, valueList, nil
}

func (rec *QueryValue) BuildSelectUseAs() (string, []interface{}, error) {
	var query string
	var valueList []interface{}
	{
		var strList []string
		for _, val := range rec.SelectList {
			str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return "", nil, err
			}
			strList = append(strList, str)
			valueList = append(valueList, valList...)
		}
		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("SELECT %v", query)
		}
	}
	return query, valueList, nil
}

func (rec *QueryValue) BuildWhere() (string, []interface{}, error) {
	var query string
	var valueList []interface{}
	{
		var strList []string
		flg := false
		for _, val := range rec.WhereList {
			prefix, str, valList, err := val.Build(rec.Meta)
			if err != nil {
				return "", nil, err
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
	}
	return query, valueList, nil
}

func (rec *QueryValue) BuildWhereUseAs() (string, []interface{}, error) {
	var query string
	var valueList []interface{}
	{
		var strList []string
		flg := false
		for _, val := range rec.WhereList {
			prefix, str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return "", nil, err
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
	}
	return query, valueList, nil
}

func (rec *QueryValue) BuildGroupByUseAs() (string, error) {
	var query string
	{
		var strList []string
		for _, val := range rec.GroupByList {
			str, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return "", err
			}
			strList = append(strList, str)
		}
		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("GROUP BY %v", query)
		}
	}
	return query, nil
}

func (rec *QueryValue) BuildHavingUseAs() (string, []interface{}, error) {
	var query string
	var valueList []interface{}
	{
		var strList []string
		flg := false
		for _, val := range rec.HavingList {
			prefix, str, valList, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return "", nil, err
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
	}
	return query, valueList, nil
}

func (rec *QueryValue) BuildOrderByUseAs() (string, error) {
	var query string
	{
		var strList []string
		for _, val := range rec.OrderByList {
			str, err := val.BuildUseAs(rec.Meta)
			if err != nil {
				return "", err
			}
			strList = append(strList, str)
		}
		if len(strList) > 0 {
			query = strings.Join(strList, ", ")
			query = fmt.Sprintf("ORDER BY %v", query)
		}
	}
	return query, nil
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

func (rec *QueryValue) AddMeta(schema string, tablePtr interface{}, as string) {
	err := rec.Meta.Add(schema, tablePtr, as)
	if err != nil {
		rec.ErrorList = append(rec.ErrorList, err)
	}
}

func (rec *QueryValue) SetTable(mode int, valueList ...interface{}) {
	data := &QueryTable{}
	data.Set(mode, valueList...)
	rec.Table = data
}

func (rec *QueryValue) AddJoin(mode int, tablePtr interface{}) {
	data := &QueryJoin{}
	data.Set(mode, tablePtr)
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
	Mode      int
	ValueList []interface{}
}

func (rec *QueryTable) Build(meta *Meta) (string, error) {
	if len(rec.ValueList) != 1 {
		return "", errors.New(fmt.Sprintf("table length must be 1"))
	}

	data := meta.Get(rec.ValueList[0])
	if data == nil {
		return "", errors.New("table meta not exist")
	}

	query := data.SchemaTable

	return query, nil
}

func (rec *QueryTable) BuildUseAs(meta *Meta) (string, error) {
	if len(rec.ValueList) != 1 {
		return "", errors.New(fmt.Sprintf("table length must be 1"))
	}

	data := meta.Get(rec.ValueList[0])
	if data == nil {
		return "", errors.New("table meta not exist")
	}

	query := data.SchemaTable
	if data.As != "" {
		query = fmt.Sprintf("%v as %v", query, data.As)
	}

	return query, nil
}

func (rec *QueryTable) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}

type QueryJoin struct {
	Mode     int
	TablePtr interface{}
}

func (rec *QueryJoin) BuildUseAs(meta *Meta) (string, error) {
	var prefix string
	switch rec.Mode {
	case QueryJoinModeInner:
		prefix = "INNER"
	case QueryJoinModeLeft:
		prefix = "LEFT"
	case QueryJoinModeRight:
		prefix = "RIGHT"
	default:
		return "", errors.New("join mode not exist")
	}

	data := meta.Get(rec.TablePtr)
	if data == nil {
		return "", errors.New("join meta not exist")
	}

	table := data.SchemaTable
	if data.As != "" {
		table = fmt.Sprintf("%v as %v", table, data.As)
	}

	query := fmt.Sprintf("%v JOIN %v ON", prefix, table)

	return query, nil
}

func (rec *QueryJoin) Set(mode int, tablePtr interface{}) {
	rec.Mode = mode
	rec.TablePtr = tablePtr
}

type QueryJoinWhere struct {
	Mode      int
	Prefix    string
	TablePtr  interface{}
	ValueList []interface{}
}

func (rec *QueryJoinWhere) BuildUseAs(meta *Meta) (string, string, []interface{}, error) {
	query, valueList, err := buildQueryWhere(meta, "joinWhere", true, rec.Mode, rec.ValueList...)
	if err != nil {
		return "", "", nil, err
	}
	return rec.Prefix, query, valueList, nil
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

	if len(rec.ValueList) != 1 {
		return "", errors.New(fmt.Sprintf("valuesColumn length must be 1"))
	}

	data := meta.Get(rec.ValueList[0])
	if data == nil {
		return "", errors.New("valuesColumn meta not exist")
	}

	query = data.Column

	return query, nil
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

	{
		var strList []string
		for _, val := range rec.ValueList {
			strList = append(strList, PlaceHolder)
			valueList = append(valueList, val)
		}
		var str string
		if len(strList) > 0 {
			str = strings.Join(strList, ", ")
		}
		query = fmt.Sprintf("(%v)", str)
	}

	return query, valueList, nil
}

func (rec *QueryValues) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}

type QuerySet struct {
	Mode      int
	ValueList []interface{}
}

func (rec *QuerySet) Build(meta *Meta) (string, interface{}, error) {
	var query string

	if len(rec.ValueList) != 2 {
		return "", nil, errors.New(fmt.Sprintf("set length must be 2"))
	}

	data := meta.Get(rec.ValueList[0])
	if data == nil {
		return "", nil, errors.New("set meta not exist")
	}

	query = fmt.Sprintf("%v = %v", data.Column, PlaceHolder)

	return query, rec.ValueList[1], nil
}

func (rec *QuerySet) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}

type QuerySelect struct {
	Mode      int
	ValueList []interface{}
}

func (rec *QuerySelect) BuildUseAs(meta *Meta) (string, []interface{}, error) {
	var query string
	var valueList []interface{}

	switch rec.Mode {
	case QueryModeDefault:
		if len(rec.ValueList) < 1 {
			return query, valueList, errors.New("select length must be greater than 1")
		}

		var strList []string
		for _, val := range rec.ValueList {
			if fmt.Sprintf("%T", val) == "[]interface {}" {
				for _, value := range val.([]interface{}) {
					valueList = append(valueList, value)
				}
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
			return query, valueList, errors.New("select length should be 1")
		}
		data := meta.Get(rec.ValueList[0])
		if data == nil {
			return query, valueList, errors.New("select not exist")
		}
		query = fmt.Sprintf("%s.*", data.SchemaTableAs)
	default:
		return query, valueList, errors.New("select mode not exist")
	}

	return query, valueList, nil
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
	if err != nil {
		return "", "", nil, err
	}
	return rec.Prefix, query, valueList, nil
}

func (rec *QueryWhere) BuildUseAs(meta *Meta) (string, string, []interface{}, error) {
	query, valueList, err := buildQueryWhere(meta, "where", true, rec.Mode, rec.ValueList...)
	if err != nil {
		return "", "", nil, err
	}
	return rec.Prefix, query, valueList, nil
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

func (rec *QueryGroupBy) BuildUseAs(meta *Meta) (string, error) {
	var query string

	if len(rec.ValueList) < 1 {
		return query, errors.New("select length must be greater than 1")
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

	return query, nil
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

func (rec *QueryHaving) BuildUseAs(meta *Meta) (string, string, []interface{}, error) {
	query, valueList, err := buildQueryWhere(meta, "having", true, rec.Mode, rec.ValueList...)
	if err != nil {
		return "", "", nil, err
	}
	return rec.Prefix, query, valueList, nil
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

func (rec *QueryOrderBy) BuildUseAs(meta *Meta) (string, error) {
	var query string

	if len(rec.ValueList) < 1 {
		return query, errors.New("orderBy length must be greater than 1")
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
		return query, errors.New("orderBy mode not exist")
	}

	return query, nil
}

func (rec *QueryOrderBy) Set(mode int, valueList ...interface{}) {
	rec.Mode = mode
	rec.ValueList = valueList
}
