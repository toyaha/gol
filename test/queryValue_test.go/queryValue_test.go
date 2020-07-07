package query_test

import (
	"fmt"
	"testing"

	"github.com/toyaha/gol"
	"github.com/toyaha/gol/test"
)

func TestQueryValue_GetInsertQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableItem4 := test.Item{}
	tableItem5 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, fromList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, conflictList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.SetTable(gol.QueryModeDefault, val[0])
		}

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		queryValue.SetLimit(limit)

		queryValue.SetOffset(offset)

		query, valueList, err := queryValue.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem1, true},
		}
		fromList := [][]interface{}{
			{&tableItem2, true},
			{&tableItem3, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem4, true}, {gol.QueryJoinModeInner, &tableItem4}},
			{{&tableItem5, true}, {gol.QueryJoinModeLeft, &tableItem5}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Num},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Num, 5},
			{gol.QueryModeDefault, &tableItem1.Str, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`INSERT INTO "item" ("id", "str") VALUES (?, ?), (?, ?)`,
			"[3 a 4 b]",
		}
		fn(t, tableList, fromList, joinList, joinWhereList, valuesColumnList, valuesList, conflictList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetInsertDoNothingQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableItem4 := test.Item{}
	tableItem5 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, fromList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, conflictList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.SetTable(gol.QueryModeDefault, val[0])
		}

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		queryValue.SetLimit(limit)

		queryValue.SetOffset(offset)

		query, valueList, err := queryValue.GetInsertDoNothingQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem1, true},
		}
		fromList := [][]interface{}{
			{&tableItem2, true},
			{&tableItem3, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem4, true}, {gol.QueryJoinModeInner, &tableItem4}},
			{{&tableItem5, true}, {gol.QueryJoinModeLeft, &tableItem5}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Num},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Num, 5},
			{gol.QueryModeDefault, &tableItem1.Str, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`INSERT INTO "item" ("id", "str") VALUES (?, ?), (?, ?) ON CONFLICT DO NOTHING`,
			"[3 a 4 b]",
		}
		fn(t, tableList, fromList, joinList, joinWhereList, valuesColumnList, valuesList, conflictList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetInsertDoUpdateQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableItem4 := test.Item{}
	tableItem5 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, fromList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, conflictList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.SetTable(gol.QueryModeDefault, val[0])
		}

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		queryValue.SetLimit(limit)

		queryValue.SetOffset(offset)

		query, valueList, err := queryValue.GetInsertDoUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem1, true},
		}
		fromList := [][]interface{}{
			{&tableItem2, true},
			{&tableItem3, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem4, true}, {gol.QueryJoinModeInner, &tableItem4}},
			{{&tableItem5, true}, {gol.QueryJoinModeLeft, &tableItem5}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Num},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Num, 5},
			{gol.QueryModeDefault, &tableItem1.Str, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`INSERT INTO "item" ("id", "str") VALUES (?, ?), (?, ?) ON CONFLICT ("id", "num") DO UPDATE SET "num" = ?, "str" = ?`,
			"[3 a 4 b 5 c]",
		}
		fn(t, tableList, fromList, joinList, joinWhereList, valuesColumnList, valuesList, conflictList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetInsertIgnoreQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableItem4 := test.Item{}
	tableItem5 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, fromList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, conflictList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.SetTable(gol.QueryModeDefault, val[0])
		}

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		queryValue.SetLimit(limit)

		queryValue.SetOffset(offset)

		query, valueList, err := queryValue.GetInsertIgnoreQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem1, true},
		}
		fromList := [][]interface{}{
			{&tableItem2, true},
			{&tableItem3, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem4, true}, {gol.QueryJoinModeInner, &tableItem4}},
			{{&tableItem5, true}, {gol.QueryJoinModeLeft, &tableItem5}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Num},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Num, 5},
			{gol.QueryModeDefault, &tableItem1.Str, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`INSERT IGNORE INTO "item" ("id", "str") VALUES (?, ?), (?, ?)`,
			"[3 a 4 b]",
		}
		fn(t, tableList, fromList, joinList, joinWhereList, valuesColumnList, valuesList, conflictList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetInsertOnDepulicateKeyUpdateQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableItem4 := test.Item{}
	tableItem5 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, fromList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, conflictList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.SetTable(gol.QueryModeDefault, val[0])
		}

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		queryValue.SetLimit(limit)

		queryValue.SetOffset(offset)

		query, valueList, err := queryValue.GetInsertOnDuplicateKeyUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem1, true},
		}
		fromList := [][]interface{}{
			{&tableItem2, true},
			{&tableItem3, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem4, true}, {gol.QueryJoinModeInner, &tableItem4}},
			{{&tableItem5, true}, {gol.QueryJoinModeLeft, &tableItem5}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Num},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Num, 5},
			{gol.QueryModeDefault, &tableItem1.Str, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`INSERT INTO "item" ("id", "str") VALUES (?, ?), (?, ?) ON DUPLICATE KEY UPDATE "num" = ?, "str" = ?`,
			"[3 a 4 b 5 c]",
		}
		fn(t, tableList, fromList, joinList, joinWhereList, valuesColumnList, valuesList, conflictList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetUpdateQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableItem4 := test.Item{}
	tableItem5 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, fromList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, conflictList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.SetTable(gol.QueryModeDefault, val[0])
		}

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		queryValue.SetLimit(limit)

		queryValue.SetOffset(offset)

		query, valueList, err := queryValue.GetUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem1, true},
		}
		fromList := [][]interface{}{
			{&tableItem2, true},
			{&tableItem3, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem4, true}, {gol.QueryJoinModeInner, &tableItem4}},
			{{&tableItem5, true}, {gol.QueryJoinModeLeft, &tableItem5}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Num},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Num, 5},
			{gol.QueryModeDefault, &tableItem1.Str, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`UPDATE "item" SET "num" = ?, "str" = ? WHERE "item"."id" = ? AND "item"."id" = ?`,
			"[5 c 6 7]",
		}
		fn(t, tableList, fromList, joinList, joinWhereList, valuesColumnList, valuesList, conflictList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetDeleteQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableItem4 := test.Item{}
	tableItem5 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, fromList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, conflictList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.SetTable(gol.QueryModeDefault, val[0])
		}

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		queryValue.SetLimit(limit)

		queryValue.SetOffset(offset)

		query, valueList, err := queryValue.GetDeleteQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem1, true},
		}
		fromList := [][]interface{}{
			{&tableItem2, true},
			{&tableItem3, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem4, true}, {gol.QueryJoinModeInner, &tableItem4}},
			{{&tableItem5, true}, {gol.QueryJoinModeLeft, &tableItem5}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Num},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Num, 5},
			{gol.QueryModeDefault, &tableItem1.Str, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`DELETE FROM "item" WHERE "item"."id" = ? AND "item"."id" = ?`,
			"[6 7]",
		}
		fn(t, tableList, fromList, joinList, joinWhereList, valuesColumnList, valuesList, conflictList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetSelectQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableItem4 := test.Item{}
	tableItem5 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, fromList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, conflictList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.SetTable(gol.QueryModeDefault, val[0])
		}

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		queryValue.SetLimit(limit)

		queryValue.SetOffset(offset)

		query, valueList, err := queryValue.GetSelectQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem1, true},
		}
		fromList := [][]interface{}{
			{&tableItem2, true},
			{&tableItem3, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem4, true}, {gol.QueryJoinModeInner, &tableItem4}},
			{{&tableItem5, true}, {gol.QueryJoinModeLeft, &tableItem5}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Num},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Num, 5},
			{gol.QueryModeDefault, &tableItem1.Str, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count("t1"."id"), "t1"."str" FROM "item" as "t1", "item" as "t2", "item" as "t3" INNER JOIN "item" as "t4" ON "t4"."id" = "t1"."id" AND "t4"."id" = ? LEFT JOIN "item" as "t5" ON "t5"."id" = "t1"."id" AND "t5"."id" = ? WHERE "t1"."id" = ? AND "t1"."id" = ? GROUP BY "t1"."id", "t1"."str" HAVING "t1"."id" = ? AND "t1"."id" = ? ORDER BY "t1"."id", "t1"."str" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, fromList, joinList, joinWhereList, valuesColumnList, valuesList, conflictList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetSelectCountQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableItem4 := test.Item{}
	tableItem5 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, fromList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, conflictList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.SetTable(gol.QueryModeDefault, val[0])
		}

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		queryValue.SetLimit(limit)

		queryValue.SetOffset(offset)

		query, valueList, err := queryValue.GetSelectCountQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem1, true},
		}
		fromList := [][]interface{}{
			{&tableItem2, true},
			{&tableItem3, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem4, true}, {gol.QueryJoinModeInner, &tableItem4}},
			{{&tableItem5, true}, {gol.QueryJoinModeLeft, &tableItem5}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem4, &tableItem4.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem5, &tableItem5.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Num},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Num, 5},
			{gol.QueryModeDefault, &tableItem1.Str, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Str},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count(*) as count FROM "item" as "t1", "item" as "t2", "item" as "t3" INNER JOIN "item" as "t4" ON "t4"."id" = "t1"."id" AND "t4"."id" = ? LEFT JOIN "item" as "t5" ON "t5"."id" = "t1"."id" AND "t5"."id" = ? WHERE "t1"."id" = ? AND "t1"."id" = ? GROUP BY "t1"."id", "t1"."str" HAVING "t1"."id" = ? AND "t1"."id" = ? ORDER BY "t1"."id", "t1"."str" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, fromList, joinList, joinWhereList, valuesColumnList, valuesList, conflictList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_BuildTable(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		queryValue.SetTable(gol.QueryModeDefault, tableList[0][0])

		query, err := queryValue.BuildTable()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem, true},
		}
		check := `"item"`
		fn(t, tableList, check)
	})
}

func TestQueryValue_BuildTableUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		queryValue.SetTable(gol.QueryModeDefault, tableList[0][0])

		query, err := queryValue.BuildTableUseAs()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem, true},
		}
		check := `"item" as "t1"`
		fn(t, tableList, check)
	})
}

func TestQueryValue_BuildFromTableUseAs(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}

	fn := func(t *testing.T, fromList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range fromList {
			queryValue.AddMeta(val[0], val[1].(bool))
			queryValue.AddFrom(gol.QueryModeDefault, val[0], val[2:]...)
		}

		query, valueList, err := queryValue.BuildFromUseAs()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		fromList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		checkList := []string{`"item" as "t1", "item" as "t2"`, "[]"}
		fn(t, fromList, checkList)
	})

	t.Run("query", func(t *testing.T) {
		fromList := [][]interface{}{
			{&tableItem1, true, "(select 1, ?)", []interface{}{2}},
			{&tableItem2, true, "(select 3, ?)", []interface{}{4}},
		}
		checkList := []string{`(select 1, ?) as "t1", (select 3, ?) as "t2"`, "[2 4]"}
		fn(t, fromList, checkList)
	})
}

func TestQueryValue_BuildJoinUseAs(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, metaList [][]interface{}, joinList [][][]interface{}, joinWhereList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range joinList {
			queryValue.AddMeta(val[0][0], val[0][1].(bool))
			queryValue.AddJoin(val[1][0].(int), val[1][1], val[1][2:]...)
		}

		for _, val := range joinWhereList {
			queryValue.AddJoinWhere(val[0].(int), val[1].(string), val[2], val[3:]...)
		}

		query, valueList, err := queryValue.BuildJoinUseAs()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
		}
		joinList := [][][]interface{}{
			{{&tableItem2, true}, {gol.QueryJoinModeInner, &tableItem2}},
			{{&tableItem3, true}, {gol.QueryJoinModeLeft, &tableItem3}},
			{{&tableTag, true}, {gol.QueryJoinModeRight, &tableTag}},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableTag, &tableTag.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableTag, &tableTag.Id, 3},
		}
		checkList := []string{
			fmt.Sprintf(
				`INNER JOIN "item" as "t2" ON "t2"."id" = "t1"."id" AND "t2"."id" = ? LEFT JOIN "item" as "t3" ON "t3"."id" = "t1"."id" AND "t3"."id" = ? RIGHT JOIN "PUBLIC"."TAG" as "t4" ON "t4"."ID" = "t1"."id" AND "t4"."ID" = ?`,
			),
			"[1 2 3]",
		}
		fn(t, metaList, joinList, joinWhereList, checkList)
	})
}

func TestQueryValue_BuildJoinWhereUseAs(t *testing.T) {
	// supported by TestQueryValue_BuildJoin
}

func TestQueryValue_BuildValuesColumn(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, valuesColumnList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range valuesColumnList {
			queryValue.AddValuesColumn(val[0].(int), val[1:]...)
		}

		query, err := queryValue.BuildValuesColumn()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Str},
		}
		check := `("id", "str")`
		fn(t, metaList, valuesColumnList, check)
	})
}

func TestQueryValue_BuildValues(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, valuesList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		query, valueList, err := queryValue.BuildValues()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 1, "a"},
			{gol.QueryModeDefault, 2, "b"},
		}
		checkList := []string{
			"(?, ?), (?, ?)",
			"[1 a 2 b]",
		}
		fn(t, metaList, valuesList, checkList)
	})
}

func TestQueryValue_BuildConflict(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, conflictList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range conflictList {
			queryValue.AddConflict(val[0].(int), val[1:]...)
		}

		query, valueList, err := queryValue.BuildConflict()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		conflictList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Str},
		}
		checkList := []string{
			`"id", "str"`,
			"[]",
		}
		fn(t, metaList, conflictList, checkList)
	})
}

func TestQueryValue_BuildSet(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, setList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range setList {
			queryValue.AddSet(val[0].(int), val[1:]...)
		}

		query, valueList, err := queryValue.BuildSet()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id, 1},
			{gol.QueryModeDefault, &tableItem.Str, "a"},
		}
		checkList := []string{
			`"id" = ?, "str" = ?`,
			"[1 a]",
		}
		fn(t, metaList, setList, checkList)
	})
}

func TestQueryValue_BuildSelect(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, selectList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		query, valueList, err := queryValue.BuildSelect()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeDefault, &tableItem.Str},
			{gol.QueryModeDefault, &tableItem.Str},
		}
		checkList := []string{
			`SELECT count("item"."id"), "item"."str", "item"."str"`,
			"[]",
		}
		fn(t, metaList, selectList, checkList)
	})

	t.Run("all", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := [][]interface{}{
			{gol.QueryModeAll, &tableItem},
		}
		checkList := []string{
			`SELECT "item".*`,
			"[]",
		}
		fn(t, metaList, selectList, checkList)
	})

	t.Run("query", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "(select a from b where c = ?)", []interface{}{1}},
			{gol.QueryModeDefault, "(select d from e where f = ?)", []interface{}{2}},
		}
		checkList := []string{
			`SELECT (select a from b where c = ?), (select d from e where f = ?)`,
			"[1 2]",
		}
		fn(t, metaList, selectList, checkList)
	})
}

func TestQueryValue_BuildSelectUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, selectList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		query, valueList, err := queryValue.BuildSelectUseAs()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeDefault, &tableItem.Str},
			{gol.QueryModeDefault, &tableItem.Str},
		}
		checkList := []string{
			`SELECT count("t1"."id"), "t1"."str", "t1"."str"`,
			"[]",
		}
		fn(t, metaList, selectList, checkList)
	})

	t.Run("all", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := [][]interface{}{
			{gol.QueryModeAll, &tableItem},
		}
		checkList := []string{
			`SELECT "t1".*`,
			"[]",
		}
		fn(t, metaList, selectList, checkList)
	})

	t.Run("query", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "(select a from b where c = ?)", []interface{}{1}},
			{gol.QueryModeDefault, "(select d from e where f = ?)", []interface{}{2}},
		}
		checkList := []string{
			`SELECT (select a from b where c = ?), (select d from e where f = ?)`,
			"[1 2]",
		}
		fn(t, metaList, selectList, checkList)
	})
}

func TestQueryValue_BuildWhere(t *testing.T) {
	tableItem := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, metaList [][]interface{}, whereList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		query, valueList, err := queryValue.BuildWhere()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
			{&tableTag, true},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Str, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Str, "b"},
			{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{3, 4, 5}},
			{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{6, 7, 8}},
			{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 9},
			{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 10},
			{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 11},
			{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 12},
			{gol.QueryModeNest, gol.QueryPrefixAnd},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, "count(", &tableItem.Id, ") = ? ?", []interface{}{13, 14}},
			{gol.QueryModeNestClose, gol.QueryPrefixAnd},
			{gol.QueryModeIs, gol.QueryPrefixOr, &tableTag.Id, 15},
			{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableTag.Id, 16},
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Str, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Str, "d"},
			{gol.QueryModeIn, gol.QueryPrefixOr, &tableTag.Id, []interface{}{17, 18, 19}},
			{gol.QueryModeInNot, gol.QueryPrefixOr, &tableTag.Id, []interface{}{20, 21, 22}},
			{gol.QueryModeGt, gol.QueryPrefixOr, &tableTag.Id, 23},
			{gol.QueryModeGte, gol.QueryPrefixOr, &tableTag.Id, 24},
			{gol.QueryModeLt, gol.QueryPrefixOr, &tableTag.Id, 25},
			{gol.QueryModeLte, gol.QueryPrefixOr, &tableTag.Id, 26},
			{gol.QueryModeNest, gol.QueryPrefixOr},
			{gol.QueryModeDefault, gol.QueryPrefixOr, "count(", &tableTag.Id, ") = ? ?", []interface{}{27, 28}},
			{gol.QueryModeNestClose, gol.QueryPrefixOr},
		}
		checkList := []string{
			`WHERE "item"."id" = ? AND "item"."id" = ? AND "item"."id" != ? AND "item"."str" LIKE ? AND "item"."str" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( count("item"."id") = ? ? ) OR "PUBLIC"."TAG"."ID" = ? OR "PUBLIC"."TAG"."ID" != ? OR "PUBLIC"."TAG"."STR" LIKE ? OR "PUBLIC"."TAG"."STR" NOT LIKE ? OR "PUBLIC"."TAG"."ID" IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" NOT IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" > ? OR "PUBLIC"."TAG"."ID" >= ? OR "PUBLIC"."TAG"."ID" < ? OR "PUBLIC"."TAG"."ID" <= ? OR ( count("PUBLIC"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, metaList, whereList, checkList)
	})
}

func TestQueryValue_BuildWhereUseAs(t *testing.T) {
	tableItem := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, metaList [][]interface{}, whereList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range whereList {
			queryValue.AddWhere(val[0].(int), val[1].(string), val[2:]...)
		}

		query, valueList, err := queryValue.BuildWhereUseAs()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
			{&tableTag, true},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Str, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Str, "b"},
			{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{3, 4, 5}},
			{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{6, 7, 8}},
			{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 9},
			{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 10},
			{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 11},
			{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 12},
			{gol.QueryModeNest, gol.QueryPrefixAnd},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, "count(", &tableItem.Id, ") = ? ?", []interface{}{13, 14}},
			{gol.QueryModeNestClose, gol.QueryPrefixAnd},
			{gol.QueryModeIs, gol.QueryPrefixOr, &tableTag.Id, 15},
			{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableTag.Id, 16},
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Str, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Str, "d"},
			{gol.QueryModeIn, gol.QueryPrefixOr, &tableTag.Id, []interface{}{17, 18, 19}},
			{gol.QueryModeInNot, gol.QueryPrefixOr, &tableTag.Id, []interface{}{20, 21, 22}},
			{gol.QueryModeGt, gol.QueryPrefixOr, &tableTag.Id, 23},
			{gol.QueryModeGte, gol.QueryPrefixOr, &tableTag.Id, 24},
			{gol.QueryModeLt, gol.QueryPrefixOr, &tableTag.Id, 25},
			{gol.QueryModeLte, gol.QueryPrefixOr, &tableTag.Id, 26},
			{gol.QueryModeNest, gol.QueryPrefixOr},
			{gol.QueryModeDefault, gol.QueryPrefixOr, "count(", &tableTag.Id, ") = ? ?", []interface{}{27, 28}},
			{gol.QueryModeNestClose, gol.QueryPrefixOr},
		}
		checkList := []string{
			`WHERE "t1"."id" = ? AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."str" LIKE ? AND "t1"."str" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( count("t1"."id") = ? ? ) OR "t2"."ID" = ? OR "t2"."ID" != ? OR "t2"."STR" LIKE ? OR "t2"."STR" NOT LIKE ? OR "t2"."ID" IN (?, ?, ?) OR "t2"."ID" NOT IN (?, ?, ?) OR "t2"."ID" > ? OR "t2"."ID" >= ? OR "t2"."ID" < ? OR "t2"."ID" <= ? OR ( count("t2"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, metaList, whereList, checkList)
	})
}

func TestQueryValue_BuildGroupBy(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, groupByList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		query, err := queryValue.BuildGroupBy()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Str},
		}
		check := `GROUP BY "item"."id", "item"."str"`
		fn(t, metaList, groupByList, check)
	})
}

func TestQueryValue_BuildGroupByUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, groupByList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range groupByList {
			queryValue.AddGroupBy(val[0].(int), val[1])
		}

		query, err := queryValue.BuildGroupByUseAs()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Str},
		}
		check := `GROUP BY "t1"."id", "t1"."str"`
		fn(t, metaList, groupByList, check)
	})
}

func TestQueryValue_BuildHaving(t *testing.T) {
	tableItem := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, metaList [][]interface{}, havingList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		query, valueList, err := queryValue.BuildHaving()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
			{&tableTag, true},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Str, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Str, "b"},
			{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{3, 4, 5}},
			{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{6, 7, 8}},
			{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 9},
			{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 10},
			{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 11},
			{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 12},
			{gol.QueryModeNest, gol.QueryPrefixAnd},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, "count(", &tableItem.Id, ") = ? ?", []interface{}{13, 14}},
			{gol.QueryModeNestClose, gol.QueryPrefixAnd},
			{gol.QueryModeIs, gol.QueryPrefixOr, &tableTag.Id, 15},
			{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableTag.Id, 16},
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Str, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Str, "d"},
			{gol.QueryModeIn, gol.QueryPrefixOr, &tableTag.Id, []interface{}{17, 18, 19}},
			{gol.QueryModeInNot, gol.QueryPrefixOr, &tableTag.Id, []interface{}{20, 21, 22}},
			{gol.QueryModeGt, gol.QueryPrefixOr, &tableTag.Id, 23},
			{gol.QueryModeGte, gol.QueryPrefixOr, &tableTag.Id, 24},
			{gol.QueryModeLt, gol.QueryPrefixOr, &tableTag.Id, 25},
			{gol.QueryModeLte, gol.QueryPrefixOr, &tableTag.Id, 26},
			{gol.QueryModeNest, gol.QueryPrefixOr},
			{gol.QueryModeDefault, gol.QueryPrefixOr, "count(", &tableTag.Id, ") = ? ?", []interface{}{27, 28}},
			{gol.QueryModeNestClose, gol.QueryPrefixOr},
		}
		checkList := []string{
			`HAVING "item"."id" = ? AND "item"."id" = ? AND "item"."id" != ? AND "item"."str" LIKE ? AND "item"."str" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( count("item"."id") = ? ? ) OR "PUBLIC"."TAG"."ID" = ? OR "PUBLIC"."TAG"."ID" != ? OR "PUBLIC"."TAG"."STR" LIKE ? OR "PUBLIC"."TAG"."STR" NOT LIKE ? OR "PUBLIC"."TAG"."ID" IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" NOT IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" > ? OR "PUBLIC"."TAG"."ID" >= ? OR "PUBLIC"."TAG"."ID" < ? OR "PUBLIC"."TAG"."ID" <= ? OR ( count("PUBLIC"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, metaList, havingList, checkList)
	})
}

func TestQueryValue_BuildHavingUseAs(t *testing.T) {
	tableItem := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, metaList [][]interface{}, havingList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range havingList {
			queryValue.AddHaving(val[0].(int), val[1].(string), val[2:]...)
		}

		query, valueList, err := queryValue.BuildHavingUseAs()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
			{&tableTag, true},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Str, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Str, "b"},
			{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{3, 4, 5}},
			{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{6, 7, 8}},
			{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 9},
			{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 10},
			{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 11},
			{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 12},
			{gol.QueryModeNest, gol.QueryPrefixAnd},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, "count(", &tableItem.Id, ") = ? ?", []interface{}{13, 14}},
			{gol.QueryModeNestClose, gol.QueryPrefixAnd},
			{gol.QueryModeIs, gol.QueryPrefixOr, &tableTag.Id, 15},
			{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableTag.Id, 16},
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Str, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Str, "d"},
			{gol.QueryModeIn, gol.QueryPrefixOr, &tableTag.Id, []interface{}{17, 18, 19}},
			{gol.QueryModeInNot, gol.QueryPrefixOr, &tableTag.Id, []interface{}{20, 21, 22}},
			{gol.QueryModeGt, gol.QueryPrefixOr, &tableTag.Id, 23},
			{gol.QueryModeGte, gol.QueryPrefixOr, &tableTag.Id, 24},
			{gol.QueryModeLt, gol.QueryPrefixOr, &tableTag.Id, 25},
			{gol.QueryModeLte, gol.QueryPrefixOr, &tableTag.Id, 26},
			{gol.QueryModeNest, gol.QueryPrefixOr},
			{gol.QueryModeDefault, gol.QueryPrefixOr, "count(", &tableTag.Id, ") = ? ?", []interface{}{27, 28}},
			{gol.QueryModeNestClose, gol.QueryPrefixOr},
		}
		checkList := []string{
			`HAVING "t1"."id" = ? AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."str" LIKE ? AND "t1"."str" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( count("t1"."id") = ? ? ) OR "t2"."ID" = ? OR "t2"."ID" != ? OR "t2"."STR" LIKE ? OR "t2"."STR" NOT LIKE ? OR "t2"."ID" IN (?, ?, ?) OR "t2"."ID" NOT IN (?, ?, ?) OR "t2"."ID" > ? OR "t2"."ID" >= ? OR "t2"."ID" < ? OR "t2"."ID" <= ? OR ( count("t2"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, metaList, havingList, checkList)
	})
}

func TestQueryValue_BuildOrderBy(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, orderByList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		query, err := queryValue.BuildOrderBy()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeAsc, &tableItem.Str},
			{gol.QueryModeDesc, &tableItem.Str},
		}
		check := `ORDER BY count("item"."id"), "item"."str" ASC, "item"."str" DESC`
		fn(t, metaList, orderByList, check)
	})
}

func TestQueryValue_BuildOrderByUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, orderByList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range orderByList {
			queryValue.AddOrderBy(val[0].(int), val[1:]...)
		}

		query, err := queryValue.BuildOrderByUseAs()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeAsc, &tableItem.Str},
			{gol.QueryModeDesc, &tableItem.Str},
		}
		check := `ORDER BY count("t1"."id"), "t1"."str" ASC, "t1"."str" DESC`
		fn(t, metaList, orderByList, check)
	})
}

func TestQueryValue_BuildLimit(t *testing.T) {
	fn := func(t *testing.T, limit int, check string) {
		queryValue := gol.NewQueryValue(nil)

		queryValue.SetLimit(limit)

		query, err := queryValue.BuildLimit()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("10", func(t *testing.T) {
		limit := 10
		check := fmt.Sprintf(`LIMIT %v`, limit)
		fn(t, limit, check)
	})

	t.Run("1000", func(t *testing.T) {
		limit := 200
		check := fmt.Sprintf(`LIMIT %v`, limit)
		fn(t, limit, check)
	})
}

func TestQueryValue_BuildOffset(t *testing.T) {
	fn := func(t *testing.T, offset int, check string) {
		queryValue := gol.NewQueryValue(nil)

		queryValue.SetOffset(offset)

		query, err := queryValue.BuildOffset()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("10", func(t *testing.T) {
		offset := 10
		check := fmt.Sprintf(`OFFSET %v`, offset)
		fn(t, offset, check)
	})

	t.Run("1000", func(t *testing.T) {
		offset := 1000
		check := fmt.Sprintf(`OFFSET %v`, offset)
		fn(t, offset, check)
	})
}

func TestQueryTable_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryTable{}
		data.Set(gol.QueryModeDefault, tableList[0][0])

		query, err := data.Build(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem, true},
		}
		check := `"item"`
		fn(t, tableList, check)
	})
}

func TestQueryTable_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryTable{}
		data.Set(gol.QueryModeDefault, tableList[0][0])

		query, err := data.BuildUseAs(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		tableList := [][]interface{}{
			{&tableItem, true},
		}
		check := `"item" as "t1"`
		fn(t, tableList, check)
	})
}

func TestQueryFrom_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, fromList [][]interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range fromList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryFrom{}
		data.Set(gol.QueryModeDefault, fromList[0][0], fromList[0][2:]...)

		query, valueList, err := data.BuildUseAs(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		fromList := [][]interface{}{
			{&tableItem, true},
		}
		checkList := []string{`"item" as "t1"`, "[]"}
		fn(t, fromList, checkList)
	})

	t.Run("query", func(t *testing.T) {
		fromList := [][]interface{}{
			{&tableItem, true, "(select 1)"},
		}
		checkList := []string{`(select 1) as "t1"`, "[]"}
		fn(t, fromList, checkList)
	})
}

func TestQueryJoin_BuildUseAs(t *testing.T) {
	tableItem1 := test.Item{}

	fn := func(t *testing.T, joinList [][][]interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range joinList {
			err := meta.Add(val[0][0], val[0][1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryJoin{}
		data.Set(joinList[0][1][0].(int), joinList[0][1][1], joinList[0][1][2:]...)

		query, valueList, err := data.BuildUseAs(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("inner", func(t *testing.T) {
		joinList := [][][]interface{}{
			{{&tableItem1, true}, {gol.QueryJoinModeInner, &tableItem1}},
		}
		checkList := []string{`INNER JOIN "item" as "t1" ON`, "[]"}
		fn(t, joinList, checkList)
	})

	t.Run("left", func(t *testing.T) {
		joinList := [][][]interface{}{
			{{&tableItem1, true}, {gol.QueryJoinModeLeft, &tableItem1}},
		}
		checkList := []string{`LEFT JOIN "item" as "t1" ON`, "[]"}
		fn(t, joinList, checkList)
	})

	t.Run("right", func(t *testing.T) {
		joinList := [][][]interface{}{
			{{&tableItem1, true}, {gol.QueryJoinModeRight, &tableItem1}},
		}
		checkList := []string{`RIGHT JOIN "item" as "t1" ON`, "[]"}
		fn(t, joinList, checkList)
	})

	t.Run("query", func(t *testing.T) {
		joinList := [][][]interface{}{
			{{&tableItem1, true}, {gol.QueryJoinModeInner, &tableItem1, "(select ?)", []interface{}{1}}},
		}
		checkList := []string{`INNER JOIN (select ?) as "t1" ON`, "[1]"}
		fn(t, joinList, checkList)
	})
}

func TestQueryJoinWhere_BuildUseAs(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, joinWhereList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryJoinWhere{}
		data.Set(joinWhereList[0].(int), joinWhereList[1].(string), joinWhereList[2], joinWhereList[3:]...)

		prefix, query, valueList, err := data.BuildUseAs(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target, check := prefix, checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := query, checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := fmt.Sprintf("%+v", valueList), checkList[2]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("and default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Str, "a"}
		checkList := []string{
			"AND",
			`"t1"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Str, "a"}
		checkList := []string{
			"AND",
			`"t1"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"t1"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"t1"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"OR",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixOr, &tableItem1, &tableItem1.Str, "a"}
		checkList := []string{
			"OR",
			`"t1"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableItem1, &tableItem1.Str, "a"}
		checkList := []string{
			"OR",
			`"t1"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"t1"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"t1"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("and nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"AND",
			`(`,
			"[]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("or nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixOr, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"OR",
			`(`,
			"[]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})

	t.Run("nest close", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem1, true},
			{&tableItem2, true},
		}
		joinWhereList := []interface{}{gol.QueryModeNestClose, gol.QueryPrefixAnd, &tableItem1, &tableItem1.Id, 1}
		checkList := []string{
			"AND",
			`)`,
			"[]",
		}
		fn(t, metaList, joinWhereList, checkList)
	})
}

func TestQueryValuesColumn_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, valuesColumnList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryValuesColumn{}
		data.Set(valuesColumnList[0].(int), valuesColumnList[1:]...)

		query, err := data.Build(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		valuesColumnList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"id"`
		fn(t, metaList, valuesColumnList, check)
	})
}

func TestQueryValues_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, valuesList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryValues{}
		data.Set(valuesList[0].(int), valuesList[1:]...)

		query, valueList, err := data.Build(nil)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		valuesList := []interface{}{gol.QueryModeDefault, 1, 2, 3}
		checkList := []string{
			"(?, ?, ?)",
			"[1 2 3]",
		}
		fn(t, metaList, valuesList, checkList)
	})
}

func TestQueryConflict_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, conflictList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryConflict{}
		data.Set(conflictList[0].(int), conflictList[1:]...)

		query, valueList, err := data.Build(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		conflictList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		checkList := []string{
			`"id"`,
			"[]",
		}
		fn(t, metaList, conflictList, checkList)
	})
}

func TestQuerySet_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, setList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QuerySet{}
		data.Set(setList[0].(int), setList[1], setList[2])

		query, value, err := data.Build(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", value)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		setList := []interface{}{gol.QueryModeDefault, &tableItem.Id, 1}
		checkList := []string{
			`"id" = ?`,
			`1`,
		}
		fn(t, metaList, setList, checkList)
	})
}

func TestQuerySelect_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, selectList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QuerySelect{}
		data.Set(selectList[0].(int), selectList[1:]...)

		query, valueList, err := data.Build(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		checkList := []string{`count("item"."id")`, "[]"}
		fn(t, metaList, selectList, checkList)
	})

	t.Run("all", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := []interface{}{gol.QueryModeAll, &tableItem}
		checkList := []string{`"item".*`, "[]"}
		fn(t, metaList, selectList, checkList)
	})

	t.Run("query", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := []interface{}{gol.QueryModeDefault, "(select a from b where c = ?)", []interface{}{1}}
		checkList := []string{`(select a from b where c = ?)`, "[1]"}
		fn(t, metaList, selectList, checkList)
	})
}

func TestQuerySelect_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, selectList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QuerySelect{}
		data.Set(selectList[0].(int), selectList[1:]...)

		query, valueList, err := data.BuildUseAs(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			check := checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}

		{
			target := fmt.Sprintf("%+v", valueList)
			check := checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		checkList := []string{`count("t1"."id")`, "[]"}
		fn(t, metaList, selectList, checkList)
	})

	t.Run("all", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := []interface{}{gol.QueryModeAll, &tableItem}
		checkList := []string{`"t1".*`, "[]"}
		fn(t, metaList, selectList, checkList)
	})

	t.Run("query", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := []interface{}{gol.QueryModeDefault, "(select a from b where c = ?)", []interface{}{1}}
		checkList := []string{`(select a from b where c = ?)`, "[1]"}
		fn(t, metaList, selectList, checkList)
	})
}

func TestQueryWhere_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, whereList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryWhere{}
		data.Set(whereList[0].(int), whereList[1].(string), whereList[2:]...)

		prefix, query, valueList, err := data.Build(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target, check := prefix, checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := query, checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := fmt.Sprintf("%+v", valueList), checkList[2]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("and default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Str, "a"}
		checkList := []string{
			"AND",
			`"item"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Str, "a"}
		checkList := []string{
			"AND",
			`"item"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixOr, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixOr, &tableItem.Str, "a"}
		checkList := []string{
			"OR",
			`"item"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableItem.Str, "a"}
		checkList := []string{
			"OR",
			`"item"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`(`,
			"[]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`(`,
			"[]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("nest close", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeNestClose, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`)`,
			"[]",
		}
		fn(t, metaList, whereList, checkList)
	})
}

func TestQueryWhere_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, whereList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryWhere{}
		data.Set(whereList[0].(int), whereList[1].(string), whereList[2:]...)

		prefix, query, valueList, err := data.BuildUseAs(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target, check := prefix, checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := query, checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := fmt.Sprintf("%+v", valueList), checkList[2]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("and default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Str, "a"}
		checkList := []string{
			"AND",
			`"t1"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Str, "a"}
		checkList := []string{
			"AND",
			`"t1"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"t1"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"t1"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixOr, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"OR",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixOr, &tableItem.Str, "a"}
		checkList := []string{
			"OR",
			`"t1"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableItem.Str, "a"}
		checkList := []string{
			"OR",
			`"t1"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"t1"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"t1"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("and nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`(`,
			"[]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("or nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`(`,
			"[]",
		}
		fn(t, metaList, whereList, checkList)
	})

	t.Run("nest close", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		whereList := []interface{}{gol.QueryModeNestClose, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`)`,
			"[]",
		}
		fn(t, metaList, whereList, checkList)
	})
}

func TestQueryGroupBy_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, groupByList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
				return
			}
		}

		data := &gol.QueryGroupBy{}
		data.Set(groupByList[0].(int), groupByList[1:]...)
		query, err := data.Build(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		groupByList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"item"."id"`
		fn(t, metaList, groupByList, check)
	})
}

func TestQueryGroupBy_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, groupByList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
				return
			}
		}

		data := &gol.QueryGroupBy{}
		data.Set(groupByList[0].(int), groupByList[1:]...)
		query, err := data.BuildUseAs(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		groupByList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"t1"."id"`
		fn(t, metaList, groupByList, check)
	})
}

func TestQueryHaving_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, havingList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryHaving{}
		data.Set(havingList[0].(int), havingList[1].(string), havingList[2:]...)

		prefix, query, valueList, err := data.Build(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target, check := prefix, checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := query, checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := fmt.Sprintf("%+v", valueList), checkList[2]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("and default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Str, "a"}
		checkList := []string{
			"AND",
			`"item"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Str, "a"}
		checkList := []string{
			"AND",
			`"item"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixOr, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIs, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLike, gol.QueryPrefixOr, &tableItem.Str, "a"}
		checkList := []string{
			"OR",
			`"item"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableItem.Str, "a"}
		checkList := []string{
			"OR",
			`"item"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIn, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeGt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeGte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeNest, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`(`,
			"[]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeNest, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`(`,
			"[]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("nest close", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeNestClose, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`)`,
			"[]",
		}
		fn(t, metaList, havingList, checkList)
	})
}

func TestQueryHaving_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, havingList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryHaving{}
		data.Set(havingList[0].(int), havingList[1].(string), havingList[2:]...)

		prefix, query, valueList, err := data.BuildUseAs(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target, check := prefix, checkList[0]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := query, checkList[1]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
		{
			target, check := fmt.Sprintf("%+v", valueList), checkList[2]
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("and default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Str, "a"}
		checkList := []string{
			"AND",
			`"t1"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Str, "a"}
		checkList := []string{
			"AND",
			`"t1"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"t1"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"t1"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"t1"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixOr, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"OR",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or is", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIs, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or is not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" != ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or like", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLike, gol.QueryPrefixOr, &tableItem.Str, "a"}
		checkList := []string{
			"OR",
			`"t1"."str" LIKE ?`,
			"[a]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or like not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableItem.Str, "a"}
		checkList := []string{
			"OR",
			`"t1"."str" NOT LIKE ?`,
			"[a]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or in", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeIn, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"t1"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or in not", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"t1"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or gt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeGt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" > ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or gte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeGte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" >= ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or lt", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" < ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or lte", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeLte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"t1"."id" <= ?`,
			"[1]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("and nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeNest, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`(`,
			"[]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("or nest", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeNest, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`(`,
			"[]",
		}
		fn(t, metaList, havingList, checkList)
	})

	t.Run("nest close", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		havingList := []interface{}{gol.QueryModeNestClose, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`)`,
			"[]",
		}
		fn(t, metaList, havingList, checkList)
	})
}

func TestQueryOrderBy_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, orderByList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
				return
			}
		}

		data := &gol.QueryOrderBy{}
		data.Set(orderByList[0].(int), orderByList[1:]...)
		query, err := data.Build(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("item"."id")`
		fn(t, metaList, orderByList, check)
	})

	t.Run("asc", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := []interface{}{gol.QueryModeAsc, "count(", &tableItem.Id, ")"}
		check := `count("item"."id") ASC`
		fn(t, metaList, orderByList, check)
	})

	t.Run("desc", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := []interface{}{gol.QueryModeDesc, "count(", &tableItem.Id, ")"}
		check := `count("item"."id") DESC`
		fn(t, metaList, orderByList, check)
	})
}

func TestQueryOrderBy_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, orderByList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range metaList {
			err := meta.Add(val[0], val[1].(bool))
			if err != nil {
				t.Error(err)
				return
			}
		}

		data := &gol.QueryOrderBy{}
		data.Set(orderByList[0].(int), orderByList[1:]...)
		query, err := data.BuildUseAs(meta)
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := query
			if target != check {
				t.Error("\ntarget:", target, "\ncheck :", check)
			}
		}
	}

	t.Run("default", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id")`
		fn(t, metaList, orderByList, check)
	})

	t.Run("asc", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := []interface{}{gol.QueryModeAsc, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id") ASC`
		fn(t, metaList, orderByList, check)
	})

	t.Run("desc", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := []interface{}{gol.QueryModeDesc, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id") DESC`
		fn(t, metaList, orderByList, check)
	})
}
