package main

import "fmt"

func main() {
	var list []int64

	l := len(list) // 0
	c := cap(list) // 0

	list = []int64{1,2,3,4,5}
	l = len(list) // 5
	c = cap(list) // 5

	list = make([]int64, 0, 5)
	l = len(list) // 0
	c = cap(list) // 5

	list = make([]int64, 5, 5) // [0,0,0,0,0]
	l = len(list) // 5
	c = cap(list) // 5

	list = []int64 {1,2,3,4}
	list = append(list, 5) // {1,2,3,4,5}

	list = make([]int64, 0, 3) // len: 0, cap: 3
	list = append(list, 1)
	list = append(list, 2)
	list = append(list, 3)
	list = append(list, 4) // len: 4, cap: 6

	l = len(list) // 5
	c = cap(list) // 8
	fmt.Println(l, c)
}