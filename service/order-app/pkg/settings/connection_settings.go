package settings

import (
	"time"
)

type SqlOption struct {
	Driver             string        `mapstructure:"driver" yaml:"driver"`
	URI                string        `mapstructure:"uri" yaml:"uri"`
	MaxConnLifeTime    time.Duration `mapstructure:"max_conn_life_time" yaml:"max_conn_life_time"`
	MaxIdleConnections int           `mapstructure:"max_idle_connections" yaml:"max_idle_connections"`
	MaxOpenConnections int           `mapstructure:"max_open_connections" yaml:"max_open_connections"`
	Enabled            bool          `mapstructure:"enabled" yaml:"enabled"`
}

type ConnectionSettings struct {
	Sql SqlOption `mapstructure:"sql"`
}
