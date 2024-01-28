package network

import (
	"fmt"
	"time"
)

type ServerOpts struct {
	Transport []Transport
}

type Server struct {
	ServerOpts
	rpcCh  chan RPC
	quitCh chan struct{}
}

func NewServer(opts ServerOpts) *Server {

	return &Server{
		ServerOpts: opts,
		rpcCh:      make(chan RPC),
		quitCh:     make(chan struct{}, 1),
	}
}

func (s *Server) Start() {

	s.initTransport()
	ticker := time.NewTicker(5 * time.Second)
free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v\n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			fmt.Printf("do somethng every 5 sec \n")
		}
	}
	fmt.Println("server_shutdown")
}

func (s *Server) initTransport() {

	for _, tr := range s.Transport {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}
}
