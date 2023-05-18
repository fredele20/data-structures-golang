package main

import "fmt"


type Stack struct {
	items []int
}

// Push adds a new item at the end of the stack
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop removes item at the end of the stack
// And RETURNS the removed item
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	removedItem := l
	s.items = s.items[:l]

	return removedItem
}


// Queues
type Queue struct {
	items []int
}

// Enqueue add a new item to the end of the list
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue removes the first item from the list
func (q *Queue) Dequeue() int {
	removedItem := q.items[0]
	q.items = q.items[1:]

	return removedItem
}


func main() {

	fmt.Println("Stacks")
	myStack := Stack{}

	fmt.Println(myStack)

	fmt.Println("Push to stack")
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)

	fmt.Println(myStack)

	fmt.Println("Pop off the stack")
	myStack.Pop()
	fmt.Println(myStack)

	fmt.Println("==========================================")

	fmt.Println("Queues")
	myQueue := Queue{}
	fmt.Println(myQueue)

	fmt.Println("Enqueue")
	myQueue.Enqueue(500)
	myQueue.Enqueue(600)
	myQueue.Enqueue(700)

	fmt.Println(myQueue)

	fmt.Println("Dequeue")
	myQueue.Dequeue()
	myQueue.Dequeue()
	fmt.Println(myQueue)



}