package setting

func New() *Config {
	return &Config{
		Http: HttpConfig{Port: 8080},
	}
}

type Config struct {
	Http HttpConfig
}
