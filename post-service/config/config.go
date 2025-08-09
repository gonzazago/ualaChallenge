package config

type Config struct {
	ServerPort string
}

// LoadConfig carga la configuración.
func LoadConfig() *Config {
	cfg := &Config{}
	cfg.ServerPort = ":8080" // Usamos un puerto diferente para no chocar con otros servicios
	return cfg
}
