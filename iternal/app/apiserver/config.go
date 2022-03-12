package apiserver

type Config struct {
	BindAddr             string `toml:"bind_addr"`
	LogLevel             string `toml:"log_level"`
	DatabaseURL          string `toml:"database_url"`
	JwtSecretKey         string `toml:"jwt_secret_key"`
	AccessTokenLifetime  int    `toml:"access_token_lifetime"`
	RefreshTokenLifetime int    `toml:"refresh_token_lifetime"`
}

func NewConfig() *Config {
	return new(Config)
}
