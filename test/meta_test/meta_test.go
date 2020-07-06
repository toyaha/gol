package meta_test

import (
	"testing"

	"github.com/toyaha/gol"
	"github.com/toyaha/gol/test"
)

func TestQueryMeta(t *testing.T) {
	tableItem := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, metaList []string, tableList []interface{}, field interface{}, checkMap map[string]string) {
		queryMeta := gol.NewMeta(nil)
		queryMeta.DatabaseType = metaList[0]
		queryMeta.NamingConventionForTable = metaList[1]
		queryMeta.NamingConventionForColumn = metaList[2]

		err := queryMeta.Add(tableList[0], tableList[1].(bool))
		if err != nil {
			t.Error(err)
			return
		}

		meta := queryMeta.Get(field)

		{
			target := meta.BaseSchema
			check, ok := checkMap["BaseSchema"]
			if !ok {
				t.Error("BaseSchema is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.BaseTable
			check, ok := checkMap["BaseTable"]
			if !ok {
				t.Error("BaseTable is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.BaseAs
			check, ok := checkMap["BaseAs"]
			if !ok {
				t.Error("BaseAs is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.BaseColumn
			check, ok := checkMap["BaseColumn"]
			if !ok {
				t.Error("BaseColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.Schema
			check, ok := checkMap["Schema"]
			if !ok {
				t.Error("Schema is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.Table
			check, ok := checkMap["Table"]
			if !ok {
				t.Error("Table is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.As
			check, ok := checkMap["As"]
			if !ok {
				t.Error("As is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.Column
			check, ok := checkMap["Column"]
			if !ok {
				t.Error("Column is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.TableColumn
			check, ok := checkMap["TableColumn"]
			if !ok {
				t.Error("TableColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.TableAs
			check, ok := checkMap["TableAs"]
			if !ok {
				t.Error("TableAs is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.TableAsColumn
			check, ok := checkMap["TableAsColumn"]
			if !ok {
				t.Error("TableAsColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.SchemaTable
			check, ok := checkMap["SchemaTable"]
			if !ok {
				t.Error("SchemaTable is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.SchemaTableColumn
			check, ok := checkMap["SchemaTableColumn"]
			if !ok {
				t.Error("SchemaTableColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.SchemaTableAs
			check, ok := checkMap["SchemaTableAs"]
			if !ok {
				t.Error("SchemaTableAs is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := meta.SchemaTableAsColumn
			check, ok := checkMap["SchemaTableAsColumn"]
			if !ok {
				t.Error("SchemaTableAsColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	}

	t.Run("mssql item", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypeMssql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableItem, false}
		var field interface{} = &tableItem.Str
		var checkMap = map[string]string{
			"BaseSchema":          "",
			"BaseTable":           "item",
			"BaseAs":              "",
			"BaseColumn":          "str",
			"Schema":              "",
			"Table":               "[item]",
			"As":                  "",
			"Column":              "[str]",
			"TableColumn":         "[item].[str]",
			"TableAs":             "[item]",
			"TableAsColumn":       "[item].[str]",
			"SchemaTable":         "[item]",
			"SchemaTableColumn":   "[item].[str]",
			"SchemaTableAs":       "[item]",
			"SchemaTableAsColumn": "[item].[str]",
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("mssql tag", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypeMssql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableTag, false}
		var field interface{} = &tableTag.Str
		var checkMap = map[string]string{
			"BaseSchema":          "PUBLIC",
			"BaseTable":           "TAG",
			"BaseAs":              "",
			"BaseColumn":          "STR",
			"Schema":              "[PUBLIC]",
			"Table":               "[TAG]",
			"As":                  "",
			"Column":              "[STR]",
			"TableColumn":         "[TAG].[STR]",
			"TableAs":             "[TAG]",
			"TableAsColumn":       "[TAG].[STR]",
			"SchemaTable":         "[PUBLIC].[TAG]",
			"SchemaTableColumn":   "[PUBLIC].[TAG].[STR]",
			"SchemaTableAs":       "[PUBLIC].[TAG]",
			"SchemaTableAsColumn": "[PUBLIC].[TAG].[STR]",
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("mssql item as", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypeMssql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableItem, true}
		var field interface{} = &tableItem.Str
		var checkMap = map[string]string{
			"BaseSchema":          "",
			"BaseTable":           "item",
			"BaseAs":              "t1",
			"BaseColumn":          "str",
			"Schema":              "",
			"Table":               "[item]",
			"As":                  "[t1]",
			"Column":              "[str]",
			"TableColumn":         "[item].[str]",
			"TableAs":             "[t1]",
			"TableAsColumn":       "[t1].[str]",
			"SchemaTable":         "[item]",
			"SchemaTableColumn":   "[item].[str]",
			"SchemaTableAs":       "[t1]",
			"SchemaTableAsColumn": "[t1].[str]",
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("mssql tag as", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypeMssql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableTag, true}
		var field interface{} = &tableTag.Str
		var checkMap = map[string]string{
			"BaseSchema":          "PUBLIC",
			"BaseTable":           "TAG",
			"BaseAs":              "t1",
			"BaseColumn":          "STR",
			"Schema":              "[PUBLIC]",
			"Table":               "[TAG]",
			"As":                  "[t1]",
			"Column":              "[STR]",
			"TableColumn":         "[TAG].[STR]",
			"TableAs":             "[t1]",
			"TableAsColumn":       "[t1].[STR]",
			"SchemaTable":         "[PUBLIC].[TAG]",
			"SchemaTableColumn":   "[PUBLIC].[TAG].[STR]",
			"SchemaTableAs":       "[t1]",
			"SchemaTableAsColumn": "[t1].[STR]",
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("mysql item", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypeMysql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableItem, false}
		var field interface{} = &tableItem.Str
		var checkMap = map[string]string{
			"BaseSchema":          "",
			"BaseTable":           "item",
			"BaseAs":              "",
			"BaseColumn":          "str",
			"Schema":              "",
			"Table":               "`item`",
			"As":                  "",
			"Column":              "`str`",
			"TableColumn":         "`item`.`str`",
			"TableAs":             "`item`",
			"TableAsColumn":       "`item`.`str`",
			"SchemaTable":         "`item`",
			"SchemaTableColumn":   "`item`.`str`",
			"SchemaTableAs":       "`item`",
			"SchemaTableAsColumn": "`item`.`str`",
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("mysql tag", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypeMysql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableTag, false}
		var field interface{} = &tableTag.Str
		var checkMap = map[string]string{
			"BaseSchema":          "PUBLIC",
			"BaseTable":           "TAG",
			"BaseAs":              "",
			"BaseColumn":          "STR",
			"Schema":              "`PUBLIC`",
			"Table":               "`TAG`",
			"As":                  "",
			"Column":              "`STR`",
			"TableColumn":         "`TAG`.`STR`",
			"TableAs":             "`TAG`",
			"TableAsColumn":       "`TAG`.`STR`",
			"SchemaTable":         "`PUBLIC`.`TAG`",
			"SchemaTableColumn":   "`PUBLIC`.`TAG`.`STR`",
			"SchemaTableAs":       "`PUBLIC`.`TAG`",
			"SchemaTableAsColumn": "`PUBLIC`.`TAG`.`STR`",
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("mysql item as", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypeMysql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableItem, true}
		var field interface{} = &tableItem.Str
		var checkMap = map[string]string{
			"BaseSchema":          "",
			"BaseTable":           "item",
			"BaseAs":              "t1",
			"BaseColumn":          "str",
			"Schema":              "",
			"Table":               "`item`",
			"As":                  "`t1`",
			"Column":              "`str`",
			"TableColumn":         "`item`.`str`",
			"TableAs":             "`t1`",
			"TableAsColumn":       "`t1`.`str`",
			"SchemaTable":         "`item`",
			"SchemaTableColumn":   "`item`.`str`",
			"SchemaTableAs":       "`t1`",
			"SchemaTableAsColumn": "`t1`.`str`",
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("mysql tag as", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypeMysql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableTag, true}
		var field interface{} = &tableTag.Str
		var checkMap = map[string]string{
			"BaseSchema":          "PUBLIC",
			"BaseTable":           "TAG",
			"BaseAs":              "t1",
			"BaseColumn":          "STR",
			"Schema":              "`PUBLIC`",
			"Table":               "`TAG`",
			"As":                  "`t1`",
			"Column":              "`STR`",
			"TableColumn":         "`TAG`.`STR`",
			"TableAs":             "`t1`",
			"TableAsColumn":       "`t1`.`STR`",
			"SchemaTable":         "`PUBLIC`.`TAG`",
			"SchemaTableColumn":   "`PUBLIC`.`TAG`.`STR`",
			"SchemaTableAs":       "`t1`",
			"SchemaTableAsColumn": "`t1`.`STR`",
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("postgresql item", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypePostgresql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableItem, false}
		var field interface{} = &tableItem.Str
		var checkMap = map[string]string{
			"BaseSchema":          ``,
			"BaseTable":           `item`,
			"BaseAs":              ``,
			"BaseColumn":          `str`,
			"Schema":              ``,
			"Table":               `"item"`,
			"As":                  ``,
			"Column":              `"str"`,
			"TableColumn":         `"item"."str"`,
			"TableAs":             `"item"`,
			"TableAsColumn":       `"item"."str"`,
			"SchemaTable":         `"item"`,
			"SchemaTableColumn":   `"item"."str"`,
			"SchemaTableAs":       `"item"`,
			"SchemaTableAsColumn": `"item"."str"`,
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("postgresql tag", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypePostgresql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableTag, false}
		var field interface{} = &tableTag.Str
		var checkMap = map[string]string{
			"BaseSchema":          `PUBLIC`,
			"BaseTable":           `TAG`,
			"BaseAs":              ``,
			"BaseColumn":          `STR`,
			"Schema":              `"PUBLIC"`,
			"Table":               `"TAG"`,
			"As":                  ``,
			"Column":              `"STR"`,
			"TableColumn":         `"TAG"."STR"`,
			"TableAs":             `"TAG"`,
			"TableAsColumn":       `"TAG"."STR"`,
			"SchemaTable":         `"PUBLIC"."TAG"`,
			"SchemaTableColumn":   `"PUBLIC"."TAG"."STR"`,
			"SchemaTableAs":       `"PUBLIC"."TAG"`,
			"SchemaTableAsColumn": `"PUBLIC"."TAG"."STR"`,
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("postgresql item as", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypePostgresql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableItem, true}
		var field interface{} = &tableItem.Str
		var checkMap = map[string]string{
			"BaseSchema":          ``,
			"BaseTable":           `item`,
			"BaseAs":              `t1`,
			"BaseColumn":          `str`,
			"Schema":              ``,
			"Table":               `"item"`,
			"As":                  `"t1"`,
			"Column":              `"str"`,
			"TableColumn":         `"item"."str"`,
			"TableAs":             `"t1"`,
			"TableAsColumn":       `"t1"."str"`,
			"SchemaTable":         `"item"`,
			"SchemaTableColumn":   `"item"."str"`,
			"SchemaTableAs":       `"t1"`,
			"SchemaTableAsColumn": `"t1"."str"`,
		}
		fn(t, metaList, tableList, field, checkMap)
	})

	t.Run("postgresql tag as", func(t *testing.T) {
		var metaList = []string{
			gol.DatabaseTypePostgresql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		var tableList = []interface{}{&tableTag, true}
		var field interface{} = &tableTag.Str
		var checkMap = map[string]string{
			"BaseSchema":          `PUBLIC`,
			"BaseTable":           `TAG`,
			"BaseAs":              `t1`,
			"BaseColumn":          `STR`,
			"Schema":              `"PUBLIC"`,
			"Table":               `"TAG"`,
			"As":                  `"t1"`,
			"Column":              `"STR"`,
			"TableColumn":         `"TAG"."STR"`,
			"TableAs":             `"t1"`,
			"TableAsColumn":       `"t1"."STR"`,
			"SchemaTable":         `"PUBLIC"."TAG"`,
			"SchemaTableColumn":   `"PUBLIC"."TAG"."STR"`,
			"SchemaTableAs":       `"t1"`,
			"SchemaTableAsColumn": `"t1"."STR"`,
		}
		fn(t, metaList, tableList, field, checkMap)
	})
}
