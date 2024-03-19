package config

import "fmt"

var C = new(Config)

type Config struct {
	App string `toml:"app"`

	AppEnv   string `toml:"app_env"`
	RunMode  string `toml:"run_mode"`
	Timezone string `toml:"timezone"`
	Log      Log    `toml:"log"`
	HTTP     HTTP   `toml:"http"`
	MySQL    MySQL  `toml:"mysql"`
	Gorm     Gorm   `toml:"gorm"`
	Redis    Redis  `toml:"redis"`
}

type Log struct {
	Level      int8
	OutputPath string
}

type HTTP struct {
	Host               string
	Port               int
	CertFile           string
	KeyFile            string
	ShutdownTimeout    int
	MaxContentLength   int64
	MaxReqLoggerLength int `default:"1024"`
	MaxResLoggerLength int
}

type Gorm struct {
	Debug             bool
	DBType            string
	MaxLifetime       int
	MaxOpenConns      int
	MaxIdleConns      int
	TablePrefix       string
	EnableAutoMigrate bool
}

type MySQL struct {
	Host       string
	Port       int
	User       string
	Password   string
	DBName     string
	Parameters string
}

func (a MySQL) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s",
		a.User, a.Password, a.Host, a.Port, a.DBName, a.Parameters)
}

type Redis struct {
	Addr     string `toml:"addr"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}
