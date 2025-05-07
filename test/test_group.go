package test

import (
	"context"

	"github.com/leulshawell/go-cluster/fs"
)

func Test_server() {

	ctx := context.Background()
	fs.StartVolumeServer("/", 8888, &ctx, make(chan struct{}))

	return
}
