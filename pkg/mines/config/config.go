package config

import (
	"net"
	"strings"
)

type Server struct {
	Host string
	Port string
}

func (c *Server) Endpoint() string {
	return net.JoinHostPort(c.Host, c.Port)
}

func (m *Memcached) Endpoints() string {
	return strings.Join(
		m.Servers,
		";",
	)
}

type Memcached struct {
	Servers []string
	Enabled bool
}

type Root struct {
	Seed      int
	TTL       int
	Server    Server
	Memcached Memcached
}

func Default() *Root {
	return &Root{
		Seed: 1234,
		TTL:  2,
		Server: Server{
			Host: "",
			Port: "65000",
		},
		Memcached: Memcached{
			Servers: []string{},
			Enabled: false,
		},
	}
}
