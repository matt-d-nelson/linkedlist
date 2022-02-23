//Temp package change
package linkedlist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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

//Add error return
func (ll *LinkedList) Add(v interface{}) error {
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

func (ll *LinkedList) Read(i int) (Node, error) {
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

type Collection interface {
	Add(v interface{}) error
	Read(i int) (Node, error)
	String() string
	Reverse()
}

type APIQueue struct {
	Store Collection
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

func (q *APIQueue) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/add":
		//Changed so all values added are strings
		v := r.URL.Query().Get("value")
		if err := q.Store.Add(v); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"value added": v})
	case "/get":
		idx, err := strconv.Atoi(r.URL.Query().Get("index"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		v, err := q.Store.Read(idx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"value": v.Value})
	case "/listAll":
		msg := q.Store.String()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"list": msg})
	case "/reverse":
		q.Store.Reverse()
		msg := q.Store.String()
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]interface{}{"list reversed": msg})
	default:
		http.Error(w, fmt.Sprintf("path %v undefined", r.URL.Path), http.StatusBadRequest)
	}
}
