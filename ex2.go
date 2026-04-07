package main

import "fmt"

func main() {
	str := "Hello"
	for _, r := range str {
		fmt.Println(string(r)+"->", int(r))
	}
}
