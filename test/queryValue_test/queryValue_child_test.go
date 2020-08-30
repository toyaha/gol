package queryValue_test

import (
	"fmt"
	"testing"

	"github.com/toyaha/gol"
	"github.com/toyaha/gol/test"
)

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
		check := `"item" as "t0"`
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
		checkList := []string{`"item" as "t0"`, "[]"}
		fn(t, fromList, checkList)
	})

	t.Run("query", func(t *testing.T) {
		fromList := [][]interface{}{
			{&tableItem, true, "(select 1)"},
		}
		checkList := []string{`(select 1) as "t0"`, "[]"}
		fn(t, fromList, checkList)
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
			`"t0"."id" = ?`,
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
			`"t0"."id" = ?`,
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
			`"t0"."id" != ?`,
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
			`"t0"."str" LIKE ?`,
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
			`"t0"."str" NOT LIKE ?`,
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
			`"t0"."id" IN (?, ?, ?)`,
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
			`"t0"."id" NOT IN (?, ?, ?)`,
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
			`"t0"."id" > ?`,
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
			`"t0"."id" >= ?`,
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
			`"t0"."id" < ?`,
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
			`"t0"."id" <= ?`,
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
			`"t0"."id" = ?`,
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
			`"t0"."id" = ?`,
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
			`"t0"."id" != ?`,
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
			`"t0"."str" LIKE ?`,
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
			`"t0"."str" NOT LIKE ?`,
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
			`"t0"."id" IN (?, ?, ?)`,
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
			`"t0"."id" NOT IN (?, ?, ?)`,
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
			`"t0"."id" > ?`,
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
			`"t0"."id" >= ?`,
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
			`"t0"."id" < ?`,
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
			`"t0"."id" <= ?`,
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
			"?, ?, ?",
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
		data.Set(setList[0].(int), setList[1:]...)

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
		setList := []interface{}{gol.QueryModeDefault, &tableItem.Num, 1}
		checkList := []string{
			`"num" = ?`,
			`[1]`,
		}
		fn(t, metaList, setList, checkList)
	})

	t.Run("column", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		setList := []interface{}{gol.QueryModeDefault, &tableItem.Num, &tableItem.Id}
		checkList := []string{
			`"num" = "id"`,
			`[]`,
		}
		fn(t, metaList, setList, checkList)
	})

	t.Run("calc", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		setList := []interface{}{gol.QueryModeDefault, &tableItem.Num, &tableItem.Id, " + ?", []interface{}{1}}
		checkList := []string{
			`"num" = "id" + ?`,
			`[1]`,
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
		checkList := []string{`count("t0"."id")`, "[]"}
		fn(t, metaList, selectList, checkList)
	})

	t.Run("all", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		selectList := []interface{}{gol.QueryModeAll, &tableItem}
		checkList := []string{`"t0".*`, "[]"}
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
			`"t0"."id" = ?`,
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
			`"t0"."id" = ?`,
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
			`"t0"."id" != ?`,
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
			`"t0"."str" LIKE ?`,
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
			`"t0"."str" NOT LIKE ?`,
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
			`"t0"."id" IN (?, ?, ?)`,
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
			`"t0"."id" NOT IN (?, ?, ?)`,
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
			`"t0"."id" > ?`,
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
			`"t0"."id" >= ?`,
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
			`"t0"."id" < ?`,
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
			`"t0"."id" <= ?`,
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
			`"t0"."id" = ?`,
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
			`"t0"."id" = ?`,
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
			`"t0"."id" != ?`,
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
			`"t0"."str" LIKE ?`,
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
			`"t0"."str" NOT LIKE ?`,
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
			`"t0"."id" IN (?, ?, ?)`,
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
			`"t0"."id" NOT IN (?, ?, ?)`,
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
			`"t0"."id" > ?`,
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
			`"t0"."id" >= ?`,
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
			`"t0"."id" < ?`,
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
			`"t0"."id" <= ?`,
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
		check := `"t0"."id"`
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
			`"t0"."id" = ?`,
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
			`"t0"."id" = ?`,
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
			`"t0"."id" != ?`,
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
			`"t0"."str" LIKE ?`,
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
			`"t0"."str" NOT LIKE ?`,
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
			`"t0"."id" IN (?, ?, ?)`,
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
			`"t0"."id" NOT IN (?, ?, ?)`,
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
			`"t0"."id" > ?`,
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
			`"t0"."id" >= ?`,
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
			`"t0"."id" < ?`,
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
			`"t0"."id" <= ?`,
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
			`"t0"."id" = ?`,
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
			`"t0"."id" = ?`,
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
			`"t0"."id" != ?`,
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
			`"t0"."str" LIKE ?`,
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
			`"t0"."str" NOT LIKE ?`,
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
			`"t0"."id" IN (?, ?, ?)`,
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
			`"t0"."id" NOT IN (?, ?, ?)`,
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
			`"t0"."id" > ?`,
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
			`"t0"."id" >= ?`,
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
			`"t0"."id" < ?`,
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
			`"t0"."id" <= ?`,
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
		check := `count("t0"."id")`
		fn(t, metaList, orderByList, check)
	})

	t.Run("asc", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := []interface{}{gol.QueryModeAsc, "count(", &tableItem.Id, ")"}
		check := `count("t0"."id") ASC`
		fn(t, metaList, orderByList, check)
	})

	t.Run("desc", func(t *testing.T) {
		metaList := [][]interface{}{
			{&tableItem, true},
		}
		orderByList := []interface{}{gol.QueryModeDesc, "count(", &tableItem.Id, ")"}
		check := `count("t0"."id") DESC`
		fn(t, metaList, orderByList, check)
	})
}
