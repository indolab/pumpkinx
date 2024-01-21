package network

import "sync"

type LocalTransport struct {
	addr      NetAddr
	consumeCh chan RPC
	lock      sync.RWMutex
	peers     map[NetAddr]*LocalTransport
}
