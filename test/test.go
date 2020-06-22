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
	TableItem1 = Item{}
	TableItem2 = Item{}
	TableItem3 = Item{}
	TableItem4 = Item{}
)

type Item struct {
	Id        int          `column:"id" json:"id"`
	CreatedAt time.Time    `column:"created_at" json:"created_at"`
	UpdatedAt time.Time    `column:"updated_at" json:"updated_at"`
	DeletedAt gol.NullTime `column:"deleted_at" json:"deleted_at"`
	Name      string       `column:"name" json:"name"`
}

type ItemMssql struct {
	Id        int          `schema:"dbo" table:"item" column:"id" json:"id"`
	CreatedAt time.Time    `schema:"dbo" table:"item" column:"created_at" json:"created_at"`
	UpdatedAt time.Time    `schema:"dbo" table:"item" column:"updated_at" json:"updated_at"`
	DeletedAt gol.NullTime `schema:"dbo" table:"item" column:"deleted_at" json:"deleted_at"`
	Name      string       `schema:"dbo" table:"item" column:"name" json:"name"`
}

type Tag struct {
	Id        int          `schema:"PUBLIC" table:"TAG" column:"ID" json:"id"`
	CreatedAt time.Time    `schema:"PUBLIC" table:"TAG" column:"CREATED_AT" json:"created_at"`
	UpdatedAt time.Time    `schema:"PUBLIC" table:"TAG" column:"UPDATED_AT" json:"updated_at"`
	DeletedAt gol.NullTime `schema:"PUBLIC" table:"TAG" column:"DELETED_AT" json:"deleted_at"`
	Name      string       `schema:"PUBLIC" table:"TAG" column:"NAME" json:"name"`
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
