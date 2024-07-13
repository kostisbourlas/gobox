package p2p

// Peer represents the remote node that is connected to us.
type Peer interface{}

// Transport handles the communication between the nodes in
// the network. Can be in the form of (TCP, UDP, Websockets, ...)
type Transport interface {
	ListenAndAccept() error
}
