package structs

import "testing"

func TestDetails(t *testing.T) {

	author1 := author{
		"Naveen",
		"Ramanathan",
		"Golang Enthusiast",
	}

	blogPost1 := blogPost{
		"Inheritance in Go",
		"Go supports composition instead of inheritance",
		author1,
	}
	blogPost1.Details()

	websites := website{
		blogPosts: []blogPost{
			{"dfd",
				"dfdsfs",
				author1,
			},
		},
	}

	websites.Contents()

}
