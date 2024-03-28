package p2p

// peer is an interface that represents the remote node.
type Peer interface {
}

// its anything that handle communication
// between the nodes in the network
type Transport interface {
	ListenAndAccept() error
}
