package config

type Config struct {
	Listener
	LogLevel string
}

type Listener struct {
	Host string
	Port int
}
