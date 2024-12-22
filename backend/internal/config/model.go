package config

type Config struct {
	Server   Server
	Postgres PostgresDatabase
}

type Server struct {
	Host         string
	Port         string
	AllowOrigins []string
}

type PostgresDatabase struct {
	DSN string
}
