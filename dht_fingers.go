package dht

import (
//	"encoding/hex"
//	"fmt"
)

const bits int = 4


type FingerTable struct {
	nodefingerlist [bits]*DHTNode
}

func (node *DHTNode) setFingers (msg *Msg){
	for i := 0; i < bits; i++ {
		adress := msg.Adress
		id := msg.Id
		node.fingers.nodefingerlist[i] = &FingerTable{adress, id}
		
	}
}


/*
func init_finger_table(n *DHTNode) [bits]*DHTNode{
	var templist [bits]*DHTNode
	for i := 0; i < bits; i++ {
		x,_ := hex.DecodeString(n.nodeId) //func DecodeString(s string) ([]byte, error)
		y, _ := calcFinger(x, (i+1), bits) // returnerar (string, []byte)
		if y == "" {
			y = "00"
		} else{}
		//fmt.Println("FINGER", i+1, "POINTS ON NODE", y)
		succ := n.lookup(y)  //ändra ej, vanlig lookup
		// calcFinger(n []byte, k int, m int) (string, []byte) {
		templist[i] = succ

	}
return templist
}
*/


/*
func updateFingers(node *DHTNode)  [bits]*DHTNode{
	//var templist [bits]*DHTNode
	for i := 0; i < bits; i++ {
		x,_ := hex.DecodeString(node.nodeId)
		y, _ := calcFinger(x, (i+1), bits)
/*		if y == "" {
			y = "00"
		}

		if (y == node.fingers.nodefingerlist[i].nodeId){
		} else {
			a := node.lookup(y)
			node.fingers.nodefingerlist[i] = a
	}


	
}
return node.fingers.nodefingerlist
}
*/