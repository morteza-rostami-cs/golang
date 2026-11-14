```go

func main() {

	var sex = "sex"

	println(sex)
	// accessing the index gets me the raw byte
	println(sex[0]) // 115

	var char byte = sex[0]// 115

	// convert byte or rune to unicode (code point) string
	println(string(char))// s => byte

	var r rune = 's'

	println(r)// 115
	// rune: is code_point (unicode)
	println(string(r))// s => rune

	// byte => is just raw bit
	// rune => is unicode code point byte/char

	// byte is and alias for uint8
	var raw_byte uint8 = 'g' // 1 bytes -- 8 bits
	var bb byte = raw_byte
	println(bb) // 103

	// rune:: is an alias for int32
	// rune:: holds a single unicode (code point), closest thing to char
	var f int32 = 'd' // 4 bytes -- 32 bits
	var rune1 rune = f
	println(f, rune1)// 100 100

	var username string = "ali fucker "

	println("======\n")

	// using range on a string -- returns runes.
	for _, char := range username {
		// println(char)
		// char = rune = int32 = unicode code point
		fmt.Printf("%d :: %s\n", char, string(char))
	}

	// even when we have a char that takes more bytes inside the string, eg: 😁 => range gives us only the starting index of 😁 and it converts the whole thing to a single rune/char
	var name = "ali😁"

	// string bytes are utf-8 encoded text, one char like this 😁 , can span on to 4 bytes or more
	for i, char := range name {
		fmt.Printf("%d :: %s\n", i, string(char))
	}
	/*
	0 :: a
	1 :: l
	2 :: i
	3 :: 😁
	*/

	// also convert string to slice of runes
	var auth_user []rune = []rune("ali😁")
	println(auth_user)// [97 108 105 128513]

	// convert slice of runes to string
	println(string(auth_user))// ali😁

	// get rune in index[0]
	runes1, size := utf8.DecodeRuneInString(name[0:])
	println(runes1, size)// 97 1
}
```
