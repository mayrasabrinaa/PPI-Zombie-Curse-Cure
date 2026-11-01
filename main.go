package main

import (
	"fmt"
)

var arena [20][50]int
var jogadorlinha [0]int = 10
var jogadorcoluna [1]int = 25

func main() {
	fmt.Println(">>>>> ZOMBIE CURSE CURE <<<<<")

	for i := 0; i < 20; i++ {
		for j := 0; j < 50; j++ {

			if i == 0 || i == 19 {
				print("-")
			} else if j == 0 || j == 49 {
				print("|")
			} else {
				print(" ")
			}

		}

		fmt.Println()
	}
}
