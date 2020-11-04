package main

import (
	"fmt"

	"./foo"
)

func main() {
	fmt.Println(foo.Max)
	fmt.Println(foo.FooFunc(1))
	fmt.Println(foo.internalFunc(1))
}
