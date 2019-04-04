package helpers

import (
	"fmt"
	"math/rand"
	"time"
)

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// List ....
type List struct {
	Length int
	Start  *ListNode
}

// GenerateListnode - Generate List containing n items
func GenerateListnode(n int) *ListNode {
	rand.Seed(time.Now().UTC().UnixNano())

	head := &ListNode{}
	current := head

	for i := 1; i < n; i++ {
		current.Next = &ListNode{Val: getRandInt()}
		current = current.Next
	}

	return head
}

func getRandInt() int {
	return rand.Intn(10)

}

// Print - Print the list in a helpful way
func (l *ListNode) Print() {
	for i := l; i != nil; i = i.Next {
		fmt.Print("[", i.Val, "]")

		if i.Next != nil {
			fmt.Print("->")
		} else {
			fmt.Println()
		}
	}

}
