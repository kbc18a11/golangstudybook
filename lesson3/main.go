package main

import "fmt"

func init() {
	fmt.Println(1)
}

func init() {
	fmt.Println(3)
}

func main() {
	n := 3

	switch {
	case n < 2:
		fmt.Println("1 or 2")
	case n < 4:
		fmt.Println("3 or 4")

	default:
		fmt.Println("ああ")
	}
}
