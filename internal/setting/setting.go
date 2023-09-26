package setting

func New(httpPort int, dbDrv, dbSource string) *Config {
	return &Config{
		Http: HttpCfg{Port: httpPort},
		DB:   DBCfg{Driver: dbDrv, Source: dbSource},
	}
}

type Config struct {
	Http HttpCfg
	DB   DBCfg
}
