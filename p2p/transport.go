package p2p

import "net"

// peer is an interface that represents the remote node.
type Peer interface {
	net.Conn
	Send([]byte) error
	CloseStream()
}

// its anything that handle communication
// between the nodes in the network
type Transport interface {
	Addr() string
	Dial(string) error
	ListenAndAccept() error
	Consume() <-chan RPC
	Close() error
}
