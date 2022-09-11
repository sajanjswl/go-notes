package structs

import (
	"fmt"
)

// #Composition Instead of Inheritance in golang
// * Whenever one struct field is embedded in another
// * Go gives us the option to access the embedded fields as if they were part of the outer struct
// * This means that p.author.fullName() in  can be replaced with p.fullName().
// * It is not possible to anonymously embed a slice. A field name is required.

type author struct {
	firstName string
	lastName  string
	bio       string
}

type blogPost struct {
	title   string
	content string
	author
}

type website struct {
	blogPosts []blogPost
}

func (a author) fullName() string {
	return fmt.Sprintf("%s %s", a.firstName, a.lastName)
}

func (b blogPost) Details() {
	fmt.Println("Title: ", b.title)
	fmt.Println("Content: ", b.content)
	fmt.Println("Author: ", b.fullName())
	fmt.Println("Bio: ", b.bio)
}

func (w website) Contents() {
	fmt.Println("Contents of Website")
	for _, v := range w.blogPosts {
		v.Details()
		fmt.Println()
	}
}
