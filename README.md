# gol

## connect
``` go
host := "localhost"
port := "5432"
user := "user"
pass := "pass"
database := config.DbDatabase
optionMap := make(map[string]string)
optionMap["sslMode"] = gol.PostgresqlSslModeDisable
db, err := gol.Open(gol.DatabaseTypePostgresql, host, port, user, pass, database, optionMap)
if err != nil {
    return nil, err
}
```

## select
``` go
type User struct {
    Id `column:"id"`
    Name `column:"name"`
}

var result []*User
table := User{}
query := db.NewQuery(&table)
query.SetSelect(&table)

// select user.* from user
err := query.Select(&result)
if err != nil {
    // error
}
```