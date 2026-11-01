package main

import (
	"fmt"
	"math/rand"
)

var arena [20][50]int
var jogadorlinha int = 10
var jogadorcoluna int = 25
var zombielinha [20]int
var zombiecoluna [20]int
var temZombieAqui bool

func main() {
	fmt.Println(">>>>> ZOMBIE CURSE CURE <<<<<")
	for k := 0; k < 20; k++ {
		zombielinha[k] = rand.Intn(18) + 1
		zombiecoluna[k] = rand.Intn(48) + 1

	}

	for i := 0; i < 20; i++ {
		for j := 0; j < 50; j++ {

			if i == 0 || i == 19 {
				fmt.Print("-")
			} else if j == 0 || j == 49 {
				fmt.Print("|")
			} else if i == jogadorlinha && j == jogadorcoluna {
				fmt.Print("@")
			} else {
				temZombieAqui = false
				for k := 0; k < 20; k++ {
					if i == zombielinha[k] && j == zombiecoluna[k] {
						fmt.Print("X")
						temZombieAqui = true
						break
					}
				}
				if !temZombieAqui {
					fmt.Print(" ")
				}
			}

		}

		fmt.Println()
	}
}
