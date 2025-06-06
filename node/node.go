package node

import (
	"os"
	"strings"

	"github.com/leulshawell/go-cluster/group"
)

type Node struct {
}

func NewNode(c group.Config) (*Node, error) {
	return &Node{}, nil

}

// Groups for a node are stored as env variable
// GO_CLUSTER_GROUPS in key1:key2:key3... format
func Groups() []*group.Group {
	groupsListEnv := os.Getenv("GO_CLUSTER_GROUPS")
	if groupsListEnv == "" {
		return []*group.Group{}
	}

	keys := strings.Split(groupsListEnv, ":")
	var groups []*group.Group

	for _, key := range keys {
		//get group config from minikey server

		c, err := group.GetGroupConfig(key)
		if err != nil {
			panic("Grooup doesn't exist but key still does")
		}

		groups = append(groups, &group.Group{
			Config: c,
		})

	}

	return groups
}
