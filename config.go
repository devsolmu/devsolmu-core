package main

// Config define DBConfig
type Config struct {
	DB *DBConfig
}

// DBConfig define database connect information
type DBConfig struct {
	Dialect  string
	Username string
	Password string
	Name     string
	Charset  string
}

// GetConfig is connecting part
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Username: "root",
			Password: "",
			Name:     "devsolmu",
			Charset:  "utf8",
		},
	}
}
