package group

import (
	"github.com/google/uuid"
	"github.com/leulshawell/go-cluster/node"
	"net"
)

type Config struct {
	Interface net.Interface
	Port      int
	Key       uuid.UUID
}

type Group struct {
	name   string
	nodes  []*node.Node
	config *Config
}

func (g *Group) Accept(n *Node) error {
	return nil
}

// Join the cluster
func (s *Node) Join(g *Group) error {
	return g.Accept(s)
}

func Start() {
	//status check for nodes

	//schedule jobs from queue on free nodes

}
