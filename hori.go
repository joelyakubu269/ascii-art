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
	s := os.Args[1:]
	t := strings.Join(s, " ")
	handle(t, lines)

}
func handle(arg string, lines []string) {
	for _, r := range arg {
		for i := 0; i < 8; i++ {
			index := int(r) - 32
			start := index * 9
			if start < len(lines) {
				fmt.Println(lines[start+i])
			}

		}
		fmt.Println()
	}

}
