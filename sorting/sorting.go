package main

import (
	"fmt"
	"strconv"
)

// Table of contents
// 1. Bubble Sort.
// 2. Recursive Bubble Sort.
// 3. Insertion Sort.
// 4. Selection Sort.
// 5. Merge Sort.
// 6. Counting Sort.

// 1. Buble Sort is the simplest sorting algorithm that works by repeatedly swapping the adjacent elements if they are in the wrong order
// 2. Requires 2 forloops to work so the worst time complexity is O(n^2)

func bubleSort(array []int) []int {

	for i := 0; i < len(array)-1; i++ {

		for j := 0; j < len(array)-i-1; j++ {

			if array[j] > array[j+1] {

				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

	return array
}

func quickSort(arr []int, low, high int) {
	if low < high {

		// Find the pivot index of an lower bound of an array
		var pivot = partition(arr, low, high)

		// Apply Divide and conquer stratagy
		// to sort the left side and right side of an array
		// respective to the pivot position

		// Left hand side array
		quickSort(arr, low, pivot)

		// Right hand side array
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {

	// Pick a lowest bound element as an pivot value
	var pivot = arr[low]

	var i = low
	var j = high

	for i < j {

		// Increment the index value of "i"
		// till the left most values should be less than or equal to the pivot value
		for arr[i] <= pivot && i < high {
			i++
		}

		// Decrement the index value of "j"
		// till the right most values should be greater than to the pivot value
		for arr[j] > pivot && j > low {
			j--
		}

		// Interchange the values of present
		// in the index i and j of an array only if i is less than j
		if i < j {
			var temp = arr[i]
			arr[i] = arr[j]
			arr[j] = temp
		}
	}

	// Interchange the element in j's poisition to the lower bound of an array
	// and place the Pivot element to the j's position
	arr[low] = arr[j]
	arr[j] = pivot

	// Finally return the appropriate index position of the pivot element
	return j
}

func printArray(arr []int) {

	// Loop the item and print each item in the console with white space seperated
	for i := 0; i < len(arr); i++ {
		fmt.Print(strconv.Itoa(arr[i]) + " ")
	}

	fmt.Println("")
}

func main() {

	var arr = []int{15, 3, 12, 6, -9, 9, 0}

	fmt.Print("Before Sorting: ")
	printArray(arr)

	quickSort(arr, 0, len(arr)-1)
	fmt.Print("After Sorting: ")
	printArray(arr)
}
