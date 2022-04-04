package main

import "fmt"

func main() {
	var v1 int
	fmt.Println(v1)

	var v2 int = 100
	fmt.Println(v2)

	v3 := 5
	fmt.Println(v3)

	v4 := 5
	v4 = 10
	fmt.Println(v4)

	var v5, v6 = 1,2
	v5, v6 = 3,4
	v5, v6 = v6, v5

	var v7 int

	v5, v6, v7 = 1,2,3

	var (
		v01 = 1
		v02 = "string"
	)
	fmt.Println(v7)
	_ = v01
	_ = v02

	// Стиль
	var EarthRadiusInMeter int = 6371000
	_ = EarthRadiusInMeter

	var earth_radius_in_meter int = 6371000 // Лучше испоьзовать CamelCase для названий переменных, подобную запись использовать для веток
	_ =	earth_radius_in_meter
}