package main

import (
	"fmt"
)

var head *LinkedList
var tail *LinkedList

func main() {

	// fmt.Println("the char is", findFistNoRepeatingCharacter("aaabsdfsfaa@45"))

	list := new(LinkedList)
	list.Add(10)
	head = list.NextNode
	tail = list.NextNode

	for i := 50; i < 10; i = i + 10 {
		tail.Add(i)
		tail = tail.NextNode
	}

	for pivot := head; pivot != nil; pivot = pivot.NextNode {
		fmt.Println("Printing pivot", pivot)
	}

}

type LinkedList struct {
	Value    int
	NextNode *LinkedList
}

func (list *LinkedList) Add(num int) {

	node := &LinkedList{
		Value: num,
	}
	list.NextNode = node

}

func findFistNoRepeatingCharacter(str string) rune {

	queue := []rune{}
	maphash := map[rune]bool{}

	for _, char := range str {

		_, found := maphash[char]

		if !found {

			queue = append(queue, char)
			maphash[char] = true
		} else {

			if len(queue) != 0 && queue[0] == char {
				queue = queue[1:]
			}
		}

	}

	fmt.Println("printing char", string(queue))

	return queue[0]

}
