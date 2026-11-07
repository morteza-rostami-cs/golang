```Go

// bubble sort in Go

/*
what is a bubble sort?

*/

func bubbleInt(s []int) []int {
	var size = len(s)

	// if a swap happens, set this to true, but: it will reset and the end of each pass i
	var swap bool = false

	// i loos for every element once, until we have no swap for one complete i pass
	for i :=0; i<size; i++ {

		// compares each index with it's neighbor
		for j :=0; j<size; j++ {

			var holder int

			// prevent out of bound
			if j + 1 == size {
				break
			}

			// compare
			if s[j] > s[j+1] {
				// found smaller, swap
				holder = s[j]
				s[j] = s[j+1]
				s[j+1] = holder

				swap = true // one swap happened in this pass
			}
		}

		// out of j loop
		if swap == true {
			// swap happened in j loop
			swap = false // reset
		} else {
			// no swap happened
			println("no swap happened.")
			break
		}
	}

	return s
}

func main() {

	var scores []int = []int{5, 10, 4, 2, 6}

	sorted := bubbleInt(scores)
	println(sorted)
}

```
