package main

type stack struct {
	array []int
}

func (s *stack) Push(a int) {
	s.array = append(s.array, a)
}
func (s *stack) Pop() int {
	popedItem := s.array[len(s.array)-1]

	s.array = s.array[:len(s.array)-1]
	return popedItem
}

type Queue struct {
	array []int
}

func (q *Queue) Add(n int) {
	q.array = append(q.array, n)
}

func (q *Queue) DeQueue() int {
	var dequedItem int
	if len(q.array) >= 0 {
		dequedItem = q.array[0]
		q.array = q.array[1:]
	}
	return dequedItem
}

func main() {

}
