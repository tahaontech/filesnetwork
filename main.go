package main

import (
	"fmt"
	"log"

	"github.com/tahaontech/filesnetwork/p2p"
)

func OnPeerFunc(peer p2p.Peer) error {
	// do logics with peer
	peer.Close()

	return nil
}

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":3000",
		HandShakeFunc: p2p.NopHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeerFunc,
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("msg: %+v\n", msg)
		}
	}()

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is running")

	select {}
}
