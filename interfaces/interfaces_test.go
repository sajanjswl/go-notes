package interfaces

import (
	"fmt"
	"testing"
)

func TestArea(t *testing.T) {

	c := circle{
		radius: 5,
	}

	r := rectangle{
		length: 4,
		bradth: 5,
	}
	var mystring myString = "StirngValue"
	fmt.Printf("area of a cirlce of radius %f is %f \n", c.radius, c.area())

	fmt.Printf("area of a rectangle of legth %f and breadth %f is %f \n", r.bradth, r.length, r.area())

	fmt.Println("lenth of string type message", mystring.area())

	fmt.Println(assert(c))
	fmt.Println(assert(r))
	fmt.Println(assert(mystring))

}
