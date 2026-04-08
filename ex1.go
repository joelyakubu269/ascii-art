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
	printC("A", lines)

}
func printC(s string, lines []string) {
	if s == "" {
		return
	}

	for i := 0; i < 8; i++ {
		for _, r := range s {
			if r < 32 || r > 126 {
				continue
			}
			index := int(r) - 32
			start := index * 9
			fmt.Print(lines[start+i])

		}
		fmt.Println()
	}

}
