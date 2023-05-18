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


func main() {
	myList := LinkedList{}

	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	fmt.Println(myList)

}