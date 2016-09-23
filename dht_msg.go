package dht

import (

)
type Msg struct {

	Key string	//värdet
	Src string	//från noden som kalla
	Dst string //destinationsadress
	Bytes []byte //transport funktionen, msg.Bytes

}


