package main

import (
	"github.com/leulshawell/go-cluster/go-cluster/fs"
	"github.com/leulshawell/go-cluster/go-cluster/node"
)

func main() {
	groups := node.Groups() //discover groups on this Node is on

	println("Hello from the cluster")
	file := fs.Open("/root")
	file.Read([]byte{})

}
