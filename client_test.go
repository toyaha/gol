package gol

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

func TestClient_Query_SelectRow(t *testing.T) {
	table := testItem{}
	tests := []struct {
		name     string
		rec      *Client
		setQuery func(*Query) *Query
		dest     interface{}
		wantErr  bool
	}{
		{
			"select min",
			nil,
			func(query *Query) *Query {
				query.SetSelect("min(", &table.Id, ") as min")
				return query
			},
			&struct {
				Min NullInt64 `column:"min" json:"min"`
			}{},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" for mssql", func(t *testing.T) {
			db, err := testNewClientMssql()
			if err != nil {
				t.Errorf("Query.SelectRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer func() {
				_ = db.Close()
			}()

			query := db.NewQuery(&table)
			query = tt.setQuery(query)
			err = query.SelectRow(tt.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query.SelectRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		t.Run(tt.name+" for mysql", func(t *testing.T) {
			db, err := testNewClientMysql()
			if err != nil {
				t.Errorf("Query.SelectRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer func() {
				_ = db.Close()
			}()

			query := db.NewQuery(&table)
			query = tt.setQuery(query)
			err = query.SelectRow(tt.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query.SelectRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		t.Run(tt.name+" for postgresql", func(t *testing.T) {
			db, err := testNewClientPostgresql()
			if err != nil {
				t.Errorf("Query.SelectRow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			defer func() {
				_ = db.Close()
			}()

			query := db.NewQuery(&table)
			query = tt.setQuery(query)
			err = query.SelectRow(tt.dest)
			if (err != nil) != tt.wantErr {
				t.Errorf("Query.SelectRow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestClient_Query_SelectResultRows(t *testing.T) {
	table := testItem{}
	tests := []struct {
		name     string
		rec      *Query
		setQuery func(*Query) *Query
		want     *sql.Rows
		wantErr  bool
	}{
		{
			"default",
			nil,
			func(query *Query) *Query {
				query.SetSelect(&table.Id)
				return query
			},
			&sql.Rows{},
			false,
		},
	}
	for _, tt := range tests {
		for _, databaseType := range []string{DatabaseTypeMssql, DatabaseTypeMysql, DatabaseTypePostgresql} {
			fmt.Printf("%+v\n", databaseType)
			t.Run(tt.name+" for "+databaseType, func(t *testing.T) {
				db, err := testNewClient(databaseType)
				if err != nil {
					t.Errorf("Query.SelectResultRows() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				defer func() {
					_ = db.Close()
				}()

				query := db.NewQuery(&table)
				query = tt.setQuery(query)
				rows, err := query.SelectResultRows()
				if fmt.Sprintf("%T", rows) != fmt.Sprintf("%T", tt.want) {
					t.Errorf("Query.SelectResultRows() error = %v, wantErr %v", err, tt.wantErr)
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("Query.SelectResultRows() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}
	}
}
