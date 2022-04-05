package main

import "fmt"

func main() {
	type Point struct {
		x int
		y int
	}

	p := Point{
		x: 5,
		y: 20,
	}

	p = Point{ 5, 15}
	fmt.Println(p.x, p.y)

	p.x = 55
	fmt.Println(p.x, p.y)
}