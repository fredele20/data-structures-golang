package main

import "fmt"


type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
	length int
}

func (l *LinkedList) prepend(node *Node) {
	second := l.head
	l.head = node
	l.head.next = second
	l.length++
}

func (l LinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Printf("\n")
}

func (l *LinkedList) deleteWithValue(value int) {
	if l.length == 0 {
		fmt.Println("There are no list in the linked list")
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
	
}


func main() {
	myList := LinkedList{}

	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}
	node4 := &Node{data: 25}
	node5 := &Node{data: 20}
	node6 := &Node{data: 2}


	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	// fmt.Println(myList)
	myList.printListData()
	myList.deleteWithValue(100)
	myList.printListData()
	myList.deleteWithValue(2)
	myList.printListData()
	emptyList := LinkedList{}
	emptyList.printListData()

}