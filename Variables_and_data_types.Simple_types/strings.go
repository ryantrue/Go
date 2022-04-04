package main

import "fmt"

func main() {
	var mySting string

	var hello string = "\tHello\n"
	var title string = "\tHello/n"

	word := "my string"

	str := "Привет"

	var b byte = 'c'
	var r rune = '学'

	fmt.Println(mySting)
	fmt.Println(hello)
	fmt.Println(title)
	fmt.Println(b)
	fmt.Println(r)

	str = "12345"

	fmt.Println(len(str))
	fmt.Println(str[0]) // 49 Первый байт, не символ
	_ = len(str) // количество байтов

	str = "строка"
	_ = len(str) // 5

	fmt.Println(len(str)) // 12
	
	newStr := str[2:4] // "34"
	fmt.Println(newStr)
	newStr = str[:4] // "1234"
	newStr = str[2:] // "345"

	word = "Hello"
	hello = word + " world!" // "Hello world!"
	// word[0] = "h" // ошибка

	// isBigger := "aba" > "abc"
	// isBigger = "aba" < "abc"
	// isEqual = "aba" == "abc"
}