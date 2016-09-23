package dht

import (
	"fmt"
	"encoding/hex"
)

//const bits int = 3


type Contact struct {
	ip   string
	port string
}

type DHTNode struct {
	nodeId      string
	successor   *DHTNode
	predecessor *DHTNode
	contact     Contact
	fingers 	*FingerTable
	transport	*Transport
	msg			*Msg
}


func makeDHTNode(nodeId *string, ip string, port string) *DHTNode {
	dhtNode := new(DHTNode)
	dhtNode.contact.ip = ip
	dhtNode.contact.port = port

	if nodeId == nil {
		genNodeId := generateNodeId()
		dhtNode.nodeId = genNodeId
	} else {
		dhtNode.nodeId = *nodeId
	}

	dhtNode.successor = nil
	dhtNode.predecessor = nil

	//KOMMENTERA DETTA SEN!...
	dhtNode.fingers = new(FingerTable)
	dhtNode.fingers.nodefingerlist = [bits]*DHTNode{}

	return dhtNode
}

func (dhtNode *DHTNode) addToRing(newDHTNode *DHTNode) { 
	//KOLLAR FÖRSTA FALLET, connectar 2 noder.
	if (dhtNode.predecessor == nil && dhtNode.successor == nil) {
		newDHTNode.predecessor = dhtNode
		newDHTNode.successor = dhtNode
		dhtNode.successor = newDHTNode
		dhtNode.predecessor = newDHTNode
	
		dhtNode.fingers.nodefingerlist = init_finger_table(dhtNode)  //Findfingers
		newDHTNode.fingers.nodefingerlist = init_finger_table(newDHTNode)
		dhtNode.stabilize(dhtNode.nodeId)
		newDHTNode.stabilize(newDHTNode.nodeId)

	
	//	dhtNode.fingers.Fingers = init_finger_table(dhtNode)
		} else if between([]byte(dhtNode.nodeId), []byte(dhtNode.successor.nodeId), []byte(newDHTNode.nodeId)){
		dhtNode.successor.predecessor = newDHTNode
		newDHTNode.successor = dhtNode.successor
		dhtNode.successor = newDHTNode
		newDHTNode.predecessor = dhtNode
		newDHTNode.fingers.nodefingerlist = init_finger_table(newDHTNode)
		newDHTNode.stabilize(newDHTNode.nodeId)
		//updateFingers(newDHTNode)
		


	}else {
		dhtNode.successor.addToRing(newDHTNode)
	}
	
}

func (dhtNode *DHTNode) lookup(key string) *DHTNode {
	if between([]byte(dhtNode.nodeId), []byte(dhtNode.successor.nodeId), []byte(key)){
		//fmt.Println("node id:",dhtNode.nodeId,"dht successor node id", dhtNode.successor.nodeId," key:", key)
		if(dhtNode.nodeId == key){
			return dhtNode
		} else{
			return dhtNode.successor
		} 
	}else{
		//distance(a, b []byte, bits int) *big.Int
		return dhtNode.successor.lookup(key)
		//return dhtNode.successor.lookup(key)
	}
}

func (dhtNode *DHTNode) acceleratedLookupUsingFingers(key string) *DHTNode {
	for i := len(dhtNode.fingers.nodefingerlist); i > 0; i-- {
		if between([]byte(dhtNode.nodeId), []byte(dhtNode.fingers.nodefingerlist[i-1].nodeId), []byte(key)){
			
		//	fmt.Println(key,"ligger mellan nod",dhtNode.nodeId, "och hans finger", dhtNode.fingers.nodefingerlist[i-1].nodeId)
			
		} else {

			var a = dhtNode.fingers.nodefingerlist[i-1]

		//	fmt.Println(key,"FANNS INTE MELLAN NOD",dhtNode.nodeId, "och hans finger", a.nodeId,
		//	", hoppar till", dhtNode.nodeId, "senast kollade finger", a.nodeId)
			
			return a.acceleratedLookupUsingFingers(key)
		}
		
	}

	return dhtNode // XXX This is not correct obviously

}
func (dhtNode *DHTNode) responsible(key string) bool {
	// TODO
	return false

}
func (dhtNode *DHTNode) printRing() {
	//fmt.Println(dhtNode.nodeId)

	for i := dhtNode; i != dhtNode.predecessor; i = i.successor {
		fmt.Println(i.nodeId)
	}
	fmt.Println(dhtNode.predecessor.nodeId)
	// TODO
}

func (dhtNode *DHTNode) testCalcFingers(m int, bits int) {
	idBytes, _ := hex.DecodeString(dhtNode.nodeId)
	fingerHex, _ := calcFinger(idBytes, m, bits)
	fingerSuccessor := dhtNode.lookup(fingerHex)
	fingerSuccessorBytes, _ := hex.DecodeString(fingerSuccessor.nodeId)
	fmt.Println("successor    " + fingerSuccessor.nodeId)

	dist := distance(idBytes, fingerSuccessorBytes, bits)
	fmt.Println("distance     " + dist.String())
}

func (dhtNode *DHTNode) printTable(){
		for i := 0; i < len(dhtNode.fingers.nodefingerlist); i++ {
		fmt.Println("Node",dhtNode.nodeId,"finger",i+1, "points to", dhtNode.fingers.nodefingerlist[i])
		
	}
}

func (dhtNode *DHTNode) stabilize(node string) {
	//n := dhtNode.successor.predecessor.nodeId
	if dhtNode.successor.nodeId != node {
	updateFingers(dhtNode.successor)
/*		for i := 0; i< bits; i++{
			if upd[i] != nil {

			}
		}
*/		dhtNode.successor.stabilize(node)
	}
/*	var a = between([]byte(dhtNode.nodeId), []byte(dhtNode.successor.nodeId), []byte(n))
		if a {
	//		fmt.Println("stabilize node:",dhtNode.nodeId,"?", "skiten ligger mellan", dhtNode.nodeId, "och", dhtNode.successor.nodeId,"gör inget")
		} else{
	//		fmt.Println("skiten ligger inte mellan, uppdatera fingers.")
			updateFingers(dhtNode)
	}
	*/
}

func (dhtNode *DHTNode) start_server(){
	go dhtNode.transport.listen()
}








/*func updateFingers(node *DHTNode){
	for i := node; i != node.predecessor; i = i.successor{
		i.fingers.nodefingerlist = init_finger_table(i)
	}
}

*/

/*func (dhtNode *DHTNode) find_successor(node *DHTNode) *DHTNode{
	predecessorNode := dhtNode.find_predecessor(node)
	return predecessorNode.successor
}

func (dhtNode *DHTNode) find_predecessor(node *DHTNode) *DHTNode{
	successorNode := dhtNode
	return successorNode

}*/
