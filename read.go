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
	s := "HELLO"
	printChar(s, lines)

}
func printChar(s string, lines []string) {
	for _, r := range s {
		index := int(r) - 32
		start := index * 9
		for i := 0; i < 8; i++ {
			fmt.Print(lines[start+i])
		}
		fmt.Println()
	}
}
