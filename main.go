package main

import (
	"fmt"
	"log"

	"github.com/tahaontech/filesnetwork/p2p"
)

func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    ":4000",
		HandShakeFunc: p2p.NopHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
	}
	tr := p2p.NewTCPTransport(tcpOpts)

	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Server is running")

	select {}
}
