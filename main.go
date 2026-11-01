package main

import (
	"fmt"
)

var arena [20][50]int

func main() {
	fmt.Println(">>>>> ZOMBIE CURSE CURE <<<<<")

	for i := 0; i < 20; i++ {
		for j := 0; j < 50; j++ {

			if i == 0 || i == 19 {
				print("-")
			} else if j == 0 || j == 49 {
				print("|")a
			} else {
				print(" ")
			}
		}

		fmt.Println()
	}
}
