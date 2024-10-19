package http_transport

type Config struct {
	Port int `yaml:"port"`
}

type ProxyConfig struct {
	Host string `yaml:"host"`
}

type Proxy struct{}
