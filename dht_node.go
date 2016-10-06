package dht

import (
	//"encoding/hex"
	"fmt"
	"time"
)

//const bits int = 3

type Contact struct {
	ip   string
	port string
}

type DHTNode struct {
	nodeId      string
	successor   *tinyNode
	predecessor *tinyNode
	contact     Contact
	fingers     *FingerTable
	transport   *Transport
	responseQ   chan *Msg
	TaskQ       chan *Task
}

type tinyNode struct {
	nodeId string
	adress string
}

type Task struct {
	message *Msg
	Type    string
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

	dhtNode.successor = &tinyNode{dhtNode.nodeId, ip + ":" + port}
	dhtNode.predecessor = nil
	dhtNode.fingers = new(FingerTable)
	//KOMMENTERA DETTA SEN
	dhtNode.fingers.nodefingerlist = [bits]*DHTNode{}
	dhtNode.createTransport()
	dhtNode.responseQ = make(chan *Msg)
	dhtNode.TaskQ = make(chan *Task)
	return dhtNode
}

func (dhtNode *DHTNode) createTransport() {
	dhtNode.transport = &Transport{dhtNode, dhtNode.contact.ip + ":" + dhtNode.contact.port, nil}
	dhtNode.transport.msgQ = make(chan *Msg)
	dhtNode.transport.initmsgQ()
}

func (dhtNode *DHTNode) join(master *tinyNode) {
	src := dhtNode.contact.ip + ":" + dhtNode.contact.port
	message := message("join", src, master.adress, src, dhtNode.nodeId, nil)
	dhtNode.transport.send(message)
	for {
		select {
		case r := <-dhtNode.responseQ:
			dhtNode.successor.adress = r.Src
			dhtNode.successor.nodeId = r.Key
			return
			//fmt.Println(dhtNode.nodeId, dhtNode.successor)
		}
	}
}

//Ligger du mellan noderna, nej. skicka join msg till nästa nod och kolla
//om han ligger mellan den noden och hans successor.
func (node *DHTNode) findSucc(msg *Msg) {
	var a = between([]byte(node.nodeId), []byte(node.successor.nodeId), []byte(msg.Key))
	if a {

		node.transport.send(message("response", msg.Dst, msg.Origin, node.successor.adress, node.successor.nodeId, nil))
		node.successor.adress = msg.Origin
		node.successor.nodeId = msg.Key
	} else {
		node.transport.send(message("join", msg.Origin, node.successor.adress, msg.Dst, msg.Key, nil))

	}
}

func (node *DHTNode) printNetworkRing(msg *Msg) {
	if msg.Origin != msg.Dst {

		fmt.Println(node.nodeId, node.successor)
		node.transport.send(printMessage(msg.Origin, node.successor.adress))
	}
}

func (dhtNode *DHTNode) start_server() {
	go dhtNode.initTaskQ()
	go dhtNode.stableTimmer()
	go dhtNode.transport.listen()
}

func (dhtNode *DHTNode) notifyNetwork(msg *Msg) {
	if (dhtNode.predecessor.adress == "") || between([]byte(dhtNode.predecessor.nodeId), []byte(dhtNode.nodeId), []byte(msg.Key)) {
		dhtNode.predecessor.adress = msg.Src
		dhtNode.predecessor.nodeId = msg.Key
	}
}

func (node *DHTNode) initTaskQ() {
	go func() {
		for {
			select {
			case t := <-node.TaskQ:
				switch t.Type {
				case "printRing": //test case
					node.printNetworkRing(t.message)
					//node.improvePrintRing(node.msg)
					//transport.send(&Msg{"printRing", "", v.Src, []byte("tjuuu")})
				case "join":
					node.findSucc(t.message)

				case "stabilize":
					fmt.Println("stabilize case: ", node.nodeId)
					node.stabilize()
				}
			}
		}
	}()
}

func (dhtnode *DHTNode) stabilize() {
	nodeAdress := dhtnode.contact.ip + ":" + dhtnode.contact.port
	predOfSucc := getNodeMessage(nodeAdress, dhtnode.successor.adress) // id eller adress?
	go func() { dhtnode.transport.send(predOfSucc) }()
	time := time.NewTimer(time.Millisecond * 3000)
	for {
		select {
		case r := <-dhtnode.responseQ:
			fmt.Println("case 1 stab: ")
			between := (between([]byte(dhtnode.nodeId), []byte(dhtnode.successor.nodeId), []byte(r.Key))) && r.Key != " " /*) && msg.Key != "" )*/ 
			if between {
				dhtnode.successor.adress = r.Src //origin eller source
				//dhtnode.successor.adress = msg.Origin
				//dhtnode.successor.nodeId = msg.Key
				dhtnode.successor.nodeId = r.Key
				fmt.Println("beetween")
				return
			}
			//ska notifymessage ha fler variabler?
			N := notifyMessage(nodeAdress, dhtnode.successor.adress)

			go func() {
				dhtnode.transport.send(N)
			}()
			fmt.Println("dhtnode id:", dhtnode.nodeId, "dhtnode successor id:", dhtnode.successor, "dhtnode predecessor id:", dhtnode.predecessor)
			return
		case timer := <-time.C: //timer
			fmt.Println("TIMER ERROR:", timer)
			return
		}
	}
}

func (dhtnode *DHTNode) stableTimmer() {
	for {
		time.Sleep(time.Millisecond * 1000)
		dhtnode.createNewTask(nil, "stabilize")
	}
}

func (dhtnode *DHTNode) createNewTask(msg *Msg, typeOfTask string) {
	task := &Task{msg, typeOfTask}
	dhtnode.TaskQ <- task
}

func (node *DHTNode) setSucc(msg *Msg) {
		node.successor.adress = msg.Src
		node.successor.nodeId = msg.Key
}

func (node *DHTNode) setPred(msg *Msg) {
		node.predecessor.adress = msg.Src
		node.predecessor.nodeId = msg.Key
}

func (node *DHTNode) getPred(msg *Msg){
	go func () {
		responseMsg := responseMsg(msg.Dst, msg.Src, node.predecessor.adress, node.predecessor.nodeId)
		node.transport.send(responseMsg)
	}()
}