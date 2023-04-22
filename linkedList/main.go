package linkedlist

import "fmt"

// Single Linked List
type Node struct {
	Key      int
	NextNode *Node
}

func NewNode(key int) *Node {

	var node Node
	node.Key = key
	node.NextNode = nil
	return &node
}

func Show(head *Node) {
	temHead := head

	for temHead != nil {
		fmt.Println(temHead.Key)
		temHead = temHead.NextNode
	}

}

func main() {
	head := NewNode(1)
	tempHead := head
	for i := 2; i < 10; i++ {
		if tempHead.NextNode == nil {
			tempHead.NextNode = NewNode(i)
		}
		tempHead = tempHead.NextNode

	}

	Show(head)
}
