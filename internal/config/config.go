package config

type Server struct {
	Metadata Metadata `toml:"metadata"`
	Log      Log      `toml:"logging"`
	Port     int      `toml:"port" env-required:"true"`
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
