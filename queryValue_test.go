package gol

import (
	"reflect"
	"testing"
)

func TestQueryJoin_BuildUseAs(t *testing.T) {
	table := testItem{}
	var meta *Meta
	{
		query := NewQuery()
		query.SetJoin(&table)
		meta = query.Value.Meta
	}
	type fields struct {
		Mode      int
		TablePtr  interface{}
		ValueList []interface{}
	}
	type args struct {
		meta *Meta
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		want1   []interface{}
		wantErr bool
	}{
		{
			name: "inner",
			fields: fields{
				Mode:      QueryJoinModeInner,
				TablePtr:  &table,
				ValueList: nil,
			},
			args:    args{meta: meta},
			want:    `INNER JOIN "item" as "t0" ON`,
			want1:   nil,
			wantErr: false,
		},
		{
			name: "inner query",
			fields: fields{
				Mode:      QueryJoinModeInner,
				TablePtr:  &table,
				ValueList: []interface{}{"(select 1)", []interface{}{1}},
			},
			args:    args{meta: meta},
			want:    `INNER JOIN (select 1) as "t0" ON`,
			want1:   []interface{}{1},
			wantErr: false,
		},
		{
			name: "left",
			fields: fields{
				Mode:      QueryJoinModeLeft,
				TablePtr:  &table,
				ValueList: nil,
			},
			args:    args{meta: meta},
			want:    `LEFT JOIN "item" as "t0" ON`,
			want1:   nil,
			wantErr: false,
		},
		{
			name: "left query",
			fields: fields{
				Mode:      QueryJoinModeLeft,
				TablePtr:  &table,
				ValueList: []interface{}{"(select 1)", []interface{}{1}},
			},
			args:    args{meta: meta},
			want:    `LEFT JOIN (select 1) as "t0" ON`,
			want1:   []interface{}{1},
			wantErr: false,
		},
		{
			name: "right",
			fields: fields{
				Mode:      QueryJoinModeRight,
				TablePtr:  &table,
				ValueList: nil,
			},
			args:    args{meta: meta},
			want:    `RIGHT JOIN "item" as "t0" ON`,
			want1:   nil,
			wantErr: false,
		},
		{
			name: "right query",
			fields: fields{
				Mode:      QueryJoinModeRight,
				TablePtr:  &table,
				ValueList: []interface{}{"(select 1)", []interface{}{1}},
			},
			args:    args{meta: meta},
			want:    `RIGHT JOIN (select 1) as "t0" ON`,
			want1:   []interface{}{1},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &QueryJoin{
				Mode:      tt.fields.Mode,
				TablePtr:  tt.fields.TablePtr,
				ValueList: tt.fields.ValueList,
			}
			got, got1, err := rec.BuildUseAs(tt.args.meta)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryJoin.BuildUseAs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QueryJoin.BuildUseAs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("QueryJoin.BuildUseAs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestQueryOrderBy_Build(t *testing.T) {
	table := testItem{}
	var meta *Meta
	{
		query := NewQuery()
		query.SetJoin(&table)
		meta = query.Value.Meta
	}
	type fields struct {
		Mode      int
		ValueList []interface{}
	}
	type args struct {
		meta *Meta
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Mode:      QueryModeDefault,
				ValueList: []interface{}{&table.Id},
			},
			args:    args{meta: meta},
			want:    `"item"."id"`,
			wantErr: false,
		},
		{
			name: "query",
			fields: fields{
				Mode:      QueryModeDefault,
				ValueList: []interface{}{"count(", &table.Id, ") desc"},
			},
			args:    args{meta: meta},
			want:    `count("item"."id") desc`,
			wantErr: false,
		},
		{
			name: "error not valueList",
			fields: fields{
				Mode:      QueryModeDefault,
				ValueList: []interface{}{},
			},
			args:    args{meta: meta},
			want:    ``,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &QueryOrderBy{
				Mode:      tt.fields.Mode,
				ValueList: tt.fields.ValueList,
			}
			got, err := rec.Build(tt.args.meta)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryOrderBy.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QueryOrderBy.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryOrderBy_BuildUseAs(t *testing.T) {
	table := testItem{}
	var meta *Meta
	{
		query := NewQuery()
		query.SetJoin(&table)
		meta = query.Value.Meta
	}
	type fields struct {
		Mode      int
		ValueList []interface{}
	}
	type args struct {
		meta *Meta
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "default",
			fields: fields{
				Mode:      QueryModeDefault,
				ValueList: []interface{}{&table.Id},
			},
			args:    args{meta: meta},
			want:    `"t0"."id"`,
			wantErr: false,
		},
		{
			name: "query",
			fields: fields{
				Mode:      QueryModeDefault,
				ValueList: []interface{}{"count(", &table.Id, ") desc"},
			},
			args:    args{meta: meta},
			want:    `count("t0"."id") desc`,
			wantErr: false,
		},
		{
			name: "error not valueList",
			fields: fields{
				Mode:      QueryModeDefault,
				ValueList: []interface{}{},
			},
			args:    args{meta: meta},
			want:    ``,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &QueryOrderBy{
				Mode:      tt.fields.Mode,
				ValueList: tt.fields.ValueList,
			}
			got, err := rec.BuildUseAs(tt.args.meta)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryOrderBy.BuildUseAs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("QueryOrderBy.BuildUseAs() = %v, want %v", got, tt.want)
			}
		})
	}
}
