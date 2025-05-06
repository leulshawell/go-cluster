package work

import "github.com/leulshawell/go-cluster/go-cluster/group"

type Work struct {
}

func (w *Work) Schedule(g *group.Group) (bool, error) {

	return true, nil

}
