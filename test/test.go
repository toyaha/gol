package test

import (
	"time"

	"github.com/toyaha/gol"
)

const (
	DbMssqlHost     = "localhost"
	DbMssqlPort     = "1433"
	DbMssqlUser     = "user"
	DbMssqlPass     = "Password!"
	DbMssqlDatabase = "db"

	DbMysqlHost     = "localhost"
	DbMysqlPort     = "3306"
	DbMysqlUser     = "user"
	DbMysqlPass     = "pass"
	DbMysqlDatabase = "db"

	DbPostgresqlHost     = "localhost"
	DbPostgresqlPort     = "5432"
	DbPostgresqlUser     = "user"
	DbPostgresqlPass     = "pass"
	DbPostgresqlDatabase = "db"

	SchemaPublic     = "public"
	SchemaTest       = "test"
	SchemaMssql      = "dbo"
	SchemaPostgresql = "public"
)

var (
	TableItem1       = Item{}
	TableItem2       = Item{}
	TableItem3       = Item{}
	TableItem4       = Item{}
	TableItem5       = Item{}
	TableItemDetail1 = ItemDetail{}
	TableItemDetail2 = ItemDetail{}
	TableItemDetail3 = ItemDetail{}
	TableItemDetail4 = ItemDetail{}
	TableItemDetail5 = ItemDetail{}
	TableTag1        = Tag{}
	TableTag2        = Tag{}
	TableTag3        = Tag{}
	TableTag4        = Tag{}
	TableTag5        = Tag{}
)

type Item struct {
	Id        int          `json:"id"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt time.Time    `json:"updated_at"`
	DeletedAt gol.NullTime `json:"deleted_at"`
	Num       int          `json:"num"`
	Str       string       `json:"str"`
}

type ItemDetail struct {
	Id        int           `json:"id"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	DeletedAt gol.NullTime  `json:"deleted_at"`
	ItemId    gol.NullInt64 `json:"item_id"`
	Num       int           `json:"num"`
	Str       string        `json:"str"`
}

type Tag struct {
	Id        int          `schema:"PUBLIC" table:"TAG" column:"ID" json:"id"`
	CreatedAt time.Time    `schema:"PUBLIC" table:"TAG" column:"CREATED_AT" json:"created_at"`
	UpdatedAt time.Time    `schema:"PUBLIC" table:"TAG" column:"UPDATED_AT" json:"updated_at"`
	DeletedAt gol.NullTime `schema:"PUBLIC" table:"TAG" column:"DELETED_AT" json:"deleted_at"`
	Num       string       `schema:"PUBLIC" table:"TAG" column:"NUM" json:"NUM"`
	Str       string       `schema:"PUBLIC" table:"TAG" column:"STR" json:"STR"`
}

func NewClientMssql() (*gol.Client, error) {
	db, err := gol.NewClient(
		gol.DatabaseTypeMssql,
		DbMssqlHost,
		DbMssqlPort,
		DbMssqlUser,
		DbMssqlPass,
		DbMssqlDatabase,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return db, err
}

func NewClientMysql() (*gol.Client, error) {
	db, err := gol.NewClient(
		gol.DatabaseTypeMysql,
		DbMysqlHost,
		DbMysqlPort,
		DbMysqlUser,
		DbMysqlPass,
		DbMysqlDatabase,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return db, err
}

func NewClientPostgresql() (*gol.Client, error) {
	db, err := gol.NewClient(
		gol.DatabaseTypePostgresql,
		DbPostgresqlHost,
		DbPostgresqlPort,
		DbPostgresqlUser,
		DbPostgresqlPass,
		DbPostgresqlDatabase,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return db, err
}
