package gol

import (
	"reflect"
	"testing"
)

func TestQuery_SelectResultRows(t *testing.T) {
	// tests := []struct {
	// 	name    string
	// 	rec     *Query
	// 	want    *sql.Rows
	// 	wantErr bool
	// }{
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := tt.rec.SelectResultRows()
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("Query.SelectResultRows() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("Query.SelectResultRows() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestQuery_SelectRow(t *testing.T) {
	// type args struct {
	// 	dest interface{}
	// }
	// tests := []struct {
	// 	name    string
	// 	rec     *Query
	// 	args    args
	// 	wantErr bool
	// }{
	// 	{
	// 	},
	// }
	// for _, tt := range tests {
	// 	tt.rec = NewQuery(nil)
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if err := tt.rec.SelectRow(tt.args.dest); (err != nil) != tt.wantErr {
	// 			t.Errorf("Query.SelectRow() error = %v, wantErr %v", err, tt.wantErr)
	// 		}
	// 	})
	// }
}

func TestQuery_SelectResultRow(t *testing.T) {
	// tests := []struct {
	// 	name    string
	// 	rec     *Query
	// 	want    *Row
	// 	wantErr bool
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		got, err := tt.rec.SelectResultRow()
	// 		if (err != nil) != tt.wantErr {
	// 			t.Errorf("Query.SelectResultRow() error = %v, wantErr %v", err, tt.wantErr)
	// 			return
	// 		}
	// 		if !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("Query.SelectResultRow() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func TestQuery_SetJoin(t *testing.T) {
	table := testItem{}
	type args struct {
		tablePtr  interface{}
		valueList []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"default",
			args{
				tablePtr:  &table,
				valueList: []interface{}{},
			},
		},
		{
			"sub query",
			args{
				tablePtr:  &table,
				valueList: []interface{}{"(select 1)"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Query{
				Client: &Client{},
				Config: NewConfig(),
				Value:  NewQueryValue(nil),
			}
			rec.SetJoin(tt.args.tablePtr, tt.args.valueList...)
			if len(rec.Value.JoinList) != 1 {
				t.Errorf("Query.SetJoin() length not 1")
			} else {
				if rec.Value.JoinList[0].TablePtr != tt.args.tablePtr {
					t.Errorf("Query.SetJoin() TablePtr not match")
				}
				if !reflect.DeepEqual(rec.Value.JoinList[0].ValueList, tt.args.valueList) {
					t.Errorf("Query.SetJoin() value=%+v , want%+v", rec.Value.JoinList[0].ValueList, tt.args.valueList)
				}
			}
		})
	}
}

func TestQuery_SetJoinLeft(t *testing.T) {
	table := testItem{}
	type args struct {
		tablePtr  interface{}
		valueList []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"default",
			args{
				tablePtr:  &table,
				valueList: []interface{}{},
			},
		},
		{
			"sub query",
			args{
				tablePtr:  &table,
				valueList: []interface{}{"(select 1)"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Query{
				Client: &Client{},
				Config: NewConfig(),
				Value:  NewQueryValue(nil),
			}
			rec.SetJoinLeft(tt.args.tablePtr, tt.args.valueList...)
			if len(rec.Value.JoinList) != 1 {
				t.Errorf("Query.SetJoinLeft() length not 1")
			} else {
				if rec.Value.JoinList[0].TablePtr != tt.args.tablePtr {
					t.Errorf("Query.SetJoinLeft() TablePtr not match")
				}
				if !reflect.DeepEqual(rec.Value.JoinList[0].ValueList, tt.args.valueList) {
					t.Errorf("Query.SetJoinLeft() value=%+v , want%+v", rec.Value.JoinList[0].ValueList, tt.args.valueList)
				}
			}
		})
	}
}

func TestQuery_SetJoinRight(t *testing.T) {
	table := testItem{}
	type args struct {
		tablePtr  interface{}
		valueList []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"default",
			args{
				tablePtr:  &table,
				valueList: []interface{}{},
			},
		},
		{
			"sub query",
			args{
				tablePtr:  &table,
				valueList: []interface{}{"(select 1)"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Query{
				Client: &Client{},
				Config: NewConfig(),
				Value:  NewQueryValue(nil),
			}
			rec.SetJoinRight(tt.args.tablePtr, tt.args.valueList...)
			if len(rec.Value.JoinList) != 1 {
				t.Errorf("Query.SetJoinRight() length not 1")
			} else {
				if rec.Value.JoinList[0].TablePtr != tt.args.tablePtr {
					t.Errorf("Query.SetJoinRight() TablePtr not match")
				}
				if !reflect.DeepEqual(rec.Value.JoinList[0].ValueList, tt.args.valueList) {
					t.Errorf("Query.SetJoinRight() value=%+v , want%+v", rec.Value.JoinList[0].ValueList, tt.args.valueList)
				}
			}
		})
	}
}

func TestQuery_SetOrderBy(t *testing.T) {
	table := testItem{}
	type args struct {
		valueList []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "default",
			args: args{
				valueList: []interface{}{&table.Id},
			},
		},
		{
			name: "query",
			args: args{
				valueList: []interface{}{"count(", &table.Id, ") desc"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &Query{
				Client: &Client{},
				Config: NewConfig(),
				Value:  NewQueryValue(nil),
			}
			rec.SetOrderBy(tt.args.valueList...)
			if len(rec.Value.OrderByList) != 1 {
				t.Errorf("Query.SetOrderBy() length not 1")
			} else {
				if !reflect.DeepEqual(rec.Value.OrderByList[0].ValueList, tt.args.valueList) {
					t.Errorf("Query.SetOrderBy() value=%+v , want%+v", rec.Value.OrderByList[0].ValueList, tt.args.valueList)
				}
			}
		})
	}
}
