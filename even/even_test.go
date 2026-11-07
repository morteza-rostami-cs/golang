package even // package to test

import (
	"fmt"
	"testing"
)

// receives a pointer
func TestEven(t *testing.T) {

	// test an even number
	if !Even(2) {
		// if it was not even
		t.Log("2 should be even!")
		t.Fail() // test failed
	}
}

// go runs this on, go test -- anc compare the output
func ExampleEven() {
	if Even(2) {
		fmt.Printf("Is even\n")
	}

	// Output:
	// Is even
}

// so:: both of above funcs work for testing
