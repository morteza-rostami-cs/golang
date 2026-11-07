```go

// package main is entry point of program, it defines a stand alone executable. vs a library.
package main

// has to come after package name
import "fmt"

// entry point of the program
func main() {

	var username string = "ali"
	//constant variable
	const final = "GOOGLE"

	username = "sara"
	// final = ""

	fmt.Println("hi, you cock sucker " + username)
	fmt.Println(final)
}


```

```go
// package main is entry point of program, it defines a stand alone executable. vs a library.
package main

// has to come after package name
import (
	"fmt"
	"os"
)

// print command line args
func printCommands(args []string) {
	var line string
	separator := " "

	// start from 1 to only get the command line args
	for i := 1; i < len(args); i++ {
		line += separator + os.Args[i]
	}

	fmt.Println(line)
}

// entry point of the program
func main() {
	/// os.Args -- is a slice of strings
	fmt.Println(os.Args)

	// path to the executable we are running for main.go
	// /tmp/go-build703579749/b001/exe/main
	fmt.Println(os.Args[0])

	// passed arguments
	fmt.Println(os.Args[1]) // second arg


	// program executable
	// fmt.Printf("program name: %s\n", os.Args[0])

	fmt.Println("=====\n")

	for i, arg := range os.Args {
		fmt.Printf(" os.Args[%d] = %q\n", i, arg)
	}

	fmt.Println(len(os.Args))// 2
	fmt.Println(os.Args[0:1]) // index[0]

	fmt.Println("=====\n")

	printCommands(os.Args)
}

```

```go
// package main is entry point of program, it defines a stand alone executable. vs a library.
package main

// has to come after package name
import (
	"fmt"
	// "os"
)

// entry point of the program
func main() {

	// map: key is string and values are int
	counts := make(map[string]int)

	counts["orange"] = 5

	counts["apple"]++
	counts["banana"]++
	counts["apple"]++

	//map[apple:2 banana:1 orange:5]
	fmt.Println(counts)

	fmt.Println(counts["apple"])
	fmt.Println(counts["banana"])
	// fmt.Println(counts["apple"])

	fmt.Println("===\n")
	// find duplicate values
	for key, value := range counts {
		// %s string -- %d int
		// fmt.Printf("%s = %d\n", key, value)

		// check if something like: 'orange'=2
		if value > 1 {
			fmt.Printf("%s = %d\n", key, value)
		} else {
			fmt.Println("=== not duplicated === \n")
			fmt.Printf("%s = %d\n", key, value)
		}
	}
}

```

```go
// package main is entry point of program, it defines a stand alone executable. vs a library.
package main

// has to come after package name
import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "os"
)

// entry point of the program
func main() {
	// create an input scanner
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter something > ")
	// loop runs, until:: there is no line to read -- return: false
	for scanner.Scan() {
		// read a line
		line := scanner.Text()

		// if: 'quit' exit this loop
		if line == "quit" {
			fmt.Println("goodbye pale.")
			break
		}

		// print a line
		fmt.Println(line)
	}

	// check for scanner error
	if err := scanner.Err(); err != nil {
		// log input scanner error
		log.Fatal("Error reading input: ", err)
	}
}
```

```go
// package main is entry point of program, it defines a stand alone executable. vs a library.
package main

// has to come after package name
import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	// "os"
)

// entry point of the program
func main() {
	// create an input scanner
	scanner := bufio.NewScanner(os.Stdin)

	// store input in a map
	lines := make(map[string]int)

	fmt.Println("Enter shit > ")
	for scanner.Scan() {
		// get the next line and push/increment in our map
		line := scanner.Text()

		// check for quit
		if line == "quit" {
			fmt.Println("BYE.")
			break
		}

		// push input to our map
		lines[line]++
		fmt.Println(line + " was added!")
	}

	// check for scanner error
	if err := scanner.Err(); err != nil {
		// log input scanner error
		log.Fatal("Error reading input: ", err)
	}

	// duplicates := []string{} // slice

	var hasDuplicates bool = false

	// check for duplicates, when we we quit the loop
	for line, count := range lines {
		// key:value
		if count > 1 {
			fmt.Println("key: " + line + " __ " + "value: " + strconv.Itoa(count))
			hasDuplicates = true//found at least one duplicate
		}
	}

	if !hasDuplicates {
		fmt.Println("no duplicate was found.")
	}
}
```

```go
package main

import (
	"fmt"
	"reflect"
)

var println = fmt.Println

// := can only be used inside functions
// fuck := ""

// a function that returns 2 values
func getValues(a int , b int) (int, int) {
	return a, b
}

func main() {
	// println("i love squeezing boobs")

	var username string
	// variables can not be unused
	// if: no value assigned -- go assigns zero values, eg: int = 0 and string = ""
	var score int
	var playing bool

	println(username, score, playing)

	if username == "" {
		println("this is an empty string")
	}

	// defer types
	enemy := "fucker"

	// print type
	fmt.Printf("%T\n", enemy) // string

	// type as a value
	println(reflect.TypeOf(34)) // int

	// check if value is int
	// skip: value, ok:: int or not (boolean)
	_, ok := any(score).(int)

	if ok {
		println("this is an integer.")
	}

	// group variable declaration
	var (
		age int
		active bool
	)
	println(age, active)


	// use blank identifier _ , skip a return value
	_, b := getValues(2, 3)
	println(b)

	// pass a struct as parameter
	greet(Options{
		Name: "al",
		Email: "al@gmail.com",
	})


}


// use a struct for function parameters
type Options struct {
	Name string
	Email string
}

func greet(options Options) {
	println(options.Name, options.Email)
}
```

```go
package main

import (
	"fmt"
)

var println = fmt.Println

// := can only be used inside functions
// fuck := ""

// a function that returns 2 values
func getValues(a int , b int) (int, int) {
	return a, b
}

func main() {
	// basic types
	var username string = "ali"
	var score int = 23
	var playing bool = false
	var points float64 = 2.3

	println(username, score, playing, points)

	// can not do operations with different number types, int + float64
	// println(score + points)
	println(float64(score) + points) // 25.3

	// constant can be: numbers, strings, booleans
	const age = 54
	// iota -- multiple constants
	const (
		a = 1
		b = 2
		c = 3
	)

	println(age, a, b, c)

	// this goes from 0, 1, 2 ...
	const (
		zero = iota
		one
		two
	)
	println(zero, one, two)

	// type for const
	const (
		count = 0
		countStr string = "0"
	)
	println(count, countStr)
}
```

```go
package main

import (
	"fmt"
)

var println = fmt.Println

func main() {
	// strings

	var fuckBoy string = "i fuck"

	println(fuckBoy)

	// strings are immutable
	// fuckBoy[0] = "G"

	// access index
	// you get bytes -- not character -- strings are UTF-8 sequences.
	// a char is 3 bytes
	println(fuckBoy[1]) // 32

	// println(rune)
	println(len(fuckBoy)) // string_length: 6

	// what the hell is runes in Go?

	// literal alias
	// type rune = int32

	r := "é"
	fmt.Printf("%s\n", r) //f
	println(r) // f

	// use runes to manipulate string
	story := "fuck this shit"
	runeStory := []rune(story)

	// rune -- should be single quote
	runeStory[0] = '$'

	// convert back, rune to string
	changed := string(runeStory)
	fmt.Printf("%s\n", changed)

	// =======
	// we can also declare our string as rune
	var gg []rune = []rune("google") // slice of runes

	gg[0] = '%'
	println(string(gg)) // %oogle
}

```

```go

func main(){
	var letter rune = 'A'

	var word string = "Ass"

	// word[0] -> gives you the raw byte
	// we need to get the rune or point_code by decoding the utf-8 string
	// []rune(word)[0] -> convert to rune and get the first char
	if letter == []rune(word)[0] {
		fmt.Printf("%d = %d\n", letter, []rune(word)[0])
		// 65 = 65
	}

	//============== search for a letter inside a string

	var secret string = "cat😀"

	// search for rune a
	var a rune = 'a'

	// range:: only goes over real (point codes), not every single bytes.
	// len(secret) = 6 => number of bytes in string
	// but: this loop runs for max 4 chars/runes in "cat😀"
	for i, r := range secret {
		// just print the last byte/index
		println(i, len(secret) - 1)
		if r == a {
			println("rune 'a' is inside secret")

			break
		} else {
			// at the last byte, print no search result, but: if we find it sooner, we never run this

			// check:: number of bytes
			// if i == len(secret) - 1 {
			// 	println("no search result.")
			// }

			// we need to check: length of runes/chars inside the string -- last index
			// also:: len([]rune(secret))
			if i == utf8.RuneCountInString(secret) - 1 {
				println("this is the last index, and no search result.")
			}
		}
	}
}
```

```go
func main() {
	// changing a string

	var name string = "ali"

	name2 := changeStr(name, 0, '$')

	fmt.Printf("%s\n", name2) // $li

	// ========

	// multi-line strings
	multiStr := "start string " +
		"end string"
	println(multiStr)

	// raw string literal
	// preserve:: white space and new line
	str := `start a string here.
	this is a story about a girl with big tits.
	`
	println(str)

	// a rune is a int32 = 4 bytes -- unicode code point -- eg: 'A' = 65
	// a string is a sequence of utf-8 encoded bytes, eg: "A😀", and A might take 1 byte, but 😀 might take 4 bytes or something

	// get runes/chars inside the string
	var story string = "A😀"

	for _, r := range story {
		// get a single rune | unicode code point | char | int32
		fmt.Printf("%d\n", r)
		// 65, 128512
	}

	// go has a type = error
	// error is also an interface
	var e error

	println(e) // init value is => <nil>
}
```

```go
func main() {
	// go operators

	a := 10
	b := 3

	println(a + b)// 13
	println(a - b) // 7
	println(a * b) // 30
	// int division
	println(a / b) // 3
	println(10.0 / 3.0) //3.3333333333333335
	println(a % b) // 1

	// bit wise operator
	gg := 6
	ff := 3
	println("====\n")

	// bitwise AND -- keeps bits that are set in both values
	println(gg & ff) // 2
	println(gg | ff) // 7
	println(gg ^ ff) // 5
	println(gg &^ ff) // 4

	//===============

	// logical operators

	println(true && false) // false
	println(true || false) // true
	println(!true) // false

}
```

```go
// function the change string
func changeStr(str string, index int, value rune) (string) {
	var slice []rune = []rune(str)
	slice[index] = value
	// convert slice of runes into string
	return string(slice)
}

// if error - exit with error
func openFile(name string) (string, error) {
	file, err := os.Open(name)

	// return with error
	if err != nil {
		return "", err
	}

	// close file
	defer file.Close()

	// println(file)
	// nil for no error
	return file.Name(), nil
}

func main() {
	// go controls

	var playing bool = true

	if playing {
		println("game is running")
	} else {
		println("game over")
	}

	age := 18

	if age < 10 {
		println("you are a kid")
	} else if age >= 10 && age < 18 {
		println("you are not an adult")
	} else {
		println("ok, you are an adult now")
	}

	// init a variable inside if
	// init; condition {}
	// if can not load the file, returns error, nil for no error
	// err -- only exists in scope of if and it's else
	if err := os.Chmod("", 0664); err != nil {
		println(err)
	}

	// go does not have falsy values
	// we have to use boolean or something that evaluate to bool
	// if 0 {
	// 	println("")
	// }

	var zero int = 0
	if zero == 0 {
		println("this is zero")
	}

	var emptyStr string = ""
	if emptyStr == "" {
		println("string is empty")
	}

	// slice declaration -- default value is nil
	var list []rune

	println(list) // []
	println(reflect.TypeOf(list)) // []int32

	// so empty rune is nil
	println(list == nil) // true

	if len(list) == 0 {
		println("our slice is empty")
	}

	// empty slice literal
	values := []int{} // this not nil

	println(len(values)) // 0

	// make
	ss := make(
		[]int, // type
		0, // size
	)

	println(ss) // []
	println(ss == nil)// false

	// the best way to just check the length
	if len(ss) == 0 {
		println("ss is empty")
	}

	// not empty
	if !(len([]rune{'A'}) == 0) {
		println("not fucking empty!")
	} else {
		println("is fucking empty")
	}

	fileName, err := openFile("/")

	if err != nil {
		println(err)
	} else {
		println(fileName)
	}
}
```

```go

// jump to a certain part of a function

func loop() int {
	i := 0

	Here: // label
		println(i)

		if i > 5 {
			return i
		}


		i++
		goto Here // jump to label
}

func main() {
	// loop()

	// normal for loop
	for i := 0; i < 5; i++ {
		println(i)
	}

	// while loop
	var count int = 0

	for count < 5 {
		println(count)
		count++
	}


	var j int = 0

	// infinite loop
	for {
		println(j)

		if j > 3 {
			fmt.Printf("stop at %d\n", j)
			break
		}

		j++

	}

	// we can label outer loop and break it
	I:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("[%d][%d]\n", i, j)

			// break the outer loop at [1][1]
			if i == 1 && j == 1 {
				// break the outer loop
				break I
			}
		}
	}
}
```

```go

/*
index:  0    1    2    3    4    5
byte:  0x74 0x25 0xF0 0x9F 0x98 0x80
char:   t    %    └──────── 😀 ────────┘
                  (4 bytes: F0 9F 98 80)
*/

func main() {
	// range
	// iterate over collection of values
	// slices, arrays, strings, maps, channels

	// range over slice
	list := []string{"ali", "sara", "john", "angelina"}

	for i, name := range list {

		if name == "john" {
			// string(rune) - this is for rune to string conversion
			println("ali is in index: " + strconv.Itoa(i)) // index = 2
		}

		fmt.Printf("%d :: %s\n", i, name)
	}

	// range over string -- returns rune|point code|char


	str := "t%😀"

	println(str[0]) // 116 -- rune
	println(string(str[0])) // -- we get t -- we pass the first byte or t
	println([]rune(str)[0]) // 116

	println(str[3]) // 159 -- this is part of emoji bytes
	println(string(str[3])) // !! this does not work for getting the emoji char
	println(string(str[2]))

	// decode rune in a string -- second method
	r, size := utf8.DecodeRuneInString(str[2:]) // start at index emoji bytes starts
	println(string(r), size) // 4 bytes

	// to get the emoji rune or char => have to pass the rune
	println(string([]rune(str)[2])) // 😀

	// pos:: is start index of bytes for each char
	for pos, char := range str {
		fmt.Printf(
			"character %c starts at byte position %d\n",
			char,
			pos,
		)
	}

	for i := 0; i < 4; i++ {

		// switch statement
		switch i {
		case 0:
			println("zero")
		case 1:
			println("one")
		default:
			println("some other value")
		}
	}

	// switch with statement
	var score int = 5

	switch {
	case score < 0:
		println("negative score is not allowed")
	case score < 5:
		println("under 5, got to do better")
	case score >= 5 && score < 10:
		println("between 5 and 10")
	default:
		println("other scores")
	}

	// fallthrough :: execute next case, even it the previous one is true
	var girl string = "hot"
	switch girl {
	case "hot":
		println("this girl is supper hot")
		fallthrough
	case "busty":
		println("she is also busty") // this also runs
	}

	// run if this case or that is true
	switch girl {
	case "hot", "busty":
		println("she is so hot and busty.")
	default:
		println("she is not hot")
	}
}
```

```go

func main() {
	// default value for int is 0
	nums := make(
		[]int,// type
		2, // elements to initialize
		4,// size
	)

	println(nums[0])// default value is zero
	println(len(nums)) // holds 2 elements
	println(cap(nums)) // underlying array can hold 4

	// append elements, returns a new slice
	nums = append(nums, 99)

	// just inserted 1 value into underlying array
	println(len(nums)) // 3
	println(cap(nums)) // 4

	// copy an slice
	var original []int = []int{1, 2, 3}

	var target []int = make([]int, 3)// 3 default values, zeros

	copy(target, original)
	println(target) // [1, 2, 3]
}
```

```go

func main() {
	// arrays and slices and maps

	// arrays

	// basically if we pass a fixed size -- then: it is an array.
	var arr [4]int // 4 places

	arr[0] = 1
	arr[1] = 12

	println(arr) // [1, 12, 0, 0]

	// declare array and initialize (composite literal)
	var values [3]int = [3]int{1, 2, 3}

	// loop over array
	for i, val := range values {
		println(i, val)
	}

	// go can also count the elements
	scores := [...]int{1, 2, 3}// fixed size 3

	println(scores)
	println(len(scores))// number of elements in array
	println(cap(scores))

	// 2d array
	grid := [2][2]int{
		{1, 2},
		{2, 3},
	}

	// loop over 2d array
	for row :=0; row < len(grid); row++ {
		for col :=0; col < len(grid[row]); col++ {
			fmt.Printf("row: %d :: col: %d\n", row, col)
			println("value: ", grid[row][col])
		}
	}
}
```

```go

func main() {
	// slices
	// can grow when we append elements.

	// init 5 elements with zeros
	values := make([]int, 5)

	// length
	println(len(values))// items stored
	println(cap(values)) // capacity

	// under laying array -- fixed sized
	var arr [4]int

	println(arr) // [0, 0, 0, 0]

	// arr -- becomes the underlying array
	slice := arr[0:2]
	println(reflect.TypeOf(slice)) // slice of int, []int
	println(slice) // [0, 0]

	println(cap(slice))// 4 --- underlying array has 4 indexes
	// just push to underlying array and return a new slice
	slice = append(slice, 99)

	println(slice)// [0, 0, 99]

	// === get part of a slice
	var scores []int = []int{1, 2, 3, 4}

	// [i:j] -- j is excluded
	println(scores[0:1]) // [1]
	println(scores[:1]) // [1]
	println(scores[2:]) // [3, 4]

	// this is not an independent array
	scores2 := scores[:]

	scores2[1] = 100

	// it changes the same underlying array
	println(scores, scores2) // [1 100 3 4] [1 100 3 4]

	println("===\n")

	scores3 := scores[:2]
	// even though underlying array has 4 slots, you can not access more that 2 elements on this slice
	//println(scores3[3]) // Error
	for i, val := range scores3 {
		println(i, val) // two loops -- 0, 1
	}

	// append another slice
	items := []string{"box", "gold"}
	questItems := []string{"sward", "wooden bow"}

	items = append(items, questItems...)

	println(items) // [box, gold, sward, wooden bow]

	// copy slice
	source := []int{1, 2, 3}
	target := make([]int, 3)

	// copy into target
	numberOfCopiedItems := copy(target, source)

	println(numberOfCopiedItems) // 3

	target[0] = 78
	// target has a different underlying array now
	println(target, source) // [78 2 3] [1 2 3]

	//Important: copy copies the minimum of the two slice lengths.
}
```

```go
func main() {
	// maps
	// dictionary or hash map

	// map[keyType]valueType

	colors := map[string]string{
		"red": "red",
		"blue": "blue",
	}

	println(colors)

	// empty map
	days := make(
		map[string]int,
	)

	println(days)
	println(len(days)) // 0

	// set key value
	days["sunday"] = 6
	// access value with key
	println(days["sunday"])

	// if the key exists -> updates it
	days["sunday"] = 66

	// loop over a map
	// order is not guaranteed
	for key, val := range days {
		println(key, val)
	}

	// get the value and check if it exists
	value, ok := days["monday"] // called:: comma-ok form

	if ok {
		println(value)
	} else {
		println("value does not exists.")
	}

	// delete map entry
	delete(days, "sunday")
	println(days) // map[]
}
```

```go

/*
# challenge

Q2. (0) For-loop
1. Create a simple loop with the for construct. Make it loop 10 times and print out
the loop counter with the fmt package.
2. Rewrite the loop from 1. to use goto. The keyword for may not be used.
3. Rewrite the loop again so that it fills an array and then prints that array to the
screen.
*/

// loop with goto
func loop() int {
	i := 0

	here:

		if i >= 10 {
			println("stop the loop")
			return i
		}
		println(i)

		i++
		goto here
}

// fill an array and print
func printArr() {
	const size = 10 // array size should be known in compile time
	var arr [size]int = [size]int{}

	for i:=0; i<size; i++ {
		// write to an array
		arr[i] = i
	}

	// we need to read the array -- if we write to it
	for i, v := range arr {
		println(i, v)
	}
}

func main() {
	for i:=0; i<10; i++ {
		fmt.Printf("%d\n", i)
	}

	println("======\n")

	loop()

	println("====\n")

	printArr()
}
```

```go

/*
Q3. (0) FizzBuzz
1. Solve this problem, called the Fizz-Buzz [23] problem:
Write a program that prints the numbers from 1 to 100. But for multiples
of three print “Fizz” instead of the number and for the multiples of five
print “Buzz”. For numbers which are multiples of both three and five print
“FizzBuzz”

*/

func main() {

	// fizzbuzz
	const min = 1
	const max = 100

	for i:=min; i <= max; i++ {

		// first check for both, then check multiples of 3 and 5
		if i % 3 == 0 && i % 5 == 0 {
			// multiples of both 3 and 5, print "FizzBuzz"
			println("FizzBuzz: " + strconv.Itoa(i))
		} else if i % 3 == 0 {
			// multiples of 3, print "Fizz"
			println("Fizz: " + strconv.Itoa(i))
		} else if i % 5 == 0 {
			// multiples of 5, print "Buzz"
			println("Buzz: " + strconv.Itoa(i))
		} else {
			// other numbers
			println("Others: " + strconv.Itoa(i))
		}
	}
}
```

```go
package main

import (
	// "bufio"
	"fmt"
	"math/rand"
	"unicode/utf8"

	// "reflect"
	// "slices"
	// "strconv"
	"strings"
	// "unicode"
	// "unicode/utf8"
	// "unsafe"
	// "os"
)

var println = fmt.Println
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
// ================ end functions

// init is a go function -- gets called auto
func init() {
}

/*
Q4. (1) Strings
1. Create a Go program that prints the following (up to 100 characters):
A
AA
AAA
AAAA
AAAAA
AAAAAA
AAAAAAA
...
2. Create a program that counts the number of characters in this string:
asSASA ddd dsjkdsjs dk
In addition, make it output the number of bytes in that string. Hint: Check out the
utf8 package.
3. Extend/change the program from the previous question to replace the three runes
at position 4 with ’abc’.
4. Write a Go program that reverses a string, so “foobar” is printed as “raboof”. Hint:
You will need to know about conversion; skip ahead to section “Conversions” on
page 56.
*/

func printHalfPyramid() {

	// char to print
	var char string = "A"

	// var sb strings.Builder

	for i:=1; i<=20; i++ {
		// repeat i number of string A
		str := strings.Repeat(char, i)
		// print it
		println(str)
	}
}

// return (num of chars/runes, num of bytes)
func countCharsInStr(str string) (int, int) {
	/*
	2. Create a program that counts the number of characters in this string:
asSASA ddd dsjkdsjs dk

	we count: whitespace and duplicates
	*/
	// var str = "asSASA ddd dsjkdsjs dk"

	bytesCount := len(str)
	runeCount := utf8.RuneCountInString(str)

	return runeCount, bytesCount
}

func replaceStrRunes() (string) {
	//replace index:4,5,6
	//          ___
	str := "asSASA ddd dsjkdsjs dk"

	var text []rune = []rune(str)
	text[4] = 'a'
	text[5] = 'b'
	text[6] = 'c'

	return string(text)
}

// reverse a string
func reverseStr(str string) string {
	s := []rune(str)

	p1 := 0
	p2 := len(s) - 1

	for {
		// stop when pointers on same index or passed each other
		if p1 >= p2 {
			println("end of our swap")
			break
		}

		// swap
		s[p1], s[p2] = s[p2], s[p1]

		// move the pointers towards each other
		p1++
		p2--
	}

	return string(s)
}

// average of a float slice
func averageFloat(f []float64) float64 {
	if len(f) == 0 {
		return 0.0
	}
	var sum float64 = 0.0
	for _, v := range f {
		sum += v
	}

	return sum / float64(len(f))
}

func main() {
	// printHalfPyramid()
	var str = "asSASA ddd dsjkdsjs dk"

	runes, bytes := countCharsInStr(str)// 22, 22
	println(bytes, runes)

	// emoji takes up more raw bytes = 5 , but we have 2 chars/runes
	r, b := countCharsInStr("A😀")// 5, 2
	println(b, r)

	println(replaceStrRunes())
	// asSAabcddd dsjkdsjs dk

	res := reverseStr("foobar")
	println(res)

	// converts both rune and slice of runes to a string
	println(string('A'))
	println(string([]rune{'A', 'B'}))

	println(averageFloat([]float64{2.3, 3.4, 5.6}))
}
```

```go
package main

import (
	// "bufio"
	"fmt"
	"math/rand"
	// "reflect"
	// "slices"
	// "strconv"
	// "unicode"
	// "unicode/utf8"
	// "unsafe"
	// "os"
)

var println = fmt.Println
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
// ================ end functions

// init is a go function -- gets called auto
func init() {
}

// functions in Go
func foo(param string) string {
	return "from foo: " + param
}

// multiple parameter shorthand
func add(a, b int) int {
	return a + b
}

// return multiple values
func calc(a, b int) (int, int) {
	return a + b, a * b
}

// in Go, everything is passed by value
func change(x int) {// func gets a copy
	x = 100
	println("Change: ")
	println(x)
}

// even slices are passed by value
func addItem(slice []int) {
	slice = append(slice, 5)
	println(slice)
}

func main() {
	println(foo("fuck"))

	res1, res2 := calc(2, 2)

	println(res1, res2)

	var x = 99
	change(x)
	println(x)//100

	var s []int = []int{1, 2}
	addItem(s)
	println(s) // [1, 2]

	greet()
}

// function can be declared after using it
func greet() {
	println("ruck , yeah!")

	rec(0)
}

// basic recursion
func rec(i int) {

	// base case, return if this case is true
	if i >= 10 {
		return
	}

	rec(i + 1) // we go into calls

	// at the end we return here, 9, 8, 7, 6, 5 ...
	fmt.Printf("%d ", i)

	// last frame to pop, has i=0
	if i == 0 {
		println("\nend\n")
	}
}
```

```go
package main

import (
	"fmt"
)

var println = fmt.Println

// global variable :: can be accessed by functions inside this package
var a = "i am global"

func main() {
	println(a)

	// assign to global
	a = "assigned from main"
	println(a)

	// new local:: declared inside a function
	a := "i am local"
	println(a)
}
```

```go
package main

import (
	"errors"
	"fmt"
)

var println = fmt.Println

// go result parameters, so the result is named

// named results parameters: 1. behave like local vars 2. they auto init by zero/default values
func divide(a, b int) (result int, err error) {

	println(result)// 0
	println(err) // <nil>
	// result = 1
	// println(result) // 1

	var Err = errors.New("divide by zero.")
	if b == 0 {
		// return 0, Err
		err = Err
		// called: naked return
		return // 0, err -- is returned auto
	}

	result = a / b

	// so, even if we just return, go returns the named results
	return

	// return a / b, nil
}

func main() {
	result, err := divide(5, 0)

	if err != nil {
		println(err)
	}

	println(result)
}
```

```go
package main

import (
	// "errors"
	"fmt"
	"os"
)

var println = fmt.Println

// more on functions

func readFile(path string) error {
	file, err := os.Open(path)

	// handle file open error
	if err != nil {
		return fmt.Errorf("opening %s: %w", path, err)
	}

	// schedule a function call, until surrounding function returns -- so, we do not have to close the file, on each error
	defer file.Close()

	// multiple defer statements run like a stack: last in first out
	// defer func somethingElse(){}

	defer func() {println("runs after ali@gmail")}()

	// after surrounding func done:: this runs first
	defer func(email string) {
		println(email)
	}("ali@gmail.com")

	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading %s: %w", path, err)
	}

	println("data: ")
	println(string(data))// data is bytes
	return nil
}

// store a function literal
var f = func(name string) {
	println("from f: " + name)
}

// this function returns 1
func goo() (result int) {
	// so, defer runs at the end of a func, just before return
	defer func() {
		// can only do this with named return values
		result++
	}()
	return 0
}

func main() {
	err := readFile("file.txt")

	if err != nil {
		println(err)
	}

	// function literal
	func() {
		println("this func runs immediately.")
	}()
	f("ali")

	res := goo()
	println(res)
}
```

```go
package main

import (
	// "errors"
	"fmt"
	"reflect"
)

var println = fmt.Println

/*
Remember the distinction:
- arg ...int declares a variadic parameter.
- arg... expands a slice when passing it to a variadic function.
*/

//variadic parameters -- accept variable number of arguments.
func f(arg ...int) {
	// arg -- []int -- slice of int
	for _, n := range arg {
		fmt.Printf("%d\n", n)
	}

	// expand the slice into another func
	func(arg ...int) {
		fmt.Printf("== %d\n", arg[0])
	}(arg...)
}

// function as value
var foo = func() {
	println("foo")
}

func main() {

	f(1, 2, 3, 4)

	foo()

	// also we can declare anonymous func inside another function
	goo :=func() {
		println("goo")
	}

	goo()

	reflect.TypeOf(goo) // goo

	// function inside a map
	// key: string, val: func
	var functions = map[string]func() string {
		//key:val
		"f": func() string {return "f"},
		"g": func() string {return "g"},
		"x": func() string {return "x"},
	}

	v1 := functions["f"]()
	v2 := functions["g"]()
	v3 := functions["x"]()

	println(v1, v2, v3) // f g x

	// callback
	add := func(a, b int) int {
		return a + b
	}

	// callback
	calc := func(
		v int,
		fn func(int, int) int, // callback
	) int {
		// calling the callback
		return fn(v, v)
	}

	res := calc(2, add)
	println(res)//4
}
```

```go
package main

import (
	// "errors"
	"fmt"
)

var println = fmt.Println

// panic and recover

func raisePanic() {
	println("A")

	// function panic and return with error
	panic("something wrong in raisePanic()")

	println("B")// unreachable code
}

func main() {
	// raisePanic()

	//stack unwinding.
	fb :=func() {
		println("start fb")

		defer func(){println("deferred in fb")}()

		// basically this bubbles back to all older functions on stack
		panic("shit happened inside fb()")

		println("end of fb()")// unreachable
	}

	fa := func() {
		println("start fa")

		// when it panics -- we start emptying the stack, but deferred functions are called.
		defer func(){println("deferred in fa")}()

		// recover should be called inside a deferred function
		// handle panic
		// during normal execution, recover() returns nil and nothing else happens
		defer func(){
			panic_err := recover()

			println("recovered-- ", panic_err)
		}()

		fb()
		println("end of fa()")
	}

	fa()

}
```

```go
package main

import (
	// "errors"
	"fmt"
)

var println = fmt.Println

// panic and recover

// run a callback, return true if it panics
// handle a panic
func throwPanic(fn func()) (b bool) { // b has default = false
	// recover() -- must be called inside a defer
	defer func() {
		// recover, returns the panic_value or error
		if err := recover(); err != nil {
			// set the named return value
			println(err)
			b = true
		}
	}()

	// call our callback, if it panics! we recover here
	fn()

	// naked return, auto returns b
	return
}

func main() {

	fa := func() {
		println("from fa")
		panic("something wrong inside fa")

		// what ever here, will be unreachable.
	}

	panicHappened := throwPanic(fa)
	// true if there is a panic somewhere down fa
	// println(res)

	if panicHappened {
		println("fa is panicking")
	}
}
```

```go
package main

import (
	// "errors"
	"fmt"
	// module/package
	// "learn/even"

	// name alias
	m "learn/even"
	// eve "learn/even/Even" // not a thing!!
)

var println = fmt.Println

func main() {

	m.Fn()
	// m.ExampleEven()

	res := m.Even(48)

	if res {
		println("this value is even.")
	}

	// even.odd() // private
}
```

```go
package main

import (
	// "errors"
	"fmt"
	// module/package
	// "learn/even"
	// eve "learn/even/Even" // not a thing!!
)

var println = fmt.Println

func main() {
	// pointers

	var username string = "ali"

	fa :=func(username string){
		username = "sara"
		println(username) // sara
	}

	fa(username)

	// outside value does not change, it is passed by a value copy
	println(username)// ali

	// use a pointer
	var pUsername *string // a pointer to a string

	// zero value or default for a pointer is nil
	println(pUsername) // <nil>

	// store username address inside pointer, address-of-operator &
	pUsername = &username

	// a func that takes a pointer
	fb :=func(username *string){
		println("raw pointer.")
		println(username)// 0x153f49290070

		// dereference the pointer for read and write
		*username = "jake"// assign: get the memory
		println(*username)// access: jake
	}

	fb(pUsername)

	println(username) // jake

	// memory allocation, new, make

	// constructor and composite literals
}
```

```go

func main() {

	var age int = 23
	var p *int = &age

	*p++ // this increments value, not pointer arithmetic

	println(*p, age) // 24 24

	// Go is garbage collected

	// allocate memory for an int32 -- 32 bits -- 4 bytes -- returns a pointer to that memory location
	var p1 *int = new(int)// int pointer, holds address


	println(p1)// 0x211a9de9e020
	// also with zero/default value
	println(*p1) // 0

	*p1 = 45

	p2 := *p1
	println(p2)


	// a normal variable, holds the value and not the address
	var score int = 23
	println(score, &score)

	// new(T) , T :: could be any type. returns:: a pointer to a zero-valued T

}
```

```go

func main() {

	//zero-value-is-useful design principle.
	// BUT: for some types, when we create a pointer using new, and it has some default value. but we can use the pointer directly without (dereference)

	// *bytes.Buffer
	pp := new(bytes.Buffer)

	// this is a pointer to buffer
	fmt.Printf("%T\n", pp) // *bytes.Buffer
	pp.WriteString("fuck!")
	println(pp.String()) // fuck!

	// a Buffer
	var b bytes.Buffer

	fmt.Printf("%T\n", b) // bytes.Buffer

	b.WriteString("hi, ")
	b.WriteString("shit")
	b.WriteByte('!')

	println(b.String())// hi, shit!
	println(b.Len()) // 9
}
```

```go
/*
new(T) returns *T, pointing to a zero-valued T.

make(T) returns an initialized T.

make only works with slices, maps, and channels.
*/

func main() {

	// it returns:: initialized value and not a pointer
	// make :: only used for, slice, map, channel

	scores := make([]int, 10)

	// 10 zeros
	println(len(scores), scores)

	// compare to new
	// you get a pointer
	// but also: at the same time, no init value! value = nil
	pointer := new([]int)
	println(pointer) // &[] -- this is a pointer to empty nil slice

	// =================
	// =================

	userIds := make(map[string]int)

	println(userIds) // empty map[]

	messages := make(chan int)
	println(messages) // 0xd78ec0e0070
}
```

```go

func main() {

	type Point struct {
		x, y int
	}

	// composite literal can be used to initialize: struct, array, slice, map
	p1 := Point{x: 1, y: 2} // keyed: order does not matter
	// when: unkeyed -> must supply every field
	p2 := Point{11, 12} // unkeyed: positional

	println(p1, p2) // {1, 2} {11, 12}

	// composite literal for slice
	var books []string = []string{"sex", "dick", "pussy"}

	println(books)

	// map
	var userIds = map[string]int{
		"ali": 1,
		"sara": 2,
	}

	for key, id := range userIds {
		println(key, id)
	}

	// constructor is just a pattern
	// a function who returns a pointer to a struct, with initialized values and validation.

	type Server struct {
		host string
		port int
		conns []string
	}

	fn :=func(host string, port int) (*Server, error) {

		// if port is negative or zero
		if port <= 0 {
			return nil, fmt.Errorf("invalid port %d", port)
		}

		// dereference and set and
		serverPointer := &Server{
			host: host,
			port: port,
			conns: make([]string, 0),// zero elements, not nil
		}

		return serverPointer, nil
	}

	p, err := fn("localhost", 3000)

	println(p)// &{localhost 3000 []} -- a pointer
	if err != nil {
		println(err)
	} else {
		fmt.Printf("%p\n", p)// 0x38141400c120
		// println -- dereferences the pointer (auto)
		println((*p).host, p.host) // localhost localhost
	}
}
```

```go

// return a reference/pointer to a local variable!! it's value stays around, even after the function pops off.

func foo() *int {
	x := 23
	return &x
}

func main() {

	type User struct {
		name string
	}

	// x is a pointer to local foo() variable
	var x *int = foo() // foo, is done here.

	// does not dereference pointer to int
	println(x)// 0x3775bbcac008

	// reassign:: we can still play around with x
	*x = 2
	println(*x) // 2

	var user User = User{name: "ali"}
	var p *User = &user // pointer to a struct
	// some how it dereference pointer to a struct
	println(p.name) // ali


}
```

```go

// define your own types

// underlying type is int -- but: foo is a separate type now.
type foo int

type User struct {
	username string
	score int
}

func main() {

	var age foo = 12
	var score int = 13
	println(age, score)
	// foo is not the same as int
	// if age == score {

	// }

	var user User = User{
		username: "ali",
		score: 23,
	}
	sara := User{username: "sara", score: 21}

	sara.score++
	println(user, sara)
	fmt.Printf("%v\n", sara)

	sex := new(User)
	println(sex)// pointer
	println(sex.username, sex.score)// "" 0

	sex.username = "fucker"
	sex.score = 89
	println(sex.username, sex.score)

	//Important: Go can automatically format structs using fmt.Printf and %v. That's why printing a gives us &{Pete 42}.

	type Animal struct{
		// field
		name string
		power float64

		// children *[]string
		children []string
		eat func()
	}

	bear := Animal{
		name: "bear",
		power: 99,
		children: []string{"a", "b"},
		eat: func(){println("bear eats.")},
	}

	println(bear)
	bear.eat()

	type T int
	// anonymous fields
	type Fuck struct {
		T// embedded field
	}

	var f = Fuck{}
	println(f)

	// so:: other packages , can only access the public fields
	type Person struct {
		Name string // public
		age int // private
	}

	person := Person{Name: "fuck", age: 2}
	println(person.Name) // can access inside the same package
}
```

```go

// methods, function that works with type

type Player struct {
	name string
	attackPower int
}

// player method
// receiver:: (player *Player)
func (player *Player) attack(enemy string) {
	// we have access to player.filed, inside this method
	println("player attacks " + enemy + " for " + strconv.Itoa(player.attackPower))
}

func main() {
	var p = Player{name: "joe", attackPower: 10}

	println(p)
	// p is not a pointer but Go does this: (&n).doSomething(2)
	p.attackPower++
	p.attack("goblin") // player attacks goblin for 11
}
```

```go
// embed type

type Gun struct{
	name string
	damage int
}

func (gun *Gun) shoot(){
	println("gun damage: " + strconv.Itoa(gun.damage))
}

type Player struct{
	// now player can also access Gun's methods
	Gun// embed type (composition)
	name string
}

func main() {

	var player = Player{name: "hero", damage: 9}

	player.shoot()
}
```

```go

func main() {

	str := "ali🙂"

	// uint8 -- 8 bits -- 1 byte
	bytes := []byte(str)
	runes := []rune(str)

	println(bytes)
	println(runes)
	/*
	# utf-8 bytes
	[97 108 105 240 159 153 130]
	# unicode code points (chars)
	[97 108 105 128578]
	*/

	// emoji needs more bytes
	// b := []byte{'a', 'l', 'i', '🙂'} wrong!!

	b := []byte{'a', 'l', 'i'}
	println(string(b)) // ali

	// store code points
	r := []rune{97, 108, 105}
	println(string(r))// ali

	// convert may own types
	type Foo int

	var score int = 3

	var foo Foo = Foo(score)
	goo := Foo(3) // convert int to Foo

	// both are Foo now
	if foo == goo {
		println("now, foo and goo are Foo")
	}
}
```

```go
// interface is the contract for methods
// interface, does not enforce implementation of other fields on our type, just the methods and their shape
type I interface {
	Get() int
	Put(int)
}

type S struct {
	i int
}

// methods
func (p *S) Get() int {
	return p.i
}

func (p *S) Put(v int) {
	p.i = v
}

// func (p *S) foo() {}

func main(){

	var s = S{i: 0}
	println(s.Get())
	s.Put(5)
	println(s.Get())

	//This is called implicit interface implementation.

	// we need to pass a pointer to S, which must have shape of I
	// whatever we pass here, has to have:: Get() and Put(), but: also it can have other stuff
	func(obj I){
		println(obj.Get())
		obj.Put(3)
		println(obj.Get())

		// obj.foo()
	}(&s)// pass a pointer to S
}
```

```go

// interfaces
type Animal interface{
	Eat(string)
	Attack()
}

type Cat struct{
	name string
	power int
}
func (c *Cat) Eat(food string){println(c.name + " eats. " + food)}
func (c *Cat) Attack(){println(c.name + " attacks " + strconv.Itoa(c.power))}


type Bear struct{
	name string
	power int
}
func (c *Bear) Eat(food string){println(c.name + " eats. " + food)}
func (c *Bear) Attack(){println(c.name + " attacks " + strconv.Itoa(c.power))}

// switch based on types
func fn(face Animal) {
	// it's called:: type switch
	switch face.(type) {
	case *Cat:// type is *Cat -- pointer_cat
		println("this is a cat")
	case *Bear:
		println("this is a bear")
	default:
		println("some other type")
	}
}

func main() {
	cat := Cat{name: "cat", power: 10}
	bear := Bear{name: "bear", power: 300}
	// always pass a pointer to whatever that is using an interface
	fn(&cat)
	fn(&bear)
}
```

```go

type User interface{
	login(string) string
}
type Guest struct{
	name string
}
func (u *Guest) login(email string) (res string) {
	println("user is in: " + u.name)
	println(email)
	return "login success"
}

func main() {

	var guest = Guest{name: "ali"}
	// we put the pointer in some interface value
	// var some_interface interface{} = &guest
	// interface{} value = any
	var some_interface any = &guest
	// user:: User interface type
	// ok:: guest is of type user
	user, ok := some_interface.(User)

	if ok {
		// guest implements interface User method
		println(user.login("ali@gmail.com"))
	} else {
		println("does not implement User interface.")
	}
}

```

```go

// empty interface
// type i interface{}

// this is a interface type with no methods, so we can pass any value into this function
func fa(some interface{}) {
	println(some)
}

type User struct{name string}

func main() {

	fa(3)
	fa("SOME")
	fa(true)
	fa(&User{name: "al"}) // struct pointer

}
```

```go

type U interface{
	login(string) string
}

type Admin struct{
	name string
	email string
}
func (u *Admin) login(email string) string {
	return email + " logged in."
}

// we use empty interface and assertion here
// so we can use, interface{} to pass any type in
func fa(some interface{}) string {
	// println(some)

	// cause: the type is interface, assert the dynamic value *Admin first
	admin, ok := some.(*Admin)
	if !ok {
		return "not an *Admin"
	}
	println(admin.name)

	// if it is not type U -- panics
	return some.(U).login(admin.email)
}

func main() {
	res := fa(&Admin{name: "jenifer", email: "jen@gmail.com"})
	println(res)
}
```

```go

type User struct{
	name string
	// getName func() string
}

// type we pass in receiver, must be inside the same package.
func (self *User) getName() string {
	return self.name
}

func main(){
	var user = User{name: "angelina"}
	println(user.getName())
}
```

```go

type Foo int

// if we send a pointer -> we access the outside foo
// without pointer -> we get a copy
// only: use a type available in this package, inside receiver. types from other packages (no), built-in types (no)
// receiver value -- can not be an interface either.
func (self *Foo) Emit() {
	// just print foo value, which is a int
	fmt.Printf("%v\n", *self) // 3
}

func main() {

	// Foo is just an int (not alias though)
	var foo Foo = 3
	foo.Emit() // 3
}
```

```go
package main

import (
	"fmt"
	// module/package
	// "learn/even"
	// eve "learn/even/Even" // not a thing!!
)

var println = fmt.Println

//interface-based polymorphism

/*
# how Go, compares strings::

“Smaller” means comes first in Go’s byte-by-byte, lexicographic order. Go compares from the start and stops at the first difference:

"apple" < "banana" is true because a comes before b.
"ant" < "anteater" is true because the shorter string comes first when the shared part is identical.
"Zebra" < "apple" is true because uppercase letters sort before lowercase ones.
*/

// so the problem is, we want to pass []int or []string , into our bubble sort function

// define an interface that both types can implement, all the methods needed for bubbleSort algo
type Sorter interface{
	Len() int // get the size of slice
	Less(i, j int) bool // compare two elements, for smaller value
	Swap(i, j int) // swap two elements
}

// now these types can implement Sorter, remember:: we can not use receiver and declare method on built-in methods, like: []int
type SliceInt []int
type SliceStr []string

// value receiver, so we can just use a value in receiver. cause: in this case:: operation is not expensive, so: it copies the slice
func (self SliceInt) Len() int {
	return len(self)
}

func (self SliceInt) Less(i, j int) bool {
	// compare two ints
	return self[j] < self[i]
}

func (self SliceInt) Swap(i, j int) {
	// swap neighbors
	holder := self[i]
	self[i] = self[j]
	self[j] = holder
}

// we need to implement the same methods for []string

func (self SliceStr) Len() int {
	// get number of elements, in slice of string
	return len(self)
}

func (self SliceStr) Less(i, j int) bool {
	// we can just compare strings using < or >
	return  self[j] < self[i]
}

// NOTICE :: in this case: slice has an underlying array -- so even though we are passing a copy of slice here, it still can change scores and names, outside!!
func (self SliceStr) Swap(i, j int) {
	holder := self[i]
	self[i] = self[j]
	self[j] = holder
}

/*
# we keep i on one element, and we move j over all the the others, find the smallest then arr[i] and swap

# when we move i forward, arr[i] already has the smallest value compare to arr[0] position
# as we move i, whatever is bubbling to the left if it, is already sorted.

 |  |
[4, 8, 3, 1]
i = 0

j = i+1
j = i+2
*/

// bubble sort
func sort(x Sorter) {
	var size = x.Len() // []int or []string

	// i goes from 0 to length - 1
	for i :=0; i < size; i++ {
		for j :=i+1; j < size; j++ {
			// compare and swap, if: a[j] < a[i]
			if x.Less(i, j) {
				x.Swap(i, j)
			}
		}
	}
}

func main() {
	var scores = SliceInt{4, 8, 3, 1}
	var names = SliceStr{"sara", "al", "boobs"}

	sort(&scores)
	sort(&names)

	println(scores)
	println(names)
}
```

```go

// inspect type and tag
func ShowTag(i interface{}) {
	t := reflect.TypeOf(i)
	println(t)//*main.Person

	v := reflect.ValueOf(i)
	println(v) //&{sal 0}

	if t.Kind() == reflect.Ptr {
		// dereference pointer type
		t = t.Elem()
		println(t) // main.Person
	}

	// get person.name, first struct field
	tag := t.Field(0).Tag
	fmt.Println(tag)
	// get a specific tag value
	println(t.Field(0).Tag.Get("namestr"))// person-name
}

type Person struct{
	// struct tag
	name string `namestr:"person-name"`
	age int
}

func main() {

	person := Person{name: "sal"}
	println(person.name)
	println(person)

	// get person.name tag
	ShowTag(&person)//namestr:"person-name"

}
```

```go

// goroutine
func great(name string) {println("hi " + name)}

func main(){

	println("before greet") // this

	// this runs outside of current flow
	go great("ali")

	println("after greet") // then this happens
}
```

```go
// goroutine

func ready(drink string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)

	println(drink, "is ready!")
}

func main(){

	go ready("tea", 2)
	go ready("coffee", 1)

	println("i am waiting.")

	// we have to wait 5 second, so terminal is open for getting the results of goroutines
	//also, when main() finishes, program terminates, including any goroutines, still running.
	time.Sleep(5 * time.Second)
}
```

```go

// concurrency

/*
With concurrency, Go can switch between them, allowing both to make progress.

With parallelism, tey can actually execute simultaneously on different CPU cores.
*/

var channelA chan string
var channelB chan int

func ready(drink string, sec int) {
	// sleep this goroutine for some seconds
	time.Sleep(time.Duration(sec) * time.Second)

	println(drink, "is ready!")

	// send one through our channel
	channelA <- drink
	channelB <- 2
}

func main() {
	// init our channel
	channelA = make(chan string)
	channelB = make(chan int)

	// run two go routines
	go ready("Tea", 2)
	go ready("Coffee", 1)

	println("i am waiting!")

	// when we have a channel, some how main waits!
	// each receive => for it's value
	// println(<-c)
	// println(<-c)

	// wait for multiple channel operations
	received :=0
	// L: // labeled loop
	for received < 4 { // 4 received values
		//select:: waits for a channel operation to happen
		select{
		case value, ok := <-channelA:
			if !ok {
				println("channelA was closed")
				return
			}
			println(value) // channel value
			// i++
			// wait for 2 operations on this channel
			// if i > 1{
			// 	break L
			// }
		case value, ok := <-channelB:
			if !ok {
				println("channelB closed")
				return
			}
			println(value)

			// i++
			// if i > 1{
			// 	break L
			// }
		}
		received++
	}
}
```

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	// module/package
	// "learn/even"
	// eve "learn/even/Even" // not a thing!!
)

var println = fmt.Println

//read a file
func readFile() {
	// 1 byte -- 8 bits -- uint8 -- 1024 bytes
	// files are read as raw data, not as strings
	buffer := make([]byte, 1024)

	// file, _ := os.Open("/etc/passwd")
	// file => is a pointer type
	file, err := os.Open("./output.txt")

	// check for file open error
	if err != nil {
		println(err)
		return
	}

	defer file.Close()

	for {
		// read file into a buffer -- up to 1024 bytes
		n, err := file.Read(buffer)

		if err != nil {
			println(err)
			return
		}

		// n :: is the number of bytes left to read
		if n == 0 {
			println("no bytes to read.")
			break
		}

		// we write buffer, to standard input (terminal)
		// [:n] -> because we might have only n bytes in the buffer, and the rest of 1024 might be empty
		os.Stdout.Write(buffer[:n]) // [start:n]
	}
}

func bufferReader() {
	buffer := make([]byte, 1024)

	file, _ := os.Open("./output.txt")

	defer file.Close()

	// buffered reader, around the file
	reader := bufio.NewReader(file)
	// buffered writer, around std
	writer := bufio.NewWriter(os.Stdout)

	//push any remaining buffered data to the underlying writer
	defer writer.Flush()

	for {
		n, _ := reader.Read(buffer)

		if n == 0 {
			break
		}

		writer.Write(buffer[:n])
	}
}

func main() {
	// readFile()
	bufferReader()
}

```

```go
// networking

func main() {

	conn, err := net.Dial(
		"tcp", // network protocol
		"192.0.32.10:80", // address
	)

	if err != nil {
		println(err)
		return
	}

	defer conn.Close()

	// make a request
	req := "GET / HTTP/1.1\r\nHost: 192.0.32.10\r\nConnection: close\r\n\r\n"

	// write the request to server
	if _, err := conn.Write([]byte(req)); err != nil {
		println(err); return
	}

	var buf = make([]byte, 1024)

	// println(conn.LocalAddr())

	n, err := conn.Read(buf)

	if err != nil {
		println(err)
		return
	}

	println(n)
	fmt.Printf("%s\n", buf[:n])

	/*
	120
	HTTP/1.0 302 Moved Temporarily
	Location: https://192.0.32.10/
	Server: Apache
	Connection: close
	Content-Length: 0
	*/
	// println(buf)
}
```

```go

func main() {
	// http networking

	response, err := http.Get("https://example.com")

	if err != nil {
		println(err); return
	}

	defer response.Body.Close()// close response body

	// read body data
	// response.Body -> implements io.Reader interface
	body, err := io.ReadAll(response.Body)

	if err != nil {println(err); return}

	// body -> slice of bytes/runes
	println(string(body))
}
```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```

```go

```
