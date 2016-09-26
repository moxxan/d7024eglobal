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
	udpAddr, err := net.ResolveUDPAddr("udp", msg.Dst)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Println("error SEND function is:", err)
	}
	encoded, err := json.Marshal(msg)
	defer conn.Close()
	_, err = conn.Write(encoded)

}

func (transport *Transport) initmsgQ() {
	go func() {
		for {
			select {
			case v := <-transport.msgQ:
				switch v.Key {
				case "hello":
					fmt.Println(string(v.Bytes))
					transport.send(&Msg{"printRing", "", v.Src, []byte("tjuuu")})
				case "reply":
					fmt.Println("hej:", string(v.Bytes))

				case "printRing":
					transport.node.printRing()
					//transport.send(&Msg{"ring", "", v.Src, []byte(transport.node.printRing())})
				}
			}
		}
	}()
}
