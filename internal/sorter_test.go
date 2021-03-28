package internal

import "testing"

func TestSorter_Bubble(t *testing.T) {
	testCases := []struct {
		head *Node
		want []int
	}{
		{&Node{value: 3, next: &Node{value: 2, next: &Node{value: 1}}}, []int{1, 2, 3}},
		{&Node{value: 99, next: &Node{value: -12, next: &Node{value: 4}}}, []int{-12, 4, 99}},
		{&Node{value: 20, next: &Node{value: 7, next: &Node{value: 7}}}, []int{7, 7, 20}},
	}
	for _, tc := range testCases {
		subject := Sorter{Head: tc.head}
		result := subject.Bubble()
		i := 0
		for result != nil {
			if got := result.value; got != tc.want[i] {
				t.Errorf("got %v; want %v", got, tc.want[i])
			}
			result = result.next
			i++
		}
	}
}
