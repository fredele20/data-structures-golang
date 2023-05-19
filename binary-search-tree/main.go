package main

import "fmt"

var count int

// Node represents the components of a binary search tree
type Node struct {
	Key int
	Left *Node
	Right *Node
}

// Insert will add a node to the tree
// The key to add should not be already in the tree
func (n *Node) Insert(key int) {
	if n.Key < key {
		// move right
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		// move left
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

// Search will take in a key value
// and RETURN true if there is a node with that value
func (n *Node) Search(key int) bool {

	count++

	if n == nil {
		return false
	}

	if key > n.Key {
		return n.Right.Search(key)
	} else if key < n.Key {
		return n.Left.Search(key)
	}

	return true
}


func main() {
	tree := &Node{Key: 100}
	tree.Insert(48)
	tree.Insert(200)
	tree.Insert(300)
	tree.Insert(50)
	tree.Insert(39)
	tree.Insert(20)

	fmt.Println(tree.Search(87))

	fmt.Println(count)

	fmt.Println(tree)
}