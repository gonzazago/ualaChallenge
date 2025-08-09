package config

// Config contiene toda la configuración para la aplicación.
// En un proyecto real, esto se cargaría desde variables de entorno o un archivo.
type Config struct {
	ServerPort string
	SQLConfig  struct {
		URLString string
	}
}

// LoadConfig carga la configuración. Por ahora, usamos valores fijos.
func LoadConfig() *Config {
	cfg := &Config{}
	cfg.ServerPort = ":8080"
	cfg.SQLConfig.URLString = "user:password@tcp(127.0.0.1:3306)/users_db?parseTime=true"
	return cfg
}
