package dht

import (
	"net"
	"encoding/json"
	"fmt"
)
type Msg struct {

	Key string	//värdet
	Src string	//från noden som kalla
	Dst string //destinationsadress
	Bytes []byte //transport funktionen, msg.Bytes

}

type Transport struct {
	node *DHTNode
	bindAddress string // rad 20, bindadress måste finnas.
}
	

func (transport *Transport) listen() {
	udpAddr, err := net.ResolveUDPAddr("udp", transport.bindAddress)
	fmt.Println("transport bindaddress:", transport.bindAddress)
	conn, err := net.ListenUDP("udp", udpAddr)
		if err != nil{
		fmt.Println("error LISTEN function is:", err)
	}
	defer conn.Close()
	dec := json.NewDecoder(conn)
	for {
		msg := Msg{}
		err = dec.Decode(&msg)

	}

}

func (transport *Transport) send(msg *Msg) {
	udpAddr, err := net.ResolveUDPAddr("udp", msg.Dst)
	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil{
		fmt.Println("error SEND function is:", err)
	}
	encoded, err := json.Marshal(msg)
	defer conn.Close()
	_, err = conn.Write(encoded)

}



















//funtionen Bytes tar ett msg.Bytes() från transport funktionen, så
//mdstrukten måste encodas till bytes, jag använde json då det är smidigt.

/*
func (msg *Msg) Bytes() []byte {
	encoded, err := json.Marshal(msg)
	if err == nil{
		fmt.Println("encoded value is:", encoded)
		return encoded
	}
	fmt.Println("error BYTES function is:", err)
	return nil
	}
*/