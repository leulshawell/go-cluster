package work

import (
	"github.com/leulshawell/go-cluster/go-cluster/group"
)

type Work struct {
	path string
	env  string //the environment
}

func (w *Work) Schedule(g *group.Group) (bool, error) {

	return true, nil

}
