package config

type Config struct {
    DB     DBConfig
    Redis  RedisConfig
    JWT    JWTConfig
}

type DBConfig struct {
    DSN string `mapstructure:"DB_DSN"`
}

type RedisConfig struct {
    Addr string `mapstructure:"REDIS_ADDR"`
}

type JWTConfig struct {
    Secret string `mapstructure:"JWT_SECRET"`
}
