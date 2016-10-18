package dht

import (
	"encoding/json"
	"fmt"
	"net"
)

type Transport struct {
	node        *DHTNode
	bindAddress string // rad 20, bindadress måste finnas.
	msgQ        chan *Msg
}

func (transport *Transport) listen() {
	udpAddr, err := net.ResolveUDPAddr("udp", transport.bindAddress)
	//	fmt.Println("transport bindaddress:", transport.bindAddress)
	conn, err := net.ListenUDP("udp", udpAddr)
	conn.SetReadBuffer(10000)
	conn.SetWriteBuffer(10000)
	if err != nil {
		fmt.Println("error LISTEN function is:", err)
	}
	defer conn.Close()
	dec := json.NewDecoder(conn)
	for {
		msg := Msg{}
		err = dec.Decode(&msg)
		go func() {
			transport.msgQ <- &msg
		}()

	}

}

func (transport *Transport) send(msg *Msg) {
	if transport.node.alive {
	//	fmt.Println("transport send msg.dst:", msg.Dst)
		udpAddr, err := net.ResolveUDPAddr("udp", msg.Dst)
		conn, err := net.DialUDP("udp", nil, udpAddr)
		if err != nil {
			fmt.Println("error SEND function is:", err)
		}
		encoded, err := json.Marshal(msg)
		defer conn.Close()
		_, err = conn.Write(encoded)
	}
}

func (transport *Transport) initmsgQ() {
	go func() {
		for {
			select {
			case msg := <-transport.msgQ:
				switch msg.Type {
				case "fingerPrint": //test case
					transport.node.printNetworkFingers(msg)
					//transport.node.TaskQ <- &Task{msg, "printRingFingers"}
				//	transport.send(&Msg{"printRing", "", v.Src, []byte("tjuuu")})
				case "reply": //test
					fmt.Println("hej:", string(msg.Bytes))
				case "printRing":
					go func() { transport.node.TaskQ <- &Task{msg, "printRing"} }() //transport.node.printRing()
					//transport.send(&Msg{"ring", "", v.Src, []byte(transport.node.printRing())})
				case "addToRing":
					transport.node.printNetworkRing(msg)
				case "response":
					go func() { transport.node.responseQ <- msg }()
				case "join":
					fmt.Println("transport join")
					go func() { transport.node.TaskQ <- &Task{msg, "join"} }()
				case "notify":
					//		fmt.Println("notify network")
					go transport.node.notifyNetwork(msg)
				case "pred":
					go transport.node.getPred(msg)
				case "lookup":
					go transport.node.improvedNetworkLookUp(msg)
					//fmt.Println("initmsgQ lookup: ")
					//go transport.node.networkLookup(msg)
				case "fingerLookup":
					//go transport.node.LookUpNetworkFinger(msg)
					//go transport.node.lookupFingers(msg)
				case "heartBeat":
					if transport.node.alive {
						transport.node.transport.send(heartBeatAnswer(msg.Origin, msg.Dst))
					}
				case "heartAnswer":
					go func() { transport.node.heartBeatQ <- msg }()
				case "isAlive":
					if transport.node.alive {
						transport.node.transport.send(responseMessage(msg.Dst, msg.Origin, transport.bindAddress, transport.node.nodeId))
					}
				case "nodeFound":
					transport.node.transport.send(ackMsg(msg.Dst, msg.Origin))
					//eller är det msg.liteNode.id i &finger!?
					go func() { transport.node.fingerQ <- msg.liteNode }()
				case "ack":
					go func() { transport.node.responseQ <- msg }()
				case "fingerStart":
					go transport.node.setNetworkFingers(msg)
				}
			}
		}
	}()
}