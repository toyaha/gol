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

	fn := func(t *testing.T, tableList [][]interface{}, joinList [][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
		}

		queryValue.SetTable(gol.QueryModeDefault, tableList[0][1])

		for _, val := range joinList {
			queryValue.AddJoin(val[0].(int), val[1])
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
			{"", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`INSERT INTO "item" ("id", "name") VALUES (?, ?), (?, ?)`,
			"[3 a 4 b]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
			{"", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`INSERT INTO "item" ("id", "name") VALUES (?, ?), (?, ?)`,
			"[3 a 4 b]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
			{"s3", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`INSERT INTO "s1"."item" ("id", "name") VALUES (?, ?), (?, ?)`,
			"[3 a 4 b]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
			{"s3", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`INSERT INTO "s1"."item" ("id", "name") VALUES (?, ?), (?, ?)`,
			"[3 a 4 b]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetUpdateQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, joinList [][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
		}

		queryValue.SetTable(gol.QueryModeDefault, tableList[0][1])

		for _, val := range joinList {
			queryValue.AddJoin(val[0].(int), val[1])
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
			{"", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`UPDATE "item" SET "id" = ?, "name" = ? WHERE "item"."id" = ? AND "item"."id" = ?`,
			"[5 c 6 7]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
			{"", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`UPDATE "item" SET "id" = ?, "name" = ? WHERE "item"."id" = ? AND "item"."id" = ?`,
			"[5 c 6 7]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
			{"s3", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`UPDATE "s1"."item" SET "id" = ?, "name" = ? WHERE "s1"."item"."id" = ? AND "s1"."item"."id" = ?`,
			"[5 c 6 7]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
			{"s3", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`UPDATE "s1"."item" SET "id" = ?, "name" = ? WHERE "s1"."item"."id" = ? AND "s1"."item"."id" = ?`,
			"[5 c 6 7]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetDeleteQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, joinList [][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
		}

		queryValue.SetTable(gol.QueryModeDefault, tableList[0][1])

		for _, val := range joinList {
			queryValue.AddJoin(val[0].(int), val[1])
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
			{"", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`DELETE FROM "item" WHERE "item"."id" = ? AND "item"."id" = ?`,
			"[6 7]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
			{"", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`DELETE FROM "item" WHERE "item"."id" = ? AND "item"."id" = ?`,
			"[6 7]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
			{"s3", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`DELETE FROM "s1"."item" WHERE "s1"."item"."id" = ? AND "s1"."item"."id" = ?`,
			"[6 7]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
			{"s3", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`DELETE FROM "s1"."item" WHERE "s1"."item"."id" = ? AND "s1"."item"."id" = ?`,
			"[6 7]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetSelectQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, joinList [][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
		}

		queryValue.SetTable(gol.QueryModeDefault, tableList[0][1])

		for _, val := range joinList {
			queryValue.AddJoin(val[0].(int), val[1])
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
			{"", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count("item"."id"), "item"."name" FROM "item" INNER JOIN "item" ON "item"."id" = "item"."id" AND "item"."id" = ? LEFT JOIN "item" ON "item"."id" = "item"."id" AND "item"."id" = ? WHERE "item"."id" = ? AND "item"."id" = ? GROUP BY "item"."id", "item"."name" HAVING "item"."id" = ? AND "item"."id" = ? ORDER BY "item"."id", "item"."name" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
			{"", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count("t1"."id"), "t1"."name" FROM "item" as "t1" INNER JOIN "item" as "t2" ON "t2"."id" = "t1"."id" AND "t2"."id" = ? LEFT JOIN "item" as "t3" ON "t3"."id" = "t1"."id" AND "t3"."id" = ? WHERE "t1"."id" = ? AND "t1"."id" = ? GROUP BY "t1"."id", "t1"."name" HAVING "t1"."id" = ? AND "t1"."id" = ? ORDER BY "t1"."id", "t1"."name" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
			{"s3", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count("s1"."item"."id"), "s1"."item"."name" FROM "s1"."item" INNER JOIN "s2"."item" ON "s2"."item"."id" = "s1"."item"."id" AND "s2"."item"."id" = ? LEFT JOIN "s3"."item" ON "s3"."item"."id" = "s1"."item"."id" AND "s3"."item"."id" = ? WHERE "s1"."item"."id" = ? AND "s1"."item"."id" = ? GROUP BY "s1"."item"."id", "s1"."item"."name" HAVING "s1"."item"."id" = ? AND "s1"."item"."id" = ? ORDER BY "s1"."item"."id", "s1"."item"."name" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
			{"s3", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count("t1"."id"), "t1"."name" FROM "s1"."item" as "t1" INNER JOIN "s2"."item" as "t2" ON "t2"."id" = "t1"."id" AND "t2"."id" = ? LEFT JOIN "s3"."item" as "t3" ON "t3"."id" = "t1"."id" AND "t3"."id" = ? WHERE "t1"."id" = ? AND "t1"."id" = ? GROUP BY "t1"."id", "t1"."name" HAVING "t1"."id" = ? AND "t1"."id" = ? ORDER BY "t1"."id", "t1"."name" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_GetSelectCountQuery(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, joinList [][]interface{}, joinWhereList [][]interface{}, valuesColumnList [][]interface{}, valuesList [][]interface{}, setList [][]interface{}, selectList [][]interface{}, whereList [][]interface{}, groupByList [][]interface{}, havingList [][]interface{}, orderByList [][]interface{}, limit int, offset int, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
		}

		queryValue.SetTable(gol.QueryModeDefault, tableList[0][1])

		for _, val := range joinList {
			queryValue.AddJoin(val[0].(int), val[1])
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
			{"", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count(*) as count FROM "item" INNER JOIN "item" ON "item"."id" = "item"."id" AND "item"."id" = ? LEFT JOIN "item" ON "item"."id" = "item"."id" AND "item"."id" = ? WHERE "item"."id" = ? AND "item"."id" = ? GROUP BY "item"."id", "item"."name" HAVING "item"."id" = ? AND "item"."id" = ? ORDER BY "item"."id", "item"."name" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
			{"", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count(*) as count FROM "item" as "t1" INNER JOIN "item" as "t2" ON "t2"."id" = "t1"."id" AND "t2"."id" = ? LEFT JOIN "item" as "t3" ON "t3"."id" = "t1"."id" AND "t3"."id" = ? WHERE "t1"."id" = ? AND "t1"."id" = ? GROUP BY "t1"."id", "t1"."name" HAVING "t1"."id" = ? AND "t1"."id" = ? ORDER BY "t1"."id", "t1"."name" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
			{"s3", &tableItem3, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count(*) as count FROM "s1"."item" INNER JOIN "s2"."item" ON "s2"."item"."id" = "s1"."item"."id" AND "s2"."item"."id" = ? LEFT JOIN "s3"."item" ON "s3"."item"."id" = "s1"."item"."id" AND "s3"."item"."id" = ? WHERE "s1"."item"."id" = ? AND "s1"."item"."id" = ? GROUP BY "s1"."item"."id", "s1"."item"."name" HAVING "s1"."item"."id" = ? AND "s1"."item"."id" = ? ORDER BY "s1"."item"."id", "s1"."item"."name" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
			{"s3", &tableItem3, "t3"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
		}
		joinWhereList := [][]interface{}{
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ", &tableItem1.Id},
			{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, " = ", &tableItem1.Id},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem3, &tableItem3.Id, 2},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 3, "a"},
			{gol.QueryModeDefault, 4, "b"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id, 5},
			{gol.QueryModeDefault, &tableItem1.Name, "c"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem1.Id, ")"},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 6},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 7},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 8},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem1.Id, 9},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem1.Id},
			{gol.QueryModeDefault, &tableItem1.Name},
		}
		limit := 10
		offset := 100
		checkList := []string{
			`SELECT count(*) as count FROM "s1"."item" as "t1" INNER JOIN "s2"."item" as "t2" ON "t2"."id" = "t1"."id" AND "t2"."id" = ? LEFT JOIN "s3"."item" as "t3" ON "t3"."id" = "t1"."id" AND "t3"."id" = ? WHERE "t1"."id" = ? AND "t1"."id" = ? GROUP BY "t1"."id", "t1"."name" HAVING "t1"."id" = ? AND "t1"."id" = ? ORDER BY "t1"."id", "t1"."name" LIMIT 10 OFFSET 100`,
			"[1 2 6 7 8 9]",
		}
		fn(t, tableList, joinList, joinWhereList, valuesColumnList, valuesList, setList, selectList, whereList, groupByList, havingList, orderByList, limit, offset, checkList)
	})
}

func TestQueryValue_BuildTable(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
		}

		queryValue.SetTable(gol.QueryModeDefault, tableList[0][1])

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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		check := `"item"`
		fn(t, tableList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		check := `"item"`
		fn(t, tableList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		check := `"s1"."item"`
		fn(t, tableList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		check := `"s1"."item"`
		fn(t, tableList, check)
	})
}

func TestQueryValue_BuildTableUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
		}

		queryValue.SetTable(gol.QueryModeDefault, tableList[0][1])

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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		check := `"item"`
		fn(t, tableList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		check := `"item" as "t1"`
		fn(t, tableList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		check := `"s1"."item"`
		fn(t, tableList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		check := `"s1"."item" as "t1"`
		fn(t, tableList, check)
	})
}

func TestQueryValue_BuildJoin(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}
	tableItem3 := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, tableList [][]interface{}, joinList [][]interface{}, joinWhereList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
		}

		for _, val := range joinList {
			queryValue.AddJoin(val[0].(int), val[1])
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
			{"", &tableItem3, ""},
			{"", &tableTag, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
			{gol.QueryJoinModeRight, &tableTag},
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
				`INNER JOIN "item" ON "item"."id" = "item"."id" AND "item"."id" = ? LEFT JOIN "item" ON "item"."id" = "item"."id" AND "item"."id" = ? RIGHT JOIN "PUBLIC"."TAG" ON "PUBLIC"."TAG"."ID" = "item"."id" AND "PUBLIC"."TAG"."ID" = ?`,
			),
			"[1 2 3]",
		}
		fn(t, tableList, joinList, joinWhereList, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
			{"", &tableItem3, "t3"},
			{"", &tableTag, "t4"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
			{gol.QueryJoinModeRight, &tableTag},
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
		fn(t, tableList, joinList, joinWhereList, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
			{"s3", &tableItem3, ""},
			{"s4", &tableTag, ""},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
			{gol.QueryJoinModeRight, &tableTag},
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
				`INNER JOIN "s2"."item" ON "s2"."item"."id" = "s1"."item"."id" AND "s2"."item"."id" = ? LEFT JOIN "s3"."item" ON "s3"."item"."id" = "s1"."item"."id" AND "s3"."item"."id" = ? RIGHT JOIN "s4"."TAG" ON "s4"."TAG"."ID" = "s1"."item"."id" AND "s4"."TAG"."ID" = ?`,
			),
			"[1 2 3]",
		}
		fn(t, tableList, joinList, joinWhereList, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
			{"s3", &tableItem3, "t3"},
			{"s4", &tableTag, "t4"},
		}
		joinList := [][]interface{}{
			{gol.QueryJoinModeInner, &tableItem2},
			{gol.QueryJoinModeLeft, &tableItem3},
			{gol.QueryJoinModeRight, &tableTag},
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
				`INNER JOIN "s2"."item" as "t2" ON "t2"."id" = "t1"."id" AND "t2"."id" = ? LEFT JOIN "s3"."item" as "t3" ON "t3"."id" = "t1"."id" AND "t3"."id" = ? RIGHT JOIN "s4"."TAG" as "t4" ON "t4"."ID" = "t1"."id" AND "t4"."ID" = ?`,
			),
			"[1 2 3]",
		}
		fn(t, tableList, joinList, joinWhereList, checkList)
	})
}

func TestQueryValue_BuildJoinWhere(t *testing.T) {
	// supported by TestQueryValue_BuildJoin
}

func TestQueryValue_BuildValuesColumn(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, valuesColumnList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `("id", "name")`
		fn(t, tableList, valuesColumnList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `("id", "name")`
		fn(t, tableList, valuesColumnList, check)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `("id", "name")`
		fn(t, tableList, valuesColumnList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		valuesColumnList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `("id", "name")`
		fn(t, tableList, valuesColumnList, check)
	})
}

func TestQueryValue_BuildValues(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, valuesList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		valuesList := [][]interface{}{
			{gol.QueryModeDefault, 1, "a"},
			{gol.QueryModeDefault, 2, "b"},
		}
		checkList := []string{
			"(?, ?), (?, ?)",
			"[1 a 2 b]",
		}
		fn(t, tableList, valuesList, checkList)
	})
}

func TestQueryValue_BuildSet(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, setList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id, 1},
			{gol.QueryModeDefault, &tableItem.Name, "a"},
		}
		checkList := []string{
			`SET "id" = ?, "name" = ?`,
			"[1 a]",
		}
		fn(t, tableList, setList, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id, 1},
			{gol.QueryModeDefault, &tableItem.Name, "a"},
		}
		checkList := []string{
			`SET "id" = ?, "name" = ?`,
			"[1 a]",
		}
		fn(t, tableList, setList, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id, 1},
			{gol.QueryModeDefault, &tableItem.Name, "a"},
		}
		checkList := []string{
			`SET "id" = ?, "name" = ?`,
			"[1 a]",
		}
		fn(t, tableList, setList, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		setList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id, 1},
			{gol.QueryModeDefault, &tableItem.Name, "a"},
		}
		checkList := []string{
			`SET "id" = ?, "name" = ?`,
			"[1 a]",
		}
		fn(t, tableList, setList, checkList)
	})
}

func TestQueryValue_BuildSelectUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, selectList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
		}

		for _, val := range selectList {
			queryValue.AddSelect(val[0].(int), val[1:]...)
		}

		query, err := queryValue.BuildSelectUseAs()
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeDefault, &tableItem.Name},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `SELECT count("item"."id"), "item"."name", "item"."name"`
		fn(t, tableList, selectList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeDefault, &tableItem.Name},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `SELECT count("t1"."id"), "t1"."name", "t1"."name"`
		fn(t, tableList, selectList, check)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeDefault, &tableItem.Name},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `SELECT count("s1"."item"."id"), "s1"."item"."name", "s1"."item"."name"`
		fn(t, tableList, selectList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeDefault, &tableItem.Name},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `SELECT count("t1"."id"), "t1"."name", "t1"."name"`
		fn(t, tableList, selectList, check)
	})

	t.Run("all table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		selectList := [][]interface{}{
			{gol.QueryModeAll, &tableItem},
		}
		check := `SELECT "item".*`
		fn(t, tableList, selectList, check)
	})

	t.Run("all table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeAll, &tableItem},
		}
		check := `SELECT "t1".*`
		fn(t, tableList, selectList, check)
	})

	t.Run("all schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		selectList := [][]interface{}{
			{gol.QueryModeAll, &tableItem},
		}
		check := `SELECT "s1"."item".*`
		fn(t, tableList, selectList, check)
	})

	t.Run("all schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		selectList := [][]interface{}{
			{gol.QueryModeAll, &tableItem},
		}
		check := `SELECT "t1".*`
		fn(t, tableList, selectList, check)
	})
}

func TestQueryValue_BuildWhere(t *testing.T) {
	tableItem := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, tableList [][]interface{}, whereList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
			{"", &tableTag, ""},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`WHERE "item"."id" = ? AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( count("item"."id") = ? ? ) OR "PUBLIC"."TAG"."ID" = ? OR "PUBLIC"."TAG"."ID" != ? OR "PUBLIC"."TAG"."NAME" LIKE ? OR "PUBLIC"."TAG"."NAME" NOT LIKE ? OR "PUBLIC"."TAG"."ID" IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" NOT IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" > ? OR "PUBLIC"."TAG"."ID" >= ? OR "PUBLIC"."TAG"."ID" < ? OR "PUBLIC"."TAG"."ID" <= ? OR ( count("PUBLIC"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
			{"", &tableTag, "t2"},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`WHERE "item"."id" = ? AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( count("item"."id") = ? ? ) OR "PUBLIC"."TAG"."ID" = ? OR "PUBLIC"."TAG"."ID" != ? OR "PUBLIC"."TAG"."NAME" LIKE ? OR "PUBLIC"."TAG"."NAME" NOT LIKE ? OR "PUBLIC"."TAG"."ID" IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" NOT IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" > ? OR "PUBLIC"."TAG"."ID" >= ? OR "PUBLIC"."TAG"."ID" < ? OR "PUBLIC"."TAG"."ID" <= ? OR ( count("PUBLIC"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
			{"s2", &tableTag, ""},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`WHERE "s1"."item"."id" = ? AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( count("s1"."item"."id") = ? ? ) OR "s2"."TAG"."ID" = ? OR "s2"."TAG"."ID" != ? OR "s2"."TAG"."NAME" LIKE ? OR "s2"."TAG"."NAME" NOT LIKE ? OR "s2"."TAG"."ID" IN (?, ?, ?) OR "s2"."TAG"."ID" NOT IN (?, ?, ?) OR "s2"."TAG"."ID" > ? OR "s2"."TAG"."ID" >= ? OR "s2"."TAG"."ID" < ? OR "s2"."TAG"."ID" <= ? OR ( count("s2"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
			{"s2", &tableTag, "t2"},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`WHERE "s1"."item"."id" = ? AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( count("s1"."item"."id") = ? ? ) OR "s2"."TAG"."ID" = ? OR "s2"."TAG"."ID" != ? OR "s2"."TAG"."NAME" LIKE ? OR "s2"."TAG"."NAME" NOT LIKE ? OR "s2"."TAG"."ID" IN (?, ?, ?) OR "s2"."TAG"."ID" NOT IN (?, ?, ?) OR "s2"."TAG"."ID" > ? OR "s2"."TAG"."ID" >= ? OR "s2"."TAG"."ID" < ? OR "s2"."TAG"."ID" <= ? OR ( count("s2"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, whereList, checkList)
	})
}

func TestQueryValue_BuildWhereUseAs(t *testing.T) {
	tableItem := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, tableList [][]interface{}, whereList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
			{"", &tableTag, ""},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`WHERE "item"."id" = ? AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( count("item"."id") = ? ? ) OR "PUBLIC"."TAG"."ID" = ? OR "PUBLIC"."TAG"."ID" != ? OR "PUBLIC"."TAG"."NAME" LIKE ? OR "PUBLIC"."TAG"."NAME" NOT LIKE ? OR "PUBLIC"."TAG"."ID" IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" NOT IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" > ? OR "PUBLIC"."TAG"."ID" >= ? OR "PUBLIC"."TAG"."ID" < ? OR "PUBLIC"."TAG"."ID" <= ? OR ( count("PUBLIC"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
			{"", &tableTag, "t2"},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`WHERE "t1"."id" = ? AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( count("t1"."id") = ? ? ) OR "t2"."ID" = ? OR "t2"."ID" != ? OR "t2"."NAME" LIKE ? OR "t2"."NAME" NOT LIKE ? OR "t2"."ID" IN (?, ?, ?) OR "t2"."ID" NOT IN (?, ?, ?) OR "t2"."ID" > ? OR "t2"."ID" >= ? OR "t2"."ID" < ? OR "t2"."ID" <= ? OR ( count("t2"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
			{"s2", &tableTag, ""},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`WHERE "s1"."item"."id" = ? AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( count("s1"."item"."id") = ? ? ) OR "s2"."TAG"."ID" = ? OR "s2"."TAG"."ID" != ? OR "s2"."TAG"."NAME" LIKE ? OR "s2"."TAG"."NAME" NOT LIKE ? OR "s2"."TAG"."ID" IN (?, ?, ?) OR "s2"."TAG"."ID" NOT IN (?, ?, ?) OR "s2"."TAG"."ID" > ? OR "s2"."TAG"."ID" >= ? OR "s2"."TAG"."ID" < ? OR "s2"."TAG"."ID" <= ? OR ( count("s2"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
			{"s2", &tableTag, "t2"},
		}
		whereList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`WHERE "t1"."id" = ? AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( count("t1"."id") = ? ? ) OR "t2"."ID" = ? OR "t2"."ID" != ? OR "t2"."NAME" LIKE ? OR "t2"."NAME" NOT LIKE ? OR "t2"."ID" IN (?, ?, ?) OR "t2"."ID" NOT IN (?, ?, ?) OR "t2"."ID" > ? OR "t2"."ID" >= ? OR "t2"."ID" < ? OR "t2"."ID" <= ? OR ( count("t2"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, whereList, checkList)
	})
}

func TestQueryValue_BuildGroupByUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, groupByList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `GROUP BY "item"."id", "item"."name"`
		fn(t, tableList, groupByList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `GROUP BY "t1"."id", "t1"."name"`
		fn(t, tableList, groupByList, check)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `GROUP BY "s1"."item"."id", "s1"."item"."name"`
		fn(t, tableList, groupByList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		groupByList := [][]interface{}{
			{gol.QueryModeDefault, &tableItem.Id},
			{gol.QueryModeDefault, &tableItem.Name},
		}
		check := `GROUP BY "t1"."id", "t1"."name"`
		fn(t, tableList, groupByList, check)
	})
}

func TestQueryValue_BuildHavingUseAs(t *testing.T) {
	tableItem := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, tableList [][]interface{}, havingList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
			{"", &tableTag, ""},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`HAVING "item"."id" = ? AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( count("item"."id") = ? ? ) OR "PUBLIC"."TAG"."ID" = ? OR "PUBLIC"."TAG"."ID" != ? OR "PUBLIC"."TAG"."NAME" LIKE ? OR "PUBLIC"."TAG"."NAME" NOT LIKE ? OR "PUBLIC"."TAG"."ID" IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" NOT IN (?, ?, ?) OR "PUBLIC"."TAG"."ID" > ? OR "PUBLIC"."TAG"."ID" >= ? OR "PUBLIC"."TAG"."ID" < ? OR "PUBLIC"."TAG"."ID" <= ? OR ( count("PUBLIC"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
			{"", &tableTag, "t2"},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`HAVING "t1"."id" = ? AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( count("t1"."id") = ? ? ) OR "t2"."ID" = ? OR "t2"."ID" != ? OR "t2"."NAME" LIKE ? OR "t2"."NAME" NOT LIKE ? OR "t2"."ID" IN (?, ?, ?) OR "t2"."ID" NOT IN (?, ?, ?) OR "t2"."ID" > ? OR "t2"."ID" >= ? OR "t2"."ID" < ? OR "t2"."ID" <= ? OR ( count("t2"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
			{"s2", &tableTag, ""},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`HAVING "s1"."item"."id" = ? AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( count("s1"."item"."id") = ? ? ) OR "s2"."TAG"."ID" = ? OR "s2"."TAG"."ID" != ? OR "s2"."TAG"."NAME" LIKE ? OR "s2"."TAG"."NAME" NOT LIKE ? OR "s2"."TAG"."ID" IN (?, ?, ?) OR "s2"."TAG"."ID" NOT IN (?, ?, ?) OR "s2"."TAG"."ID" > ? OR "s2"."TAG"."ID" >= ? OR "s2"."TAG"."ID" < ? OR "s2"."TAG"."ID" <= ? OR ( count("s2"."TAG"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
			{"s2", &tableTag, "t2"},
		}
		havingList := [][]interface{}{
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 0},
			{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1},
			{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 2},
			{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"},
			{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "b"},
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
			{gol.QueryModeLike, gol.QueryPrefixOr, &tableTag.Name, "c"},
			{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableTag.Name, "d"},
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
			`HAVING "t1"."id" = ? AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( count("t1"."id") = ? ? ) OR "t2"."ID" = ? OR "t2"."ID" != ? OR "t2"."NAME" LIKE ? OR "t2"."NAME" NOT LIKE ? OR "t2"."ID" IN (?, ?, ?) OR "t2"."ID" NOT IN (?, ?, ?) OR "t2"."ID" > ? OR "t2"."ID" >= ? OR "t2"."ID" < ? OR "t2"."ID" <= ? OR ( count("t2"."ID") = ? ? )`,
			"[0 1 2 a b 3 4 5 6 7 8 9 10 11 12 13 14 15 16 c d 17 18 19 20 21 22 23 24 25 26 27 28]",
		}
		fn(t, tableList, havingList, checkList)
	})
}

func TestQueryValue_BuildOrderByUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, orderByList [][]interface{}, check string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range tableList {
			queryValue.AddMeta(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeAsc, &tableItem.Name},
			{gol.QueryModeDesc, &tableItem.Name},
		}
		check := `ORDER BY count("item"."id"), "item"."name" ASC, "item"."name" DESC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeAsc, &tableItem.Name},
			{gol.QueryModeDesc, &tableItem.Name},
		}
		check := `ORDER BY count("t1"."id"), "t1"."name" ASC, "t1"."name" DESC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeAsc, &tableItem.Name},
			{gol.QueryModeDesc, &tableItem.Name},
		}
		check := `ORDER BY count("s1"."item"."id"), "s1"."item"."name" ASC, "s1"."item"."name" DESC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		orderByList := [][]interface{}{
			{gol.QueryModeDefault, "count(", &tableItem.Id, ")"},
			{gol.QueryModeAsc, &tableItem.Name},
			{gol.QueryModeDesc, &tableItem.Name},
		}
		check := `ORDER BY count("t1"."id"), "t1"."name" ASC, "t1"."name" DESC`
		fn(t, tableList, orderByList, check)
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
			err := meta.Add(val[0].(string), val[1], val[2].(string))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryTable{}
		data.Set(gol.QueryModeDefault, tableList[0][1])

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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		check := `"item"`
		fn(t, tableList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		check := `"item"`
		fn(t, tableList, check)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		check := `"s1"."item"`
		fn(t, tableList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		check := `"s1"."item"`
		fn(t, tableList, check)
	})
}

func TestQueryTable_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryTable{}
		data.Set(gol.QueryModeDefault, tableList[0][1])

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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		check := `"item"`
		fn(t, tableList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		check := `"item" as "t1"`
		fn(t, tableList, check)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		check := `"s1"."item"`
		fn(t, tableList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		check := `"s1"."item" as "t1"`
		fn(t, tableList, check)
	})
}

func TestQueryJoin_BuildUseAs(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, joinList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QueryJoin{}
		data.Set(joinList[0].(int), joinList[1])

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

	t.Run("inner table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinList := []interface{}{gol.QueryJoinModeInner, &tableItem2}
		check := `INNER JOIN "item" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("inner table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
		}
		joinList := []interface{}{gol.QueryJoinModeInner, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `INNER JOIN "item" as "t2" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("inner schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
		}
		joinList := []interface{}{gol.QueryJoinModeInner, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `INNER JOIN "s2"."item" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("inner schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
		}
		joinList := []interface{}{gol.QueryJoinModeInner, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `INNER JOIN "s2"."item" as "t2" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("left table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinList := []interface{}{gol.QueryJoinModeLeft, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `LEFT JOIN "item" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("left table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
		}
		joinList := []interface{}{gol.QueryJoinModeLeft, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `LEFT JOIN "item" as "t2" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("left schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
		}
		joinList := []interface{}{gol.QueryJoinModeLeft, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `LEFT JOIN "s2"."item" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("left schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
		}
		joinList := []interface{}{gol.QueryJoinModeLeft, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `LEFT JOIN "s2"."item" as "t2" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("right table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinList := []interface{}{gol.QueryJoinModeRight, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `RIGHT JOIN "item" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("right table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
		}
		joinList := []interface{}{gol.QueryJoinModeRight, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `RIGHT JOIN "item" as "t2" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("right schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
		}
		joinList := []interface{}{gol.QueryJoinModeRight, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `RIGHT JOIN "s2"."item" ON`
		fn(t, tableList, joinList, check)
	})

	t.Run("right schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
		}
		joinList := []interface{}{gol.QueryJoinModeRight, &tableItem2, &tableItem2.Id, &tableItem1.Id}
		check := `RIGHT JOIN "s2"."item" as "t2" ON`
		fn(t, tableList, joinList, check)
	})
}

func TestQueryJoinWhere_BuildUseAs(t *testing.T) {
	tableItem1 := test.Item{}
	tableItem2 := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, joinWhereList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
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

	t.Run("and default table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and default table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, "t1"},
			{"", &tableItem2, "t2"},
		}
		joinWhereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"t2"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and default schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, ""},
			{"s2", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"s2"."item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and default schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem1, "t1"},
			{"s2", &tableItem2, "t2"},
		}
		joinWhereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"t2"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and default", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and is", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and is not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and like", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Name, "a"}
		checkList := []string{
			"AND",
			`"item"."name" LIKE ?`,
			"[a]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and like not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Name, "a"}
		checkList := []string{
			"AND",
			`"item"."name" NOT LIKE ?`,
			"[a]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and in", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and in not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and gt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and gte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and lt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and lte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or default", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or is", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or is not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or like", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixOr, &tableItem2, &tableItem2.Name, "a"}
		checkList := []string{
			"OR",
			`"item"."name" LIKE ?`,
			"[a]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or like not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableItem2, &tableItem2.Name, "a"}
		checkList := []string{
			"OR",
			`"item"."name" NOT LIKE ?`,
			"[a]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or in", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or in not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or gt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or gte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or lt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or lte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("and nest", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"AND",
			`(`,
			"[]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("or nest", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixOr, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"OR",
			`(`,
			"[]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})

	t.Run("nest close", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem1, ""},
			{"", &tableItem2, ""},
		}
		joinWhereList := []interface{}{gol.QueryModeNestClose, gol.QueryPrefixAnd, &tableItem2, &tableItem2.Id, 1}
		checkList := []string{
			"AND",
			`)`,
			"[]",
		}
		fn(t, tableList, joinWhereList, checkList)
	})
}

func TestQueryValuesColumn_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, valuesColumnList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		valuesColumnList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"id"`
		fn(t, tableList, valuesColumnList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		valuesColumnList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"id"`
		fn(t, tableList, valuesColumnList, check)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		valuesColumnList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"id"`
		fn(t, tableList, valuesColumnList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		valuesColumnList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"id"`
		fn(t, tableList, valuesColumnList, check)
	})
}

func TestQueryValues_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, valuesList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		valuesList := []interface{}{gol.QueryModeDefault, 1, 2, 3}
		checkList := []string{
			"(?, ?, ?)",
			"[1 2 3]",
		}
		fn(t, tableList, valuesList, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		valuesList := []interface{}{gol.QueryModeDefault, 1, 2, 3}
		checkList := []string{
			"(?, ?, ?)",
			"[1 2 3]",
		}
		fn(t, tableList, valuesList, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		valuesList := []interface{}{gol.QueryModeDefault, 1, 2, 3}
		checkList := []string{
			"(?, ?, ?)",
			"[1 2 3]",
		}
		fn(t, tableList, valuesList, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		valuesList := []interface{}{gol.QueryModeDefault, 1, 2, 3}
		checkList := []string{
			"(?, ?, ?)",
			"[1 2 3]",
		}
		fn(t, tableList, valuesList, checkList)
	})
}

func TestQuerySet_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, setList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		setList := []interface{}{gol.QueryModeDefault, &tableItem.Id, 1}
		checkList := []string{
			`"id" = ?`,
			`1`,
		}
		fn(t, tableList, setList, checkList)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		setList := []interface{}{gol.QueryModeDefault, &tableItem.Id, 1}
		checkList := []string{
			`"id" = ?`,
			`1`,
		}
		fn(t, tableList, setList, checkList)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		setList := []interface{}{gol.QueryModeDefault, &tableItem.Id, 1}
		checkList := []string{
			`"id" = ?`,
			`1`,
		}
		fn(t, tableList, setList, checkList)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		setList := []interface{}{gol.QueryModeDefault, &tableItem.Id, 1}
		checkList := []string{
			`"id" = ?`,
			`1`,
		}
		fn(t, tableList, setList, checkList)
	})
}

func TestQuerySelect_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, selectList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
			if err != nil {
				t.Error(err)
			}
		}

		data := &gol.QuerySelect{}
		data.Set(selectList[0].(int), selectList[1:]...)

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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		selectList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("item"."id")`
		fn(t, tableList, selectList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		selectList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id")`
		fn(t, tableList, selectList, check)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		selectList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("s1"."item"."id")`
		fn(t, tableList, selectList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		selectList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id")`
		fn(t, tableList, selectList, check)
	})

	t.Run("all table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		selectList := []interface{}{gol.QueryModeAll, &tableItem}
		check := `"item".*`
		fn(t, tableList, selectList, check)
	})

	t.Run("all table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		selectList := []interface{}{gol.QueryModeAll, &tableItem}
		check := `"t1".*`
		fn(t, tableList, selectList, check)
	})

	t.Run("all schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		selectList := []interface{}{gol.QueryModeAll, &tableItem}
		check := `"s1"."item".*`
		fn(t, tableList, selectList, check)
	})

	t.Run("all schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		selectList := []interface{}{gol.QueryModeAll, &tableItem}
		check := `"t1".*`
		fn(t, tableList, selectList, check)
	})
}

func TestQueryWhere_Build(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, whereList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
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

	t.Run("and default table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and default table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and default schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"s1"."item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and default schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"s1"."item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and default", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and is", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and is not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and like", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"}
		checkList := []string{
			"AND",
			`"item"."name" LIKE ?`,
			"[a]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and like not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "a"}
		checkList := []string{
			"AND",
			`"item"."name" NOT LIKE ?`,
			"[a]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and in", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and in not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and gt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and gte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and lt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and lte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or default", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixOr, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or is", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or is not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or like", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixOr, &tableItem.Name, "a"}
		checkList := []string{
			"OR",
			`"item"."name" LIKE ?`,
			"[a]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or like not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableItem.Name, "a"}
		checkList := []string{
			"OR",
			`"item"."name" NOT LIKE ?`,
			"[a]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or in", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or in not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or gt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or gte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or lt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or lte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and nest", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`(`,
			"[]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or nest", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`(`,
			"[]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("nest close", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeNestClose, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`)`,
			"[]",
		}
		fn(t, tableList, whereList, checkList)
	})
}

func TestQueryWhere_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, whereList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
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

	t.Run("and default table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and default table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and default schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"s1"."item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and default schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and default", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and is", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and is not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and like", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"}
		checkList := []string{
			"AND",
			`"item"."name" LIKE ?`,
			"[a]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and like not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "a"}
		checkList := []string{
			"AND",
			`"item"."name" NOT LIKE ?`,
			"[a]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and in", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and in not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and gt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and gte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and lt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and lte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or default", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixOr, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or is", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIs, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or is not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or like", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLike, gol.QueryPrefixOr, &tableItem.Name, "a"}
		checkList := []string{
			"OR",
			`"item"."name" LIKE ?`,
			"[a]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or like not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableItem.Name, "a"}
		checkList := []string{
			"OR",
			`"item"."name" NOT LIKE ?`,
			"[a]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or in", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeIn, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or in not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or gt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeGt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or gte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeGte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or lt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or lte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeLte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("and nest", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`(`,
			"[]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("or nest", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeNest, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`(`,
			"[]",
		}
		fn(t, tableList, whereList, checkList)
	})

	t.Run("nest close", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		whereList := []interface{}{gol.QueryModeNestClose, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`)`,
			"[]",
		}
		fn(t, tableList, whereList, checkList)
	})
}

func TestQueryGroupBy_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, groupByList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
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

	t.Run("table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		groupByList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"item"."id"`
		fn(t, tableList, groupByList, check)
	})

	t.Run("table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		groupByList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"t1"."id"`
		fn(t, tableList, groupByList, check)
	})

	t.Run("schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		groupByList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"s1"."item"."id"`
		fn(t, tableList, groupByList, check)
	})

	t.Run("schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		groupByList := []interface{}{gol.QueryModeDefault, &tableItem.Id}
		check := `"t1"."id"`
		fn(t, tableList, groupByList, check)
	})
}

func TestQueryHaving_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, havingList []interface{}, checkList []string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
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

	t.Run("and default table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and default table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and default schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"s1"."item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and default schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"t1"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and default", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixAnd, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and is", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeIs, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and is not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and like", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeLike, gol.QueryPrefixAnd, &tableItem.Name, "a"}
		checkList := []string{
			"AND",
			`"item"."name" LIKE ?`,
			"[a]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and like not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixAnd, &tableItem.Name, "a"}
		checkList := []string{
			"AND",
			`"item"."name" NOT LIKE ?`,
			"[a]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and in", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeIn, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and in not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixAnd, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"AND",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and gt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeGt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and gte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeGte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and lt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeLt, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and lte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeLte, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or default", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeDefault, gol.QueryPrefixOr, &tableItem.Id, " = ?", []interface{}{1}}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or is", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeIs, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" = ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or is not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeIsNot, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" != ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or like", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeLike, gol.QueryPrefixOr, &tableItem.Name, "a"}
		checkList := []string{
			"OR",
			`"item"."name" LIKE ?`,
			"[a]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or like not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeLikeNot, gol.QueryPrefixOr, &tableItem.Name, "a"}
		checkList := []string{
			"OR",
			`"item"."name" NOT LIKE ?`,
			"[a]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or in", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeIn, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or in not", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeInNot, gol.QueryPrefixOr, &tableItem.Id, []interface{}{1, 2, 3}}
		checkList := []string{
			"OR",
			`"item"."id" NOT IN (?, ?, ?)`,
			"[1 2 3]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or gt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeGt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" > ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or gte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeGte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" >= ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or lt", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeLt, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" < ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or lte", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeLte, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`"item"."id" <= ?`,
			"[1]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("and nest", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeNest, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`(`,
			"[]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("or nest", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeNest, gol.QueryPrefixOr, &tableItem.Id, 1}
		checkList := []string{
			"OR",
			`(`,
			"[]",
		}
		fn(t, tableList, havingList, checkList)
	})

	t.Run("nest close", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		havingList := []interface{}{gol.QueryModeNestClose, gol.QueryPrefixAnd, &tableItem.Id, 1}
		checkList := []string{
			"AND",
			`)`,
			"[]",
		}
		fn(t, tableList, havingList, checkList)
	})
}

func TestQueryOrderBy_BuildUseAs(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, tableList [][]interface{}, orderByList []interface{}, check string) {
		meta := gol.NewMeta(nil)

		for _, val := range tableList {
			err := meta.Add(val[0].(string), val[1], val[2].(string))
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

	t.Run("default table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		orderByList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("item"."id")`
		fn(t, tableList, orderByList, check)
	})

	t.Run("default table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		orderByList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id")`
		fn(t, tableList, orderByList, check)
	})

	t.Run("default schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		orderByList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("s1"."item"."id")`
		fn(t, tableList, orderByList, check)
	})

	t.Run("default schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		orderByList := []interface{}{gol.QueryModeDefault, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id")`
		fn(t, tableList, orderByList, check)
	})

	t.Run("asc table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		orderByList := []interface{}{gol.QueryModeAsc, "count(", &tableItem.Id, ")"}
		check := `count("item"."id") ASC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("asc table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		orderByList := []interface{}{gol.QueryModeAsc, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id") ASC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("asc schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		orderByList := []interface{}{gol.QueryModeAsc, "count(", &tableItem.Id, ")"}
		check := `count("s1"."item"."id") ASC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("asc schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		orderByList := []interface{}{gol.QueryModeAsc, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id") ASC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("desc table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, ""},
		}
		orderByList := []interface{}{gol.QueryModeDesc, "count(", &tableItem.Id, ")"}
		check := `count("item"."id") DESC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("desc table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"", &tableItem, "t1"},
		}
		orderByList := []interface{}{gol.QueryModeDesc, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id") DESC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("desc schema table", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, ""},
		}
		orderByList := []interface{}{gol.QueryModeDesc, "count(", &tableItem.Id, ")"}
		check := `count("s1"."item"."id") DESC`
		fn(t, tableList, orderByList, check)
	})

	t.Run("desc schema table as", func(t *testing.T) {
		tableList := [][]interface{}{
			{"s1", &tableItem, "t1"},
		}
		orderByList := []interface{}{gol.QueryModeDesc, "count(", &tableItem.Id, ")"}
		check := `count("t1"."id") DESC`
		fn(t, tableList, orderByList, check)
	})
}
