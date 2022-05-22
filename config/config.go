package config

type Config struct {
	BindAddr string `env:"BIND_ADDR" env-default:":8080"`
	LogLevel string `env:"LOD_LEVEL" env-default:"debug"`
	Postgres
	Minio
	JWT
}
type Postgres struct {
	User     string `env:"USER" env-default:"postgres"`
	Password string `env:"PASSWORD" env-default:"Jad108fsdlknzc"`
	Host     string `env:"HOST" env-default:"host.docker.internal:5436"`
	DBName   string `env:"DB_NAME" env-default:"postgres"`
	SSL      string `env:"SSL" env-default:"disable"`
}

type Minio struct {
	EndPoint        string `env:"ENDPOINT" env-default:"cis-software.cf:9000"`
	BucketName      string `env:"BUCKET_NAME" env-default:"min1"`
	AccessKeyID     string `env:"ACCESS_KEY_ID" env-default:"minio"`
	SecretAccessKey string `env:"SECRET_ACCESS_KEY" env-default:"minio123"`
	UseSSL          bool   `env:"USE_SSL" env-default:"false"`
}

type JWT struct {
	SecretKey            string `env:"SECRET_KEY" env-default:"key123"`
	AccessTokenLifetime  int    `env:"ACCESS_TOKEN_LIFETIME" env-default:"20"`
	RefreshTokenLifetime int    `env:"REFRESH_TOKEN_LIFETIME" env-default:"525600"`
}

func NewConfig() *Config {
	return new(Config)
}
