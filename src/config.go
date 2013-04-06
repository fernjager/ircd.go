package main

type Config struct {
	options map[string]string
}

func ConfigInit() *Config {
	return &Config{make(map[string]string)}
}

func (conf *Config) getOption(key string) string {
	return conf.options[key]
}

func (conf *Config) setOption(key string, value string) {
	conf.options[key] = value
}
