package wiring

import (
	"sepehrsoh/go-http-server/config"
)

type Wire struct {
	Configurations *config.Configurations
}

var HttpWire *Wire = nil

func NewWiring(configurations *config.Configurations) *Wire {
	if HttpWire == nil {
		wire := &Wire{
			Configurations: configurations,
		}

		HttpWire = wire
	}
	return HttpWire
}
