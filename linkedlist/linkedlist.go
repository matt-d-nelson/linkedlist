package linkedlist

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
	Prev  *Node
}

func (n Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (ll *LinkedList) Add(v interface{}) {
	node := Node{Value: v}
	if ll.Head == nil {
		ll.Head = &node
		ll.Tail = &node
		return
	}
	node.Prev = ll.Tail
	ll.Tail.Next = &node
	ll.Tail = &node
}

func (ll *LinkedList) Index(i int) (Node, error) {
	curr := 0
	n := ll.Head
	for ; curr < i && n != nil; curr++ {
		n = n.Next
	}
	if curr != i || n == nil {
		return Node{}, fmt.Errorf("linkedlist: %d is beyond length of list (%d)", i, curr)
	}
	return *n, nil
}

func (ll *LinkedList) Reverse() {
	curr := ll.Head
	for curr != nil {
		next := curr.Next
		curr.Next, curr.Prev = curr.Prev, curr.Next
		curr = next
	}
	ll.Head, ll.Tail = ll.Tail, ll.Head
}

func (ll LinkedList) String() string {
	msg := "{"
	for curr := ll.Head; curr != nil; curr = curr.Next {
		if msg != "{" {
			msg += ", "
		}
		msg += curr.String()
	}
	msg += "}"
	return msg
}
