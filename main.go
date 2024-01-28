package main

import (
	"time"

	"github.com/indolab/pumpkinx/network"
)

func main() {

	trLocal := network.NewLocalTransport("Local")
	trRemote := network.NewLocalTransport("Remote")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.SendMessage(trLocal.Addr(), []byte("hello Local"))
			time.Sleep(1 * time.Second)
		}

	}()

	opts := network.ServerOpts{
		Transport: []network.Transport{trLocal},
	}
	s := network.NewServer(opts)
	s.Start()
}
