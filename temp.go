package main

import "fmt"

func main() {
	node :=&Node{}

	AddDataToList(node)
}

type Node struct {
	Head *Node
	Key  interface{}
	Tail *Node
}

func AddDataToList(node *Node) {

	for i := 1; i < 10; i++ {
		if node.Head == nil {
			node.Key = i
		} else {
			node.Tail = &Node{Key: i}
		}
	}

}

func show(node *Node) {

	for node.Tail != nil {
		fmt.Println(node.Key)
		node.Tail=node.
	}
}
