package config

// TODO: Use Viper to load env variables (i.e. for client directory and appname)

type Config struct {
	AppName     string
	Parallel    int
	BaseURL     string
	LogLevel    string
	EnableLocal bool
}

func LoadConfig(appName string) *Config {
	return &Config{
		AppName:     appName,
		Parallel:    64,
		BaseURL:     "https://raw.githubusercontent.com/luan/tibia-client/main/",
		LogLevel:    "info",
		EnableLocal: false,
	}
}
