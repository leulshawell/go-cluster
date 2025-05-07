package main

import (
	"github.com/leulshawell/go-cluster/go-cluster/group"
)

func main() {

	key := "some-group-key"
	// groups := node.Groups() //discover groups on this Node is on

	_, err := group.GetGroupConfig(key)
	if err != nil {
		panic("Group not found")
	}

}
