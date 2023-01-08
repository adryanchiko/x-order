package settings

type ServerOptions struct {
	Port       string `mapstructure:"port" yaml:"port"`
	APIBase    string `mapstructure:"api_base" yaml:"api_base"`
	DomainName string `mapstructure:"domain_name" yaml:"domain_name"`
	LogLevel   int    `mapstructure:"log_level" yaml:"log_level"`
}

type AppSettings struct {
	Name        string        `mapstructure:"name" yaml:"name"`
	Version     string        `mapstructure:"version" yaml:"version"`
	Description string        `mapstructure:"description" yaml:"description"`
	Server      ServerOptions `mapstructure:"server" yaml:"server"`
}
