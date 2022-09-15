package strings

import "fmt"

// Array
// * Array are of fixed size ex type Array []int
// * The most common purpose of Array  in Go is to hold storage for a slice


// Slice
// * A slice is a data structure describing a contiguous section of an array stored separately from the slice variable itself
// * A slice is not an array
// * A slice describes a piece of an array
// * It’s important to understand that even though a slice contains a pointer, it is itself a value.
// * Under the covers, it is a struct value holding a pointer and a length
// *  It is not a pointer to a struct.
// *  slice := make([]int, 10, 15)
// * copy(newSlice, slice)

// Illustration Hearder not visible of slice
type sliceHeader struct {
	Length        int
	ZerothElement *byte
}

// slice := sliceHeader{
//     Length:        50,
//     ZerothElement: &buffer[100],
// }




// Strings
// * In Go, a string is in effect a read-only slice of bytes
// * A "string literal" is a sequence of characters from the source character set enclosed in double quotation marks
// * Unicode UT8-encoded
// *  fmt.Printf(“%c”,str[I]).   For printing charter of a string
// * Its risky to access  characters like this because a character can be of more than one but
// * fmt.Sprintf("%s %s", string1, string2)
// * String are immutable  to mutate a string first we have to convert it into a slice of rune and mutate it and convert back to string


// Note: Difference between a byte, a character, and a rune, the difference between Unicode and UTF-8, the difference between a string and a string literal.
// Symbol:   A 65(dec)   41(Hex) ,     Z   90(Dec)    5A(Hex)
// Symbol:   a  97(dec)    61(Hex),      z.   122(Dex)  7A(Hex)



// Rune
// * A rune is a builtin type in Go and it's the alias of int32. Rune represents a Unicode code point in Go
// * It doesn't matter how many bytes the code point occupies, it can be represented by a rune
// * String to rune    runes := []rune(str)
// * 'a'   any valid unicode character within single quote is a rune

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}




