package internal

type Sorter struct {
	Head *Node
}

func (s Sorter) Bubble() *Node {
	for {
		current := s.Head
		swapped := false
		for current != nil {
			if current.next != nil && current.value > current.next.value {
				a := current.value
				b := current.next.value
				current.value = b
				current.next.value = a
				swapped = true
			}
			current = current.next
		}
		if swapped == false {
			break
		}
	}
	return s.Head
}
