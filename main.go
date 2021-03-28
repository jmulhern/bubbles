package main

import . "github.com/jmulhern/bubbles/internal"

func main() {
	head := RandomLinkedList(10)
	PrintLinkedList(head)
	sorter := Sorter{Head: head}
	head = sorter.Bubble()
	PrintLinkedList(head)
}
