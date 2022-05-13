package configuration

type Config struct {
	Port int
	Env  string
}

func NewConfig() Config {
	return Config{}
}
