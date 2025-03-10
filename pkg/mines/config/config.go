package config

import (
	"net"
)

type Server struct {
	Host string
	Port string
}

func (c *Server) Endpoint() string {
	return net.JoinHostPort(c.Host, c.Port)
}

type Reaper struct {
	FrequencySeconds int
	TimeoutSeconds   int
}

type Root struct {
	Seed   int
	Server Server
	Reaper Reaper
}

func Default() *Root {
	return &Root{
		Seed: 1234,
		Server: Server{
			Host: "",
			Port: "65000",
		},
		Reaper: Reaper{
			FrequencySeconds: 60,
			TimeoutSeconds:   10,
		},
	}
}
