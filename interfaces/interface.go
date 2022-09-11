package interfaces

import "fmt"

// Interfaces
// * Its and abstract data type
// * Holds Method Signature
// * When a type provides definition for all the methods in the interface, it is said to implement the interface.
// * Type variable struct/Interface/string/int data type
// *  Go interfaces are implemented implicitly if a type contains all the methods declared in the interface.
// * If multiple type implements the same interface they become the same type
// * An interface can be thought of as being represented internally by a tuple (type, value)
// * type is the underlying concrete type of the interface and value holds the value of the concrete type
// * All type implements empty interface
// * An interface with no underlying concrete type and value is called zero value interface
// * We can assert type of interface by interfacevariable.(type)

type Geometry interface {
	area() float64
	// perimenter() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	length float64
	bradth float64
}

func (c circle) area() float64 {

	return 3.14 * c.radius * c.radius
}

func (r rectangle) area() float64 {

	return r.length * r.bradth
}

type myString string

func (ms myString) area() float64 {

	return float64(len(ms))
}

// # Polymorphism
// * In GoLang, polymorphism is achieved mainly using interfaces.
// * A type implementing a function defined in interface becomes the type defined as an interface.
// * Polymorphism is used to reduce code in general
// * A single function can be used to do the same thing on multiple different objects.
// * Go not being a strict OO-language achieves polymorphism in an elegant way

func assert(i Geometry) interface{} {

	switch i.(type) {
	case myString:
		return fmt.Sprintf(" the interface is of type myString")
	case circle:
		return fmt.Sprintf("the interface has concreate type as circle")
	case rectangle:
		return fmt.Sprintf("the interface has concreate type as recatnage")
	default:
		return fmt.Sprintf("unknow interface")

	}
}
