package main

import (
	"fmt"
	"math/rand"

	"github.com/eiannone/keyboard"
)

var jogadorlinha int = 10
var jogadorcoluna int = 25
var zombielinha [20]int
var zombiecoluna [20]int
var temZombieAqui bool
var camundongolinha [10]int
var camundongocoluna [10]int
var temCamundongoAqui bool
var camundongoAtual int = 0
var passos int = 0
var melhorPontuacao int = 0
var partidaAtual int = 1
var gameOver bool = false

func main() {

	for k := 0; k < 20; k++ {
		zombielinha[k] = rand.Intn(18) + 1
		zombiecoluna[k] = rand.Intn(48) + 1

	}
	for k := 0; k < 10; k++ {
		camundongolinha[k] = rand.Intn(18) + 1
		camundongocoluna[k] = rand.Intn(48) + 1
	}

	err := keyboard.Open()
	if err != nil {
		panic(err)

	}
	defer keyboard.Close()

	for {
		fmt.Print("\033[H\033[2J")
		if gameOver {
			if passos > melhorPontuacao {
				melhorPontuacao = passos
			}

			fmt.Println("\n>>>>> ZOMBIE CURSE CURE <<<<<")
			fmt.Println("        🧟-🐭-🧪           \n")

			fmt.Println("--------------------------------------------------")
			fmt.Println("                   GAME 0VER!                     ")
			fmt.Println("--------------------------------------------------\n")

			fmt.Println("GAME OVER! Voce foi pego por um Zumbi!")
			fmt.Println("Deseja jogar novamente? (S/N)")

			for {
				charGO, keyGO, _ := keyboard.GetKey()
				if keyGO == keyboard.KeyEsc || charGO == 'n' || charGO == 'N' {
					return
				} else if charGO == 's' || charGO == 'S' {
					partidaAtual++
					passos = 0
					camundongoAtual = 0
					jogadorlinha = 10
					jogadorcoluna = 25
					gameOver = false

					for k := 0; k < 20; k++ {
						zombielinha[k] = rand.Intn(18) + 1
						zombiecoluna[k] = rand.Intn(48) + 1
					}
					for k := 0; k < 10; k++ {
						camundongolinha[k] = rand.Intn(18) + 1
						camundongocoluna[k] = rand.Intn(48) + 1
					}
					break
				}
			}
			continue
		}

		fmt.Println("\n>>>>> ZOMBIE CURSE CURE <<<<<")
		fmt.Println("         🧟-🐭-🧪           \n")
		if melhorPontuacao == 0 {
			fmt.Printf("%-40s ? passos\n", "melhor pontuação:")
		} else {
			fmt.Printf("%-40s %d passos\n", "melhor pontuação:", melhorPontuacao)
		}

		fmt.Print("camundongos capturados: ")
		for c := 0; c < camundongoAtual; c++ {
			fmt.Printf("%c ", 'A'+c)
		}
		fmt.Println()

		textoPartida := fmt.Sprintf("partida atual: %d", partidaAtual)
		fmt.Printf("%-40s %d passos\n", textoPartida, passos)
		fmt.Println("-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-")

		fmt.Println("Pressione W, A, S, D para mover. Pressione ESC para sair.")

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
						temCamundongoAqui = false

						if camundongoAtual < 10 && i == camundongolinha[camundongoAtual] && j == camundongocoluna[camundongoAtual] {
							fmt.Printf("%c", 'A'+camundongoAtual)
							temCamundongoAqui = true
						}

						if !temCamundongoAqui {
							fmt.Print(" ")
						}
					}
				}

			}

			fmt.Println()
		}

		char, key, err := keyboard.GetKey()
		if err != nil || key == keyboard.KeyEsc {
			break
		}

		moveu := false

		if (char == 'w' || char == 'W') && jogadorlinha > 1 {
			jogadorlinha--
			moveu = true
		} else if (char == 's' || char == 'S') && jogadorlinha < 18 {
			jogadorlinha++
			moveu = true
		} else if (char == 'a' || char == 'A') && jogadorcoluna > 1 {
			jogadorcoluna--
			moveu = true
		} else if (char == 'd' || char == 'D') && jogadorcoluna < 48 {
			jogadorcoluna++
			moveu = true
		}

		if camundongoAtual < 10 && jogadorlinha == camundongolinha[camundongoAtual] && jogadorcoluna == camundongocoluna[camundongoAtual] {
			camundongoAtual++
		}

		if moveu {
			passos++

			if camundongoAtual < 10 {
				passoRatoLinha := rand.Intn(3) - 1
				passoRatoColuna := rand.Intn(3) - 1

				novaLinhaRato := camundongolinha[camundongoAtual] + passoRatoLinha
				novaColunaRato := camundongocoluna[camundongoAtual] + passoRatoColuna

				if novaLinhaRato >= 1 && novaLinhaRato <= 18 {
					camundongolinha[camundongoAtual] = novaLinhaRato
				}
				if novaColunaRato >= 1 && novaColunaRato <= 48 {
					camundongocoluna[camundongoAtual] = novaColunaRato
				}
			}

			for k := 0; k < 20; k++ {
				passoLinha := rand.Intn(3) - 1
				passoColuna := rand.Intn(3) - 1

				novaLinha := zombielinha[k] + passoLinha
				novaColuna := zombiecoluna[k] + passoColuna

				if novaLinha >= 1 && novaLinha <= 18 {
					zombielinha[k] = novaLinha
				}
				if novaColuna >= 1 && novaColuna <= 48 {
					zombiecoluna[k] = novaColuna
				}
				if jogadorlinha == zombielinha[k] && jogadorcoluna == zombiecoluna[k] {
					gameOver = true
				}
			}
		}
	}
}
