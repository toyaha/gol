package client_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/toyaha/gol"
	"github.com/toyaha/gol/test"
)

var (
	timeNow = time.Now()
)

func TestClient_Query(t *testing.T) {
	fn := func(db *gol.Client, checkList []string) {
		rows, err := db.Query(`SELECT 1 AS id UNION SELECT 2`)
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = rows.Close()
		}()

		key := 0
		for rows.Next() {
			var row int
			err := rows.Scan(&row)
			if err != nil {
				t.Errorf("\nerror: %v", err)
				return
			}

			target := fmt.Sprintf("%+v", row)
			check := checkList[key]
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
				return
			}

			key++
		}
	}

	t.Run("mssql", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var checkList = []string{"1", "2"}
		fn(db, checkList)
	})

	t.Run("mysql", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var checkList = []string{"1", "2"}
		fn(db, checkList)
	})

	t.Run("postgresql", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var checkList = []string{"1", "2"}
		fn(db, checkList)
	})
}
func TestClient_QueryRow(t *testing.T) {
	t.Run("mssql", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var str string
		err = db.QueryRow(`SELECT 'abcde'`).Scan(&str)
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		target := "abcde"
		check := str
		if target != check {
			t.Errorf("\ntarget: %v\ncheck : %v", target, check)
		}
	})

	t.Run("mysql", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var str string
		err = db.QueryRow(`SELECT 'abcde'`).Scan(&str)
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		target := "abcde"
		check := str
		if target != check {
			t.Errorf("\ntarget: %v\ncheck : %v", target, check)
		}
	})

	t.Run("postgresql", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var str string
		err = db.QueryRow(`SELECT 'abcde'`).Scan(&str)
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		target := "abcde"
		check := str
		if target != check {
			t.Errorf("\ntarget: %v\ncheck : %v", target, check)
		}
	})
}

func TestClient_Find(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList int
		err = db.Find(&resultList, "SELECT 1 AS id UNION SELECT 2")
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*int", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList int
		err = db.Find(&resultList, "SELECT 1 AS id UNION SELECT 2")
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList test.Item
		err = db.Find(&resultList, "SELECT 1 AS id UNION SELECT 2")
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList test.Item
		err = db.Find(&resultList, "SELECT 1 AS id UNION SELECT 2")
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*[]struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList []test.Item
		err = db.Find(&resultList, "SELECT 1 AS id UNION SELECT 2")
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		checkList := []string{"1", "2"}
		for key, val := range resultList {
			target := fmt.Sprintf("%v", val.Id)
			check := checkList[key]
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("*[]*struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList []*test.Item
		err = db.Find(&resultList, "SELECT 1 AS id UNION SELECT 2")
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		checkList := []string{"1", "2"}
		for key, val := range resultList {
			target := fmt.Sprintf("%v", val.Id)
			check := checkList[key]
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("*[]map[string]interface{}", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList = make([]map[string]interface{}, 0)
		err = db.Find(&resultList, "SELECT 1 AS id UNION SELECT 2")
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		checkList := []string{"1", "2"}
		for key, val := range resultList {
			target := fmt.Sprintf("%v", val["id"])
			check := checkList[key]
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})
}

func TestClient_FindRow(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var result = &test.Item{}
		err = db.FindRow(result, "SELECT 1 AS id")
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		target := fmt.Sprintf("%v", result.Id)
		check := fmt.Sprintf("%v", 1)
		if target != check {
			t.Errorf("\ntarget: %v\ncheck : %v", target, check)
		}
	})
}

func TestClient_ExtractRows(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList int
		rows, err := db.Query("SELECT 1 AS id UNION SELECT 2")
		err = db.ExtractRows(resultList, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*int", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList int
		rows, err := db.Query("SELECT 1 AS id UNION SELECT 2")
		err = db.ExtractRows(&resultList, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList test.Item
		rows, err := db.Query("SELECT 1 AS id UNION SELECT 2")
		err = db.ExtractRows(resultList, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList test.Item
		rows, err := db.Query("SELECT 1 AS id UNION SELECT 2")
		err = db.ExtractRows(&resultList, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*[]struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList []test.Item
		rows, err := db.Query("SELECT 1 AS id UNION SELECT 2")
		err = db.ExtractRows(&resultList, rows)
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		checkList := []string{"1", "2"}
		for key, val := range resultList {
			target := fmt.Sprintf("%v", val.Id)
			check := checkList[key]
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("*[]*struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList []*test.Item
		rows, err := db.Query("SELECT 1 AS id UNION SELECT 2")
		err = db.ExtractRows(&resultList, rows)
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		checkList := []string{"1", "2"}
		for key, val := range resultList {
			target := fmt.Sprintf("%v", val.Id)
			check := checkList[key]
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})

	t.Run("*[]map[string]interface{}", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList = make([]map[string]interface{}, 0)
		rows, err := db.Query("SELECT 1 AS id UNION SELECT 2")
		err = db.ExtractRows(&resultList, rows)
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		checkList := []string{"1", "2"}
		for key, val := range resultList {
			target := fmt.Sprintf("%v", val["id"])
			check := checkList[key]
			if target != check {
				t.Errorf("\ntarget: %v\ncheck : %v", target, check)
			}
		}
	})
}

func TestClient_ExtractRow(t *testing.T) {
	t.Run("int", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList int
		rows, err := db.Query("SELECT 1 AS id")
		err = db.ExtractRow(resultList, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*int", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var result int
		rows, err := db.Query("SELECT 1 AS id")
		err = db.ExtractRow(&result, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var result test.Item
		rows, err := db.Query("SELECT 1 AS id")
		err = db.ExtractRow(result, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var result test.Item
		rows, err := db.Query("SELECT 1 AS id")
		err = db.ExtractRow(&result, rows)
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}

		target := fmt.Sprintf("%v", result.Id)
		check := fmt.Sprintf("%v", 1)
		if target != check {
			t.Errorf("\ntarget: %v\ncheck : %v", target, check)
		}
	})

	t.Run("*[]struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList []test.Item
		rows, err := db.Query("SELECT 1 AS id")
		err = db.ExtractRow(&resultList, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*[]*struct", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList []*test.Item
		rows, err := db.Query("SELECT 1 AS id")
		err = db.ExtractRow(&resultList, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("*[]map[string]interface{}", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var resultList = make([]map[string]interface{}, 0)
		rows, err := db.Query("SELECT 1 AS id")
		err = db.ExtractRow(&resultList, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})

	t.Run("2line", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		var result = &test.Item{}
		rows, err := db.Query("SELECT 1 AS id UNION SELECT 2")
		err = db.ExtractRow(result, rows)
		if err == nil {
			t.Errorf("\nerror not found")
			return
		}
	})
}

func TestClient_Insert(t *testing.T) {
	t.Run("Insert mssql", func(t *testing.T) {
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

	t.Run("Insert mssql multiple lines", func(t *testing.T) {
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

	t.Run("Insert mysql", func(t *testing.T) {
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

	t.Run("Insert mysql multiple lines", func(t *testing.T) {
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

	t.Run("Insert postgresql", func(t *testing.T) {
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

	t.Run("Insert postgresql multiple lines", func(t *testing.T) {
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
}

func TestClient_InsertDoNothing(t *testing.T) {
	t.Run("InsertDoNothing postgresql", func(t *testing.T) {
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
		_, err = query.InsertDoNothing()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("InsertDoNothing postgresql multiple lines", func(t *testing.T) {
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
		_, err = query.InsertDoNothing()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})
}

func TestClient_InsertDoUpdate(t *testing.T) {
	t.Run("InsertDoUpdate postgresql", func(t *testing.T) {
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
			&table.Id,
			&table.Str,
		)
		query.SetValues(
			1,
			table.Str,
		)
		query.SetConflict(&table.Id)
		query.SetSet(&table.Str, "conflict")
		_, err = query.InsertDoUpdate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("InsertDoUpdate postgresql multiple lines", func(t *testing.T) {
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
			&table.Id,
			&table.Str,
		)
		query.SetValues(
			1,
			table.Str,
		)
		query.SetValues(
			2,
			table.Str,
		)
		query.SetConflict(&table.Id)
		query.SetSet(&table.Str, "conflict")
		_, err = query.InsertDoUpdate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})
}

func TestClient_InsertIgnore(t *testing.T) {
	t.Run("InsertIgnore mysql", func(t *testing.T) {
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
		_, err = query.InsertIgnore()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("InsertIgnore mysql multiple lines", func(t *testing.T) {
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
		_, err = query.InsertIgnore()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})
}

func TestClient_InsertOnDuplicateKeyUpdate(t *testing.T) {
	t.Run("InsertOnDuplicateKeyUpdate mysql", func(t *testing.T) {
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
			&table.Id,
			&table.Str,
		)
		query.SetValues(
			1,
			table.Str,
		)
		query.SetSet(&table.Str, "duplicate")
		_, err = query.InsertOnDuplicateKeyUpdate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("InsertOnDuplicateKeyUpdate mysql multiple lines", func(t *testing.T) {
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
			&table.Id,
			&table.Str,
		)
		query.SetValues(
			1,
			table.Str,
		)
		query.SetValues(
			2,
			table.Str,
		)
		query.SetSet(&table.Str, "duplicate")
		_, err = query.InsertOnDuplicateKeyUpdate()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})
}

func TestClient_InsertSelect(t *testing.T) {
	t.Run("InsertSelect mssql", func(t *testing.T) {
		db, err := test.NewClientMssql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		table2 := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetFrom(&table2)
		query.SetSelect(&table2.Str)
		_, err = query.InsertSelect()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("InsertSelect mysql", func(t *testing.T) {
		db, err := test.NewClientMysql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		table2 := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetFrom(&table2)
		query.SetSelect(&table2.Str)
		_, err = query.InsertSelect()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("InsertSelect postgresql", func(t *testing.T) {
		db, err := test.NewClientPostgresql()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
		defer func() {
			_ = db.Close()
		}()

		table := test.Item{}
		table2 := test.Item{}
		query := db.NewQuery(&table)
		query.SetValuesColumn(
			&table.Str,
		)
		query.SetFrom(&table2)
		query.SetSelect(&table2.Str)
		_, err = query.InsertSelect()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})
}

func TestClient_InsertSelectUnion(t *testing.T) {
	t.Run("InsertSelectUnion mssql", func(t *testing.T) {
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

	t.Run("InsertSelectUnion mssql multiple lines", func(t *testing.T) {
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

	t.Run("InsertSelectUnion mysql", func(t *testing.T) {
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

	t.Run("InsertSelectUnion mysql multiple lines", func(t *testing.T) {
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

	t.Run("InsertSelectUnion postgresql", func(t *testing.T) {
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

	t.Run("InsertSelectUnion postgresql multiple lines", func(t *testing.T) {
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
}

func TestClient_Update(t *testing.T) {
	t.Run("Update mssql", func(t *testing.T) {
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
		query.SetSet(&table.Num, &table.Id, " + ?", []interface{}{1})
		query.SetSet(&table.Str, "update")
		query.SetWhereIs(&table.Str, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("Update mysql", func(t *testing.T) {
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
		query.SetSet(&table.Num, &table.Id, " + ?", []interface{}{1})
		query.SetSet(&table.Str, "update")
		query.SetWhereIs(&table.Str, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})

	t.Run("Update postgresql", func(t *testing.T) {
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
		query.SetSet(&table.Num, &table.Id, " + ?", []interface{}{1})
		query.SetSet(&table.Str, "update")
		query.SetWhereIs(&table.Str, "")
		_, err = query.Update()
		if err != nil {
			t.Errorf("\nerror: %v", err)
			return
		}
	})
}

func TestClient_Delete(t *testing.T) {
	t.Run("Delete mssql", func(t *testing.T) {
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

	t.Run("Delete mysql", func(t *testing.T) {
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

	t.Run("Delete postgresql", func(t *testing.T) {
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
}

func TestClient_Truncate(t *testing.T) {
	t.Run("Truncate mssql", func(t *testing.T) {
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

	t.Run("Truncate mysql", func(t *testing.T) {
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

	t.Run("Truncate postgresql", func(t *testing.T) {
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
}

func TestClient_TruncateRestartIdentity(t *testing.T) {
	t.Run("TruncateRestartIdentity postgresql", func(t *testing.T) {
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

	t.Run("Meta item", func(t *testing.T) {
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

	t.Run("Meta tag", func(t *testing.T) {
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
