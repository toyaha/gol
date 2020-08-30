package gol

const (
	DatabaseTypeMssql      = "sqlserver"
	DatabaseTypeMysql      = "mysql"
	DatabaseTypePostgresql = "postgres"

	NamingConvention             = "none"
	NamingConventionPascalCase   = "PascalCase"
	NamingConventionCamelCase    = "camelCase"
	NamingConventionConstantCase = "CONSTANT_CASE"
	NamingConventionSnakeCase    = "snake_case"

	PlaceHolder = "?"

	// Not Use ssl, Not check the certificate.
	PostgresqlSslModeDisable = "disable"
	// Use ssl, Not check the certificate.
	PostgresqlSslModeRequire = "require"
	// Use ssl, Check the certificate
	PostgresqlSslModeVerifyCa = "verify-ca"
	// Use ssl, Check the certificate and confirm that it is on the server.
	PostgresqlSslModeVerifyFull = "verify-full"

	QueryJoinModeInner = iota
	QueryJoinModeLeft
	QueryJoinModeRight

	QueryModeDefault = iota
	QueryModeOne
	QueryModeAll
	QueryModeIs
	QueryModeIsNot
	QueryModeLike
	QueryModeLikeNot
	QueryModeIn
	QueryModeInNot
	QueryModeGt
	QueryModeGte
	QueryModeLt
	QueryModeLte
	QueryModeNest
	QueryModeNestClose

	QueryPrefixNone = ""
	QueryPrefixAnd  = "AND"
	QueryPrefixOr   = "OR"

	StructFieldTagNameSchema = "schema"
	StructFieldTagNameTable  = "table"
	StructFieldTagNameColumn = "column"
)

func NewConfig() *Config {
	config := &Config{
		BulkInsertCount:  500,
		DatabaseType:     DatabaseTypePostgresql,
		Log:              false,
		ResultKey:        NamingConvention,
		NamingConvention: NamingConventionSnakeCase,
		Test:             false,
	}

	return config
}

type Config struct {
	BulkInsertCount  int
	DatabaseType     string
	Log              bool
	ResultKey        string
	NamingConvention string
	Test             bool
}

func (rec *Config) InitMssql() {
	rec.SetDatabaseTypeMssql()
}

func (rec *Config) InitMysql() {
	rec.SetDatabaseTypeMysql()
}

func (rec *Config) InitPostgresql() {
	rec.SetDatabaseTypePostgresql()
}

func (rec *Config) SetBulkInsertCount(count int) {
	rec.BulkInsertCount = count
}

func (rec *Config) SetDatabaseTypeMssql() {
	rec.DatabaseType = DatabaseTypeMssql
}

func (rec *Config) SetDatabaseTypeMysql() {
	rec.DatabaseType = DatabaseTypeMysql
}

func (rec *Config) SetDatabaseTypePostgresql() {
	rec.DatabaseType = DatabaseTypePostgresql
}

func (rec *Config) SetLog(mode bool) {
	rec.Log = mode
}

func (rec *Config) SetResultKey() {
	rec.ResultKey = NamingConvention
}

func (rec *Config) SetResultKeyCamelCase() {
	rec.ResultKey = NamingConventionCamelCase
}

func (rec *Config) SetResultKeySnakeCase() {
	rec.ResultKey = NamingConventionSnakeCase
}

func (rec *Config) SetNamingConventionConstantCase() {
	rec.NamingConvention = NamingConventionConstantCase
}

func (rec *Config) SetNamingConventionSnakeCase() {
	rec.NamingConvention = NamingConventionSnakeCase
}

func (rec *Config) SetTest(mode bool) {
	rec.Test = mode
}
