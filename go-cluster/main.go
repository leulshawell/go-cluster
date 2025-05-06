package main

import (
	"flag"
	"fmt"

	"github.com/google/uuid"
	"github.com/leulshawell/go-cluster/go-cluster/fs"
	"github.com/leulshawell/go-cluster/go-cluster/group"
	"github.com/leulshawell/go-cluster/go-cluster/work"
)

func main() {

	key := flag.String("key", "", "--key=groupkey")
	fmt.Println(key)
	// groups := node.Groups() //discover groups on this Node is on

	myGroup, err := group.GetGroupConfig(*key)
	if err != nil {
		panic("Group not found")
	}

	work := work.Work{}

	work.Schedule(myGroup)

	println("Hello from the cluster")
	file := fs.Open("/root")
	file.Read([]byte{})

}
