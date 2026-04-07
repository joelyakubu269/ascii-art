package main

import "fmt"

func main() {
	art(7)
}
func art(rows int) {
	for i := 0; i < rows; i++ {
		// for k := 0; k < 1; k++ {
		// 	fmt.Print(" ")
		// }
		for j := 0; j < rows; j++ {
			if i == (rows)/2 {
				if j > 0 && j < rows-1 {
					fmt.Print(".")
				} else if j == 0 || j == rows-1 {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}
			} else {

				if j == 0 || j == 1 || j == rows-2 || j == rows-1 {
					fmt.Print("|")
				} else {
					fmt.Print(" ")
				}

			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
