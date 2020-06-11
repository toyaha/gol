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

		err := queryMeta.Add(tableList[0].(string), tableList[1], tableList[2].(string))
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

	t.Run("mysql table", func(t *testing.T) {
		metaList := []string{
			gol.DatabaseTypeMysql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		{
			tableList := []interface{}{"", &tableItem, ""}
			var field interface{} = &tableItem.Name
			checkMap := map[string]string{
				"BaseSchema":          ``,
				"BaseTable":           `item`,
				"BaseAs":              ``,
				"BaseColumn":          `name`,
				"Schema":              ``,
				"Table":               `item`,
				"As":                  ``,
				"Column":              `name`,
				"TableColumn":         `item.name`,
				"TableAs":             `item`,
				"TableAsColumn":       `item.name`,
				"SchemaTable":         `item`,
				"SchemaTableColumn":   `item.name`,
				"SchemaTableAs":       `item`,
				"SchemaTableAsColumn": `item.name`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
		{
			tableList := []interface{}{"", &tableTag, ""}
			var field interface{} = &tableTag.Name
			checkMap := map[string]string{
				"BaseSchema":          `PUBLIC`,
				"BaseTable":           `TAG`,
				"BaseAs":              ``,
				"BaseColumn":          `NAME`,
				"Schema":              `PUBLIC`,
				"Table":               `TAG`,
				"As":                  ``,
				"Column":              `NAME`,
				"TableColumn":         `TAG.NAME`,
				"TableAs":             `TAG`,
				"TableAsColumn":       `TAG.NAME`,
				"SchemaTable":         `PUBLIC.TAG`,
				"SchemaTableColumn":   `PUBLIC.TAG.NAME`,
				"SchemaTableAs":       `PUBLIC.TAG`,
				"SchemaTableAsColumn": `PUBLIC.TAG.NAME`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
	})

	t.Run("mysql table as", func(t *testing.T) {
		metaList := []string{
			gol.DatabaseTypeMysql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		{
			tableList := []interface{}{"", &tableItem, "t1"}
			var field interface{} = &tableItem.Name
			checkMap := map[string]string{
				"BaseSchema":          ``,
				"BaseTable":           `item`,
				"BaseAs":              `t1`,
				"BaseColumn":          `name`,
				"Schema":              ``,
				"Table":               `item`,
				"As":                  `t1`,
				"Column":              `name`,
				"TableColumn":         `item.name`,
				"TableAs":             `t1`,
				"TableAsColumn":       `t1.name`,
				"SchemaTable":         `item`,
				"SchemaTableColumn":   `item.name`,
				"SchemaTableAs":       `t1`,
				"SchemaTableAsColumn": `t1.name`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
		{
			tableList := []interface{}{"", &tableTag, "t1"}
			var field interface{} = &tableTag.Name
			checkMap := map[string]string{
				"BaseSchema":          `PUBLIC`,
				"BaseTable":           `TAG`,
				"BaseAs":              `t1`,
				"BaseColumn":          `NAME`,
				"Schema":              `PUBLIC`,
				"Table":               `TAG`,
				"As":                  `t1`,
				"Column":              `NAME`,
				"TableColumn":         `TAG.NAME`,
				"TableAs":             `t1`,
				"TableAsColumn":       `t1.NAME`,
				"SchemaTable":         `PUBLIC.TAG`,
				"SchemaTableColumn":   `PUBLIC.TAG.NAME`,
				"SchemaTableAs":       `t1`,
				"SchemaTableAsColumn": `t1.NAME`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
	})

	t.Run("postgresql table", func(t *testing.T) {
		metaList := []string{
			gol.DatabaseTypePostgresql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		{
			tableList := []interface{}{"", &tableItem, ""}
			var field interface{} = &tableItem.Name
			checkMap := map[string]string{
				"BaseSchema":          ``,
				"BaseTable":           `item`,
				"BaseAs":              ``,
				"BaseColumn":          `name`,
				"Schema":              ``,
				"Table":               `"item"`,
				"As":                  ``,
				"Column":              `"name"`,
				"TableColumn":         `"item"."name"`,
				"TableAs":             `"item"`,
				"TableAsColumn":       `"item"."name"`,
				"SchemaTable":         `"item"`,
				"SchemaTableColumn":   `"item"."name"`,
				"SchemaTableAs":       `"item"`,
				"SchemaTableAsColumn": `"item"."name"`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
		{
			tableList := []interface{}{"", &tableTag, ""}
			var field interface{} = &tableTag.Name
			checkMap := map[string]string{
				"BaseSchema":          `PUBLIC`,
				"BaseTable":           `TAG`,
				"BaseAs":              ``,
				"BaseColumn":          `NAME`,
				"Schema":              `"PUBLIC"`,
				"Table":               `"TAG"`,
				"As":                  ``,
				"Column":              `"NAME"`,
				"TableColumn":         `"TAG"."NAME"`,
				"TableAs":             `"TAG"`,
				"TableAsColumn":       `"TAG"."NAME"`,
				"SchemaTable":         `"PUBLIC"."TAG"`,
				"SchemaTableColumn":   `"PUBLIC"."TAG"."NAME"`,
				"SchemaTableAs":       `"PUBLIC"."TAG"`,
				"SchemaTableAsColumn": `"PUBLIC"."TAG"."NAME"`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
	})

	t.Run("postgresql table as", func(t *testing.T) {
		metaList := []string{
			gol.DatabaseTypePostgresql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		{
			tableList := []interface{}{"", &tableItem, "t1"}
			var field interface{} = &tableItem.Name
			checkMap := map[string]string{
				"BaseSchema":          ``,
				"BaseTable":           `item`,
				"BaseAs":              `t1`,
				"BaseColumn":          `name`,
				"Schema":              ``,
				"Table":               `"item"`,
				"As":                  `"t1"`,
				"Column":              `"name"`,
				"TableColumn":         `"item"."name"`,
				"TableAs":             `"t1"`,
				"TableAsColumn":       `"t1"."name"`,
				"SchemaTable":         `"item"`,
				"SchemaTableColumn":   `"item"."name"`,
				"SchemaTableAs":       `"t1"`,
				"SchemaTableAsColumn": `"t1"."name"`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
		{
			tableList := []interface{}{"", &tableTag, "t1"}
			var field interface{} = &tableTag.Name
			checkMap := map[string]string{
				"BaseSchema":          `PUBLIC`,
				"BaseTable":           `TAG`,
				"BaseAs":              `t1`,
				"BaseColumn":          `NAME`,
				"Schema":              `"PUBLIC"`,
				"Table":               `"TAG"`,
				"As":                  `"t1"`,
				"Column":              `"NAME"`,
				"TableColumn":         `"TAG"."NAME"`,
				"TableAs":             `"t1"`,
				"TableAsColumn":       `"t1"."NAME"`,
				"SchemaTable":         `"PUBLIC"."TAG"`,
				"SchemaTableColumn":   `"PUBLIC"."TAG"."NAME"`,
				"SchemaTableAs":       `"t1"`,
				"SchemaTableAsColumn": `"t1"."NAME"`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
	})

	t.Run("postgresql schema table", func(t *testing.T) {
		metaList := []string{
			gol.DatabaseTypePostgresql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		{
			tableList := []interface{}{"s1", &tableItem, ""}
			var field interface{} = &tableItem.Name
			checkMap := map[string]string{
				"BaseSchema":          `s1`,
				"BaseTable":           `item`,
				"BaseAs":              ``,
				"BaseColumn":          `name`,
				"Schema":              `"s1"`,
				"Table":               `"item"`,
				"As":                  ``,
				"Column":              `"name"`,
				"TableColumn":         `"item"."name"`,
				"TableAs":             `"item"`,
				"TableAsColumn":       `"item"."name"`,
				"SchemaTable":         `"s1"."item"`,
				"SchemaTableColumn":   `"s1"."item"."name"`,
				"SchemaTableAs":       `"s1"."item"`,
				"SchemaTableAsColumn": `"s1"."item"."name"`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
		{
			tableList := []interface{}{"s1", &tableTag, ""}
			var field interface{} = &tableTag.Name
			checkMap := map[string]string{
				"BaseSchema":          `s1`,
				"BaseTable":           `TAG`,
				"BaseAs":              ``,
				"BaseColumn":          `NAME`,
				"Schema":              `"s1"`,
				"Table":               `"TAG"`,
				"As":                  ``,
				"Column":              `"NAME"`,
				"TableColumn":         `"TAG"."NAME"`,
				"TableAs":             `"TAG"`,
				"TableAsColumn":       `"TAG"."NAME"`,
				"SchemaTable":         `"s1"."TAG"`,
				"SchemaTableColumn":   `"s1"."TAG"."NAME"`,
				"SchemaTableAs":       `"s1"."TAG"`,
				"SchemaTableAsColumn": `"s1"."TAG"."NAME"`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
	})

	t.Run("postgresql schema table as", func(t *testing.T) {
		metaList := []string{
			gol.DatabaseTypePostgresql,
			gol.NamingConventionSnakeCase,
			gol.NamingConventionSnakeCase,
		}
		{
			tableList := []interface{}{"s1", &tableItem, "t1"}
			var field interface{} = &tableItem.Name
			checkMap := map[string]string{
				"BaseSchema":          `s1`,
				"BaseTable":           `item`,
				"BaseAs":              `t1`,
				"BaseColumn":          `name`,
				"Schema":              `"s1"`,
				"Table":               `"item"`,
				"As":                  `"t1"`,
				"Column":              `"name"`,
				"TableColumn":         `"item"."name"`,
				"TableAs":             `"t1"`,
				"TableAsColumn":       `"t1"."name"`,
				"SchemaTable":         `"s1"."item"`,
				"SchemaTableColumn":   `"s1"."item"."name"`,
				"SchemaTableAs":       `"t1"`,
				"SchemaTableAsColumn": `"t1"."name"`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
		{
			tableList := []interface{}{"s1", &tableTag, "t1"}
			var field interface{} = &tableTag.Name
			checkMap := map[string]string{
				"BaseSchema":          `s1`,
				"BaseTable":           `TAG`,
				"BaseAs":              `t1`,
				"BaseColumn":          `NAME`,
				"Schema":              `"s1"`,
				"Table":               `"TAG"`,
				"As":                  `"t1"`,
				"Column":              `"NAME"`,
				"TableColumn":         `"TAG"."NAME"`,
				"TableAs":             `"t1"`,
				"TableAsColumn":       `"t1"."NAME"`,
				"SchemaTable":         `"s1"."TAG"`,
				"SchemaTableColumn":   `"s1"."TAG"."NAME"`,
				"SchemaTableAs":       `"t1"`,
				"SchemaTableAsColumn": `"t1"."NAME"`,
			}
			fn(t, metaList, tableList, field, checkMap)
		}
	})
}
