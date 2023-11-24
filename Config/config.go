package config

import "time"

type Server struct {
	IP   string `yaml:"ip"`
	Port string `yaml:"port"`
}

type SQLiteConfig struct {
	DBfile string `yaml:"dbfile"`
}

type LogicConfig struct {
	PasswordRange   [2]int `yaml:"passwordrange"`
	UsernameRange   [2]int `yaml:"maxusernamelen"`
	NDuplicateLogin int    `yaml:"nduplicatelogin"`
	MaxPageSize     int64  `yaml:"maxpagesize"`
}

type RedisConfig struct {
	Addr                          string `yaml:"addr"`
	Password                      string `yaml:"password"`
	TokenDB                       int    `yaml:"tokendb"`
	PostDB                        int    `yaml:"postdb"`
	RateLimitDB                   int    `yaml:"ratelimitdb"`
	TTL_POST_INORDER_OF_COMMUNITY int    `yaml:"ttl_post_inorder_of_community"`
}

type RateLimitConfig struct {
	Rate    int `yaml:"rate"`
	NBucket int `yaml:"nbucket"`
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

type LogConfig struct {
	Logfile  string `yaml:"logfile"`
	Loglevel string `yaml:"loglevel"`
}

type Config struct {
	// Server
	Server Server `yaml:"server"`

	// SQLite
	SQLite SQLiteConfig `yaml:"sqlite"`

	// Logic
	Logic LogicConfig `yaml:"logic"`

	// Redis
	Redis RedisConfig `yaml:"redis"`

	// RateLimit
	RateLimit RateLimitConfig `yaml:"ratelimit"`

	// Jwt
	Jwt JwtConfig `yaml:"jwt"`

	// Snowflake
	Snowflake SnowflakeConfig `yaml:"snowflake"`

	// Log
	Log LogConfig `yaml:"log"`
}

var Cfg *Config = &Config{
	Server: Server{
		IP:   "localhost",
		Port: "8080",
	},

	SQLite: SQLiteConfig{
		DBfile: "bluebell.db",
	},

	Logic: LogicConfig{
		UsernameRange:   [2]int{4, 16},
		PasswordRange:   [2]int{8, 20},
		NDuplicateLogin: 3,
		MaxPageSize:     10,
	},

	Redis: RedisConfig{
		Addr:                          "localhost:6379",
		Password:                      "",
		TokenDB:                       0,
		PostDB:                        1,
		RateLimitDB:                   2,
		TTL_POST_INORDER_OF_COMMUNITY: 20, // seconds
	},

	RateLimit: RateLimitConfig{
		Rate:    20,  // per second
		NBucket: 800, // bucket number
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
