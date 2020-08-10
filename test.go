package gol

import (
	"errors"
	"time"
)

const (
	testDbMssqlHost     = "localhost"
	testDbMssqlPort     = "1433"
	testDbMssqlUser     = "user"
	testDbMssqlPass     = "Password!"
	testDbMssqlDatabase = "db"

	testDbMysqlHost     = "localhost"
	testDbMysqlPort     = "3306"
	testDbMysqlUser     = "user"
	testDbMysqlPass     = "pass"
	testDbMysqlDatabase = "db"

	testDbPostgresqlHost     = "localhost"
	testDbPostgresqlPort     = "5432"
	testDbPostgresqlUser     = "user"
	testDbPostgresqlPass     = "pass"
	testDbPostgresqlDatabase = "db"

	testSchemaPublic     = "public"
	testSchemaTest       = "test"
	testSchemaMssql      = "dbo"
	testSchemaPostgresql = "public"
)

var (
	testTableItem1       = testItem{}
	testTableItem2       = testItem{}
	testTableItem3       = testItem{}
	testTableItem4       = testItem{}
	testTableItem5       = testItem{}
	testTableItemDetail1 = testItemDetail{}
	testTableItemDetail2 = testItemDetail{}
	testTableItemDetail3 = testItemDetail{}
	testTableItemDetail4 = testItemDetail{}
	testTableItemDetail5 = testItemDetail{}
	testTableTag1        = testTag{}
	testTableTag2        = testTag{}
	testTableTag3        = testTag{}
	testTableTag4        = testTag{}
	testTableTag5        = testTag{}
)

type testItem struct {
	Id       int       `table:"item" column:"id" json:"id"`
	CreateAt time.Time `table:"item" column:"create_at" json:"create_at"`
	UpdateAt time.Time `table:"item" column:"update_at" json:"update_at"`
	DeleteAt NullTime  `table:"item" column:"delete_at" json:"delete_at"`
	Num      int       `table:"item" column:"num" json:"num"`
	Str      string    `table:"item" column:"str" json:"str"`
}

type testItemDetail struct {
	Id       int       `table:"item" column:"id" json:"id"`
	CreateAt time.Time `table:"item" column:"create_at" json:"create_at"`
	UpdateAt time.Time `table:"item" column:"update_at" json:"update_at"`
	DeleteAt NullTime  `table:"item" column:"delete_at" json:"delete_at"`
	ItemId   NullInt64 `table:"item" column:"item_id" json:"item_id"`
	Num      int       `table:"item" column:"num" json:"num"`
	Str      string    `table:"item" column:"str" json:"str"`
}

type testTag struct {
	Id       int       `schema:"PUBLIC" table:"TAG" column:"ID" json:"id"`
	CreateAt time.Time `schema:"PUBLIC" table:"TAG" column:"CREATE_AT" json:"create_at"`
	UpdateAt time.Time `schema:"PUBLIC" table:"TAG" column:"UPDATE_AT" json:"update_at"`
	DeleteAt NullTime  `schema:"PUBLIC" table:"TAG" column:"DELETE_AT" json:"delete_at"`
	Num      string    `schema:"PUBLIC" table:"TAG" column:"NUM" json:"NUM"`
	Str      string    `schema:"PUBLIC" table:"TAG" column:"STR" json:"STR"`
}

func testNewClient(databaseType string) (*Client, error) {
	switch databaseType {
	case DatabaseTypeMssql:
		return testNewClientMssql()
	case DatabaseTypeMysql:
		return testNewClientMysql()
	case DatabaseTypePostgresql:
		return testNewClientPostgresql()
	}
	return nil, errors.New("none database")
}

func testNewClientMssql() (*Client, error) {
	db, err := NewClient(
		DatabaseTypeMssql,
		testDbMssqlHost,
		testDbMssqlPort,
		testDbMssqlUser,
		testDbMssqlPass,
		testDbMssqlDatabase,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return db, err
}

func testNewClientMysql() (*Client, error) {
	db, err := NewClient(
		DatabaseTypeMysql,
		testDbMysqlHost,
		testDbMysqlPort,
		testDbMysqlUser,
		testDbMysqlPass,
		testDbMysqlDatabase,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return db, err
}

func testNewClientPostgresql() (*Client, error) {
	db, err := NewClient(
		DatabaseTypePostgresql,
		testDbPostgresqlHost,
		testDbPostgresqlPort,
		testDbPostgresqlUser,
		testDbPostgresqlPass,
		testDbPostgresqlDatabase,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return db, err
}
