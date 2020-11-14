package main

import "fmt"

func homo(a [3]int) {
	for i, v := range a {
		a[i] = v * v
	}

	return
}

func main() {
	a := [3]int{1, 2, 3}

	homo(a)

	fmt.Println(a)
}
