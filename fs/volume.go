package fs

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"os"

	quic "github.com/quic-go/quic-go"
)

type Volume struct {
}

func handleNewConn(con quic.Connection) {
	return
}

// start an independent volume server
func StartVolumeServer(rooPath string, port int, ctx *context.Context, clusterChan chan struct{}) error {

	if _, err := os.Stat(rooPath); os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.MkdirAll(rooPath, 0755)
		if err != nil {
			log.Fatal("failed to create directory: %w", err)
		}
	}

	var (
		tlsConf  *tls.Config
		quicConf *quic.Config
	)

	udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: port})
	if err != nil {
		return nil
	}

	tr := quic.Transport{
		Conn: udpConn,
	}

	ln, err := tr.Listen(tlsConf, quicConf)

	for {
		conn, err := ln.Accept(*ctx)
		if err != nil {
			fmt.Printf(err.Error())
		}
		go handleNewConn(conn)
	}
}
