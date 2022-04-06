package main

import "fmt"

func main() {
	
	sl := []int64{9,8,7}
	for key, value := range sl {
		fmt.Printf("key: %v, val: %v \n", key, value)
	}
	// key: 0, val: 9
	// key: 1, val: 8
	// key: 2, val: 7

	for _, value := range sl {
		fmt.Println(value)
	}
	// 9
	// 8
	// 7

	ages := map[string]int {
		"Andrey": 30,
		"Anastasiya": 25,
	}

for key, value := range ages {
fmt.Println(key)
fmt.Println(value)
}

word := "слёрм"

for i := 0; i < len(word); i++ {
	fmt.Println(word[i])
	fmt.Printf("%T", word[i])
}

for key, value := range word {
	fmt.Println(key)
	fmt.Println(value)
	fmt.Printf("%T", value)
}


}