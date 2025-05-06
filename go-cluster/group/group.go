package group

import (
	"github.com/google/uuid"
	"net"
)

const (
	FIRST_UP = iota
	HIGH
	Completed
	Failed
)

type Job struct {
	path string
	env  string
}

type Config struct {
	Interface          net.Interface
	Port               int
	Key                uuid.UUID
	MasterSelectionRle int
}

type Group struct {
	name    string
	config  *Config
	jobPool []Job
}

func (g *Group) Up() {
	//discover

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
