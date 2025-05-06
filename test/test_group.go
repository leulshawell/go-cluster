package main

import (
	"net"

	"github.com/google/uuid"
	"github.com/leulshawell/go-cluster/go-cluster/group"
)

func main() {

	interfaces, _ := net.Interfaces()

	config := &group.Config{
		Key:       uuid.New(),
		Interface: interfaces[0],
		Port:      7000,
	}

	group := &group.Group{
		Name:   "trainer",
		Config: config,
	}

	group.Up()
	return
}
