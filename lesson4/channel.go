package main

import "fmt"

func main() {
	ch := make(chan int, 10)

	ch <- 5
	ch <- 77

	i := <-ch
	j := <-ch

	fmt.Println(i)
	fmt.Println(j)
	//fmt.Println(ch)
}
