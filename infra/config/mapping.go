package config

// Config app config
type Config struct {
	Global GlobalConfig `mapstructure:"global"`
	Db     DBConfig     `mapstructure:"db"`
	Tests  TestsConfig  `mapstructure:"tests"`
}

// GlobalConfig global app configs
type GlobalConfig struct {
	SecretKey string `mapstructure:"secret-key"`
	TokenName string `mapstructure:"token-name"`
}

// DBConfig database configs
type DBConfig struct {
	User      string `mapstructure:"user"`
	Name      string `mapstructure:"name"`
	Password  string `mapstructure:"password"`
	Port      int    `mapstructure:"port"`
	Host      string `mapstructure:"host"`
	SilentLog bool   `mapstructure:"silent-log-mode"`
}

// TestsConfig tests configs
type TestsConfig struct {
	Db TestsDBConfig `mapstructure:"db"`
}

// TestsDBConfig database configs
type TestsDBConfig struct {
	User      string `mapstructure:"user"`
	Name      string `mapstructure:"name"`
	Password  string `mapstructure:"password"`
	Port      int    `mapstructure:"port"`
	Host      string `mapstructure:"host"`
	SilentLog bool   `mapstructure:"silent-log-mode"`
}
