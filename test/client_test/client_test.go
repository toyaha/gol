package client_test

import (
	"testing"
	"time"

	"github.com/toyaha/gol/test"
)

var (
	timeNow = time.Now()
)

func TestClient_Exec(t *testing.T) {
	t.Run("mssql insert", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql insert multiple lines", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql insertSelectUnion", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.InsertSelectUnion()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql insertSelectUnion multiple lines", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.InsertSelectUnion()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql update", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetSet(&table.Str, "update")
		query.SetWhereIs(&table.Str, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql delete", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetWhereIs(&table.Str, "update")
		_, err = query.Delete()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql truncate", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		_, err = query.Truncate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql insert", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql insert multiple lines", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql insertSelectUnion", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.InsertSelectUnion()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql insertSelectUnion multiple lines", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.InsertSelectUnion()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql update", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetSet(&table.Str, "update")
		query.SetWhereIs(&table.Str, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql delete", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetWhereIs(&table.Str, "update")
		_, err = query.Delete()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql truncate", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		_, err = query.Truncate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql insert", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql insert multiple lines", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql insertSelectUnion", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.InsertSelectUnion()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql insertSelectUnion multiple lines", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetValues(
			table.Str,
		)
		query.SetValues(
			table.Str,
		)
		_, err = query.InsertSelectUnion()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql update", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetSet(&table.Str, "update")
		query.SetWhereIs(&table.Str, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql delete", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		query.SetWhereIs(&table.Str, "update")
		_, err = query.Delete()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql truncate", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		_, err = query.Truncate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql truncate restart identity", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQuery(&table)
		_, err = query.TruncateRestartIdentity()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})
}

func TestClient_Meta(t *testing.T) {
	tableItem := test.Item{}
	tableTag := test.Tag{}

	fn := func(t *testing.T, table interface{}, field interface{}, checkMap map[string]string) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		db.AddMeta(table)

		{
			target := db.GetBaseSchema(field)
			check, ok := checkMap["BaseSchema"]
			if !ok {
				t.Error("BaseSchema is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetBaseTable(field)
			check, ok := checkMap["BaseTable"]
			if !ok {
				t.Error("BaseTable is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetBaseAs(field)
			check, ok := checkMap["BaseAs"]
			if !ok {
				t.Error("BaseAs is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetBaseColumn(field)
			check, ok := checkMap["BaseColumn"]
			if !ok {
				t.Error("BaseColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchema(field)
			check, ok := checkMap["Schema"]
			if !ok {
				t.Error("Schema is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTable(field)
			check, ok := checkMap["Table"]
			if !ok {
				t.Error("Table is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetAs(field)
			check, ok := checkMap["As"]
			if !ok {
				t.Error("As is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetColumn(field)
			check, ok := checkMap["Column"]
			if !ok {
				t.Error("Column is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTableColumn(field)
			check, ok := checkMap["TableColumn"]
			if !ok {
				t.Error("TableColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTableAs(field)
			check, ok := checkMap["TableAs"]
			if !ok {
				t.Error("TableAs is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTableAsColumn(field)
			check, ok := checkMap["TableAsColumn"]
			if !ok {
				t.Error("TableAsColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTable(field)
			check, ok := checkMap["SchemaTable"]
			if !ok {
				t.Error("SchemaTable is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTableColumn(field)
			check, ok := checkMap["SchemaTableColumn"]
			if !ok {
				t.Error("SchemaTableColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTableAs(field)
			check, ok := checkMap["SchemaTableAs"]
			if !ok {
				t.Error("SchemaTableAs is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTableAsColumn(field)
			check, ok := checkMap["SchemaTableAsColumn"]
			if !ok {
				t.Error("SchemaTableAsColumn is not exist")
			}
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	}

	t.Run("item", func(t *testing.T) {
		var table interface{} = &tableItem
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
		fn(t, table, field, checkMap)
	})

	t.Run("tag", func(t *testing.T) {
		var table interface{} = &tableTag
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
		fn(t, table, field, checkMap)
	})
}
