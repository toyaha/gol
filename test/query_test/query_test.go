package query_test

import (
	"fmt"
	"testing"

	"github.com/toyaha/gol"
	"github.com/toyaha/gol/test"
)

func Query_setAllTypeA() *gol.Query {
	query := gol.NewQuery(nil)

	query.SetTable(&test.TableItem1)

	query.SetFrom(&test.TableItem2)
	query.SetFrom(&test.TableTag1)

	query.SetJoin(&test.TableItem3)
	query.SetJoinWhere(&test.TableItem3, &test.TableItem3.Id, " = ", &test.TableItem1.Id)

	query.SetJoinLeft(&test.TableTag2)
	query.SetJoinWhere(&test.TableTag2, &test.TableTag2.Id, " = ", &test.TableItem1.Id)

	query.SetJoinRight(&test.TableTag3)
	query.SetJoinWhere(&test.TableTag3, &test.TableTag3.Id, " = ", &test.TableItem1.Id)

	query.SetJoinWhere(&test.TableItem3, "1 = 1")
	query.SetJoinWhereIs(&test.TableItem3, &test.TableItem3.Id, 101)
	query.SetJoinWhereIsNot(&test.TableItem3, &test.TableItem3.Id, 102)
	query.SetJoinWhereLike(&test.TableItem3, &test.TableItem3.Str, "ja")
	query.SetJoinWhereLikeNot(&test.TableItem3, &test.TableItem3.Str, "jb")
	query.SetJoinWhereIn(&test.TableItem3, &test.TableItem3.Id, []interface{}{103, 104, 105})
	query.SetJoinWhereInNot(&test.TableItem3, &test.TableItem3.Id, []interface{}{106, 107, 108})
	query.SetJoinWhereGt(&test.TableItem3, &test.TableItem3.Id, 109)
	query.SetJoinWhereGte(&test.TableItem3, &test.TableItem3.Id, 110)
	query.SetJoinWhereLt(&test.TableItem3, &test.TableItem3.Id, 111)
	query.SetJoinWhereLte(&test.TableItem3, &test.TableItem3.Id, 112)

	query.SetJoinWhereNest(&test.TableItem3)
	query.SetJoinWhereOr(&test.TableItem3, "1 = 1")
	query.SetJoinWhereOrIs(&test.TableItem3, &test.TableItem3.Id, 113)
	query.SetJoinWhereOrIsNot(&test.TableItem3, &test.TableItem3.Id, 114)
	query.SetJoinWhereOrLike(&test.TableItem3, &test.TableItem3.Str, "jc")
	query.SetJoinWhereOrLikeNot(&test.TableItem3, &test.TableItem3.Str, "jd")
	query.SetJoinWhereOrIn(&test.TableItem3, &test.TableItem3.Id, []interface{}{115, 116, 117})
	query.SetJoinWhereOrInNot(&test.TableItem3, &test.TableItem3.Id, []interface{}{118, 119, 120})
	query.SetJoinWhereOrGt(&test.TableItem3, &test.TableItem3.Id, 121)
	query.SetJoinWhereOrGte(&test.TableItem3, &test.TableItem3.Id, 122)
	query.SetJoinWhereOrLt(&test.TableItem3, &test.TableItem3.Id, 123)
	query.SetJoinWhereOrLte(&test.TableItem3, &test.TableItem3.Id, 124)
	query.SetJoinWhereNestClose(&test.TableItem3)

	query.SetJoinWhereOrNest(&test.TableItem3)
	query.SetJoinWhereOr(&test.TableItem3, "1 = 1")
	query.SetJoinWhereNestClose(&test.TableItem3)

	query.SetValuesColumn(
		&test.TableItem1.CreateAt,
		&test.TableItem1.UpdateAt,
		&test.TableItem1.DeleteAt,
		&test.TableItem1.Str,
	)
	query.SetValues(
		test.TableItem1.CreateAt,
		test.TableItem1.UpdateAt,
		test.TableItem1.DeleteAt,
		test.TableItem1.Str,
	)

	query.SetConflict(&test.TableItem1.Id)

	query.SetSet(&test.TableItem1.CreateAt, test.TableItem1.CreateAt)
	query.SetSet(&test.TableItem1.UpdateAt, test.TableItem1.UpdateAt)
	query.SetSet(&test.TableItem1.DeleteAt, test.TableItem1.DeleteAt)
	query.SetSet(&test.TableItem1.Str, test.TableItem1.Str)

	query.SetSelect(&test.TableItem1.Id)
	query.SetSelect(&test.TableItem1.Str)
	query.SetSelectAll(&test.TableItem2)

	query.SetWhere("1 = 1")
	query.SetWhereIs(&test.TableItem1.Id, 201)
	query.SetWhereIsNot(&test.TableItem1.Id, 202)
	query.SetWhereLike(&test.TableItem1.Str, "wa")
	query.SetWhereLikeNot(&test.TableItem1.Str, "wb")
	query.SetWhereIn(&test.TableItem1.Id, []interface{}{203, 204, 205})
	query.SetWhereInNot(&test.TableItem1.Id, []interface{}{206, 207, 208})
	query.SetWhereGt(&test.TableItem1.Id, 209)
	query.SetWhereGte(&test.TableItem1.Id, 210)
	query.SetWhereLt(&test.TableItem1.Id, 211)
	query.SetWhereLte(&test.TableItem1.Id, 212)

	query.SetWhereNest()
	query.SetWhereOr("1 = 1")
	query.SetWhereOrIs(&test.TableItem1.Id, 213)
	query.SetWhereOrIsNot(&test.TableItem1.Id, 214)
	query.SetWhereOrLike(&test.TableItem1.Str, "wc")
	query.SetWhereOrLikeNot(&test.TableItem1.Str, "wd")
	query.SetWhereOrIn(&test.TableItem1.Id, []interface{}{215, 216, 217})
	query.SetWhereOrInNot(&test.TableItem1.Id, []interface{}{218, 219, 220})
	query.SetWhereOrGt(&test.TableItem1.Id, 221)
	query.SetWhereOrGte(&test.TableItem1.Id, 222)
	query.SetWhereOrLt(&test.TableItem1.Id, 223)
	query.SetWhereOrLte(&test.TableItem1.Id, 224)
	query.SetWhereNestClose()

	query.SetWhereOrNest()
	query.SetWhereOr("1 = 1")
	query.SetWhereNestClose()

	query.SetHaving("1 = 1")
	query.SetHavingIs(&test.TableItem1.Id, 301)
	query.SetHavingIsNot(&test.TableItem1.Id, 302)
	query.SetHavingLike(&test.TableItem1.Str, "ha")
	query.SetHavingLikeNot(&test.TableItem1.Str, "hb")
	query.SetHavingIn(&test.TableItem1.Id, []interface{}{303, 304, 305})
	query.SetHavingInNot(&test.TableItem1.Id, []interface{}{306, 307, 308})
	query.SetHavingGt(&test.TableItem1.Id, 309)
	query.SetHavingGte(&test.TableItem1.Id, 310)
	query.SetHavingLt(&test.TableItem1.Id, 311)
	query.SetHavingLte(&test.TableItem1.Id, 312)

	query.SetHavingNest()
	query.SetHavingOr("1 = 1")
	query.SetHavingOrIs(&test.TableItem1.Id, 313)
	query.SetHavingOrIsNot(&test.TableItem1.Id, 314)
	query.SetHavingOrLike(&test.TableItem1.Str, "hc")
	query.SetHavingOrLikeNot(&test.TableItem1.Str, "hd")
	query.SetHavingOrIn(&test.TableItem1.Id, []interface{}{315, 316, 317})
	query.SetHavingOrInNot(&test.TableItem1.Id, []interface{}{318, 319, 320})
	query.SetHavingOrGt(&test.TableItem1.Id, 321)
	query.SetHavingOrGte(&test.TableItem1.Id, 322)
	query.SetHavingOrLt(&test.TableItem1.Id, 323)
	query.SetHavingOrLte(&test.TableItem1.Id, 324)
	query.SetHavingNestClose()

	query.SetGroupBy(&test.TableItem1.Id)
	query.SetGroupBy(&test.TableItem1.Id)
	query.SetGroupBy(&test.TableItem1.Id)

	query.SetHavingOrNest()
	query.SetHavingOr("1 = 1")
	query.SetHavingNestClose()

	query.SetOrderBy(&test.TableItem1.Id)
	query.SetOrderByAsc(&test.TableItem1.Id)
	query.SetOrderByDesc(&test.TableItem1.Id)

	query.SetLimit(10)
	query.SetOffset(20)

	return query
}

func Query_setAllTypeB() *gol.Query {
	query := gol.NewQuery(nil)

	query.SetTable(&test.TableItem1)

	query.SetFrom(&test.TableItem2)
	query.SetFrom(&test.TableTag1)

	query.SetJoin(&test.TableItem3)
	query.SetJoinWhere(&test.TableItem3, &test.TableItem3.Id, " = ", &test.TableItem1.Id)

	query.SetJoinLeft(&test.TableTag2)
	query.SetJoinWhere(&test.TableTag2, &test.TableTag2.Id, " = ", &test.TableItem1.Id)

	query.SetJoinRight(&test.TableTag3)
	query.SetJoinWhere(&test.TableTag3, &test.TableTag3.Id, " = ", &test.TableItem1.Id)

	query.SetJoinWhere(&test.TableItem3, "count(", &test.TableItem3.Id, ") = count(", &test.TableItem3.Id, ")")
	query.SetJoinWhereIs(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereIsNot(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereLike(&test.TableItem3, "count(", &test.TableItem3.Str, ")", "ja")
	query.SetJoinWhereLikeNot(&test.TableItem3, "count(", &test.TableItem3.Str, ")", "jb")
	query.SetJoinWhereIn(&test.TableItem3, "count(", &test.TableItem3.Id, ")", []interface{}{103, 104, 105})
	query.SetJoinWhereInNot(&test.TableItem3, "count(", &test.TableItem3.Id, ")", []interface{}{106, 107, 108})
	query.SetJoinWhereGt(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereGte(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereLt(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereLte(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)

	query.SetJoinWhereNest(&test.TableItem3)
	query.SetJoinWhereOr(&test.TableItem3, "count(", &test.TableItem3.Id, ") = count(", &test.TableItem3.Id, ")")
	query.SetJoinWhereOrIs(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereOrIsNot(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereOrLike(&test.TableItem3, "count(", &test.TableItem3.Str, ")", "jc")
	query.SetJoinWhereOrLikeNot(&test.TableItem3, "count(", &test.TableItem3.Str, ")", "jd")
	query.SetJoinWhereOrIn(&test.TableItem3, "count(", &test.TableItem3.Id, ")", []interface{}{115, 116, 117})
	query.SetJoinWhereOrInNot(&test.TableItem3, "count(", &test.TableItem3.Id, ")", []interface{}{118, 119, 120})
	query.SetJoinWhereOrGt(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereOrGte(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereOrLt(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereOrLte(&test.TableItem3, "count(", &test.TableItem3.Id, ")", &test.TableItem3.Id)
	query.SetJoinWhereNestClose(&test.TableItem3)

	query.SetJoinWhereOrNest(&test.TableItem3)
	query.SetJoinWhereOr(&test.TableItem3, "count(", &test.TableItem3.Id, ") = count(", &test.TableItem3.Id, ")")
	query.SetJoinWhereNestClose(&test.TableItem3)

	query.SetValuesColumn(
		&test.TableItem1.CreateAt,
		&test.TableItem1.UpdateAt,
		&test.TableItem1.DeleteAt,
		&test.TableItem1.Str,
	)
	query.SetValues(
		test.TableItem1.CreateAt,
		test.TableItem1.UpdateAt,
		test.TableItem1.DeleteAt,
		test.TableItem1.Str,
	)
	query.SetValues(
		test.TableItem1.CreateAt,
		test.TableItem1.UpdateAt,
		test.TableItem1.DeleteAt,
		test.TableItem1.Str,
	)

	query.SetConflict(&test.TableItem1.Id)

	query.SetSet(&test.TableItem1.CreateAt, test.TableItem1.CreateAt)
	query.SetSet(&test.TableItem1.UpdateAt, test.TableItem1.UpdateAt)
	query.SetSet(&test.TableItem1.DeleteAt, test.TableItem1.DeleteAt)
	query.SetSet(&test.TableItem1.Str, test.TableItem1.Str)

	query.SetSelect("count(", &test.TableItem1.Id, ")")
	query.SetSelect("count(", &test.TableItem1.Str, ")")
	query.SetSelectAll(&test.TableItem2)

	query.SetWhere("count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetWhereIs("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereIsNot("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereLike("count(", &test.TableItem1.Str, ")", "wa")
	query.SetWhereLikeNot("count(", &test.TableItem1.Str, ")", "wb")
	query.SetWhereIn("count(", &test.TableItem1.Id, ")", []interface{}{203, 204, 205})
	query.SetWhereInNot("count(", &test.TableItem1.Id, ")", []interface{}{206, 207, 208})
	query.SetWhereGt("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereGte("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereLt("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereLte("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)

	query.SetWhereNest()
	query.SetWhereOr("count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetWhereOrIs("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereOrIsNot("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereOrLike("count(", &test.TableItem1.Str, ")", "wc")
	query.SetWhereOrLikeNot("count(", &test.TableItem1.Str, ")", "wd")
	query.SetWhereOrIn("count(", &test.TableItem1.Id, ")", []interface{}{215, 216, 217})
	query.SetWhereOrInNot("count(", &test.TableItem1.Id, ")", []interface{}{218, 219, 220})
	query.SetWhereOrGt("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereOrGte("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereOrLt("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereOrLte("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereNestClose()

	query.SetWhereOrNest()
	query.SetWhereOr("count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetWhereNestClose()

	query.SetHaving("count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetHavingIs("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingIsNot("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingLike("count(", &test.TableItem1.Str, ")", "ha")
	query.SetHavingLikeNot("count(", &test.TableItem1.Str, ")", "hb")
	query.SetHavingIn("count(", &test.TableItem1.Id, ")", []interface{}{303, 304, 305})
	query.SetHavingInNot("count(", &test.TableItem1.Id, ")", []interface{}{306, 307, 308})
	query.SetHavingGt("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingGte("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingLt("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingLte("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)

	query.SetHavingNest()
	query.SetHavingOr("count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetHavingOrIs("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingOrIsNot("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingOrLike("count(", &test.TableItem1.Str, ")", "hc")
	query.SetHavingOrLikeNot("count(", &test.TableItem1.Str, ")", "hd")
	query.SetHavingOrIn("count(", &test.TableItem1.Id, ")", []interface{}{315, 316, 317})
	query.SetHavingOrInNot("count(", &test.TableItem1.Id, ")", []interface{}{318, 319, 320})
	query.SetHavingOrGt("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingOrGte("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingOrLt("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingOrLte("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetHavingNestClose()

	query.SetGroupBy("count(", &test.TableItem1.Id, ")")
	query.SetGroupBy("count(", &test.TableItem1.Id, ")")
	query.SetGroupBy("count(", &test.TableItem1.Id, ")")

	query.SetHavingOrNest()
	query.SetHavingOr("count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetHavingNestClose()

	query.SetOrderBy("count(", &test.TableItem1.Id, ")")
	query.SetOrderByAsc("count(", &test.TableItem1.Id, ")")
	query.SetOrderByDesc("count(", &test.TableItem1.Id, ")")

	query.SetLimit(10)
	query.SetOffset(20)

	return query
}

func TestQuery_GetInsertQuery(t *testing.T) {
	t.Run("insert a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("insert b", func(t *testing.T) {
		query := Query_setAllTypeB()

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?), (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetInsertQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetInsertDoNothingQuery(t *testing.T) {
	t.Run("insert a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetInsertDoNothingQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?) ON CONFLICT DO NOTHING`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("insert b", func(t *testing.T) {
		query := Query_setAllTypeB()

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?), (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetInsertQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetInsertDoUpdateQuery(t *testing.T) {
	t.Run("insert a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetInsertDoUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?) ON CONFLICT ("id") DO UPDATE SET "create_at" = ?, "update_at" = ?, "delete_at" = ?, "str" = ?`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("insert b", func(t *testing.T) {
		query := Query_setAllTypeB()

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?), (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetInsertQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetInsertIgnoreQuery(t *testing.T) {
	t.Run("insert a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetInsertIgnoreQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT IGNORE INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("insert b", func(t *testing.T) {
		query := Query_setAllTypeB()

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?), (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetInsertQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetInsertOnDuplicateKeyUpdateQuery(t *testing.T) {
	t.Run("insert a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetInsertOnDuplicateKeyUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?) ON DUPLICATE KEY UPDATE "create_at" = ?, "update_at" = ?, "delete_at" = ?, "str" = ?`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("insert b", func(t *testing.T) {
		query := Query_setAllTypeB()

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") VALUES (?, ?, ?, ?), (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetInsertQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetInsertSelectUnionQuery(t *testing.T) {
	t.Run("insert a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetInsertSelectUnionQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") SELECT ?, ?, ?, ?`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("insert b", func(t *testing.T) {
		query := Query_setAllTypeB()

		str, valueList, err := query.GetInsertSelectUnionQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("create_at", "update_at", "delete_at", "str") SELECT ?, ?, ?, ? UNION SELECT ?, ?, ?, ?`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetInsertQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetUpdateQuery(t *testing.T) {
	t.Run("update a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `UPDATE "item" SET "create_at" = ?, "update_at" = ?, "delete_at" = ?, "str" = ? WHERE 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."str" LIKE ? AND "item"."str" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."str" LIKE ? OR "item"."str" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 )`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreateAt, test.TableItem1.UpdateAt, test.TableItem1.DeleteAt, test.TableItem1.Str)
			checkList = append(checkList, 201, 202, "wa", "wb", 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, "wc", "wd", 215, 216, 217, 218, 219, 220, 221, 222, 223, 224)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetUpdateQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})

	t.Run("error set not exist", func(t *testing.T) {
		table := test.Item{}
		query := gol.NewQuery(nil)
		query.SetTable(&table)
		_, _, err := query.GetUpdateQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `set not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetDeleteQuery(t *testing.T) {
	t.Run("delete a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetDeleteQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `DELETE FROM "item" WHERE 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."str" LIKE ? AND "item"."str" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."str" LIKE ? OR "item"."str" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 )`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, 201, 202, "wa", "wb", 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, "wc", "wd", 215, 216, 217, 218, 219, 220, 221, 222, 223, 224)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetDeleteQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetTruncateQuery(t *testing.T) {
	t.Run("truncate a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetTruncateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `TRUNCATE TABLE "item"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})
}

func TestQuery_GetTruncateRestartIdentityQuery(t *testing.T) {
	t.Run("truncate a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetTruncateRestartIdentityQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `TRUNCATE TABLE "item" RESTART IDENTITY`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetTruncateRestartIdentityQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetSelectQuery(t *testing.T) {
	t.Run("select a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetSelectQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT "t1"."id", "t1"."str", "t2".* FROM "item" as "t1", "item" as "t2", "PUBLIC"."TAG" as "t3" INNER JOIN "item" as "t4" ON "t4"."id" = "t1"."id" AND 1 = 1 AND "t4"."id" = ? AND "t4"."id" != ? AND "t4"."str" LIKE ? AND "t4"."str" NOT LIKE ? AND "t4"."id" IN (?, ?, ?) AND "t4"."id" NOT IN (?, ?, ?) AND "t4"."id" > ? AND "t4"."id" >= ? AND "t4"."id" < ? AND "t4"."id" <= ? AND ( 1 = 1 OR "t4"."id" = ? OR "t4"."id" != ? OR "t4"."str" LIKE ? OR "t4"."str" NOT LIKE ? OR "t4"."id" IN (?, ?, ?) OR "t4"."id" NOT IN (?, ?, ?) OR "t4"."id" > ? OR "t4"."id" >= ? OR "t4"."id" < ? OR "t4"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "PUBLIC"."TAG" as "t5" ON "t5"."ID" = "t1"."id" RIGHT JOIN "PUBLIC"."TAG" as "t6" ON "t6"."ID" = "t1"."id" WHERE 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."str" LIKE ? AND "t1"."str" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."str" LIKE ? OR "t1"."str" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "t1"."id", "t1"."id", "t1"."id" HAVING 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."str" LIKE ? AND "t1"."str" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."str" LIKE ? OR "t1"."str" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "t1"."id", "t1"."id" ASC, "t1"."id" DESC LIMIT 10 OFFSET 20`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, 101, 102, "ja", "jb", 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, "jc", "jd", 115, 116, 117, 118, 119, 120, 121, 122, 123, 124)
			checkList = append(checkList, 201, 202, "wa", "wb", 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, "wc", "wd", 215, 216, 217, 218, 219, 220, 221, 222, 223, 224)
			checkList = append(checkList, 301, 302, "ha", "hb", 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, "hc", "hd", 315, 316, 317, 318, 319, 320, 321, 322, 323, 324)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("select b", func(t *testing.T) {
		query := Query_setAllTypeB()

		str, valueList, err := query.GetSelectQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT count("t1"."id"), count("t1"."str"), "t2".* FROM "item" as "t1", "item" as "t2", "PUBLIC"."TAG" as "t3" INNER JOIN "item" as "t4" ON "t4"."id" = "t1"."id" AND count("t4"."id") = count("t4"."id") AND count("t4"."id") = "t4"."id" AND count("t4"."id") != "t4"."id" AND count("t4"."str") LIKE ? AND count("t4"."str") NOT LIKE ? AND count("t4"."id") IN (?, ?, ?) AND count("t4"."id") NOT IN (?, ?, ?) AND count("t4"."id") > "t4"."id" AND count("t4"."id") >= "t4"."id" AND count("t4"."id") < "t4"."id" AND count("t4"."id") <= "t4"."id" AND ( count("t4"."id") = count("t4"."id") OR count("t4"."id") = "t4"."id" OR count("t4"."id") != "t4"."id" OR count("t4"."str") LIKE ? OR count("t4"."str") NOT LIKE ? OR count("t4"."id") IN (?, ?, ?) OR count("t4"."id") NOT IN (?, ?, ?) OR count("t4"."id") > "t4"."id" OR count("t4"."id") >= "t4"."id" OR count("t4"."id") < "t4"."id" OR count("t4"."id") <= "t4"."id" ) OR ( count("t4"."id") = count("t4"."id") ) LEFT JOIN "PUBLIC"."TAG" as "t5" ON "t5"."ID" = "t1"."id" RIGHT JOIN "PUBLIC"."TAG" as "t6" ON "t6"."ID" = "t1"."id" WHERE count("t2"."id") = count("t2"."id") AND count("t1"."id") = "t1"."id" AND count("t1"."id") != "t1"."id" AND count("t1"."str") LIKE ? AND count("t1"."str") NOT LIKE ? AND count("t1"."id") IN (?, ?, ?) AND count("t1"."id") NOT IN (?, ?, ?) AND count("t1"."id") > "t1"."id" AND count("t1"."id") >= "t1"."id" AND count("t1"."id") < "t1"."id" AND count("t1"."id") <= "t1"."id" AND ( count("t2"."id") = count("t2"."id") OR count("t1"."id") = "t1"."id" OR count("t1"."id") != "t1"."id" OR count("t1"."str") LIKE ? OR count("t1"."str") NOT LIKE ? OR count("t1"."id") IN (?, ?, ?) OR count("t1"."id") NOT IN (?, ?, ?) OR count("t1"."id") > "t1"."id" OR count("t1"."id") >= "t1"."id" OR count("t1"."id") < "t1"."id" OR count("t1"."id") <= "t1"."id" ) OR ( count("t2"."id") = count("t2"."id") ) GROUP BY count("t1"."id"), count("t1"."id"), count("t1"."id") HAVING count("t2"."id") = count("t2"."id") AND count("t1"."id") = "t1"."id" AND count("t1"."id") != "t1"."id" AND count("t1"."str") LIKE ? AND count("t1"."str") NOT LIKE ? AND count("t1"."id") IN (?, ?, ?) AND count("t1"."id") NOT IN (?, ?, ?) AND count("t1"."id") > "t1"."id" AND count("t1"."id") >= "t1"."id" AND count("t1"."id") < "t1"."id" AND count("t1"."id") <= "t1"."id" AND ( count("t2"."id") = count("t2"."id") OR count("t1"."id") = "t1"."id" OR count("t1"."id") != "t1"."id" OR count("t1"."str") LIKE ? OR count("t1"."str") NOT LIKE ? OR count("t1"."id") IN (?, ?, ?) OR count("t1"."id") NOT IN (?, ?, ?) OR count("t1"."id") > "t1"."id" OR count("t1"."id") >= "t1"."id" OR count("t1"."id") < "t1"."id" OR count("t1"."id") <= "t1"."id" ) OR ( count("t2"."id") = count("t2"."id") ) ORDER BY count("t1"."id"), count("t1"."id") ASC, count("t1"."id") DESC LIMIT 10 OFFSET 20`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, "ja", "jb", 103, 104, 105, 106, 107, 108, "jc", "jd", 115, 116, 117, 118, 119, 120)
			checkList = append(checkList, "wa", "wb", 203, 204, 205, 206, 207, 208, "wc", "wd", 215, 216, 217, 218, 219, 220)
			checkList = append(checkList, "ha", "hb", 303, 304, 305, 306, 307, 308, "hc", "hd", 315, 316, 317, 318, 319, 320)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error select not exist", func(t *testing.T) {
		table := test.Item{}
		query := gol.NewQuery(nil)
		query.SetTable(&table)
		_, _, err := query.GetSelectQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `select not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetSelectQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}

func TestQuery_GetSelectCountQuery(t *testing.T) {
	t.Run("select a", func(t *testing.T) {
		query := Query_setAllTypeA()

		str, valueList, err := query.GetSelectCountQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT count(*) as count FROM "item" as "t1", "item" as "t2", "PUBLIC"."TAG" as "t3" INNER JOIN "item" as "t4" ON "t4"."id" = "t1"."id" AND 1 = 1 AND "t4"."id" = ? AND "t4"."id" != ? AND "t4"."str" LIKE ? AND "t4"."str" NOT LIKE ? AND "t4"."id" IN (?, ?, ?) AND "t4"."id" NOT IN (?, ?, ?) AND "t4"."id" > ? AND "t4"."id" >= ? AND "t4"."id" < ? AND "t4"."id" <= ? AND ( 1 = 1 OR "t4"."id" = ? OR "t4"."id" != ? OR "t4"."str" LIKE ? OR "t4"."str" NOT LIKE ? OR "t4"."id" IN (?, ?, ?) OR "t4"."id" NOT IN (?, ?, ?) OR "t4"."id" > ? OR "t4"."id" >= ? OR "t4"."id" < ? OR "t4"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "PUBLIC"."TAG" as "t5" ON "t5"."ID" = "t1"."id" RIGHT JOIN "PUBLIC"."TAG" as "t6" ON "t6"."ID" = "t1"."id" WHERE 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."str" LIKE ? AND "t1"."str" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."str" LIKE ? OR "t1"."str" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "t1"."id", "t1"."id", "t1"."id" HAVING 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."str" LIKE ? AND "t1"."str" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."str" LIKE ? OR "t1"."str" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "t1"."id", "t1"."id" ASC, "t1"."id" DESC LIMIT 10 OFFSET 20`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, 101, 102, "ja", "jb", 103, 104, 105, 106, 107, 108, 109, 110, 111, 112, 113, 114, "jc", "jd", 115, 116, 117, 118, 119, 120, 121, 122, 123, 124)
			checkList = append(checkList, 201, 202, "wa", "wb", 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, "wc", "wd", 215, 216, 217, 218, 219, 220, 221, 222, 223, 224)
			checkList = append(checkList, 301, 302, "ha", "hb", 303, 304, 305, 306, 307, 308, 309, 310, 311, 312, 313, 314, "hc", "hd", 315, 316, 317, 318, 319, 320, 321, 322, 323, 324)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("select b", func(t *testing.T) {
		query := Query_setAllTypeB()

		str, valueList, err := query.GetSelectCountQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT count(*) as count FROM "item" as "t1", "item" as "t2", "PUBLIC"."TAG" as "t3" INNER JOIN "item" as "t4" ON "t4"."id" = "t1"."id" AND count("t4"."id") = count("t4"."id") AND count("t4"."id") = "t4"."id" AND count("t4"."id") != "t4"."id" AND count("t4"."str") LIKE ? AND count("t4"."str") NOT LIKE ? AND count("t4"."id") IN (?, ?, ?) AND count("t4"."id") NOT IN (?, ?, ?) AND count("t4"."id") > "t4"."id" AND count("t4"."id") >= "t4"."id" AND count("t4"."id") < "t4"."id" AND count("t4"."id") <= "t4"."id" AND ( count("t4"."id") = count("t4"."id") OR count("t4"."id") = "t4"."id" OR count("t4"."id") != "t4"."id" OR count("t4"."str") LIKE ? OR count("t4"."str") NOT LIKE ? OR count("t4"."id") IN (?, ?, ?) OR count("t4"."id") NOT IN (?, ?, ?) OR count("t4"."id") > "t4"."id" OR count("t4"."id") >= "t4"."id" OR count("t4"."id") < "t4"."id" OR count("t4"."id") <= "t4"."id" ) OR ( count("t4"."id") = count("t4"."id") ) LEFT JOIN "PUBLIC"."TAG" as "t5" ON "t5"."ID" = "t1"."id" RIGHT JOIN "PUBLIC"."TAG" as "t6" ON "t6"."ID" = "t1"."id" WHERE count("t2"."id") = count("t2"."id") AND count("t1"."id") = "t1"."id" AND count("t1"."id") != "t1"."id" AND count("t1"."str") LIKE ? AND count("t1"."str") NOT LIKE ? AND count("t1"."id") IN (?, ?, ?) AND count("t1"."id") NOT IN (?, ?, ?) AND count("t1"."id") > "t1"."id" AND count("t1"."id") >= "t1"."id" AND count("t1"."id") < "t1"."id" AND count("t1"."id") <= "t1"."id" AND ( count("t2"."id") = count("t2"."id") OR count("t1"."id") = "t1"."id" OR count("t1"."id") != "t1"."id" OR count("t1"."str") LIKE ? OR count("t1"."str") NOT LIKE ? OR count("t1"."id") IN (?, ?, ?) OR count("t1"."id") NOT IN (?, ?, ?) OR count("t1"."id") > "t1"."id" OR count("t1"."id") >= "t1"."id" OR count("t1"."id") < "t1"."id" OR count("t1"."id") <= "t1"."id" ) OR ( count("t2"."id") = count("t2"."id") ) GROUP BY count("t1"."id"), count("t1"."id"), count("t1"."id") HAVING count("t2"."id") = count("t2"."id") AND count("t1"."id") = "t1"."id" AND count("t1"."id") != "t1"."id" AND count("t1"."str") LIKE ? AND count("t1"."str") NOT LIKE ? AND count("t1"."id") IN (?, ?, ?) AND count("t1"."id") NOT IN (?, ?, ?) AND count("t1"."id") > "t1"."id" AND count("t1"."id") >= "t1"."id" AND count("t1"."id") < "t1"."id" AND count("t1"."id") <= "t1"."id" AND ( count("t2"."id") = count("t2"."id") OR count("t1"."id") = "t1"."id" OR count("t1"."id") != "t1"."id" OR count("t1"."str") LIKE ? OR count("t1"."str") NOT LIKE ? OR count("t1"."id") IN (?, ?, ?) OR count("t1"."id") NOT IN (?, ?, ?) OR count("t1"."id") > "t1"."id" OR count("t1"."id") >= "t1"."id" OR count("t1"."id") < "t1"."id" OR count("t1"."id") <= "t1"."id" ) OR ( count("t2"."id") = count("t2"."id") ) ORDER BY count("t1"."id"), count("t1"."id") ASC, count("t1"."id") DESC LIMIT 10 OFFSET 20`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, "ja", "jb", 103, 104, 105, 106, 107, 108, "jc", "jd", 115, 116, 117, 118, 119, 120)
			checkList = append(checkList, "wa", "wb", 203, 204, 205, 206, 207, 208, "wc", "wd", 215, 216, 217, 218, 219, 220)
			checkList = append(checkList, "ha", "hb", 303, 304, 305, 306, 307, 308, "hc", "hd", 315, 316, 317, 318, 319, 320)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("success", func(t *testing.T) {
		table := test.Item{}
		query := gol.NewQuery(nil)
		query.SetTable(&table)
		query.SetSelectAll(&table)
		str, valueList, err := query.GetSelectQuery()
		if err != nil {
			t.Error(err)
			return
		}
		{
			target := str
			check := `SELECT "item".* FROM "item"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, "")
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("error table not exist", func(t *testing.T) {
		query := gol.NewQuery(nil)
		_, _, err := query.GetSelectCountQuery()
		{
			target := fmt.Sprintf("%v", err)
			check := `table not exist`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}
		}
	})
}
