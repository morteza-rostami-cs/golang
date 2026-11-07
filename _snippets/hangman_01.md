```go
package main

import (
	"bufio"
	"fmt"
	"math/rand"

	// "reflect"
	"slices"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	// "unsafe"
	"os"
)

var println = fmt.Println

var hangman = []string{
    // 0 wrong guesses
    `
  +---+
  |   |
      |
      |
      |
      |
=========`,
    // 1
    `
  +---+
  |   |
  O   |
      |
      |
      |
=========`,
    // 2
    `
  +---+
  |   |
  O   |
  |   |
      |
      |
=========`,
    // 3
    `
  +---+
  |   |
  O   |
 /|   |
      |
      |
=========`,
    // 4
    `
  +---+
  |   |
  O   |
 /|\  |
      |
      |
=========`,
    // 5
    `
  +---+
  |   |
  O   |
 /|\  |
 /    |
      |
=========`,
    // 6 — dead
    `
  +---+
  |   |
  O   |
 /|\  |
 / \  |
      |
=========`,
}

// secret words
var secrets = []string{
	"cat",
	"garden",
	"house",
	"fuck",
	"rock",
}

// ================= functions

// generate a random int within a range
func randomRange(max int, min int) (int) {

	// handle min, max -- instead of max, min
	if min > max {
		min, max = max, min
	}

	// there is no range
	if min == max {
		return min
	}

	// rand.Intn(n) -- returns a random int in [0, n) -- so: 0 up to n-1, never n
	// it panics if n <= 0 -- +1 ....
	// min= 2, max= 5 -- 5 - 2 = 3 -- then generate a random from: 0 to 2 + 2 = 4
	// so: max result is n-1
	// if we need 5 included: rand.Intn(max - min +1) + min
	return rand.Intn(max - min + 1) + min
}

// generate the placeholder string, based on secret
func createPlaceholder(secret string) ([]rune) {
	var size int = utf8.RuneCountInString(secret)
	// placeholder := []rune(secret)
	var placeholder []rune = []rune{}

	for i := 0; i < size; i++ {
		placeholder = append(placeholder, '_')
	}
	return placeholder
}

// get a random word
func makeSecret() (string) {
	random := randomRange(0, len(secrets) - 1)
	return secrets[random]
}

// ================ end functions

// init is a go function -- gets called auto
func init() {
	// reset game state
	secret = makeSecret()

	// set the place holder based on the secret
	revealed = createPlaceholder(secret)

	// println(guesses)
	// println(guess)
	println(secret)
	// println(string(revealed))
	res := formatPlaceholder(revealed)
	println(res)
	println("\n")
}

func formatPlaceholder(word []rune) (string) {
	res := strings.Join(strings.Split(string(word), ""), " ")
	return res
}

// function the change string
func changeStr(slice []rune, index int, value rune) (string) {
	// var slice []rune = []rune(str)
	slice[index] = value
	// convert slice of runes into string
	return string(slice)
}

func searchSecret(secret string, guess rune) ([]int) {
	// convert to slice of rune, so range returns the index of rune, not index of individual bytes
	var word []rune = []rune(secret)
	// var isCorrect bool = false
	var instances = []int{}

	for pos, r := range word {
		if r == guess {
			instances = append(instances, pos)
		}
	}
	//fmt.Errorf("letter %q not found.", guess)
	return instances
}

// ================ game state

var MAX_GUESSES = len(hangman)

var wrongGuesses = 0
var secret string
var guesses []rune = []rune{} // all guesses made
// length of hangman drawing slice
// var guess rune
// the word we display
var revealed []rune

func main() {

	// input scanner
	scanner := bufio.NewReader(os. Stdin)

	// game loop
	for wrongGuesses < MAX_GUESSES {
		// get user rune input
		// rune, byte size, err
		println("guesses left: " + strconv.Itoa(MAX_GUESSES - wrongGuesses))

		// check for win
		if secret == string(revealed) {
			println("win!")
			break
		}

		println("Enter your guess. or '%' to exit. >")
		r, _, err := scanner.ReadRune()

		// skip all white spaces
		if unicode.IsSpace(r) {
			continue
		}

		if err != nil {
			println("input error")
			println(err)
			break
		}

		// check r for quit
		if r == '%' {
			println("goodbye")
			break
		}

		// check for repeated guess and reject it
		if slices.Contains(guesses, r) {
			println("you already tried this guess.")
			continue
		} else {
			// push any guess we make
			guesses = append(guesses, r)
		}

		println("=====\n")
		// draw hangman based on guesses
		println(hangman[wrongGuesses])
		println("=====\n")

		// check for correct guess
		instances := searchSecret(secret, r)
		println(instances)

		if len(instances) != 0 {
			// reveal part of the placeholder
			// revealed[runeIndex] = []rune(secret)[runeIndex]

			for _, index := range instances {
				// reveal each index
				changeStr(revealed, index, []rune(secret)[index])
			}
		} else {
			// show one more part of hangman by
			wrongGuesses++
		}
	}

	if secret == string(revealed) && wrongGuesses <= MAX_GUESSES {
		// game result:
		println("You Win!")
		println("secret was: " + secret)
		println("number of wrong guesses: " + strconv.Itoa(wrongGuesses))
	} else {
		println("Game Over, you lose!")
		println("secret was: " + secret)
		println("number of wrong guesses: " + strconv.Itoa(wrongGuesses))
	}
}
```
