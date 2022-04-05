package main

import "fmt"

func addPrefix(origin string) string {
	return "prefix_" + origin
}

func addPrefixWithErr(origin string) (string, error) {
	return "prefix_" + origin, nil
}

func addPrefixWithLen(origin string) (res string, length int) {
	res = "prefix_" + origin
	length = len(res)
	return res, length
}

func main() {
	mySting := addPrefix("my_string")
	fmt.Println(mySting) // prefix_my_string

	mySting, err := addPrefixWithErr("error_string")
	fmt.Println(err) // nil
	fmt.Println(mySting) // prefix_error_string
}
