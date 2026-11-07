<!--

# hangman:: go + terminal

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

// number of allowed wrong guesses = length of hangman drawings
var guesses = len(hangman)

// current guess
var guess = ""
// secret word
var secret = ""

var placeholder = "- - - - "

// game state
var gameState = "playing" | "stopped"
var gameResult = "win" | "lose"

// game

func init() {
   // reset game state
   guess = ""
   guesses = len(hangman)

   // generate a secret word

}


-->

```go
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

// secret words
var secrets = []string{
	"cat",
	"garden",
	"house",
	"fuck",
	"rock",
}


// game state
var secret string
// length of hangman drawing slice
var guesses int = len(hangman)
var guess rune
// the word we display
var revealed []rune

// get a random word
func makeSecret() (string) {
	random := randomRange(0, len(secrets) - 1)

	return secrets[random]
}

// init is a go function -- gets called auto
func init() {
	// reset game state
	secret = makeSecret()

	// set the place holder based on the secret
	revealed = make([]rune, len([]rune(secret)))

	// set the empty _ placeholders
	for i := range revealed {
		revealed[i] = '_'
	}

	println(guesses)
	println(guess)
	println(secret)
	// println(string(revealed))
	res := getRevealedWord(revealed)
	println(res)
	println("\n")
}

func getRevealedWord(word []rune) (string) {
	res := strings.Join(strings.Split(string(word), ""), " ")
	return res
}

func main() {
	// hangman game in go

	reader := bufio.NewReader(os.Stdin)

	println("Enter a letter or 'q' to Exit.")

	for guesses > 0 {

		// read one character
		// default: runs, for 't' and '/n' (twice)
		letter, _, err := reader.ReadRune()

		// handle read error
		if err != nil {
			println(err.Error())
			// continue for the next loop
			continue
		}
		// skip the loop for newlines or whitespace
		if letter == '\n' || letter == '\r' {
			continue
		}

		if letter == 'q' {
			println("Exit the game")
			break
		}

		// print game state
		println("guesses left: " + strconv.Itoa(guesses))
		println(getRevealedWord(revealed))

		println("you entered this: " + string(letter))
	}

	// println(randomRange(5, 2))
}

```
