package config

type Config struct {
	Port     int
	DbDriver string
	DbSource string
}

//var lock = &sync.Mutex{}

var c *Config

func SetConfig(port int, driver, source string) {
	if c == nil {
		c = &Config{
			port,
			driver,
			source,
		}
	}
}

func GetConfig() *Config {
	return c
}
