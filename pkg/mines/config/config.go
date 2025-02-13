package config

import "net"

type Server struct {
	Host string
	Port string
}

func (c *Server) Endpoint() string {
	return net.JoinHostPort(c.Host, c.Port)
}

type Reaper struct {
	Interval int
	Bundled  bool
	Timeout  int
}

type Root struct {
	Seed   int
	Server Server
	Reaper Reaper
}
