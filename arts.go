package main

import "fmt"

func main() {
	arts(5)
}

func arts(rows int) {
	mid := rows / 2

	for i := 0; i < rows; i++ {
		for j := 0; j < 5; j++ {
			if j == 0 || j == 4 || i == mid {
				fmt.Print("|")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
