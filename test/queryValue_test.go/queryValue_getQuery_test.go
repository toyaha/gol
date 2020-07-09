package queryValue_test

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

func TestQueryValue_GetInsertSelectUnionQuery(t *testing.T) {
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

		query, valueList, err := queryValue.GetInsertSelectUnionQuery()
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
			`INSERT INTO "item" ("id", "str") SELECT ?, ? UNION SELECT ?, ?`,
			"[3 a 4 b]",
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
