package config

type Config struct {
	Mode     string
	Server   Server
	Postgres PostgresDatabase
}

type Server struct {
	Host         string
	Port         string
	AllowOrigins string
}

type PostgresDatabase struct {
	DSN string
}

func (cfg *Config) IsDev() bool {
	return cfg.Mode == "DEV"
}
