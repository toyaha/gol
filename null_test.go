package gol

import (
	"database/sql"
	"testing"
	"time"
)

func TestNullBool_IsNull(t *testing.T) {
	type fields struct {
		NullBool sql.NullBool
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ok want true",
			fields: fields{
				NullBool: sql.NullBool{
					Bool:  false,
					Valid: false,
				},
			},
			want: true,
		},
		{
			name: "ok want false",
			fields: fields{
				NullBool: sql.NullBool{
					Bool:  true,
					Valid: true,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &NullBool{
				NullBool: tt.fields.NullBool,
			}
			if got := rec.IsNull(); got != tt.want {
				t.Errorf("NullBool.IsNull()\ngot  = %v\nwant = %v", got, tt.want)
			}
		})
	}
}

func TestNullFloat64_IsNull(t *testing.T) {
	type fields struct {
		NullFloat64 sql.NullFloat64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ok want true",
			fields: fields{
				NullFloat64: sql.NullFloat64{
					Float64: 0,
					Valid:   false,
				},
			},
			want: true,
		},
		{
			name: "ok want false",
			fields: fields{
				NullFloat64: sql.NullFloat64{
					Float64: 1.123,
					Valid:   true,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &NullFloat64{
				NullFloat64: tt.fields.NullFloat64,
			}
			if got := rec.IsNull(); got != tt.want {
				t.Errorf("NullFloat64.IsNull()\ngot  = %v\nwant = %v", got, tt.want)
			}
		})
	}
}

func TestNullInt32_IsNull(t *testing.T) {
	type fields struct {
		NullInt32 sql.NullInt32
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ok want true",
			fields: fields{
				NullInt32: sql.NullInt32{
					Int32: 0,
					Valid: false,
				},
			},
			want: true,
		},
		{
			name: "ok want false",
			fields: fields{
				NullInt32: sql.NullInt32{
					Int32: 1,
					Valid: true,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &NullInt32{
				NullInt32: tt.fields.NullInt32,
			}
			if got := rec.IsNull(); got != tt.want {
				t.Errorf("NullInt32.IsNull()\ngot  = %v\nwant = %v", got, tt.want)
			}
		})
	}
}

func TestNullInt64_IsNull(t *testing.T) {
	type fields struct {
		NullInt64 sql.NullInt64
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ok want true",
			fields: fields{
				NullInt64: sql.NullInt64{
					Int64: 0,
					Valid: false,
				},
			},
			want: true,
		},
		{
			name: "ok want false",
			fields: fields{
				NullInt64: sql.NullInt64{
					Int64: 1,
					Valid: true,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &NullInt64{
				NullInt64: tt.fields.NullInt64,
			}
			if got := rec.IsNull(); got != tt.want {
				t.Errorf("NullInt64.IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullString_IsNull(t *testing.T) {
	type fields struct {
		NullString sql.NullString
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ok want true",
			fields: fields{
				NullString: sql.NullString{
					String: "",
					Valid:  false,
				},
			},
			want: true,
		},
		{
			name: "ok want false",
			fields: fields{
				NullString: sql.NullString{
					String: "a",
					Valid:  true,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &NullString{
				NullString: tt.fields.NullString,
			}
			if got := rec.IsNull(); got != tt.want {
				t.Errorf("NullString.IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNullTime_IsNull(t *testing.T) {
	type fields struct {
		NullTime sql.NullTime
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name: "ok want true",
			fields: fields{
				NullTime: sql.NullTime{
					Time:  time.Time{},
					Valid: false,
				},
			},
			want: true,
		},
		{
			name: "ok want false",
			fields: fields{
				NullTime: sql.NullTime{
					Time:  time.Time{},
					Valid: true,
				},
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := &NullTime{
				NullTime: tt.fields.NullTime,
			}
			if got := rec.IsNull(); got != tt.want {
				t.Errorf("NullTime.IsNull() = %v, want %v", got, tt.want)
			}
		})
	}
}
