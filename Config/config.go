package config

import "time"

type SQLiteConfig struct {
	DBfile string `yaml:"dbfile"`
}

type RedisConfig struct {
	Addr                          string `yaml:"addr"`
	Password                      string `yaml:"password"`
	TokenDB                       int    `yaml:"tokendb"`
	RateLimitDB                   int    `yaml:"ratelimitdb"`
	TTL_POST_INORDER_OF_COMMUNITY int    `yaml:"ttl_post_inorder_of_community"`
}

type RateLimitConfig struct {
	Rate    int `yaml:"rate"`
	NBucket int `yaml:"nbucket"`
}

type LogConfig struct {
	Logfile  string `yaml:"logfile"`
	Loglevel string `yaml:"loglevel"`
}

type JwtConfig struct {
	Issuer     string        `yaml:"issuer"`
	ExpireTime time.Duration `yaml:"expiretime"`
}

type SnowflakeConfig struct {
	StartTime    string `yaml:"starttime"`
	DataCenterID int64  `yaml:"datacenterid"`
	MachineID    int64  `yaml:"machineid"`
}

type Config struct {
	// Server
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`

	// SQLite
	SQLite SQLiteConfig `yaml:"sqlite"`

	// Redis
	Redis RedisConfig `yaml:"redis"`

	// RateLimit
	RateLimit RateLimitConfig `yaml:"ratelimit"`

	// Log
	Log LogConfig `yaml:"log"`

	// Jwt
	Jwt JwtConfig `yaml:"jwt"`

	// Snowflake
	Snowflake SnowflakeConfig `yaml:"snowflake"`
}

var Cfg *Config = &Config{
	IP:   "localhost",
	Port: "6500",

	SQLite: SQLiteConfig{
		DBfile: "bluebell.db",
	},

	Redis: RedisConfig{
		Addr:                          "localhost:6379",
		Password:                      "",
		TokenDB:                       0,
		RateLimitDB:                   1,
		TTL_POST_INORDER_OF_COMMUNITY: 20, // seconds
	},

	RateLimit: RateLimitConfig{
		Rate:    10,   // per second
		NBucket: 1000, // bucket number
	},

	Jwt: JwtConfig{
		Issuer:     "Bluebell",
		ExpireTime: 2 * time.Hour, // seconds
	},

	Snowflake: SnowflakeConfig{
		StartTime:    "2020-01-01 00:00:00",
		DataCenterID: 0,
		MachineID:    0,
	},

	Log: LogConfig{
		Logfile:  "bluebell.log",
		Loglevel: "info",
	},
}
