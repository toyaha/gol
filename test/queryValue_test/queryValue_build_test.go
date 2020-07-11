package queryValue_test

import (
	"fmt"
	"testing"

	"github.com/toyaha/gol"
	"github.com/toyaha/gol/test"
)

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
		check := `"item" as "t0"`
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
		checkList := []string{`"item" as "t0", "item" as "t1"`, "[]"}
		fn(t, fromList, checkList)
	})

	t.Run("query", func(t *testing.T) {
		fromList := [][]interface{}{
			{&tableItem1, true, "(select 1, ?)", []interface{}{2}},
			{&tableItem2, true, "(select 3, ?)", []interface{}{4}},
		}
		checkList := []string{`(select 1, ?) as "t0", (select 3, ?) as "t1"`, "[2 4]"}
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
				`INNER JOIN "item" as "t1" ON "t1"."id" = "t0"."id" AND "t1"."id" = ? LEFT JOIN "item" as "t2" ON "t2"."id" = "t0"."id" AND "t2"."id" = ? RIGHT JOIN "PUBLIC"."TAG" as "t3" ON "t3"."ID" = "t0"."id" AND "t3"."ID" = ?`,
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

func TestQueryValue_BuildValuesSelectUnion(t *testing.T) {
	tableItem := test.Item{}

	fn := func(t *testing.T, metaList [][]interface{}, valuesList [][]interface{}, checkList []string) {
		queryValue := gol.NewQueryValue(nil)

		for _, val := range metaList {
			queryValue.AddMeta(val[0], val[1].(bool))
		}

		for _, val := range valuesList {
			queryValue.AddValues(val[0].(int), val[1:]...)
		}

		query, valueList, err := queryValue.BuildValuesSelectUnion()
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
			"SELECT ?, ? UNION SELECT ?, ?",
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
			`SELECT count("t0"."id"), "t0"."str", "t0"."str"`,
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
			`SELECT "t0".*`,
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
			`WHERE "t0"."id" = ? AND "t0"."id" = ? AND "t0"."id" != ? AND "t0"."str" LIKE ? AND "t0"."str" NOT LIKE ? AND "t0"."id" IN (?, ?, ?) AND "t0"."id" NOT IN (?, ?, ?) AND "t0"."id" > ? AND "t0"."id" >= ? AND "t0"."id" < ? AND "t0"."id" <= ? AND ( count("t0"."id") = ? ? ) OR "t1"."ID" = ? OR "t1"."ID" != ? OR "t1"."STR" LIKE ? OR "t1"."STR" NOT LIKE ? OR "t1"."ID" IN (?, ?, ?) OR "t1"."ID" NOT IN (?, ?, ?) OR "t1"."ID" > ? OR "t1"."ID" >= ? OR "t1"."ID" < ? OR "t1"."ID" <= ? OR ( count("t1"."ID") = ? ? )`,
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
		check := `GROUP BY "t0"."id", "t0"."str"`
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
			`HAVING "t0"."id" = ? AND "t0"."id" = ? AND "t0"."id" != ? AND "t0"."str" LIKE ? AND "t0"."str" NOT LIKE ? AND "t0"."id" IN (?, ?, ?) AND "t0"."id" NOT IN (?, ?, ?) AND "t0"."id" > ? AND "t0"."id" >= ? AND "t0"."id" < ? AND "t0"."id" <= ? AND ( count("t0"."id") = ? ? ) OR "t1"."ID" = ? OR "t1"."ID" != ? OR "t1"."STR" LIKE ? OR "t1"."STR" NOT LIKE ? OR "t1"."ID" IN (?, ?, ?) OR "t1"."ID" NOT IN (?, ?, ?) OR "t1"."ID" > ? OR "t1"."ID" >= ? OR "t1"."ID" < ? OR "t1"."ID" <= ? OR ( count("t1"."ID") = ? ? )`,
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
		check := `ORDER BY count("t0"."id"), "t0"."str" ASC, "t0"."str" DESC`
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
