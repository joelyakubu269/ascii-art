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
	input := os.Args[1]
	text := string(data)
	text1 := strings.Split(strings.ReplaceAll(text, "\r", ""), "\n")
	m := buildMap(text1)
	printAscii(input, m)
}
func buildMap(lines []string) map[rune][]string {
	m := make(map[rune][]string)
	for i := 32; i <= 126; i++ {
		start := (i - 32) * 9
		charlines := lines[start : start+8]
		m[rune(i)] = charlines
	}
	return m
}
func printAscii(word string, asciiMap map[rune][]string) {
	words := strings.ReplaceAll(word, `\n`, "\n")
	wordss := strings.Split(words, "\n")
	for _, r := range wordss {
		for i := 0; i < 8; i++ {
			for _, c := range r {
				if val, ok := asciiMap[c]; ok {
					fmt.Print(val[i])
				}
			}
			fmt.Println()

		}
	}
}
