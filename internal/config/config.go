package config

type Server struct {
	Port     int      `toml:"port" env-required:"true"`
	Log      Log      `toml:"logging"`
	Metadata Metadata `toml:"metadata"`
	Postgres Postgres `toml:"postgres"`
}

type Log struct {
	Level  string `toml:"level" env-required:"true"`
	APIKey string `env:"LOGGING_API_KEY" env-required:"true"`
}

type Metadata struct {
	ServiceName string `toml:"service_name" env-required:"true"`
	InstanceID  string
	Version     string
	BuildTime   int64
}

type Postgres struct {
	URL string `env:"POSTGRES_URL" env-required:"true"`
}
