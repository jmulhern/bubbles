package internal

import "testing"

func TestRandomLinkedList(t *testing.T) {
	testCases := []struct {
		size int
		want int
	}{
		{1, 1},
		{10, 10},
		{42, 42},
	}

	for _, tc := range testCases {
		result := RandomLinkedList(tc.size)
		size := 0
		for result != nil {
			size++
			result = result.next
		}
		if got := size; got != tc.want {
			t.Errorf("got %v; want %v", got, tc.want)
		}
	}
}
