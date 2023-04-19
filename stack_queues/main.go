package main

import "fmt"

type stack struct {
	items []int
}

type queue struct {
	items []int
}

func (q *queue) Enqueue(v int) {

	q.items = append(q.items, v)

}

func (q *queue) Dequeue() int {
	itemsToBeRemoved := q.items[0]
	q.items = q.items[1:]
	return itemsToBeRemoved
}

func (s *stack) Push(v int) {

	s.items = append(s.items, v)

}

func (s *stack) Pop() int {
	l := len(s.items) - 1

	toBeRemoved := s.items[l]
	s.items = s.items[:l]
	return toBeRemoved
}

func main() {

	// stack1 := &stack{}
	// for i := 1; i < 5; i++ {
	// 	stack1.Push(i)
	// }
	// fmt.Println(stack1)

	// fmt.Println("Printing popped items", stack1.Pop())
	// fmt.Println("Printing stack", stack1)

	q := &queue{}

	for i := 1; i < 5; i++ {
		q.Enqueue(i)
	}
	fmt.Println("Printing queue", q)

	fmt.Println("Printing dqueu ite", q.Dequeue())
	fmt.Println("Printing queue", q)
}
