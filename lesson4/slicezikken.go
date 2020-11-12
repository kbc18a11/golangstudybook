package main

import "fmt"

func main() {
	s := make([]int, 10, 20)
	fmt.Println(s)
	fmt.Println(cap(s))
	for i := 0; i < 20; i++ {
		s = append(s, i)
	}
	fmt.Println(s)
	fmt.Println(cap(s))
}
