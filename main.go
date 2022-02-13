package main

import (
	"fmt"

	"github.com/matt-d-nelson/linkedlist/linkedlist"
)

func main() {
	var ll linkedlist.LinkedList
	fmt.Println(ll.String())
	ll.Add("foo")
	fmt.Println(ll.String())
	ll.Add("bar")
	fmt.Println(ll.String())
	ll.Add(212)
	fmt.Println(ll)

	fmt.Println(ll.Index(2))
	fmt.Println(ll.Index(220))

	ll.Reverse()
	fmt.Println(ll)
}
