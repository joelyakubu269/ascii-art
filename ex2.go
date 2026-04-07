package main

import "fmt"

func main() {
	str := "Hello"
	for _, r := range str {
		index := int(r) - 32
		fmt.Println(string(r), "->", int(r), "index:", index)
	}
}
