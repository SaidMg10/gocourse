package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	message := "Hello, \nGo!"
	message1 := "Hello, \tGo!"
	message2 := "Hello, \rGo!" // Go!lo
	rawMessage := `Hello\nGo`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(rawMessage)

	fmt.Println("Length of message variable is", len(rawMessage))

	fmt.Println("The first character in message var is", message[0]) // ASCII

	greeting := "Hello"
	name := "Alice"
	fmt.Println(greeting + name)
	// msg := greeting + name
	str1 := "Apple"  // A has an ASCII value of 65
	str := "apple"   // a has an ASCII value of 97
	str2 := "banana" // b has an ASCII value of 98
	str3 := "app"    // a has an ASCII value of 97
	fmt.Println(str1 < str2)
	fmt.Println(str3 < str2)
	fmt.Println(str > str1)
	fmt.Println(str > str3)

	for _, char := range message {
		// fmt.Printf("Character at index %d is %c\n", i, char)
		fmt.Printf("%v\n", char)
	}

	fmt.Println("Rune count:", utf8.RuneCountInString(greeting))

	greetingWithName := greeting + name

	fmt.Println(greetingWithName)

	ch := 'a'
	jch := 'ñ'
	fmt.Println(ch)
	fmt.Println(jch)

	fmt.Printf("%c\n", ch)
	fmt.Printf("%c\n", jch)

	cstr := string(ch)
	fmt.Println(cstr)

	fmt.Printf("Type of cstr %T\n", cstr)

	const NIHONGO = "幸運"
	fmt.Println(NIHONGO)

	jhello := "こんにちは"
	for _, runeValue := range jhello {
		fmt.Printf("%c\n", runeValue)
	}

	r := '👀'
	fmt.Printf("%v\n", r)
	fmt.Printf("%c\n", r)
}
