# gol
gol is a wrapper for database/sql rather than orm.

It is a tool to help you create simple queries.




# null structure
Use null structure for the column that allows null. This is a structure that has transferred the functionality of the sql.Null structure.


## The following methods are provided to change the value
- NullBool, NullInt32, NullInt64, NullFloat64, NullString
    - Set(value [bool | int | float64 | string])
    - Get() [bool | int | float64 | string]]
    - GetValue() [bool | int | float64 | string]
    - GetValueWithDefault(default [bool | int | float64 | string])
    - GetString() string
    - GetStringWithDefault(default [bool | int | float64 | string]) string
    - Delete()
- NullTime
    - Set(value time.Time)
    - Get() time.Time]
    - GetValue() time.Time
    - GetValueWithDefault(default time.Time)
    - GetString(format string) string
    - GetStringWithDefault(format string, default time.Time) string
    - Delete()




# transaction
Gol can be used for the following when using transactions.

``` go
func Sample() error {
    db, err := New()
    if err != nil {
        return err
    }
    defer func() {
        if p := recover(); p != nil {
            _ = db.Close()
            panic(p)
        }
        _ = db.Close()
    }()

    err = func() error {
        tx, err := db.Begin()
        if err != nil {
            return err
        }
        defer func() {
            if p := recover(); p != nil {
                _ = tx.Rollback()
                panic(p)
            }
            _ = tx.Rollback()
        }()

        // query...

        return tx.Commit()
    }()
    if err != nil {
        return err
    }

    return nil
}
```




# sample
It is a sample when writing each SQL with gol.

``` go
import (
  "github.com/toyaha/gol"
)

// item
type Item struct {
    Id  int    `column:"id" json:"id"`
    Num int    `column:"num" json:"num"`
    Str string `column:"str" json:"str"`
}

// item_detail
type Detail struct {
    Id     int    `schema:"public" table:"item_detail" column:"id" json:"id"`
    ItemId int    `schema:"public" table:"item_detail" column:"item_id" json:"item_id"`
    Num    int    `schema:"public" table:"item_detail" column:"num" json:"num"`
    Str    string `schema:"public" table:"item_detail" column:"str" json:"str"`
}

func New() (*gol.Client, error) {
    // database config
    // databaseType := gol.DatabaseTypeMssql
    // databaseType := gol.DatabaseTypeMysql
    databaseType := gol.DatabaseTypePostgresql
    host := "localhost"
    port := "5432"
    user := "username"
    pass := "password"
    database := "database"
    // postgresql options
    optionMap := map[string]string{
        "sslMode": gol.PostgresqlSslModeDisable,
        // "sslMode": gol.PostgresqlSslModeRequire,
        // "sslMode": gol.PostgresqlSslModeVerifyCa,
        // "sslMode": gol.PostgresqlSslModeVerifyFull,
    }

    db, err := gol.NewClient(databaseType, host, port, user, pass, database, optionMap)
    if err != nil {
        return nil, err
    }

    // TRUE if you want to output log
    // db.SetModeLog(true)

    return db, nil
}

func Sample() error {
    db, err := New()
    if err != nil {
        return err
    }
    defer func() {
        if p := recover(); p != nil {
            _ = db.Close()
            panic(p)
        }
        _ = db.Close()
    }()

    // insert
    {
        data := Item{}
        data.Num = 1
        data.Str = "sample"

        table := Item{}
        query := db.NewQuery(&table)
        query.SetValuesColumn(
            &table.Num,
            &table.Str,
        )
        query.SetValues(
            data.Num,
            data.Str,
        )

        // query: INSERT INTO "item" ("num", "str") VALUES (?)
        // values: [1 "sample"]
        _, err = query.Insert()
        if err != nil {
            return err
        }
    }

    // insert bulk
    {
        // If you want to set the number of bulk inserts, do as follows.
        db.Config.SetBulkInsertCount(2) // default 500

        table := Item{}
        query := db.NewQuery(&table)
        query.SetValuesColumn(
            &table.Num,
            &table.Str,
        )
        for i := 0; i < 3; i++ {
            data := Item{}
            data.Num = i
            data.Str = fmt.Sprintf("sample_%v", i)

            query.SetValues(
                data.Num,
                data.Str,
            )

            // query: INSERT INTO "item" ("num", "str") VALUES (?, ?),(?, ?)
            // values: [0 "sample_0" 1 "sample_1"]
            _, err := query.BulkInsert()
            if err != nil {
                return err
            }
        }

        // query: INSERT INTO "item" ("num", "str") VALUES (?, ?)
        // values: [2 "sample_2"]
        _, err := query.BulkInsertFinish()
        if err != nil {
            return err
        }
    }

    // update
    {
        data := Item{}
        data.Num = 1
        data.Str = "sample"

        table := Item{}
        query := db.NewQuery(&table)
        query.SetSet(&table.Str, data.Str)
        query.SetWhereIs(&table.Num, data.Num)

        // query: UPDATE "item" SET "str" = ? WHERE "item"."num" = 1
        // values: ["sample"]
        _, err = query.Update()
        if err != nil {
            return err
        }
    }

    // delete
    {
        table := Item{}
        query := db.NewQuery(&table)
        query.SetWhereIs(&table.Id, 1)

        // query: DELETE FROM "item" WHERE "item"."id" = $1
        // values: [1]
        _, err = query.Delete()
        if err != nil {
            return err
        }
    }

    // select
    {
        var resultList []*Item{}
        // var resultList []map[string]interface{}
        table := Item{}
        query := db.NewQuery(&table)
        query.SetSelectAll(&table)
        query.SetWhereIs(&table.Id, 1)
        // query: SELECT "item".* FROM "item" WHERE "item"."id" = $1
        // values: [1]
        err = query.Select(&resultList)
        if err != nil {
            return err
        }

        // any...
    }

    // select join
    {
        var resultList []*struct{
            Item
            Sample `column:"sample" json:"sample"`
        }
        table := Item{}
        tableDetail := UserDetail{}
        query := db.NewQuery()
        query.SetTable(&table)
        query.SetJoin(&tableDetail)
        query.SetJoinWhere(&tableDetail, &tableDetail.UserId, " = ", &table.Id)
        query.SetSelect(&table)
        query.SetSelect(
            &table.Id,
            &tableDetail.Sample,
        )
        query.SetWhereIs(&table.Id, 1)

        // query: SELECT "item"."id", "public"."item_detail"."sample" FROM "item" JOIN "public"."item_detail"."id" ON "public"."item_detail"."user_id" = "item"."id" WHERE "item"."id" = $1
        // values: [1]
        err = query.Select(&resultList)
        if err != nil {
            return err
        }
    }

    // database/sql.DB.Exec
    {
        // All placeholders are ?, no problem
        _, err := db.Exec(`insert into item ("str") values (?)`, "sample")
        if err != nil {
            return err
        }
    }

    // database/sql.DB.Query
    {
        // All placeholders are ?, no problem
        rows, err := db.Query(`select "item"."str" from "item" where "item"."id" = ?`, 1)
        if err != nil {
            return err
        }

        // Please refer to database/sql.
        for rows.Next() {
            // rows.Scan(....)
        }
    }

    // database/sql.DB.Query + 
    {
        // Client.Find() is Extract after executing query
        var resultList []*Item{}
        err := db.Find(&resultList, `select "item"."str" from "item" where "item"."id" = ?`, 1)
        if err != nil {
            return err
        }

        for _, val := range resultList {
            // any...
        }
    }

    return nil
}
```




# query methods

## table
|method|sql|
|---|---|
|SetTable(tablePtr interface{})|FROM tablePtr|
``` go
table := Item{}
query := db.NewQuery()
query.SetTable(&table) // query: "item"
```


## from
|method|sql|
|---|---|
|SetFrom(tablePtr interface{}, valueList ...interface{})|FROM [tablePtr or valueList...]|
``` go
table := Item{}
query := db.NewQuery()
query.SetFrom(&table) // query: "item" as "t1"
query.SetFrom(&table, `(SELECT * FROM "item" WHERE "id" = ?)`, []interface{}{1]}) // query: (SELECT * FROM "item" WHERE "id" = ?) as "t1" values: [1]
```


## join
|method|sql|
|---|---|
|SetJoin(tablePtr interface{}, valueList ...interface{})|JOIN [tablePtr or valueList...] ON|
|SetJoinLeft(tablePtr interface{}, valueList ...interface{})|LEFT JOIN [tablePtr or valueList...] ON|
|SetJoinRight(tablePtr interface{}, valueList ...interface{})|RIGHT JOIN [tablePtr or valueList...] ON|
``` go
table := Item{}
join := Detail{}
query := db.NewQuery()
query.SetJoin(&join) // query: INNER JOIN "public"."item_detail" as "t1" ON
query.SetJoinLeft(&join) // query: LEFT JOIN "public"."item_detail" as "t1" ON
query.SetJoinRigth(&join) // query: RIGHT JOIN "public"."item_detail" as "t1" ON
query.SetJoin(&join, "(select ?)", []interface{}{1}) // query: INNER JOIN (select ?) as "t1" ON values: [1]
```


## join where
|method|sql|
|---|---|
|SetJoinWhere(tablePtr interface{}, valueList ...interface{})|ON [and] ...|
|SetJoinWhereIs(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr = ?|
|SetJoinWhereIsNot(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr IS NOT ?|
|SetJoinWhereLike(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr LIKE ?|
|SetJoinWhereLikeNot(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr NOT LIKE ?|
|SetJoinWhereIn(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr IN (?)|
|SetJoinWhereInNot(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr NOT IN (?)|
|SetJoinWhereGt(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr > ?|
|SetJoinWhereGte(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr >= ?|
|SetJoinWhereLt(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr < ?|
|SetJoinWhereLte(tablePtr interface{}, valueList ...interface{})|ON [and] columnPtr <= ?|
|SetJoinWhereOr(tablePtr interface{}, valueList ...interface{})|ON [or] ...|
|SetJoinWhereOrIs(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr = ?|
|SetJoinWhereOrIsNot(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr IS NOT ?|
|SetJoinWhereOrLike(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr LIKE ?|
|SetJoinWhereOrLikeNot(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr NOT LIKE ?|
|SetJoinWhereOrIn(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr IN (?)|
|SetJoinWhereOrInNot(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr NOT IN (?)|
|SetJoinWhereOrGt(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr > ?|
|SetJoinWhereOrGte(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr >= ?|
|SetJoinWhereOrLt(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr < ?|
|SetJoinWhereOrLte(tablePtr interface{}, valueList ...interface{})|ON [or] columnPtr =< ?|
|SetJoinWhereNest()|ON ? [and] (|
|SetJoinWhereOrNest()|ON ? [or] (|
|SetJoinWhereNestClose()|ON ? )|
``` go
table := Item{}
join := Detail{}
query := db.NewQuery(&table)
query.SetJoin(&join)
query.SetJoinWhere(&join, &table.Id, "= ? or 1 = ?", []interface{}{1, 2}) // query: "item_detail"."id" = ? or 1 = ? values: [1 2]
query.SetJoinWhereIs(&join, &table.Id, 1) // query: "item_detail"."id" = 1 values: [1]
query.SetJoinWhereIs(&join, "max(", &table.Id, ")", 1) // query: max("item_detail"."id") = 1 values: [1]
query.SetJoinWhereIs(&join, &table.Id, nil) // query: "item_detail"."id" IS NULL
query.SetJoinWhereIsNot(&join, &table.Id, 1) // query: "item_detail"."id" != 1 values: [1]
query.SetJoinWhereIsNot(&join, &table.Id, nil) // query: "item_detail"."id" IS NOT NULL
query.SetJoinWhereLike(&join, &table.Str, "abc%") // query: "item_detail"."str" like 'abc%' values: ["abc%"]
query.SetJoinWhereLikeNot(&join, &table.Str, "abc%") // query: "item_detail"."str" like not 'abc%' values: ["abc%"]
query.SetJoinWhereIn(&join, &table.Id, []interface{}{1, 2, 3}) // query: "item_detail"."id" IN (1, 2, 3) values: [1 2 3]
query.SetJoinWhereInNot(&join, &table.Id, []interface{}{1, 2, 3}) // query: "item_detail"."id" IN NOT (1, 2, 3) values: [1 2 3]
query.SetJoinWhereGt(&join, &table.Id, 1) // query: "item_detail"."id" > 1 values: [1]
query.SetJoinWhereGte(&join, &table.Id, 1) // query: "item_detail"."id" => 1 values: [1]
query.SetJoinWhereLt(&join, &table.Id, 1) // query: "item_detail"."id" < 1 values: [1]
query.SetJoinWhereLte(&join, &table.Id, 1) // query: "item_detail"."id" =< 1 values: [1]
query.SetJoinWhereNest(&join) // query: (
query.SetJoinWhereNestClose(&join) // query: )
```


## insert into
|method|sql|
|---|---|
|SetValuesColumn(columnPtrList ...interface{})|INTO ? (columnPtrList...)|
|SetValues(valueList ...interface{})|VALUES (valueList...)|
|SetValuesClear()|is values clear|
``` go
table := Item{}
query := db.NewQuery(&table)
query.SetValuesColumn(&table.Id, &talbe.Str) // query: INTO "item" ("id", "str")
query.SetValues(1, "abc") // query: VALUES (?, ?) values: [1, "abc"]
query.SetValuesClear() // query: values delete
```


## set
|method|sql|
|---|---|
|SetSet(columnPtr interface{}, value interface{})|SET columnPtr = value|
``` go
table := Item{}
query := db.NewQuery(&table)
query.SetSet(&table.Id, 1) // query: "id" = 1
```


## select
|method|sql|
|---|---|
|SetSelect(valueList ...interface{})|SELECT valueList...|
|SetSelectAll(tablePtr interface{})|SELECT tablePtr.*|
``` go
table := Item{}
query := db.NewQuery(&table)
query.SetSelect("count(", &table, ")", []interface{}{1}) // query: count("item"), values: [1]
query.SetSelectAll(&table) // query: "item".*
```


## where
|method|sql|
|---|---|
|SetWhere(valueList ...interface{})|WHERE [and] ...|
|SetWhereIs(valueList ...interface{})|WHERE [and] columnPtr = ?|
|SetWhereIsNot(valueList ...interface{})|WHERE [and] columnPtr IS NOT ?|
|SetWhereLike(valueList ...interface{})|WHERE [and] columnPtr LIKE ?|
|SetWhereLikeNot(valueList ...interface{})|WHERE [and] columnPtr NOT LIKE ?|
|SetWhereIn(valueList ...interface{})|WHERE [and] columnPtr IN (?)|
|SetWhereInNot(valueList ...interface{})|WHERE [and] columnPtr NOT IN (?)|
|SetWhereGt(valueList ...interface{})|WHERE [and] columnPtr > ?|
|SetWhereGte(valueList ...interface{})|WHERE [and] columnPtr >= ?|
|SetWhereLt(valueList ...interface{})|WHERE [and] columnPtr < ?|
|SetWhereLte(valueList ...interface{})|WHERE [and] columnPtr <= ?|
|SetWhereOr(valueList ...interface{})|WHERE [or] ...|
|SetWhereOrIs(valueList ...interface{})|WHERE [or] columnPtr = ?|
|SetWhereOrIsNot(valueList ...interface{})|WHERE [or] columnPtr IS NOT ?|
|SetWhereOrLike(valueList ...interface{})|WHERE [or] columnPtr LIKE ?|
|SetWhereOrLikeNot(valueList ...interface{})|WHERE [or] columnPtr NOT LIKE ?|
|SetWhereOrIn(valueList ...interface{})|WHERE [or] columnPtr IN (?)|
|SetWhereOrInNot(valueList ...interface{})|WHERE [or] columnPtr NOT IN (?)|
|SetWhereOrGt(valueList ...interface{})|WHERE [or] columnPtr > ?|
|SetWhereOrGte(valueList ...interface{})|WHERE [or] columnPtr >= ?|
|SetWhereOrLt(valueList ...interface{})|WHERE [or] columnPtr < ?|
|SetWhereOrLte(valueList ...interface{})|WHERE [or] columnPtr =< ?|
|SetWhereNest()|WHERE [and] (|
|SetWhereOrNest()|WHERE [or] (|
|SetWhereNestClose()|WHERE )|
``` go
table := Item{}
query := db.NewQuery(&table)
query.SetWhere(&table.Id, "= ? or 1 = ?", []interface{}{1, 2}) // query: "item"."id" = ? or 1 = ? values: [1 2]
query.SetWhereIs(&table.Id, 1) // query: "item"."id" = 1 values: [1]
query.SetWhereIs("max(", &table.Id, ")", 1) // query: max("item"."id") = 1 values: [1]
query.SetWhereIs(&table.Id, nil) // query: "item"."id" IS NULL
query.SetWhereIsNot(&table.Id, 1) // query: "item"."id" != 1 values: [1]
query.SetWhereIsNot(&table.Id, nil) // query: "item"."id" IS NOT NULL
query.SetWhereLike(&table.Str, "abc%") // query: "item"."str" like 'abc%' values: ["abc%"]
query.SetWhereLikeNot(&table.Str, "abc%") // query: "item"."str" like not 'abc%' values: ["abc%"]
query.SetWhereIn(&table.Id, []interface{}{1, 2, 3}) // query: "item"."id" IN (1, 2, 3) values: [1 2 3]
query.SetWhereInNot(&table.Id, []interface{}{1, 2, 3}) // query: "item"."id" IN NOT (1, 2, 3) values: [1 2 3]
query.SetWhereGt(&table.Id, 1) // query: "item"."id" > 1 values: [1]
query.SetWhereGte(&table.Id, 1) // query: "item"."id" => 1 values: [1]
query.SetWhereLt(&table.Id, 1) // query: "item"."id" < 1 values: [1]
query.SetWhereLte(&table.Id, 1) // query: "item"."id" =< 1 values: [1]
query.SetWhereNest() // query: (
query.SetWhereNestClose() // query: )
```


## group by
|method|sql|
|---|---|
|SetGroupBy(valueList ...interface{})|GROUP BY valueList|
``` go
table := Item{}
query := db.NewQuery(&table)
query.SetGroupBy("count(", &table.Id, ")") // query: GROUP BY count("item"."id")
```


## having
|method|sql|
|---|---|
|SetHaving(valueList ...interface{})|Having [and] ...|
|SetHavingIs(valueList ...interface{})|Having [and] columnPtr = ?|
|SetHavingIsNot(valueList ...interface{})|Having [and] columnPtr IS NOT ?|
|SetHavingLike(valueList ...interface{})|Having [and] columnPtr LIKE ?|
|SetHavingLikeNot(valueList ...interface{})|Having [and] columnPtr NOT LIKE ?|
|SetHavingIn(valueList ...interface{})|Having [and] columnPtr IN (?)|
|SetHavingInNot(valueList ...interface{})|Having [and] columnPtr NOT IN (?)|
|SetHavingGt(valueList ...interface{})|Having [and] columnPtr > ?|
|SetHavingGte(valueList ...interface{})|Having [and] columnPtr >= ?|
|SetHavingLt(valueList ...interface{})|Having [and] columnPtr < ?|
|SetHavingLte(valueList ...interface{})|Having [and] columnPtr <= ?|
|SetHavingOr(valueList ...interface{})|Having [or] ...|
|SetHavingOrIs(valueList ...interface{})|Having [or] columnPtr = ?|
|SetHavingOrIsNot(valueList ...interface{})|Having [or] columnPtr IS NOT ?|
|SetHavingOrLike(valueList ...interface{})|Having [or] columnPtr LIKE ?|
|SetHavingOrLikeNot(valueList ...interface{})|Having [or] columnPtr NOT LIKE ?|
|SetHavingOrIn(valueList ...interface{})|Having [or] columnPtr IN (?)|
|SetHavingOrInNot(valueList ...interface{})|Having [or] columnPtr NOT IN (?)|
|SetHavingOrGt(valueList ...interface{})|Having [or] columnPtr > ?|
|SetHavingOrGte(valueList ...interface{})|Having [or] columnPtr >= ?|
|SetHavingOrLt(valueList ...interface{})|Having [or] columnPtr < ?|
|SetHavingOrLte(valueList ...interface{})|Having [or] columnPtr =< ?|
|SetHavingNest()|Having ? [and] (|
|SetHavingOrNest()|Having ? [or] (|
|SetHavingNestClose()|Having ? )|
``` go
table := Item{}
query := db.NewQuery(&table)
query.SetHaving(&table.Id, "= ? or 1 = ?", []interface{}{1, 2}) // query: "item"."id" = ? or 1 = ? values: [1 2]
query.SetHavingIs(&table.Id, 1) // query: "item"."id" = 1 values: [1]
query.SetHavingIs("max(", &table.Id, ")", 1) // query: max("item"."id") = 1 values: [1]
query.SetHavingIs(&table.Id, nil) // query: "item"."id" IS NULL
query.SetHavingIsNot(&table.Id, 1) // query: "item"."id" != 1 values: [1]
query.SetHavingIsNot(&table.Id, nil) // query: "item"."id" IS NOT NULL
query.SetHavingLike(&table.Str, "abc%") // query: "item"."str" like 'abc%' values: ["abc%"]
query.SetHavingLikeNot(&table.Str, "abc%") // query: "item"."str" like not 'abc%' values: ["abc%"]
query.SetHavingIn(&table.Id, []interface{}{1, 2, 3}) // query: "item"."id" IN (1, 2, 3) values: [1 2 3]
query.SetHavingInNot(&table.Id, []interface{}{1, 2, 3}) // query: "item"."id" IN NOT (1, 2, 3) values: [1 2 3]
query.SetHavingGt(&table.Id, 1) // query: "item"."id" > 1 values: [1]
query.SetHavingGte(&table.Id, 1) // query: "item"."id" => 1 values: [1]
query.SetHavingLt(&table.Id, 1) // query: "item"."id" < 1 values: [1]
query.SetHavingLte(&table.Id, 1) // query: "item"."id" =< 1 values: [1]
query.SetHavingNest() // query: (
query.SetHavingNestClose() // query: )
```


## order by
|method|sql|
|---|---|
|SetOrderBy(valueList ...interface{})|ORDER BY valueList|
|SetOrderByAsc(valueList ...interface{})|ORDER BY valueList|
|SetOrderByDesc(valueList ...interface{})|ORDER BY valueList DESC|
``` go
table := Item{}
query := db.NewQuery(&table)
query.SetOrderBy(&table.Id) // query: ORDER BY count("item"."id")
query.SetOrderByAsc("count(", &table.Id, ")") // query: ORDER BY count("item"."id") ASC
query.SetOrderByDesc("count(", &table.Id, ")") // query: ORDER BY count("item"."id") DESC
```


## limit
|method|sql|
|---|---|
|SetLimit(num int)|LIMIT num|
``` go
table := Item{}
query := db.NewQuery(&table)
query.SetLimit(1) // query: LIMIT 1
```


## offset
|method|sql|
|---|---|
|SetOffset(num int)|OFFSET num|
``` go
table := Item{}
query := db.NewQuery(&table)
query.SetOffset(1) // query: OFFSET 1
```
