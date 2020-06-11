package query_test

import (
	"fmt"
	"testing"

	"github.com/toyaha/gol"
	"github.com/toyaha/gol/test"
)

func Query_setAllTypeA(schemaMode bool, asMode bool) *gol.Query {
	query := gol.NewQuery(nil)

	if schemaMode && asMode {
		query.SetTableAsWithSchema("s1", &test.TableItem1, "t1")
		query.SetJoinAsWithSchema("s2", &test.TableItem2, "t2", &test.TableItem2.Id, " = ", &test.TableItem1.Id)
		query.SetJoinLeftAsWithSchema("s3", &test.TableItem3, "t3", &test.TableItem3.Id, " = ", &test.TableItem1.Id)
		query.SetJoinRightAsWithSchema("s4", &test.TableItem4, "t4", &test.TableItem4.Id, " = ", &test.TableItem1.Id)
	} else if schemaMode {
		query.SetTableWithSchema("s1", &test.TableItem1)
		query.SetJoinWithSchema("s2", &test.TableItem2, &test.TableItem2.Id, " = ", &test.TableItem1.Id)
		query.SetJoinLeftWithSchema("s3", &test.TableItem3, &test.TableItem3.Id, " = ", &test.TableItem1.Id)
		query.SetJoinRightWithSchema("s4", &test.TableItem4, &test.TableItem4.Id, " = ", &test.TableItem1.Id)
	} else if asMode {
		query.SetTableAs(&test.TableItem1, "t1")
		query.SetJoinAs(&test.TableItem2, "t2", &test.TableItem2.Id, " = ", &test.TableItem1.Id)
		query.SetJoinLeftAs(&test.TableItem3, "t3", &test.TableItem3.Id, " = ", &test.TableItem1.Id)
		query.SetJoinRightAs(&test.TableItem4, "t4", &test.TableItem4.Id, " = ", &test.TableItem1.Id)
	} else {
		query.SetTable(&test.TableItem1)
		query.SetJoin(&test.TableItem2, &test.TableItem2.Id, " = ", &test.TableItem1.Id)
		query.SetJoinLeft(&test.TableItem3, &test.TableItem3.Id, " = ", &test.TableItem1.Id)
		query.SetJoinRight(&test.TableItem4, &test.TableItem4.Id, " = ", &test.TableItem1.Id)
	}

	query.SetJoinWhere(&test.TableItem2, "1 = 1")
	query.SetJoinWhereIs(&test.TableItem2, &test.TableItem2.Id, 101)
	query.SetJoinWhereIsNot(&test.TableItem2, &test.TableItem2.Id, 102)
	query.SetJoinWhereLike(&test.TableItem2, &test.TableItem2.Name, "ja")
	query.SetJoinWhereLikeNot(&test.TableItem2, &test.TableItem2.Name, "jb")
	query.SetJoinWhereIn(&test.TableItem2, &test.TableItem2.Id, []interface{}{103, 104, 105})
	query.SetJoinWhereInNot(&test.TableItem2, &test.TableItem2.Id, []interface{}{106, 107, 108})
	query.SetJoinWhereGt(&test.TableItem2, &test.TableItem2.Id, 109)
	query.SetJoinWhereGte(&test.TableItem2, &test.TableItem2.Id, 110)
	query.SetJoinWhereLt(&test.TableItem2, &test.TableItem2.Id, 111)
	query.SetJoinWhereLte(&test.TableItem2, &test.TableItem2.Id, 112)

	query.SetJoinWhereNest(&test.TableItem2)
	query.SetJoinWhereOr(&test.TableItem2, "1 = 1")
	query.SetJoinWhereOrIs(&test.TableItem2, &test.TableItem2.Id, 113)
	query.SetJoinWhereOrIsNot(&test.TableItem2, &test.TableItem2.Id, 114)
	query.SetJoinWhereOrLike(&test.TableItem2, &test.TableItem2.Name, "jc")
	query.SetJoinWhereOrLikeNot(&test.TableItem2, &test.TableItem2.Name, "jd")
	query.SetJoinWhereOrIn(&test.TableItem2, &test.TableItem2.Id, []interface{}{115, 116, 117})
	query.SetJoinWhereOrInNot(&test.TableItem2, &test.TableItem2.Id, []interface{}{118, 119, 120})
	query.SetJoinWhereOrGt(&test.TableItem2, &test.TableItem2.Id, 121)
	query.SetJoinWhereOrGte(&test.TableItem2, &test.TableItem2.Id, 122)
	query.SetJoinWhereOrLt(&test.TableItem2, &test.TableItem2.Id, 123)
	query.SetJoinWhereOrLte(&test.TableItem2, &test.TableItem2.Id, 124)
	query.SetJoinWhereNestClose(&test.TableItem2)

	query.SetJoinWhereOrNest(&test.TableItem2)
	query.SetJoinWhereOr(&test.TableItem2, "1 = 1")
	query.SetJoinWhereNestClose(&test.TableItem2)

	query.SetValuesColumn(
		&test.TableItem1.CreatedAt,
		&test.TableItem1.UpdatedAt,
		&test.TableItem1.DeletedAt,
		&test.TableItem1.Name,
	)
	query.SetValues(
		test.TableItem1.CreatedAt,
		test.TableItem1.UpdatedAt,
		test.TableItem1.DeletedAt,
		test.TableItem1.Name,
	)

	query.SetSet(&test.TableItem1.CreatedAt, test.TableItem1.CreatedAt)
	query.SetSet(&test.TableItem1.UpdatedAt, test.TableItem1.UpdatedAt)
	query.SetSet(&test.TableItem1.DeletedAt, test.TableItem1.DeletedAt)
	query.SetSet(&test.TableItem1.Name, test.TableItem1.Name)

	query.SetSelect(&test.TableItem1.Id)
	query.SetSelect(&test.TableItem1.Name)
	query.SetSelectAll(&test.TableItem2)

	query.SetWhere("1 = 1")
	query.SetWhereIs(&test.TableItem1.Id, 201)
	query.SetWhereIsNot(&test.TableItem1.Id, 202)
	query.SetWhereLike(&test.TableItem1.Name, "wa")
	query.SetWhereLikeNot(&test.TableItem1.Name, "wb")
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
	query.SetWhereOrLike(&test.TableItem1.Name, "wc")
	query.SetWhereOrLikeNot(&test.TableItem1.Name, "wd")
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
	query.SetHavingLike(&test.TableItem1.Name, "ha")
	query.SetHavingLikeNot(&test.TableItem1.Name, "hb")
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
	query.SetHavingOrLike(&test.TableItem1.Name, "hc")
	query.SetHavingOrLikeNot(&test.TableItem1.Name, "hd")
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

func Query_setAllTypeB(schemaMode bool, asMode bool) *gol.Query {
	query := gol.NewQuery(nil)

	if schemaMode && asMode {
		query.SetTableAsWithSchema("s1", &test.TableItem1, "t1")
		query.SetJoinAsWithSchema("s2", &test.TableItem2, "t2", &test.TableItem2.Id, " = ", &test.TableItem1.Id)
		query.SetJoinLeftAsWithSchema("s3", &test.TableItem3, "t3", &test.TableItem3.Id, " = ", &test.TableItem1.Id)
		query.SetJoinRightAsWithSchema("s4", &test.TableItem4, "t4", &test.TableItem4.Id, " = ", &test.TableItem1.Id)
	} else if schemaMode {
		query.SetTableWithSchema("s1", &test.TableItem1)
		query.SetJoinWithSchema("s2", &test.TableItem2, &test.TableItem2.Id, " = ", &test.TableItem1.Id)
		query.SetJoinLeftWithSchema("s3", &test.TableItem3, &test.TableItem3.Id, " = ", &test.TableItem1.Id)
		query.SetJoinRightWithSchema("s4", &test.TableItem4, &test.TableItem4.Id, " = ", &test.TableItem1.Id)
	} else if asMode {
		query.SetTableAs(&test.TableItem1, "t1")
		query.SetJoinAs(&test.TableItem2, "t2", &test.TableItem2.Id, " = ", &test.TableItem1.Id)
		query.SetJoinLeftAs(&test.TableItem3, "t3", &test.TableItem3.Id, " = ", &test.TableItem1.Id)
		query.SetJoinRightAs(&test.TableItem4, "t4", &test.TableItem4.Id, " = ", &test.TableItem1.Id)
	} else {
		query.SetTable(&test.TableItem1)
		query.SetJoin(&test.TableItem2, &test.TableItem2.Id, " = ", &test.TableItem1.Id)
		query.SetJoinLeft(&test.TableItem3, &test.TableItem3.Id, " = ", &test.TableItem1.Id)
		query.SetJoinRight(&test.TableItem4, &test.TableItem4.Id, " = ", &test.TableItem1.Id)
	}

	query.SetJoinWhere(&test.TableItem2, "count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetJoinWhereIs(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereIsNot(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereLike(&test.TableItem2, "count(", &test.TableItem2.Name, ")", "ja")
	query.SetJoinWhereLikeNot(&test.TableItem2, "count(", &test.TableItem2.Name, ")", "jb")
	query.SetJoinWhereIn(&test.TableItem2, "count(", &test.TableItem2.Id, ")", []interface{}{103, 104, 105})
	query.SetJoinWhereInNot(&test.TableItem2, "count(", &test.TableItem2.Id, ")", []interface{}{106, 107, 108})
	query.SetJoinWhereGt(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereGte(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereLt(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereLte(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)

	query.SetJoinWhereNest(&test.TableItem2)
	query.SetJoinWhereOr(&test.TableItem2, "count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetJoinWhereOrIs(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereOrIsNot(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereOrLike(&test.TableItem2, "count(", &test.TableItem2.Name, ")", "jc")
	query.SetJoinWhereOrLikeNot(&test.TableItem2, "count(", &test.TableItem2.Name, ")", "jd")
	query.SetJoinWhereOrIn(&test.TableItem2, "count(", &test.TableItem2.Id, ")", []interface{}{115, 116, 117})
	query.SetJoinWhereOrInNot(&test.TableItem2, "count(", &test.TableItem2.Id, ")", []interface{}{118, 119, 120})
	query.SetJoinWhereOrGt(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereOrGte(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereOrLt(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereOrLte(&test.TableItem2, "count(", &test.TableItem2.Id, ")", &test.TableItem2.Id)
	query.SetJoinWhereNestClose(&test.TableItem2)

	query.SetJoinWhereOrNest(&test.TableItem2)
	query.SetJoinWhereOr(&test.TableItem2, "count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetJoinWhereNestClose(&test.TableItem2)

	query.SetValuesColumn(
		&test.TableItem1.CreatedAt,
		&test.TableItem1.UpdatedAt,
		&test.TableItem1.DeletedAt,
		&test.TableItem1.Name,
	)
	query.SetValues(
		test.TableItem1.CreatedAt,
		test.TableItem1.UpdatedAt,
		test.TableItem1.DeletedAt,
		test.TableItem1.Name,
	)

	query.SetSet(&test.TableItem1.CreatedAt, test.TableItem1.CreatedAt)
	query.SetSet(&test.TableItem1.UpdatedAt, test.TableItem1.UpdatedAt)
	query.SetSet(&test.TableItem1.DeletedAt, test.TableItem1.DeletedAt)
	query.SetSet(&test.TableItem1.Name, test.TableItem1.Name)

	query.SetSelect("count(", &test.TableItem1.Id, ")")
	query.SetSelect("count(", &test.TableItem1.Name, ")")
	query.SetSelectAll(&test.TableItem2)

	query.SetWhere("count(", &test.TableItem2.Id, ") = count(", &test.TableItem2.Id, ")")
	query.SetWhereIs("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereIsNot("count(", &test.TableItem1.Id, ")", &test.TableItem1.Id)
	query.SetWhereLike("count(", &test.TableItem1.Name, ")", "wa")
	query.SetWhereLikeNot("count(", &test.TableItem1.Name, ")", "wb")
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
	query.SetWhereOrLike("count(", &test.TableItem1.Name, ")", "wc")
	query.SetWhereOrLikeNot("count(", &test.TableItem1.Name, ")", "wd")
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
	query.SetHavingLike("count(", &test.TableItem1.Name, ")", "ha")
	query.SetHavingLikeNot("count(", &test.TableItem1.Name, ")", "hb")
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
	query.SetHavingOrLike("count(", &test.TableItem1.Name, ")", "hc")
	query.SetHavingOrLikeNot("count(", &test.TableItem1.Name, ")", "hd")
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
	t.Run("insert table", func(t *testing.T) {
		query := Query_setAllTypeA(false, false)

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("created_at", "updated_at", "deleted_at", "name") VALUES (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreatedAt, test.TableItem1.UpdatedAt, test.TableItem1.DeletedAt, test.TableItem1.Name)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("insert table as", func(t *testing.T) {
		query := Query_setAllTypeA(false, true)

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "item" ("created_at", "updated_at", "deleted_at", "name") VALUES (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreatedAt, test.TableItem1.UpdatedAt, test.TableItem1.DeletedAt, test.TableItem1.Name)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("insert schema table", func(t *testing.T) {
		query := Query_setAllTypeA(true, false)

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "s1"."item" ("created_at", "updated_at", "deleted_at", "name") VALUES (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreatedAt, test.TableItem1.UpdatedAt, test.TableItem1.DeletedAt, test.TableItem1.Name)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("insert schema table as", func(t *testing.T) {
		query := Query_setAllTypeA(true, true)

		str, valueList, err := query.GetInsertQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `INSERT INTO "s1"."item" ("created_at", "updated_at", "deleted_at", "name") VALUES (?, ?, ?, ?)`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreatedAt, test.TableItem1.UpdatedAt, test.TableItem1.DeletedAt, test.TableItem1.Name)
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
	t.Run("update table", func(t *testing.T) {
		query := Query_setAllTypeA(false, false)

		str, valueList, err := query.GetUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `UPDATE "item" SET "created_at" = ?, "updated_at" = ?, "deleted_at" = ?, "name" = ? WHERE 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 )`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreatedAt, test.TableItem1.UpdatedAt, test.TableItem1.DeletedAt, test.TableItem1.Name)
			checkList = append(checkList, 201, 202, "wa", "wb", 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, "wc", "wd", 215, 216, 217, 218, 219, 220, 221, 222, 223, 224)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("update table as", func(t *testing.T) {
		query := Query_setAllTypeA(false, true)

		str, valueList, err := query.GetUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `UPDATE "item" SET "created_at" = ?, "updated_at" = ?, "deleted_at" = ?, "name" = ? WHERE 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 )`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreatedAt, test.TableItem1.UpdatedAt, test.TableItem1.DeletedAt, test.TableItem1.Name)
			checkList = append(checkList, 201, 202, "wa", "wb", 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, "wc", "wd", 215, 216, 217, 218, 219, 220, 221, 222, 223, 224)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("update schema table", func(t *testing.T) {
		query := Query_setAllTypeA(true, false)

		str, valueList, err := query.GetUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `UPDATE "s1"."item" SET "created_at" = ?, "updated_at" = ?, "deleted_at" = ?, "name" = ? WHERE 1 = 1 AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( 1 = 1 OR "s1"."item"."id" = ? OR "s1"."item"."id" != ? OR "s1"."item"."name" LIKE ? OR "s1"."item"."name" NOT LIKE ? OR "s1"."item"."id" IN (?, ?, ?) OR "s1"."item"."id" NOT IN (?, ?, ?) OR "s1"."item"."id" > ? OR "s1"."item"."id" >= ? OR "s1"."item"."id" < ? OR "s1"."item"."id" <= ? ) OR ( 1 = 1 )`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreatedAt, test.TableItem1.UpdatedAt, test.TableItem1.DeletedAt, test.TableItem1.Name)
			checkList = append(checkList, 201, 202, "wa", "wb", 203, 204, 205, 206, 207, 208, 209, 210, 211, 212, 213, 214, "wc", "wd", 215, 216, 217, 218, 219, 220, 221, 222, 223, 224)
			check := fmt.Sprintf("%v", checkList)
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("update schema table as", func(t *testing.T) {
		query := Query_setAllTypeA(true, true)

		str, valueList, err := query.GetUpdateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `UPDATE "s1"."item" SET "created_at" = ?, "updated_at" = ?, "deleted_at" = ?, "name" = ? WHERE 1 = 1 AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( 1 = 1 OR "s1"."item"."id" = ? OR "s1"."item"."id" != ? OR "s1"."item"."name" LIKE ? OR "s1"."item"."name" NOT LIKE ? OR "s1"."item"."id" IN (?, ?, ?) OR "s1"."item"."id" NOT IN (?, ?, ?) OR "s1"."item"."id" > ? OR "s1"."item"."id" >= ? OR "s1"."item"."id" < ? OR "s1"."item"."id" <= ? ) OR ( 1 = 1 )`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
		{
			target := fmt.Sprintf("%v", valueList)
			var checkList []interface{}
			checkList = append(checkList, test.TableItem1.CreatedAt, test.TableItem1.UpdatedAt, test.TableItem1.DeletedAt, test.TableItem1.Name)
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
	t.Run("delete table", func(t *testing.T) {
		query := Query_setAllTypeA(false, false)

		str, valueList, err := query.GetDeleteQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `DELETE FROM "item" WHERE 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 )`
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

	t.Run("delete table as", func(t *testing.T) {
		query := Query_setAllTypeA(false, true)

		str, valueList, err := query.GetDeleteQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `DELETE FROM "item" WHERE 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 )`
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

	t.Run("delete schema table", func(t *testing.T) {
		query := Query_setAllTypeA(true, false)

		str, valueList, err := query.GetDeleteQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `DELETE FROM "s1"."item" WHERE 1 = 1 AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( 1 = 1 OR "s1"."item"."id" = ? OR "s1"."item"."id" != ? OR "s1"."item"."name" LIKE ? OR "s1"."item"."name" NOT LIKE ? OR "s1"."item"."id" IN (?, ?, ?) OR "s1"."item"."id" NOT IN (?, ?, ?) OR "s1"."item"."id" > ? OR "s1"."item"."id" >= ? OR "s1"."item"."id" < ? OR "s1"."item"."id" <= ? ) OR ( 1 = 1 )`
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

	t.Run("delete schema table as", func(t *testing.T) {
		query := Query_setAllTypeA(true, true)

		str, valueList, err := query.GetDeleteQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `DELETE FROM "s1"."item" WHERE 1 = 1 AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( 1 = 1 OR "s1"."item"."id" = ? OR "s1"."item"."id" != ? OR "s1"."item"."name" LIKE ? OR "s1"."item"."name" NOT LIKE ? OR "s1"."item"."id" IN (?, ?, ?) OR "s1"."item"."id" NOT IN (?, ?, ?) OR "s1"."item"."id" > ? OR "s1"."item"."id" >= ? OR "s1"."item"."id" < ? OR "s1"."item"."id" <= ? ) OR ( 1 = 1 )`
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
	t.Run("truncate table", func(t *testing.T) {
		query := Query_setAllTypeA(false, false)

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

	t.Run("truncate table as", func(t *testing.T) {
		query := Query_setAllTypeA(false, true)

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

	t.Run("truncate schema table", func(t *testing.T) {
		query := Query_setAllTypeA(true, false)

		str, valueList, err := query.GetTruncateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `TRUNCATE TABLE "s1"."item"`
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

	t.Run("truncate schema table as", func(t *testing.T) {
		query := Query_setAllTypeA(true, true)

		str, valueList, err := query.GetTruncateQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `TRUNCATE TABLE "s1"."item"`
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
		_, _, err := query.GetTruncateQuery()
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

func TestQuery_GetTruncateRestartIdentityQuery(t *testing.T) {
	t.Run("truncate table", func(t *testing.T) {
		query := Query_setAllTypeA(false, false)

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

	t.Run("truncate table as", func(t *testing.T) {
		query := Query_setAllTypeA(false, true)

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

	t.Run("truncate schema table", func(t *testing.T) {
		query := Query_setAllTypeA(true, false)

		str, valueList, err := query.GetTruncateRestartIdentityQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `TRUNCATE TABLE "s1"."item" RESTART IDENTITY`
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

	t.Run("truncate schema table as", func(t *testing.T) {
		query := Query_setAllTypeA(true, true)

		str, valueList, err := query.GetTruncateRestartIdentityQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `TRUNCATE TABLE "s1"."item" RESTART IDENTITY`
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
	t.Run("select table", func(t *testing.T) {
		query := Query_setAllTypeA(false, false)

		str, valueList, err := query.GetSelectQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT "item"."id", "item"."name", "item".* FROM "item" INNER JOIN "item" ON "item"."id" = "item"."id" AND 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "item" ON "item"."id" = "item"."id" RIGHT JOIN "item" ON "item"."id" = "item"."id" WHERE 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "item"."id", "item"."id", "item"."id" HAVING 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "item"."id", "item"."id" ASC, "item"."id" DESC LIMIT 10 OFFSET 20`
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

	t.Run("select table as", func(t *testing.T) {
		query := Query_setAllTypeA(false, true)

		str, valueList, err := query.GetSelectQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT "t1"."id", "t1"."name", "t2".* FROM "item" as "t1" INNER JOIN "item" as "t2" ON "t2"."id" = "t1"."id" AND 1 = 1 AND "t2"."id" = ? AND "t2"."id" != ? AND "t2"."name" LIKE ? AND "t2"."name" NOT LIKE ? AND "t2"."id" IN (?, ?, ?) AND "t2"."id" NOT IN (?, ?, ?) AND "t2"."id" > ? AND "t2"."id" >= ? AND "t2"."id" < ? AND "t2"."id" <= ? AND ( 1 = 1 OR "t2"."id" = ? OR "t2"."id" != ? OR "t2"."name" LIKE ? OR "t2"."name" NOT LIKE ? OR "t2"."id" IN (?, ?, ?) OR "t2"."id" NOT IN (?, ?, ?) OR "t2"."id" > ? OR "t2"."id" >= ? OR "t2"."id" < ? OR "t2"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "item" as "t3" ON "t3"."id" = "t1"."id" RIGHT JOIN "item" as "t4" ON "t4"."id" = "t1"."id" WHERE 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."name" LIKE ? OR "t1"."name" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "t1"."id", "t1"."id", "t1"."id" HAVING 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."name" LIKE ? OR "t1"."name" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "t1"."id", "t1"."id" ASC, "t1"."id" DESC LIMIT 10 OFFSET 20`
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

	t.Run("select schema table", func(t *testing.T) {
		query := Query_setAllTypeA(true, false)

		str, valueList, err := query.GetSelectQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT "s1"."item"."id", "s1"."item"."name", "s2"."item".* FROM "s1"."item" INNER JOIN "s2"."item" ON "s2"."item"."id" = "s1"."item"."id" AND 1 = 1 AND "s2"."item"."id" = ? AND "s2"."item"."id" != ? AND "s2"."item"."name" LIKE ? AND "s2"."item"."name" NOT LIKE ? AND "s2"."item"."id" IN (?, ?, ?) AND "s2"."item"."id" NOT IN (?, ?, ?) AND "s2"."item"."id" > ? AND "s2"."item"."id" >= ? AND "s2"."item"."id" < ? AND "s2"."item"."id" <= ? AND ( 1 = 1 OR "s2"."item"."id" = ? OR "s2"."item"."id" != ? OR "s2"."item"."name" LIKE ? OR "s2"."item"."name" NOT LIKE ? OR "s2"."item"."id" IN (?, ?, ?) OR "s2"."item"."id" NOT IN (?, ?, ?) OR "s2"."item"."id" > ? OR "s2"."item"."id" >= ? OR "s2"."item"."id" < ? OR "s2"."item"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "s3"."item" ON "s3"."item"."id" = "s1"."item"."id" RIGHT JOIN "s4"."item" ON "s4"."item"."id" = "s1"."item"."id" WHERE 1 = 1 AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( 1 = 1 OR "s1"."item"."id" = ? OR "s1"."item"."id" != ? OR "s1"."item"."name" LIKE ? OR "s1"."item"."name" NOT LIKE ? OR "s1"."item"."id" IN (?, ?, ?) OR "s1"."item"."id" NOT IN (?, ?, ?) OR "s1"."item"."id" > ? OR "s1"."item"."id" >= ? OR "s1"."item"."id" < ? OR "s1"."item"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "s1"."item"."id", "s1"."item"."id", "s1"."item"."id" HAVING 1 = 1 AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( 1 = 1 OR "s1"."item"."id" = ? OR "s1"."item"."id" != ? OR "s1"."item"."name" LIKE ? OR "s1"."item"."name" NOT LIKE ? OR "s1"."item"."id" IN (?, ?, ?) OR "s1"."item"."id" NOT IN (?, ?, ?) OR "s1"."item"."id" > ? OR "s1"."item"."id" >= ? OR "s1"."item"."id" < ? OR "s1"."item"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "s1"."item"."id", "s1"."item"."id" ASC, "s1"."item"."id" DESC LIMIT 10 OFFSET 20`
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

	t.Run("select schema table as", func(t *testing.T) {
		query := Query_setAllTypeA(true, true)

		str, valueList, err := query.GetSelectQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT "t1"."id", "t1"."name", "t2".* FROM "s1"."item" as "t1" INNER JOIN "s2"."item" as "t2" ON "t2"."id" = "t1"."id" AND 1 = 1 AND "t2"."id" = ? AND "t2"."id" != ? AND "t2"."name" LIKE ? AND "t2"."name" NOT LIKE ? AND "t2"."id" IN (?, ?, ?) AND "t2"."id" NOT IN (?, ?, ?) AND "t2"."id" > ? AND "t2"."id" >= ? AND "t2"."id" < ? AND "t2"."id" <= ? AND ( 1 = 1 OR "t2"."id" = ? OR "t2"."id" != ? OR "t2"."name" LIKE ? OR "t2"."name" NOT LIKE ? OR "t2"."id" IN (?, ?, ?) OR "t2"."id" NOT IN (?, ?, ?) OR "t2"."id" > ? OR "t2"."id" >= ? OR "t2"."id" < ? OR "t2"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "s3"."item" as "t3" ON "t3"."id" = "t1"."id" RIGHT JOIN "s4"."item" as "t4" ON "t4"."id" = "t1"."id" WHERE 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."name" LIKE ? OR "t1"."name" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "t1"."id", "t1"."id", "t1"."id" HAVING 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."name" LIKE ? OR "t1"."name" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "t1"."id", "t1"."id" ASC, "t1"."id" DESC LIMIT 10 OFFSET 20`
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

	t.Run("select b table", func(t *testing.T) {
		query := Query_setAllTypeB(false, false)

		str, valueList, err := query.GetSelectQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT count("item"."id"), count("item"."name"), "item".* FROM "item" INNER JOIN "item" ON "item"."id" = "item"."id" AND count("item"."id") = count("item"."id") AND count("item"."id") = "item"."id" AND count("item"."id") != "item"."id" AND count("item"."name") LIKE ? AND count("item"."name") NOT LIKE ? AND count("item"."id") IN (?, ?, ?) AND count("item"."id") NOT IN (?, ?, ?) AND count("item"."id") > "item"."id" AND count("item"."id") >= "item"."id" AND count("item"."id") < "item"."id" AND count("item"."id") <= "item"."id" AND ( count("item"."id") = count("item"."id") OR count("item"."id") = "item"."id" OR count("item"."id") != "item"."id" OR count("item"."name") LIKE ? OR count("item"."name") NOT LIKE ? OR count("item"."id") IN (?, ?, ?) OR count("item"."id") NOT IN (?, ?, ?) OR count("item"."id") > "item"."id" OR count("item"."id") >= "item"."id" OR count("item"."id") < "item"."id" OR count("item"."id") <= "item"."id" ) OR ( count("item"."id") = count("item"."id") ) LEFT JOIN "item" ON "item"."id" = "item"."id" RIGHT JOIN "item" ON "item"."id" = "item"."id" WHERE count("item"."id") = count("item"."id") AND count("item"."id") = "item"."id" AND count("item"."id") != "item"."id" AND count("item"."name") LIKE ? AND count("item"."name") NOT LIKE ? AND count("item"."id") IN (?, ?, ?) AND count("item"."id") NOT IN (?, ?, ?) AND count("item"."id") > "item"."id" AND count("item"."id") >= "item"."id" AND count("item"."id") < "item"."id" AND count("item"."id") <= "item"."id" AND ( count("item"."id") = count("item"."id") OR count("item"."id") = "item"."id" OR count("item"."id") != "item"."id" OR count("item"."name") LIKE ? OR count("item"."name") NOT LIKE ? OR count("item"."id") IN (?, ?, ?) OR count("item"."id") NOT IN (?, ?, ?) OR count("item"."id") > "item"."id" OR count("item"."id") >= "item"."id" OR count("item"."id") < "item"."id" OR count("item"."id") <= "item"."id" ) OR ( count("item"."id") = count("item"."id") ) GROUP BY count("item"."id"), count("item"."id"), count("item"."id") HAVING count("item"."id") = count("item"."id") AND count("item"."id") = "item"."id" AND count("item"."id") != "item"."id" AND count("item"."name") LIKE ? AND count("item"."name") NOT LIKE ? AND count("item"."id") IN (?, ?, ?) AND count("item"."id") NOT IN (?, ?, ?) AND count("item"."id") > "item"."id" AND count("item"."id") >= "item"."id" AND count("item"."id") < "item"."id" AND count("item"."id") <= "item"."id" AND ( count("item"."id") = count("item"."id") OR count("item"."id") = "item"."id" OR count("item"."id") != "item"."id" OR count("item"."name") LIKE ? OR count("item"."name") NOT LIKE ? OR count("item"."id") IN (?, ?, ?) OR count("item"."id") NOT IN (?, ?, ?) OR count("item"."id") > "item"."id" OR count("item"."id") >= "item"."id" OR count("item"."id") < "item"."id" OR count("item"."id") <= "item"."id" ) OR ( count("item"."id") = count("item"."id") ) ORDER BY count("item"."id"), count("item"."id") ASC, count("item"."id") DESC LIMIT 10 OFFSET 20`
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
		query := gol.NewQuery(nil)
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
		table := test.Item{}
		query := gol.NewQuery(nil)
		query.SetSelectAll(&table.Name)
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
}

func TestQuery_GetSelectCountQuery(t *testing.T) {
	t.Run("select table", func(t *testing.T) {
		query := Query_setAllTypeA(false, false)

		str, valueList, err := query.GetSelectCountQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT count(*) as count FROM "item" INNER JOIN "item" ON "item"."id" = "item"."id" AND 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "item" ON "item"."id" = "item"."id" RIGHT JOIN "item" ON "item"."id" = "item"."id" WHERE 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "item"."id", "item"."id", "item"."id" HAVING 1 = 1 AND "item"."id" = ? AND "item"."id" != ? AND "item"."name" LIKE ? AND "item"."name" NOT LIKE ? AND "item"."id" IN (?, ?, ?) AND "item"."id" NOT IN (?, ?, ?) AND "item"."id" > ? AND "item"."id" >= ? AND "item"."id" < ? AND "item"."id" <= ? AND ( 1 = 1 OR "item"."id" = ? OR "item"."id" != ? OR "item"."name" LIKE ? OR "item"."name" NOT LIKE ? OR "item"."id" IN (?, ?, ?) OR "item"."id" NOT IN (?, ?, ?) OR "item"."id" > ? OR "item"."id" >= ? OR "item"."id" < ? OR "item"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "item"."id", "item"."id" ASC, "item"."id" DESC LIMIT 10 OFFSET 20`
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

	t.Run("select table as", func(t *testing.T) {
		query := Query_setAllTypeA(false, true)

		str, valueList, err := query.GetSelectCountQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT count(*) as count FROM "item" as "t1" INNER JOIN "item" as "t2" ON "t2"."id" = "t1"."id" AND 1 = 1 AND "t2"."id" = ? AND "t2"."id" != ? AND "t2"."name" LIKE ? AND "t2"."name" NOT LIKE ? AND "t2"."id" IN (?, ?, ?) AND "t2"."id" NOT IN (?, ?, ?) AND "t2"."id" > ? AND "t2"."id" >= ? AND "t2"."id" < ? AND "t2"."id" <= ? AND ( 1 = 1 OR "t2"."id" = ? OR "t2"."id" != ? OR "t2"."name" LIKE ? OR "t2"."name" NOT LIKE ? OR "t2"."id" IN (?, ?, ?) OR "t2"."id" NOT IN (?, ?, ?) OR "t2"."id" > ? OR "t2"."id" >= ? OR "t2"."id" < ? OR "t2"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "item" as "t3" ON "t3"."id" = "t1"."id" RIGHT JOIN "item" as "t4" ON "t4"."id" = "t1"."id" WHERE 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."name" LIKE ? OR "t1"."name" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "t1"."id", "t1"."id", "t1"."id" HAVING 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."name" LIKE ? OR "t1"."name" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "t1"."id", "t1"."id" ASC, "t1"."id" DESC LIMIT 10 OFFSET 20`
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

	t.Run("select schema table", func(t *testing.T) {
		query := Query_setAllTypeA(true, false)

		str, valueList, err := query.GetSelectCountQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT count(*) as count FROM "s1"."item" INNER JOIN "s2"."item" ON "s2"."item"."id" = "s1"."item"."id" AND 1 = 1 AND "s2"."item"."id" = ? AND "s2"."item"."id" != ? AND "s2"."item"."name" LIKE ? AND "s2"."item"."name" NOT LIKE ? AND "s2"."item"."id" IN (?, ?, ?) AND "s2"."item"."id" NOT IN (?, ?, ?) AND "s2"."item"."id" > ? AND "s2"."item"."id" >= ? AND "s2"."item"."id" < ? AND "s2"."item"."id" <= ? AND ( 1 = 1 OR "s2"."item"."id" = ? OR "s2"."item"."id" != ? OR "s2"."item"."name" LIKE ? OR "s2"."item"."name" NOT LIKE ? OR "s2"."item"."id" IN (?, ?, ?) OR "s2"."item"."id" NOT IN (?, ?, ?) OR "s2"."item"."id" > ? OR "s2"."item"."id" >= ? OR "s2"."item"."id" < ? OR "s2"."item"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "s3"."item" ON "s3"."item"."id" = "s1"."item"."id" RIGHT JOIN "s4"."item" ON "s4"."item"."id" = "s1"."item"."id" WHERE 1 = 1 AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( 1 = 1 OR "s1"."item"."id" = ? OR "s1"."item"."id" != ? OR "s1"."item"."name" LIKE ? OR "s1"."item"."name" NOT LIKE ? OR "s1"."item"."id" IN (?, ?, ?) OR "s1"."item"."id" NOT IN (?, ?, ?) OR "s1"."item"."id" > ? OR "s1"."item"."id" >= ? OR "s1"."item"."id" < ? OR "s1"."item"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "s1"."item"."id", "s1"."item"."id", "s1"."item"."id" HAVING 1 = 1 AND "s1"."item"."id" = ? AND "s1"."item"."id" != ? AND "s1"."item"."name" LIKE ? AND "s1"."item"."name" NOT LIKE ? AND "s1"."item"."id" IN (?, ?, ?) AND "s1"."item"."id" NOT IN (?, ?, ?) AND "s1"."item"."id" > ? AND "s1"."item"."id" >= ? AND "s1"."item"."id" < ? AND "s1"."item"."id" <= ? AND ( 1 = 1 OR "s1"."item"."id" = ? OR "s1"."item"."id" != ? OR "s1"."item"."name" LIKE ? OR "s1"."item"."name" NOT LIKE ? OR "s1"."item"."id" IN (?, ?, ?) OR "s1"."item"."id" NOT IN (?, ?, ?) OR "s1"."item"."id" > ? OR "s1"."item"."id" >= ? OR "s1"."item"."id" < ? OR "s1"."item"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "s1"."item"."id", "s1"."item"."id" ASC, "s1"."item"."id" DESC LIMIT 10 OFFSET 20`
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

	t.Run("select schema table as", func(t *testing.T) {
		query := Query_setAllTypeA(true, true)

		str, valueList, err := query.GetSelectCountQuery()
		if err != nil {
			t.Error(err)
			return
		}

		{
			target := str
			check := `SELECT count(*) as count FROM "s1"."item" as "t1" INNER JOIN "s2"."item" as "t2" ON "t2"."id" = "t1"."id" AND 1 = 1 AND "t2"."id" = ? AND "t2"."id" != ? AND "t2"."name" LIKE ? AND "t2"."name" NOT LIKE ? AND "t2"."id" IN (?, ?, ?) AND "t2"."id" NOT IN (?, ?, ?) AND "t2"."id" > ? AND "t2"."id" >= ? AND "t2"."id" < ? AND "t2"."id" <= ? AND ( 1 = 1 OR "t2"."id" = ? OR "t2"."id" != ? OR "t2"."name" LIKE ? OR "t2"."name" NOT LIKE ? OR "t2"."id" IN (?, ?, ?) OR "t2"."id" NOT IN (?, ?, ?) OR "t2"."id" > ? OR "t2"."id" >= ? OR "t2"."id" < ? OR "t2"."id" <= ? ) OR ( 1 = 1 ) LEFT JOIN "s3"."item" as "t3" ON "t3"."id" = "t1"."id" RIGHT JOIN "s4"."item" as "t4" ON "t4"."id" = "t1"."id" WHERE 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."name" LIKE ? OR "t1"."name" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) GROUP BY "t1"."id", "t1"."id", "t1"."id" HAVING 1 = 1 AND "t1"."id" = ? AND "t1"."id" != ? AND "t1"."name" LIKE ? AND "t1"."name" NOT LIKE ? AND "t1"."id" IN (?, ?, ?) AND "t1"."id" NOT IN (?, ?, ?) AND "t1"."id" > ? AND "t1"."id" >= ? AND "t1"."id" < ? AND "t1"."id" <= ? AND ( 1 = 1 OR "t1"."id" = ? OR "t1"."id" != ? OR "t1"."name" LIKE ? OR "t1"."name" NOT LIKE ? OR "t1"."id" IN (?, ?, ?) OR "t1"."id" NOT IN (?, ?, ?) OR "t1"."id" > ? OR "t1"."id" >= ? OR "t1"."id" < ? OR "t1"."id" <= ? ) OR ( 1 = 1 ) ORDER BY "t1"."id", "t1"."id" ASC, "t1"."id" DESC LIMIT 10 OFFSET 20`
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
