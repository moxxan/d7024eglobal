package dht

import ()

type Msg struct {
	Key   string //värdet
	Src   string //från noden som kalla
	Dst   string //destinationsadress
	Bytes []byte //transport funktionen, msg.Bytes
	//Type  string // type of message thats is being sent
}

/*func initPringRingMessage(dst, src string) *Msg {
	msg := Msg{}
	msg.Type = "printRing"
	msg.Key = ""
	msg.Src = ""
	msg.Dst = dst
	msg.Bytes = nil

	return *msg
}*/

//func init
