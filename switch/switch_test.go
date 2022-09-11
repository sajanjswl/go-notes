package switch1

import (
	"fmt"
	"testing"
)

func TestSwitch(t *testing.T) {

	test_cases := []string{"Hello", "World", "1"}

	if Switch(test_cases[0]) == "Hello World from Switch" {
		fmt.Println("Hello Passed fmt")
		t.Log("Hello Passed")
	}

	if Switch(test_cases[2]) == "Hello Switch World" {
		t.Log("World Passed")
	}

}
