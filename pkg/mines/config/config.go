package config

import "net"

type Server struct {
	Host string
	Port string
}

func (c *Server) Endpoint() string {
	return net.JoinHostPort(c.Host, c.Port)
}

type Root struct {
	Seed   int
	Server Server
}
