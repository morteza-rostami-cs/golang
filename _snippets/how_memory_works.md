```go

/*
# bits and bytes

a bit
	- is the smallest unit -- off | on
	0 = off
	1 = on

so: if binary is a number system.
	a single binary digit is called a bit.

one bit::
	0 or 1

#================
# now what is a byte?

byte is 8 bits
	# each bit has 2 states
	8 * 2 = 16 -- 16 states

1 byte = 8 bites

# each position has two possible states:: 0 or 1
0 0 0 0 0 1 0 1

# so:: 8 bits gives us -- 256 combinations

# it can represent integers from 0 to 255
smallest value = 0
largest value = 255
total possible values = 256

#=======================

# permutation order matters

now each place/bit can have 2 possible values: 0 or 1

so:: 8 places and 2 possible values

2 ^ 8

# this is counting -- we multiply the
# counting:: total number of outcomes = product of the -- number of options/slots at each choice.

#  2 * 2  * 2 = 4 * 2 = 8
# [ ] [ ] [ ]
2 * 2 * 2 * 2 * 2 * 2 * 2 * 2 = 256

"256 permutations" just means "256 distinct bit-patterns." The count is what matters: 2 choices per position, 8 positions, multiply → 2⁸ = 256.

#=======================



*/

func main() {

	// this has 8 bites = 1 bytes = can hold 0 to 255

	// this actually overflows
	// var age uint8 = 280
	// var age uint8 =  -1 // negative also overflows

	// 8 bites unsigned int
	var age uint8 = 255 // 0 to 255

	println(age, reflect.TypeOf(age)) // 255, uint8

	// binary form of int
	fmt.Printf("Binary: %#b\n", age) // 11111111 or 0b11111111 (0b for binary)

	// 16 bits = 65,536 (0 to 65, 536) -- 2^16
	// var score uint16 = 100000 // overflow

	// 32 bits = 4,294,967,296 -- 2^32
	// var g uint32 = 5000000000 // over flow

	var score uint32 = 1_000_000

	println(score)

	var foo uint8
	// how many bites allocates?
	println(unsafe.Sizeof(foo)) // 1 bytes = 8 bites

	println(unsafe.Sizeof(score)) // 32 / 8 = 4 bytes

}
```

```go

/*
Type	Memory	Range
int8	1 byte	-128 to 127
int16	2 bytes	-32,768 to 32,767
int32	4 bytes	-2.1 billion to 2.1 billion
int64	8 bytes	Approximately -9.2 quintillion to 9.2 quintillion

There are also unsigned equivalents:
Type		Memory	Range
uint8		1 byte	0 to 255
uint16	2 bytes	0 to 65,535
uint32	4 bytes	0 to 4.29 billion
uint64	8 bytes	0 to approximately 18.4 quintillion
*/


func main() {

	// how go stores negative integers?

	// 256 / 2 = 128
	var fuck int8
	println(unsafe.Sizeof(fuck)) //1 byte = 8bites -- -128 to 127

	// -128 to 127 => so negatives get a bit more space

	var num1 int8 = 5
	var num2 int8 = -5

	fmt.Printf("num1: %08b\n", uint8(num1))
	fmt.Printf("num2: %08b\n", uint8(num2))
	/*
	num1: 00000101
	num2: 11111011
How do we get from positive 5 to negative 5?

Two's complement works like this:
1. Start with the positive binary number.
2. Flip every bit (zero becomes one, one becomes zero).
3. Add one.

Positive 5:

00000101

Flip every bit:

11111010

Add one:

11111011
	*/

	var x uint8 = 255

	// range is 0 to 255 , add 1 and it wraps around to zero
	x++
	println(x) // 0

	// also run time overflow can happen in Go! and you get compile error when assigning something out of range.

	// int in go => alias for int32 or int64 (depends on the operating system)
	var boob int = 256
	println(reflect.TypeOf(boob)) // int
	println(unsafe.Sizeof(boob)) // 8 bytes - 64 bites 8 * 8 = 64
}

```

```go

func main() {
	// floating point numbers

	// 4 bytes -- 6 to 7 decimal digits
	var g float32 = 3.141592653589793

	// bytes -- 15 to 16 decimal digits
	var f float64 = 3.141592653589793


	println(g) // 3.1415927 -> it cuts off the rest
	println(f) // 3.141592653589793

	println(unsafe.Sizeof(g)) // 4 bytes
	println(unsafe.Sizeof(f)) // 8 bytes

	// decimals are stored using a method called: use a standard called IEEE 754
	fmt.Printf("%#b\n", g) //13176795p-22 -> binary

	jo := 0.1
	mo := 0.2

	// many fractions can not be represented exactly -- using: binary floating-point numbers.
	if jo + mo == 0.3 {
		println("0.3")
	} else {
		println(jo + mo) // 0.30000000000000004
	}

	// do not compare floating point values with ==
	// instead check for approaching values

	// if a >= 0.299999

	// integer division vs floating point division

	println(5 / 2) // 2
	println(5.0 / 2.0) // 2.5

	println(5.0 / 2)// 2.5

	// there is no implicit type conversion
	var gg int = 10
	var bb float64 = float64(gg)

	println(bb)

	// float to int
	var dd int = int(bb)
	println(dd)
}

```

```go

/*
# some code points

Character	Code point	Decimal
A				U+0041		65
a				U+0061		97
€				U+20AC		8364
中				U+4E2D		20013
😀				U+1F600		128512

# code point is a code => mapped to a character

# char encoding is how that code is encoded and stored in bytes

ASCII: only 0–127, one byte each. A → byte 0x41.

UTF-8: variable length (1–4 bytes). A → 0x41 (1 byte), but 😀 → 4 bytes.

UTF-16: 2 or 4 bytes.

UTF-32: always 4 bytes.

#=========================

Same character, different representation, different byte cost.

As a rune/int32: you store the number (the code point) directly → always 4 bytes, no decoding needed.

As UTF-8 in a string: you store an encoding of that number → 1–4 bytes, and you must decode to get the number back.

So when a character is "big" (like 😀, code point 128512), the UTF-8 representation can take just as much or more space, and the tradeoff is:

Rune/int32 → fast to process as a number, but wastes space for small chars ('A' still costs 4 bytes).

UTF-8 string → compact for common text ('A' = 1 byte), but you pay a decoding cost and variable size.

*/


func main() {
	// go strings are sequence of bytes.
	// a rune represents a unicode code point

	// "Hello" => 5 bytes

	// so we represent strings as int values
	// for example: Ascii:: A = 65
	// ascii :: uses 7 bits -- 128 possible values/characters

	// byte is alias for uint8 -- 1 byte -- 8 bits
	var letter byte = 65

	println(letter) // 65
	// byte to ascii char
	fmt.Printf("%c\n", letter) // A

	// convert int to unicode string char
	println(string(letter)) // A -> ascii A is included in unicode chars also

	// convert to string containing the same digits
	println(strconv.Itoa(int(letter))) //65

	// unicode
	// A = U+0041
	// 😊 = U+1F60A
	/*
	Unicode defines which numeric value represents a character.

	UTF-8 defines how those numeric values are encoded into bytes.
	*/

	words := []string{
		"A",
		"ش",
		"中",
		"😊",
	}

	/*
	lent(string) -> returns number of bytes ,not characters

	// unicode point like 65 = A
	// utf-8 is a protocol for converting unicode point to corresponding char
	// Code point 0x1F600 (128512) → 😀

	A = 1 bytes
	ش = 2 bytes
	中 = 3 bytes
	😊 = 4 bytes
	*/
	for _, word := range words {
		// string = bytes
		fmt.Printf("%s = %d bytes\n", word, len(word))
	}

	// go string
	var name string = "😀"
	var user rune = '😀'

	// get the byte size
	// go string is UTF-8 encoded bytes. so: it takes more space.
	println(unsafe.Sizeof(name)) // string:: 16 bytes
	println(unsafe.Sizeof(user)) // rune:: 4 bytes

	// get the type
	println(reflect.TypeOf(name)) // string
	println(reflect.TypeOf(user)) // int32
	println(name, user) // 😀 128512

	// how many rune in a string? (rune is alias for int32) -- 32 bits -- 4 bytes
	println(utf8.RuneCountInString(name))//1
	// RuneLen returns the number of bytes in the UTF-8 encoding of the rune

	// rune is a int32 => always takes up 4 bytes, no matter what -- RuneLen -> gives us how much bytes the rune takes if encoded into a string using UTF-8
	println(utf8.RuneLen(user)) // 4 UTF-8 bytes

	println(utf8.RuneLen('A')) // 1 UTF-8 byte

	// when we use unicode to store 😀 inside a string and utf-8 to decode it => the representation of this that we store in memory might take more bytes to store.
	// basically something like this:: F0 9F 98 80  →  (decode UTF-8)  →  128512

	// but a rune = int32 -- always takes 4 bytes
}
```

```go

func main() {
	// strings are immutable -- so to change it -- create a new one
	var message string = "fuck"

	message = "G" + message[1:]
	println(message) // Guck

	// access individual bytes -- string is a sequence of bytes
	println(message[0]) // 71
	println(message[2]) // 99

	// 😀
	var face string = "😀"
	var face2 rune = '😀'

	println(unsafe.Sizeof(face)) // 16 bytes
	println(unsafe.Sizeof(face2)) // 4 bytes

	println("A", 'A') // A, 65
	println(len("A"))

	// string takes 16 bytes -> because of UTF-8 encoding
	// 'A'  →  code point 65  →  UTF-8 bytes: 01000001
	println(unsafe.Sizeof("A"))

	// so:: rune stores the code point 65 -- but string stores the encoded version of 65
	/*
	Rune 'A' stores the code point 65 as an int32 → 4 bytes, fixed.

	String "A" stores the UTF-8 encoding of 65 → 1 byte: 01000001.
	*/

	// compare rune with rune
	var a_rune rune = 'A'
	if a_rune == 'A' {
		println("A -> code point = 65")
		println(a_rune) // 65
	}

	var a_str string = "A"
	if "A" == "A" {
		println("A string:: ", a_str) // A
	}

	// compare rune and string
	if string(a_rune) == a_str {
		println(a_rune, a_str) // 65 A
	}

	// slice of rune (here :: decoding to runes happens)
	var a_slice_runes []rune = []rune(a_str)

	// slice of bytes/runes
	println(a_slice_runes)// [65]

	// compare rune to string
	if a_rune == a_slice_runes[0] {
		println(a_rune, a_slice_runes[0]) // 65 65
	}

	/*
	string('A') → encodes 65 as UTF-8 → "A" ✅

	[]rune("A")[0] → decodes "A" → 65 ✅
	*/

	var str string = "A😀"

	// if we use simple loop -- we go over individual raw bytes
	// i -> is the index of each byte
	for i := 0; i < len(str); i++ {
		res := str[i]
		println(res)
		/*
		65 => A
		😀 => emoji takes more bytes
		240
		159
		152
		128

		*/
	}

	// if we use range -- we get runes -- we get the actual code point
	// i => is the byte index, where that rune/code point starts
	for i, r := range str {
		println(i, r)
		/*
		# now we get -- decoded or rune -- for each char
		0 65
		1 128512
		*/
	}

	// so can compare to runes
	for i, r := range str {
		if r == 'A' {
			println(str[i]) // 65
		}
		if r == '😀' {
			println(str[i]) // get the byte (240) where emoji starts
			println(r) // 128512
			println(string(r)) // gets you the emoji
		}
	}

	/*
	Rune holds the raw code point (e.g. 65 for 'A', 128512 for 😀) as an int32.

	To compare a string's character to a rune, you must decode first — []rune(s)[0] == 'A' or for _, r := range s { r == 'A' }. The decode step turns UTF-8 bytes back into the code point number.

	The plain for i := 0; i < len(s); i++ loop walks raw bytes, so a 4-byte emoji shows up as 4 separate byte values, each holding a fragment of the encoding — meaningless on their own. That's exactly why range exists.
	*/


}
```
