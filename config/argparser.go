package config

import (
	"flag"
	"fmt"
)

type Config struct {
	Host string
	Port string
}

func ParseArgs() *Config {
	host := flag.String("host", "0.0.0.0", "REST API host")
	port := flag.String("port", "8080", "REST API port")

	flag.Parse()

	fmt.Printf("Running on %s:%s\n", *host, *port)
	return &Config{
		Host: *host,
		Port: *port,
	}
}
