package main

import "fmt"

// 1. Stack, Queue, fibonacci, palindromeCheck, Prime Numbers

var arr []int

func GetNthFibonacii(n int) int {

	if n == 2 {
		return 1
	}
	if n == 1 {
		return 0
	}

	// fmt.Println(n)
	return GetNthFibonacii(n-1) + GetNthFibonacii(n-2)

}

func palindromeCheck(str string) bool {

	strarray := []rune(str)

	for i, j := 0, len(strarray)-1; i <= j; {

		if strarray[i] == strarray[j] {
			i++
			j--
		} else {
			return false
		}

	}

	return true
}

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

func isPrime(n int) bool {

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {

	// fibonacci
	// for i := 1; i < 10; i++ {
	// 	arr = append(arr, GetNthFibonacii(i))
	// }
	// fmt.Println(arr)

	// fmt.Println(palindromeCheck("ABAC"))

	// stack1 := stack{}

	// for i := 1; i <= 10; i++ {
	// 	stack1.Push(i)

	// }
	// fmt.Println("Printing stack", stack1)
	// fmt.Println("Printing popped Item", stack1.Pop())
	// fmt.Println("printing stack after poping", stack1)

	// Queue1 := Queue{}
	// for i := 1; i <= 10; i++ {
	// 	Queue1.Add(i)
	// }
	// fmt.Println("Printing stack", Queue1.array)
	// fmt.Println("Printing popped Item", Queue1.DeQueue())
	// fmt.Println("printing stack after poping", Queue1.array)

	var primeArray []int

	for i := 2; len(primeArray) < 10; i++ {

		if isPrime(i) {
			primeArray = append(primeArray, i)
		}

	}

	fmt.Println("Printing Prime Array", primeArray)
}
