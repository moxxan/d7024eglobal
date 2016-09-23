package dht
//go test -test.run TestDHT1

import (
	"fmt"
	"testing"
)

func TestDHT1(t *testing.T) {
	id0 := "00"
	id1 := "01"
	id2 := "02"
	id3 := "03"
	id4 := "04"
	id5 := "05"
	id6 := "06"
	id7 := "07"

	id8 := "08"
	id9 := "09"
	id10 := "10"
	id11 := "11"
	id12 := "12"
	id13 := "13"
	id14 := "14"



	node0b := makeDHTNode(&id0, "localhost", "1111")
	node1b := makeDHTNode(&id1, "localhost", "1112")
	node2b := makeDHTNode(&id2, "localhost", "1113")
	node3b := makeDHTNode(&id3, "localhost", "1114")
	node4b := makeDHTNode(&id4, "localhost", "1115")
	node5b := makeDHTNode(&id5, "localhost", "1116")
	node6b := makeDHTNode(&id6, "localhost", "1117")
	node7b := makeDHTNode(&id7, "localhost", "1118")


	node0b.addToRing(node1b)
	node1b.addToRing(node2b)
	node1b.addToRing(node3b)
	node1b.addToRing(node4b)
	node2b.addToRing(node5b)
	node3b.addToRing(node6b)
	node3b.addToRing(node7b)
	

	node2b.printTable()
	node2b.printRing()
	fmt.Println("print ring:")

	//PÅHITTAT FÖR EXPERIMENTERING
	node8b := makeDHTNode(&id8, "localhost", "1119")
	node9b := makeDHTNode(&id9, "localhost", "1120")
	node10b := makeDHTNode(&id10, "localhost", "1121")
	node11b := makeDHTNode(&id11, "localhost", "1122")
	node12b := makeDHTNode(&id12, "localhost", "1123")
	node13b := makeDHTNode(&id13, "localhost", "1124")
	node14b := makeDHTNode(&id14, "localhost", "1125")

	
//PÅHITTAT FÖR EXPERIMENTERING
	node3b.addToRing(node8b)
	node2b.addToRing(node9b)
	node2b.addToRing(node10b)
	node5b.addToRing(node11b)
	node9b.addToRing(node12b)
	node12b.addToRing(node13b)
	node13b.addToRing(node14b)
	
fmt.Println("ny table:-------------")
node2b.printTable()
node2b.printRing()



	//fmt.Println("-> ring structure")
	

/*node0 := makeDHTNode(&id0,"","")
node1 := makeDHTNode(&id1,"","")
node2 := makeDHTNode(&id2,"","")
node3 := makeDHTNode(&id3,"","")
node4 := makeDHTNode(&id4,"","")
node5 := makeDHTNode(&id5,"","")
node6 := makeDHTNode(&id6,"","")
node7 := makeDHTNode(&id7,"","")

	

node0.addToRing(node1)
node1.addToRing(node2)
node1.addToRing(node3)
node1.addToRing(node4)
node4.addToRing(node5)
node3.addToRing(node6)
node3.addToRing(node7)
*/

/*
har skrivit ett testcase, där vi lägger in 8 noder o skriver ut en table
lägger sedan till 8 nya noder och skriver ut den nya tabeln, denna nya table är fel.
varför? felsök och kolla upp detta med ragget...
*/



//fmt.Println(node2b.acceleratedLookupUsingFingers(node1b.nodeId))


//	fmt.Println(node10b.successor)
//		node10b.testCalcFingers(0, 4)
/*	node3b.testCalcFingers(1, 3)
	node3b.testCalcFingers(2, 3)
	node3b.testCalcFingers(3, 3)
*/
}

func TestDHT2(t *testing.T) {
	node1 := makeDHTNode(nil, "localhost", "1111")
	node2 := makeDHTNode(nil, "localhost", "1112")
	node3 := makeDHTNode(nil, "localhost", "1113")
	node4 := makeDHTNode(nil, "localhost", "1114")
	node5 := makeDHTNode(nil, "localhost", "1115")
	node6 := makeDHTNode(nil, "localhost", "1116")
	node7 := makeDHTNode(nil, "localhost", "1117")
	node8 := makeDHTNode(nil, "localhost", "1118")
	node9 := makeDHTNode(nil, "localhost", "1119")

//	key1 := "2b230fe12d1c9c60a8e489d028417ac89de57635"
//	key2 := "87adb987ebbd55db2c5309fd4b23203450ab0083"
//	key3 := "74475501523a71c34f945ae4e87d571c2c57f6f3"

	node1.addToRing(node2)
	node1.addToRing(node3)
	node1.addToRing(node4)
	node4.addToRing(node5)
	node3.addToRing(node6)
	node3.addToRing(node7)
	node3.addToRing(node8)
	node7.addToRing(node9)

//	fmt.Println("TEST: " + node1.lookup(key1).nodeId + " is responsible for " + key1)
//	fmt.Println("TEST: " + node1.lookup(key2).nodeId + " is responsible for " + key2)
//	fmt.Println("TEST: " + node1.lookup(key3).nodeId + " is responsible for " + key3)


node1.start_server()
node2.start_server()

src := node1.contact.ip +":" + node1.contact.port
dst := node2.contact.ip +":" + node2.contact.port

node1.transport.send(&Msg{"hello",src,dst,[]byte("hello world")})
//node2.transport.send(&Msg{"nod2 sending",dst,src,[]byte("node2 rocky balboa")})
node3.transport.listen()

	

//	fmt.Println("-> ring structure")
//	node1.printRing()




}

func TestDHT4(t *testing.T) {
        id1 := "01"
        id8 := "08"
        id32 := "32"
        id67 := "67"
        id72 := "72"
        id82 := "82"
        id86 := "86"
        id87 := "87"

        node1 := makeDHTNode(&id1, "localhost", "1111")
        node8 := makeDHTNode(&id8, "localhost", "1112")
        node32 := makeDHTNode(&id32, "localhost", "1113")
        node67 := makeDHTNode(&id67, "localhost", "1114")
        node72 := makeDHTNode(&id72, "localhost", "1115")
        node82 := makeDHTNode(&id82, "localhost", "1116")
        node86 := makeDHTNode(&id86, "localhost", "1117")
        node87 := makeDHTNode(&id87, "localhost", "1118")


	node87.addToRing(node1)
	node87.addToRing(node8)
	node8.addToRing(node32)
	node8.addToRing(node67)
	node67.addToRing(node72)
	node32.addToRing(node82)
	node1.addToRing(node86)

	fmt.Println("-> ring structure")
       	node67.printRing()


       	node67.printTable()
      // 	fmt.Println(node82.lookup("210").nodeId)

       	//fmt.Println(node1.acceleratedLookupUsingFingers("05").nodeId)

}
