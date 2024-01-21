package network

type NetAddr string

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}

type RPC struct {
	From    NetAddr
	Payload []byte
}