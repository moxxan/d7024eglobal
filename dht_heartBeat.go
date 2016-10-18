package dht

import (
	"fmt"
	"time"
)

func (dhtnode *DHTNode) heartBeat() {
	nodeAdress := dhtnode.contact.ip + ":" + dhtnode.contact.port
	heartMsg := heartBeatMessage(nodeAdress, dhtnode.predecessor.adress)
	waitTimer := time.NewTimer(time.Second * 3)
	go func() { dhtnode.transport.send(heartMsg) }()
	for {
		select {
		case <-dhtnode.heartBeatQ:
			fmt.Println("stil alive baby", dhtnode.predecessor.adress)
			return

		case <-waitTimer.C:
			fmt.Println("heartstop nodeAress", nodeAdress)
			dhtnode.predecessor.adress = ""
			dhtnode.predecessor.nodeId = ""
			fmt.Println("heartstop STRING")
			dhtnode.stabilize()
			return
		}
	}
}

func (dhtnode *DHTNode) heartTimer() {
	for {
	//	fmt.Println("heartbeat")
		time.Sleep(time.Second * 4)
		dhtnode.createNewTask(nil, "heartBeat")
	}
}