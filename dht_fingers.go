package dht

import (
	"encoding/hex"
	"fmt"
)

const bits int = 7


type FingerTable struct {
	nodefingerlist [bits]*DHTNode
}


func init_finger_table(n *DHTNode) [bits]*DHTNode{
	var templist [bits]*DHTNode
	for i := 0; i < bits; i++ {
		x,_ := hex.DecodeString(n.nodeId) //func DecodeString(s string) ([]byte, error)
		y, _ := calcFinger(x, (i+1), bits) // returnerar (string, []byte)
		//fmt.Println("NODE:",n.nodeId ,"Y is now:",y)
		if y == "" {
			y = "00"
		} else{}
		//fmt.Println("FINGER", i+1, "POINTS ON NODE", y)
		succ := n.lookup(y)  //ändra ej, vanlig lookup
		// calcFinger(n []byte, k int, m int) (string, []byte) {
		templist[i] = succ

	}
	//fmt.Println("TEEMPLIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIST", templist)
	//fmt.Println(templist)
	return templist
}



func updateFingers(node *DHTNode)  [bits]*DHTNode{
	//var templist [bits]*DHTNode
	for i := 0; i < bits; i++ {
		x,_ := hex.DecodeString(node.nodeId)
		y, _ := calcFinger(x, (i+1), bits)
//		fmt.Println("NODE:",node.nodeId ,"FINGER2 is now:",y)
/*		if y == "" {
			y = "00"
		}
*/
		if (y == node.fingers.nodefingerlist[i].nodeId){
			//fmt.Println(y,"=", node.fingers.nodefingerlist[i].nodeId)
		} else {
			
			//fmt.Println(y,"!=", node.fingers.nodefingerlist[i].nodeId)
			//fmt.Println("replacing y")
			a := node.lookup(y)
			if (i == 3 && node.nodeId == "67"){
				fmt.Println("a:",a)

//				fmt.Println("NODE:",node.nodeId ,"FINGER2 is now:",y)
		//	fmt.Println("A = ",a)	
			}
			

		//	fmt.Println("a = ",a)
			node.fingers.nodefingerlist[i] = a
	}


	
}
return node.fingers.nodefingerlist
}

/*func update_finger_table(s int, i int){
}
*/

