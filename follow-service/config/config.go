package config

// Config contiene toda la configuración para la aplicación.
type Config struct {
	ServerPort string
}

// LoadConfig carga la configuración.
func LoadConfig() *Config {
	cfg := &Config{}
	cfg.ServerPort = ":8080"
	return cfg
}
