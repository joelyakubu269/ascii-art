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
	s := "THere for\nyou"
	printChar(s, lines)

}
func printChar(s string, lines []string) {
	newS := strings.Split(s, "\n")

	for _, r := range newS {
		for i := 0; i < 8; i++ {

			for _, c := range r {
				if c == ' ' {
					fmt.Print("        ")
					continue
				}
				index := int(c) - 32
				start := index * 9
				fmt.Print(lines[start+i])
			}
			fmt.Println()

		}

	}

}
