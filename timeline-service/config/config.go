package config

type Config struct {
	ServerPort       string
	FollowServiceURL string
	PostServiceURL   string
}

func LoadConfig() *Config {
	cfg := &Config{}
	cfg.ServerPort = ":8080" // Puerto para el timeline-service
	cfg.FollowServiceURL = "http://follow-service:8080"
	cfg.PostServiceURL = "http://post-service:8080"
	return cfg
}
