package main

import (
	"fmt"

	"go.etcd.io/raft/v3"
)

func main() {
	storage := raft.NewMemoryStorage()
	c := &raft.Config{
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}

	n := raft.StartNode(c, []raft.Peer{{ID: 0x02}, {ID: 0x03}})
	fmt.Printf("start successd: %+v\n", n)
}
