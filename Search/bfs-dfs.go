package main

type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) DepthFirstSearch(array []string) []string {

	array = append(array, n.Name)

	for _, n := range n.Children {

		array = n.DepthFirstSearch(array)
	}

	// Write your code here.
	return array
}

func (n *Node) BreadthFirstSearch(array []string) []string {

	queue := []*Node{}
	queue = append(queue, n)

	for len(queue) > 0 {
		currentNode := queue[0]
		array = append(array, currentNode.Name)

		for _, v := range currentNode.Children {
			queue = append(queue, v)
		}

		queue = queue[1:]

	}
	// Write your code here.
	return array
}
