package cluster

import (
	//std
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"

	//dependecies
	quic "github.com/quic-go/quic-go"
	"github.com/syndtr/goleveldb/leveldb"

	"gihub.com/leulshawell/go-cluster/go-cluster/config"
	"github.com/leulshawell/go-cluster/go-cluster/fs"
	"github.com/leulshawell/go-cluster/go-cluster/group"
)

type Cluster struct {
	Server        *quic.Transport
	Ctx           context.Context
	clusterDataDb *leveldb.DB
	Stop          context.CancelFunc
	VolumeChan    chan struct{}
	Groups        []group.Group
}

func handleNewConn(con quic.Connection) {
	return
}

func NewCluster(ctx_ *context.Context, clusterConf *config.Config) (*Cluster, error) {
	ctx, cancel := context.WithCancel(*ctx_)

	//start a the dist. fs
	go fs.StartVolumeServer(clusterConf.RootPath, clusterConf.Port, &ctx, make(chan struct{}))

	//create the keyvalue store for the cluster
	//used for storing things like exec environments...

	//everthing happens over quic but not http3
	if ctx == nil {
		ctx = context.Background()
	}

	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: clusterConf.Port})
	if err != nil {
		return nil, err
	}

	tr := quic.Transport{
		Conn: udpConn,
	}

	return &Cluster{
		Server: &tr,
		Ctx:    ctx,
		Stop:   cancel,
	}, nil
}

func (c *Cluster) Start() {

	var (
		tlsConf  *tls.Config
		quicConf *quic.Config
	)

	ln, err := c.Server.Listen(tlsConf, quicConf)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error statrtign listener %w", err))
	}

	for {
		conn, err := ln.Accept(c.Ctx)
		if err != nil {
			fmt.Println("Error accepting conn")
		}

		go handleNewConn(conn)
		// handle the connection, usually in a new Go routine
	}

}
