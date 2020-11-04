package main

import (
	"fmt"

	"./animals"
)

var n = 10

func main() {
	//n = 11
	//fmt.Println(n)
	fmt.Println(AppNeme())
	fmt.Println(animals.ElephantFeed())
	fmt.Println(animals.RabbitFeed())
	fmt.Println(animals.MonkeyFeed())
}
