package Node

import (
	"github.com/leulshawell/go-cluster/go-cluster/group"
)

type Node struct {
}

func NewNode(c group.Config) (*Node, error) {
	return &Node{}, nil

}
