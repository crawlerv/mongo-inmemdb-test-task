package config

type Config struct {
	Version               string `yaml:"version"`
	Log                   Log    `yaml:"log"`
	Mongo                 Mongo  `mapstructure:"mongo"`
	TickerIntervalSeconds uint32 `yaml:"tickerIntervalSeconds"`
}

type Log struct {
	IsProduction bool   `yaml:"is_production"`
	Dir          string `yaml:"dir"`
	File         string `yaml:"file"`
	Encoding     string `yaml:"encoding"`
}

type NameValue struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type Mongo struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
}
