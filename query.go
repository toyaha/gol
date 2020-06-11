package gol

import (
	"database/sql"
	"errors"
)

func NewQuery(config *Config) *Query {
	if config == nil {
		config = NewConfig()
	}

	query := &Query{
		Client: nil,
		Config: config,
		Value:  NewQueryValue(config),
	}

	return query
}

type Query struct {
	Client *Client
	Config *Config
	Value  *QueryValue
}

func (rec *Query) BulkInsert() (sql.Result, error) {
	var err error
	var result sql.Result

	if rec.Client == nil {
		return nil, errors.New("database does not exist")
	}

	if rec.Config.BulkInsertCount > 0 && rec.Value.GetValuesCount() >= rec.Config.BulkInsertCount {
		result, err = rec.Insert()
		if err != nil {
			return nil, err
		}

		rec.Value.ClearValues()
	}

	return result, nil
}

func (rec *Query) BulkInsertFinish() (sql.Result, error) {
	var err error
	var result sql.Result

	if rec.Client == nil {
		return nil, errors.New("database does not exist")
	}

	if rec.Config.BulkInsertCount > 0 && rec.Value.GetValuesCount() > 0 {
		result, err = rec.Insert()
		if err != nil {
			return nil, err
		}

		rec.Value.ClearValues()
	}

	return result, nil
}

func (rec *Query) Insert() (sql.Result, error) {
	if rec.Client == nil {
		return nil, errors.New("database does not exist")
	}

	query, valueList, err := rec.GetInsertQuery()
	if err != nil {
		return nil, err
	}

	result, err := rec.Client.Exec(query, valueList...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rec *Query) Update() (sql.Result, error) {
	if rec.Client == nil {
		return nil, errors.New("database does not exist")
	}

	query, valueList, err := rec.GetUpdateQuery()
	if err != nil {
		return nil, err
	}

	result, err := rec.Client.Exec(query, valueList...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rec *Query) Delete() (sql.Result, error) {
	if rec.Client == nil {
		return nil, errors.New("database does not exist")
	}

	query, valueList, err := rec.GetDeleteQuery()
	if err != nil {
		return nil, err
	}

	result, err := rec.Client.Exec(query, valueList...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rec *Query) Truncate() (sql.Result, error) {
	if rec.Client == nil {
		return nil, errors.New("database does not exist")
	}

	query, valueList, err := rec.GetTruncateQuery()
	if err != nil {
		return nil, err
	}

	result, err := rec.Client.Exec(query, valueList...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rec *Query) TruncateRestartIdentity() (sql.Result, error) {
	if rec.Client == nil {
		return nil, errors.New("database does not exist")
	}

	query, valueList, err := rec.GetTruncateRestartIdentityQuery()
	if err != nil {
		return nil, err
	}

	result, err := rec.Client.Exec(query, valueList...)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (rec *Query) Select(dest interface{}) error {
	if rec.Client == nil {
		return errors.New("database does not exist")
	}

	query, valueList, err := rec.GetSelectQuery()
	if err != nil {
		return err
	}

	err = rec.Client.Find(dest, query, valueList...)
	if err != nil {
		return err
	}

	return nil
}

func (rec *Query) SelectCount(dest interface{}) error {
	if rec.Client == nil {
		return errors.New("database does not exist")
	}

	query, valueList, err := rec.GetSelectCountQuery()
	if err != nil {
		return err
	}

	err = rec.Client.Find(dest, query, valueList...)
	if err != nil {
		return err
	}

	return nil
}

func (rec *Query) GetInsertQuery() (string, []interface{}, error) {
	return rec.Value.GetInsertQuery()
}

func (rec *Query) GetUpdateQuery() (string, []interface{}, error) {
	return rec.Value.GetUpdateQuery()
}

func (rec *Query) GetDeleteQuery() (string, []interface{}, error) {
	return rec.Value.GetDeleteQuery()
}

func (rec *Query) GetTruncateQuery() (string, []interface{}, error) {
	return rec.Value.GetTruncateQuery()
}

func (rec *Query) GetTruncateRestartIdentityQuery() (string, []interface{}, error) {
	return rec.Value.GetTruncateRestartIdentityQuery()
}

func (rec *Query) GetSelectQuery() (string, []interface{}, error) {
	return rec.Value.GetSelectQuery()
}

func (rec *Query) GetSelectCountQuery() (string, []interface{}, error) {
	return rec.Value.GetSelectCountQuery()
}

func (rec *Query) SetClient(client *Client) {
	rec.Client = client

	if rec.Value != nil {
		return
	}

	if rec.Value.Meta == nil {
		rec.Value.Meta = NewMeta(rec.Config)
	}

	for key, val := range client.Meta.Value {
		data := *val
		rec.Value.Meta.Value[key] = &data
	}
}

func (rec *Query) AddMeta(tablePtr interface{}) {
	rec.Value.AddMeta("", tablePtr, "")
}

func (rec *Query) AddMetaAs(tablePtr interface{}, alias string) {
	rec.Value.AddMeta("", tablePtr, alias)
}

func (rec *Query) AddMetaSchema(schema string, tablePtr interface{}) {
	rec.Value.AddMeta(schema, tablePtr, "")
}

func (rec *Query) AddMetaAsSchema(schema string, tablePtr interface{}, alias string) {
	rec.Value.AddMeta(schema, tablePtr, alias)
}

func (rec *Query) SetTable(tablePtr interface{}) {
	rec.Value.AddMeta("", tablePtr, "")
	rec.Value.SetTable(QueryModeDefault, tablePtr)
}

func (rec *Query) SetTableAs(tablePtr interface{}, tableAs string) {
	rec.Value.AddMeta("", tablePtr, tableAs)
	rec.Value.SetTable(QueryModeDefault, tablePtr)
}

func (rec *Query) SetTableWithSchema(schema string, tablePtr interface{}) {
	rec.Value.AddMeta(schema, tablePtr, "")
	rec.Value.SetTable(QueryModeDefault, tablePtr)
}

func (rec *Query) SetTableAsWithSchema(schema string, tablePtr interface{}, tableAs string) {
	rec.Value.AddMeta(schema, tablePtr, tableAs)
	rec.Value.SetTable(QueryModeDefault, tablePtr)
}

func (rec *Query) SetJoin(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddMeta("", tablePtr, "")
	rec.Value.AddJoin(QueryJoinModeInner, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinWithSchema(schema string, tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddMeta(schema, tablePtr, "")
	rec.Value.AddJoin(QueryJoinModeInner, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinAs(tablePtr interface{}, tableAs string, valueList ...interface{}) {
	rec.Value.AddMeta("", tablePtr, tableAs)
	rec.Value.AddJoin(QueryJoinModeInner, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinAsWithSchema(schema string, tablePtr interface{}, tableAs string, valueList ...interface{}) {
	rec.Value.AddMeta(schema, tablePtr, tableAs)
	rec.Value.AddJoin(QueryJoinModeInner, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinLeft(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddMeta("", tablePtr, "")
	rec.Value.AddJoin(QueryJoinModeLeft, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinLeftWithSchema(schema string, tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddMeta(schema, tablePtr, "")
	rec.Value.AddJoin(QueryJoinModeLeft, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinLeftAs(tablePtr interface{}, tableAs string, valueList ...interface{}) {
	rec.Value.AddMeta("", tablePtr, tableAs)
	rec.Value.AddJoin(QueryJoinModeLeft, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinLeftAsWithSchema(schema string, tablePtr interface{}, tableAs string, valueList ...interface{}) {
	rec.Value.AddMeta(schema, tablePtr, tableAs)
	rec.Value.AddJoin(QueryJoinModeLeft, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinRight(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddMeta("", tablePtr, "")
	rec.Value.AddJoin(QueryJoinModeRight, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinRightWithSchema(schema string, tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddMeta(schema, tablePtr, "")
	rec.Value.AddJoin(QueryJoinModeRight, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinRightAs(tablePtr interface{}, tableAs string, valueList ...interface{}) {
	rec.Value.AddMeta("", tablePtr, tableAs)
	rec.Value.AddJoin(QueryJoinModeRight, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinRightAsWithSchema(schema string, tablePtr interface{}, tableAs string, valueList ...interface{}) {
	rec.Value.AddMeta(schema, tablePtr, tableAs)
	rec.Value.AddJoin(QueryJoinModeRight, tablePtr)
	if len(valueList) > 0 {
		rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
	}
}

func (rec *Query) SetJoinWhere(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereIs(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeIs, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereIsNot(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeIsNot, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereLike(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeLike, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereLikeNot(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeLikeNot, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereIn(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeIn, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereInNot(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeInNot, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereGt(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeGt, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereGte(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeGte, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereLt(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeLt, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereLte(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeLte, QueryPrefixAnd, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOr(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeDefault, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrIs(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeIs, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrIsNot(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeIsNot, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrLike(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeLike, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrLikeNot(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeLikeNot, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrIn(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeIn, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrInNot(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeInNot, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrGt(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeGt, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrGte(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeGte, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrLt(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeLt, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereOrLte(tablePtr interface{}, valueList ...interface{}) {
	rec.Value.AddJoinWhere(QueryModeLte, QueryPrefixOr, tablePtr, valueList...)
}

func (rec *Query) SetJoinWhereNest(tablePtr interface{}) {
	rec.Value.AddJoinWhere(QueryModeNest, QueryPrefixAnd, tablePtr)
}

func (rec *Query) SetJoinWhereOrNest(tablePtr interface{}) {
	rec.Value.AddJoinWhere(QueryModeNest, QueryPrefixOr, tablePtr)
}

func (rec *Query) SetJoinWhereNestClose(tablePtr interface{}) {
	rec.Value.AddJoinWhere(QueryModeNestClose, QueryPrefixNone, tablePtr)
}

func (rec *Query) SetValuesColumn(columnPtrList ...interface{}) {
	for _, val := range columnPtrList {
		rec.Value.AddValuesColumn(QueryModeDefault, val)
	}
}

func (rec *Query) SetValues(valueList ...interface{}) {
	rec.Value.AddValues(QueryModeDefault, valueList...)
}

func (rec *Query) SetValuesClear() {
	rec.Value.ClearValues()
}

func (rec *Query) SetSet(columnPtr interface{}, value interface{}) {
	rec.Value.AddSet(QueryModeDefault, columnPtr, value)
}

func (rec *Query) SetSelect(valueList ...interface{}) {
	rec.Value.AddSelect(QueryModeDefault, valueList...)
}

func (rec *Query) SetSelectAll(tablePtr interface{}) {
	rec.Value.AddSelect(QueryModeAll, tablePtr)
}

func (rec *Query) SetWhere(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeDefault, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereIs(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeIs, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereIsNot(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeIsNot, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereLike(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeLike, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereLikeNot(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeLikeNot, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereIn(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeIn, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereInNot(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeInNot, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereGt(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeGt, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereGte(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeGte, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereLt(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeLt, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereLte(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeLte, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetWhereOr(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeDefault, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrIs(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeIs, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrIsNot(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeIsNot, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrLike(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeLike, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrLikeNot(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeLikeNot, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrIn(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeIn, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrInNot(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeInNot, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrGt(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeGt, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrGte(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeGte, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrLt(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeLt, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereOrLte(valueList ...interface{}) {
	rec.Value.AddWhere(QueryModeLte, QueryPrefixOr, valueList...)
}

func (rec *Query) SetWhereNest() {
	rec.Value.AddWhere(QueryModeNest, QueryPrefixAnd)
}

func (rec *Query) SetWhereOrNest() {
	rec.Value.AddWhere(QueryModeNest, QueryPrefixOr)
}

func (rec *Query) SetWhereNestClose() {
	rec.Value.AddWhere(QueryModeNestClose, QueryPrefixNone)
}

func (rec *Query) SetGroupBy(valueList ...interface{}) {
	rec.Value.AddGroupBy(QueryModeDefault, valueList...)
}

func (rec *Query) SetHaving(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeDefault, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingIs(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeIs, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingIsNot(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeIsNot, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingLike(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeLike, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingLikeNot(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeLikeNot, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingIn(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeIn, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingInNot(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeInNot, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingGt(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeGt, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingGte(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeGte, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingLt(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeLt, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingLte(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeLte, QueryPrefixAnd, valueList...)
}

func (rec *Query) SetHavingOr(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeDefault, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrIs(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeIs, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrIsNot(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeIsNot, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrLike(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeLike, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrLikeNot(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeLikeNot, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrIn(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeIn, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrInNot(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeInNot, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrGt(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeGt, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrGte(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeGte, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrLt(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeLt, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingOrLte(valueList ...interface{}) {
	rec.Value.AddHaving(QueryModeLte, QueryPrefixOr, valueList...)
}

func (rec *Query) SetHavingNest() {
	rec.Value.AddHaving(QueryModeNest, QueryPrefixAnd)
}

func (rec *Query) SetHavingOrNest() {
	rec.Value.AddHaving(QueryModeNest, QueryPrefixOr)
}

func (rec *Query) SetHavingNestClose() {
	rec.Value.AddHaving(QueryModeNestClose, QueryPrefixNone)
}

func (rec *Query) SetOrderBy(valueList ...interface{}) {
	rec.Value.AddOrderBy(QueryModeDefault, valueList...)
}

func (rec *Query) SetOrderByAsc(valueList ...interface{}) {
	rec.Value.AddOrderBy(QueryModeAsc, valueList...)
}

func (rec *Query) SetOrderByDesc(valueList ...interface{}) {
	rec.Value.AddOrderBy(QueryModeDesc, valueList...)
}

func (rec *Query) SetLimit(limit int) {
	rec.Value.SetLimit(limit)
}

func (rec *Query) SetOffset(offset int) {
	rec.Value.SetOffset(offset)
}
