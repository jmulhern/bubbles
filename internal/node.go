package internal

import (
	"math/rand"
	"time"
)

type Node struct {
	value int
	next  *Node
}

func RandomLinkedList(size int) *Node {
	rand.Seed(time.Now().UnixNano())
	current := &Node{value: rand.Intn(100)}
	for i := 0; i < size-1; i++ {
		current = &Node{value: rand.Intn(100), next: current}
	}
	return current
}

func PrintLinkedList(head *Node) {
	current := head
	for current != nil {
		print(current.value, " ")
		current = current.next
	}
	println()
}
