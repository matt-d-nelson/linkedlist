//Temp package change
package linkedlist

import (
	"fmt"
)

//-----------------------------OBJECTS-----------------------------//

type Node struct {
	Value int
	Next  *Node
	Prev  *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

//-----------------------------STRING-----------------------------//

func (n Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

//-----------------------------ADD-----------------------------//

func (ll *LinkedList) Add(v int) error {
	node := Node{Value: v}
	if ll.Head == nil {
		ll.Head = &node
		ll.Tail = &node
		return nil
	}
	node.Prev = ll.Tail
	ll.Tail.Next = &node
	ll.Tail = &node
	return nil
}

//-----------------------------INDEX-----------------------------//

func (ll *LinkedList) Index(i int) (int, error) {
	curr := 0
	n := ll.Head
	for ; curr < i && n != nil; curr++ {
		n = n.Next
	}
	if curr != i || n == nil {
		return -1, fmt.Errorf("linkedlist: %d is beyond length of list (%d)", i, curr)
	}
	return n.Value, nil
}

//-----------------------------REVERSE-----------------------------//

func (ll *LinkedList) Reverse() {
	curr := ll.Head
	for curr != nil {
		next := curr.Next
		curr.Next, curr.Prev = curr.Prev, curr.Next
		curr = next
	}
	ll.Head, ll.Tail = ll.Tail, ll.Head
}

//-----------------------------STRING-----------------------------//

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
