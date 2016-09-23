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