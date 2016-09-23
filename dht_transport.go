package dht

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



func (transport *Transport) initmsgQ{
	for(){
		select{
			case 
		}
	}
}