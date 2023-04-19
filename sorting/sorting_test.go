package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {

	array := []int{2, 7, 4, 1, 5, 3}

	result2 := bubleSort(array)

	fmt.Println("Array", result2)

	result := []int{1, 2, 3, 4, 5, 7}

	if reflect.DeepEqual(result, result2) {
		t.Log("array is sorted")
	} else {
		t.Log("array  is not sorted")
	}

}
