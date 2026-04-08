package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var arg = os.Args[1:]
	arg1 := strings.Join(arg, " ")
	// flag.StringVar(&arg, "val", "", "put the text you want to use")
	// flag.Parse()
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	printC(arg1, lines)

}
func printC(s string, lines []string) {
	if s == "" {
		return
	}
	parts := strings.ReplaceAll(s, `\n`, "\n")
	parts1 := strings.Split(parts, "\n")

	for _, r := range parts1 {
		for i := 0; i < 8; i++ {
			for _, c := range r {
				if c < 32 || c > 126 {
					continue
				}
				index := int(c) - 32
				start := index * 9
				if start+i < len(lines) {
					fmt.Print(lines[start+i])

				}

			}
			fmt.Println()

		}
	}

}
