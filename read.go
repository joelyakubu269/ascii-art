package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")

	index := 'E' - 32

	start := int(index * 9)
	for i := 0; i < 8; i++ {
		fmt.Println(lines[start+i])
	}
}
