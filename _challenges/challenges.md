```go

// go coding challenges from code challenge sites:

package main

import (
	"fmt"
	"strings"
	// module/package
	// "learn/even"
	// eve "learn/even/Even" // not a thing!!
)

/*
			spinning		words
# stop gninnips my sdrow!

Write a function

that takes in
	a string of one or more words,

and returns
	the same string,
	but with all words that have (five or more letters) __reversed__ (just like the name of this kata).

Strings passed in will consist of only letters and spaces.

Words will be separated by exactly one space.

There will be no leading or trailing spaces.

Examples:

"Hey fellow warriors"  --> "Hey wollef sroirraw"
"This is a test        --> "This is a test"
"This is another test" --> "This is rehtona test"
StringsAlgorithms

*/

/*
# solution:

# get each word
	- each word is separated with one space.
# get each word.length
# word_length >= 5

if word_length >= 5 {
	// reverse the word
}

# also, we want to reverse the word, inside the str

# so: we really reversing a range withing the str

*/


var println = fmt.Println

func SpinWords(str string) string {

	// length of chars in original string
	var size = len([]rune(str))

	// if string is empty
	if size == 0 { return "" }

	var lastIndex = size - 1

	// slice of strings
	var words []string

	// loop over
	var word []rune
	for i, char := range str {

		// skip:: space
		if char == 32 {
			// push and reset word
			words = append(words, string(word))
			word = []rune{}
			continue// skip space
		}

		word = append(word, char)

		// also push a word, on last index
		if i == lastIndex {
			words = append(words, string(word))
			word = []rune{}
		}
	}

	for i, str := range words {
		// if: str.length >= 5 => reverse it
		if len([]rune(str)) >= 5 {
			words[i] = reverseStr(str)
		}
	}

	return strings.Join(words, " ")
}

// reverse the string using two pointer
func reverseStr(str string) string {
	var slice = []rune(str)
	var size = len(slice)

	// 0 - last_index
	var l, r = 0, size - 1

	// loop, until pointers pass each other int the center
	for l < r {
		// swap l and r
		slice[l], slice[r] = slice[r], slice[l]

		// then: move pointers towards each other
		l++
		r--
	}

	return string(slice)
}

func main() {

	res := SpinWords("Hey fellow warriors")
	println(res)

	res2 := SpinWords("This is a test")
	println(res2)

	res3 := SpinWords("This is another test")
	println(res3)

	// var r = []rune{}
	// println(r)

	// now this has zero values
	// var r = make([]rune, 1)
	// println(r)
}

```
