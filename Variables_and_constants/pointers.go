package main

import "fmt"

func main() {
	answer := 42
fmt.Println(&answer) // Указатель на адррес памяти
fmt.Println(*&answer) // Указатель на значение в ячейке памяти
}