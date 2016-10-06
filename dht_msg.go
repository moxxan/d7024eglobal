package dht

import ()

type Msg struct {
	Origin string
	Key    string //värdet
	Src    string //från noden som kalla
	Dst    string //destinationsadress
	Bytes  []byte //transport funktionen, msg.Bytes
	Adress string
	Id	   string
	Type   string // type of message thats is being sent
}

func message(t, origin, dst, src, key string, bytes []byte) *Msg {
	msg := &Msg{}
	msg.Type = t
	msg.Adress = adress
	msg.Id = id
	msg.Origin = origin
	msg.Src = src
	msg.Dst = dst
	msg.Bytes = bytes
	msg.Key = key
	return msg
}

func joinMessage(dst string) *Msg {
	msg := &Msg{}
	msg.Type = "addToRing"
	msg.Adress = adress
	msg.Id = id
	msg.Origin = origin
	msg.Src = ""
	msg.Dst = dst
	msg.Bytes = nil
	//msg.Key = key
	return msg
}

func printMessage(origin, dst string) *Msg {
	msg := &Msg{}
	msg.Type = "printRing"
	msg.Adress = adress
	msg.Id = id
	msg.Origin = origin
	msg.Src = ""
	msg.Dst = dst
	msg.Bytes = nil
	//msg.Key = key
	return msg
}

func notifyMessage(src, dst string) *Msg {
	msg := &Msg{}
	msg.Type = "notify"
	msg.Adress = adress
	msg.Id = id
	msg.Origin = ""
	msg.Key = ""
	msg.Src = src
	msg.Dst = dst
	msg.Bytes = nil
	return msg
}

func getNodeMessage(src, dst string) *Msg {
	msg := &Msg{}
	msg.Type = "pred"
	msg.Adress = adress
	msg.Id = id
	msg.Origin = origin
	msg.Src = src
	msg.Dst = dst
	msg.Bytes = nil
	return msg
}

func responseMsg(src, dst, adress, id string) *Msg {
	msg := &Msg{}
	msg.Type = "response"
	msg.Adress = adress
	msg.Id = id
	msg.Origin = origin
	msg.Src = src
	msg.Dst = dst
	msg.Bytes = nil
	return msg}