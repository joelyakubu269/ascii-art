package main

import (
	"log"
	"os"
	"strings"
)

func main() {

}
func printC(s string) {
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	for _, r := range s {
		index := int(r) - 32
		start := index * 9

	}
}
