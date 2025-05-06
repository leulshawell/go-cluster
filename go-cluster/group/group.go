package group

import (
	// "github.com/geohot/minikeyvalue"
	"github.com/google/uuid"
	"net"
)

const (
	FIRST_UP = iota
	HIGH
	Completed
	Failed
)

type Config struct {
	Interface          net.Interface
	Port               int
	Key                uuid.UUID
	MasterSelectionRle int
}

type Group struct {
	volume string
	Name   string
	Config *Config
}

func (g *Group) Up() {
	//start minikeyvalue instance for group data

	//schedule

	//start start work

}

// Join the cluster
func (g *Group) Accept() error {
	return nil
}

func Start() {
	//status check for nodes

	//schedule jobs from queue on free nodes

}
