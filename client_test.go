package gol

import (
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
