/*
this is just a stupid package for test.
*/
package even

// all files in even/ , should have the same package name

// this is a documentation comment.
func Fn() {
	println("this is Fn")
}

// check to see if an int is even.
// cap names -- are exported, public
func Even(i int) bool {
	return i % 2 == 0
}

// lowercase -- private
func odd(i int) bool {
	// all even numbers are divisible by 2
	// also: all odd numbers, give 1 remainder.
	return i % 2 == 1
}
