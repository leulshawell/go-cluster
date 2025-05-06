package group

import (
	"github.com/google/uuid"
	"net"
)

type Config struct {
	Interface net.Interface
	Port      int
	Key       uuid.UUID
}

type Group struct {
	name   string
	config *Config
}

// Join the cluster
func (g *Group) Accept() error {
	return nil
}

func Start() {
	//status check for nodes

	//schedule jobs from queue on free nodes

}
