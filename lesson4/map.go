package main

import "fmt"

func main() {
	m := map[string]string{
		"aewad": "sgsergserg",
	}

	m["homo"] = "810"

	fmt.Println(m)

	if _, ok := m["ch"]; !ok {
		fmt.Println("ないよ")
	}

	for key, value := range m {
		fmt.Print(key + ":")
		fmt.Println(value)
	}

	delete(m, "homo")
	fmt.Println(len(m))
}
