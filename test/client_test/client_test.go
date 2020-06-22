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
	t.Run("mssql table insert", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryWithSchema(test.SchemaMssql, &table)
		query.SetValuesColumn(
			&table.Name,
		)
		query.SetValues(
			table.Name,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql table update", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryWithSchema(test.SchemaMssql, &table)
		query.SetSet(&table.Name, "update")
		query.SetWhereIs(&table.Name, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql table delete", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryWithSchema(test.SchemaMssql, &table)
		query.SetWhereIs(&table.Name, "update")
		_, err = query.Delete()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql table truncate", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryWithSchema(test.SchemaMssql, &table)
		_, err = query.Truncate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql table as insert", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAsWithSchema(test.SchemaMssql, &table, "test")
		query.SetValuesColumn(
			&table.Name,
		)
		query.SetValues(
			table.Name,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql table as update", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAsWithSchema(test.SchemaMssql, &table, "alias")
		query.SetSet(&table.Name, "update")
		query.SetWhereIs(&table.Name, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql table as delete", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAsWithSchema(test.SchemaMssql, &table, "alias")
		query.SetWhereIs(&table.Name, "update")
		_, err = query.Delete()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mssql table as truncate", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAsWithSchema(test.SchemaMssql, &table, "alias")
		_, err = query.Truncate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql table insert", func(t *testing.T) {
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
			&table.Name,
		)
		query.SetValues(
			table.Name,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql table update", func(t *testing.T) {
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
		query.SetSet(&table.Name, "update")
		query.SetWhereIs(&table.Name, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql table delete", func(t *testing.T) {
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
		query.SetWhereIs(&table.Name, "update")
		_, err = query.Delete()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql table truncate", func(t *testing.T) {
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

	t.Run("mysql table as insert", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAs(&table, "alias")
		query.SetValuesColumn(
			&table.Name,
		)
		query.SetValues(
			table.Name,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql table as update", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAs(&table, "alias")
		query.SetSet(&table.Name, "update")
		query.SetWhereIs(&table.Name, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql table as delete", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAs(&table, "alias")
		query.SetWhereIs(&table.Name, "update")
		_, err = query.Delete()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("mysql table as truncate", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAs(&table, "alias")
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
		query := db.NewQueryWithSchema(test.SchemaPostgresql, &table)
		query.SetValuesColumn(
			&table.Name,
		)
		query.SetValues(
			table.Name,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql table update", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryWithSchema(test.SchemaPostgresql, &table)
		query.SetSet(&table.Name, "update")
		query.SetWhereIs(&table.Name, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql table delete", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryWithSchema(test.SchemaPostgresql, &table)
		query.SetWhereIs(&table.Name, "update")
		_, err = query.Delete()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql table truncate", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryWithSchema(test.SchemaPostgresql, &table)
		_, err = query.Truncate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql table truncate restart identity", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryWithSchema(test.SchemaPostgresql, &table)
		_, err = query.TruncateRestartIdentity()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql table as insert", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAsWithSchema(test.SchemaPostgresql, &table, "alias")
		query.SetValuesColumn(
			&table.Name,
		)
		query.SetValues(
			table.Name,
		)
		_, err = query.Insert()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql table as update", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAsWithSchema(test.SchemaPostgresql, &table, "alias")
		query.SetSet(&table.Name, "update")
		query.SetWhereIs(&table.Name, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql table as delete", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAsWithSchema(test.SchemaPostgresql, &table, "alias")
		query.SetWhereIs(&table.Name, "update")
		_, err = query.Delete()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql table as truncate", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAsWithSchema(test.SchemaPostgresql, &table, "alias")
		_, err = query.Truncate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("postgresql table as truncate restart identity", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		query := db.NewQueryAsWithSchema(test.SchemaPostgresql, &table, "alias")
		_, err = query.TruncateRestartIdentity()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})
}

func TestClient_Meta(t *testing.T) {
	t.Run("meta item", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		db.AddMeta(&table)

		{
			target := db.GetBaseSchema(&table.Id)
			check := ``
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetBaseTable(&table.Id)
			check := `item`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetBaseAs(&table.Id)
			check := ``
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetBaseColumn(&table.Id)
			check := `id`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchema(&table.Id)
			check := ``
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTable(&table.Id)
			check := `"item"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetAs(&table.Id)
			check := ``
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetColumn(&table.Id)
			check := `"id"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTableAs(&table.Id)
			check := `"item"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTable(&table.Id)
			check := `"item"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTableColumn(&table.Id)
			check := `"item"."id"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTableAs(&table.Id)
			check := `"item"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTableAsColumn(&table.Id)
			check := `"item"."id"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTableAs(&table.Id)
			check := `"item"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTableAsColumn(&table.Id)
			check := `"item"."id"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("meta tag", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Tag{}
		db.AddMeta(&table)

		{
			target := db.GetBaseSchema(&table.Id)
			check := `PUBLIC`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetBaseTable(&table.Id)
			check := `TAG`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetBaseAs(&table.Id)
			check := ``
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetBaseColumn(&table.Id)
			check := `ID`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchema(&table.Id)
			check := `"PUBLIC"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTable(&table.Id)
			check := `"TAG"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetAs(&table.Id)
			check := ``
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetColumn(&table.Id)
			check := `"ID"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTableAs(&table.Id)
			check := `"TAG"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTable(&table.Id)
			check := `"PUBLIC"."TAG"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTableColumn(&table.Id)
			check := `"PUBLIC"."TAG"."ID"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTableAs(&table.Id)
			check := `"PUBLIC"."TAG"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetSchemaTableAsColumn(&table.Id)
			check := `"PUBLIC"."TAG"."ID"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTableAs(&table.Id)
			check := `"TAG"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}

		{
			target := db.GetTableAsColumn(&table.Id)
			check := `"TAG"."ID"`
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})
}
