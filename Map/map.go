package main

import "fmt"

func main() {
	var m1 map[int32]bool
	var m2 map[string]string

	m3 := make(map[int]int)

	ages := map[string]int {
		"Andrey": 30,
		"Anastasiya": 25,
	}

	age := ages["Andrey"] // 30
	ages["Nataly"] = 31 // [Andrey: 30, Anastasiya: 25, Nataly: 31]
	fmt.Println(ages["Michael"]) // 0
	ages["Michael"]++ // [Andrey: 30, Anastasiya: 25, Nataly: 31, Michael: 1]

	fmt.Println(m1, m2, m3, age, ages)
}