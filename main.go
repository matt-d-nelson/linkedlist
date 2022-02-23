package main

import (
	"net/http"

	"github.com/matt-d-nelson/linkedlist/linkedlist"
)

func main() {
	var coll linkedlist.LinkedList
	api := &linkedlist.APIQueue{
		Store: &coll,
	}
	http.ListenAndServe(":8080", api)
}
