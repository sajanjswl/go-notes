package main

import (
	"fmt"
	"sort"
	"sync"
)

func main() {

	// ch := make(chan int, 1)
	// ch2 := make(chan int)

	// var wg *sync.WaitGroup

	// wg.Add(1)

	// for i := 1; i <= 100; i++ {

	// 	if i%2 == 0 {
	// 		ch <- i
	// 	} else {
	// 		ch2 <- i
	// 	}

	// }

	// go print(wg, ch, ch2)
	// wg.Wait()

	arr := []int{2, 2, 1, 2, 1, 1, 1, 5, 1, 5, 10}
	var arr2 []int
	fmt.Println(frequencySort(arr, arr2))

}

func print(wg *sync.WaitGroup, ch <-chan int, ch2 <-chan int) {

	select {

	case even := <-ch:
		fmt.Println(even)

	case odd := <-ch2:
		fmt.Println(odd)
	}

	defer wg.Done()

}

//

// [2, 2, 1,2,1,1,1,5,1,5,10]

// [10, 5 , 2, 1]

func frequencySort(a, b []int) []int {

	hashMap := make(map[int]int)

	for _, v := range a {

		value, found := hashMap[v]

		if !found {
			hashMap[v] = 1
		} else {

			hashMap[v] = value + 1
		}

	}

	for value, _ := range hashMap {

		b = append(b, value)
	}

	sort.Ints(b)

	return b

}
