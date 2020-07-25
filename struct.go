package gol

import (
	"database/sql"
	"errors"
)

func NewRow() *Row {
	return &Row{}
}

type Row struct {
	Row *sql.Row
}

func (rec *Row) Scan(dest ...interface{}) error {
	if rec.Row == nil {
		return errors.New("row not found")
	}

	return rec.Row.Scan(dest...)
}
