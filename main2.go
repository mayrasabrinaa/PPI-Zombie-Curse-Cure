package main

import (
	"fmt"
	"math/rand"

	"github.com/eiannone/keyboard"
)

var jogadorlinha int = 10
var jogadorcoluna int = 25

var zombielinha [35]int
var zombiecoluna [35]int

var quantidadeZumbis int = 20

var temZombieAqui bool

var camundongolinha [10]int
var camundongocoluna [10]int

var temCamundongoAqui bool

var camundongoAtual int = 0

var passos int = 0
var melhorPontuacao int = 0
var partidaAtual int = 1

var gameOver bool = false

var faseFinal bool = false

var linhasMapa int = 20
var colunasMapa int = 50

func main() {

	for k := 0; k < quantidadeZumbis; k++ {
		zombielinha[k] = rand.Intn(linhasMapa-2) + 1
		zombiecoluna[k] = rand.Intn(colunasMapa-2) + 1
	}

	for k := 0; k < 10; k++ {
		camundongolinha[k] = rand.Intn(linhasMapa-2) + 1
		camundongocoluna[k] = rand.Intn(colunasMapa-2) + 1
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
			fmt.Println("                   GAME OVER!                    ")
			fmt.Println("--------------------------------------------------\n")

			fmt.Println("GAME OVER! Voce foi pego por um Zumbi!")
			fmt.Println("Deseja jogar novamente? (S/N)")

			for {

				charGO, keyGO, _ := keyboard.GetKey()

				if keyGO == keyboard.KeyEsc ||
					charGO == 'n' ||
					charGO == 'N' {

					return

				} else if charGO == 's' ||
					charGO == 'S' {

					partidaAtual++
					passos = 0
					camundongoAtual = 0

					jogadorlinha = 10
					jogadorcoluna = 25

					gameOver = false
					faseFinal = false

					linhasMapa = 20
					colunasMapa = 50

					quantidadeZumbis = 20

					for k := 0; k < quantidadeZumbis; k++ {
						zombielinha[k] =
							rand.Intn(linhasMapa-2) + 1

						zombiecoluna[k] =
							rand.Intn(colunasMapa-2) + 1
					}

					for k := 0; k < 10; k++ {
						camundongolinha[k] =
							rand.Intn(linhasMapa-2) + 1

						camundongocoluna[k] =
							rand.Intn(colunasMapa-2) + 1
					}

					break
				}
			}

			continue
		}

		fmt.Println("\n>>>>> ZOMBIE CURSE CURE <<<<<")
		fmt.Println("         🧟-🐭-🧪           \n")

		if faseFinal {
			fmt.Println("==================================================")
			fmt.Println("             A FASE FINAL CHEGOU!")
			fmt.Println("==================================================")
			fmt.Println()
		}

		if melhorPontuacao == 0 {

			fmt.Printf(
				"%-40s ? passos\n",
				"melhor pontuação:",
			)

		} else {

			fmt.Printf(
				"%-40s %d passos\n",
				"melhor pontuação:",
				melhorPontuacao,
			)
		}

		fmt.Print("camundongos capturados: ")

		for c := 0; c < camundongoAtual; c++ {
			fmt.Printf("%c ", 'A'+c)
		}

		fmt.Println()

		textoPartida :=
			fmt.Sprintf("partida atual: %d", partidaAtual)

		fmt.Printf(
			"%-40s %d passos\n",
			textoPartida,
			passos,
		)

		fmt.Printf(
			"zumbis: %d | mapa: %dx%d\n",
			quantidadeZumbis,
			colunasMapa,
			linhasMapa,
		)

		fmt.Println(
			"-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-🐭-",
		)

		fmt.Println(
			"Pressione W, A, S, D para mover. Pressione ESC para sair.",
		)

		for i := 0; i < linhasMapa; i++ {

			for j := 0; j < colunasMapa; j++ {

				if i == 0 || i == linhasMapa-1 {

					fmt.Print("-")

				} else if j == 0 || j == colunasMapa-1 {

					fmt.Print("|")

				} else if i == jogadorlinha &&
					j == jogadorcoluna {

					fmt.Print("@")

				} else {

					temZombieAqui = false

					for k := 0; k < quantidadeZumbis; k++ {

						if i == zombielinha[k] &&
							j == zombiecoluna[k] {

							fmt.Print("X")

							temZombieAqui = true

							break
						}
					}

					if !temZombieAqui {

						temCamundongoAqui = false

						if camundongoAtual < 10 &&
							i == camundongolinha[camundongoAtual] &&
							j == camundongocoluna[camundongoAtual] {

							fmt.Printf(
								"%c",
								'A'+camundongoAtual,
							)

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

		if err != nil ||
			key == keyboard.KeyEsc {

			break
		}

		moveu := false

		if (char == 'w' || char == 'W') &&
			jogadorlinha > 1 {

			jogadorlinha--
			moveu = true

		} else if (char == 's' || char == 'S') &&
			jogadorlinha < linhasMapa-2 {

			jogadorlinha++
			moveu = true

		} else if (char == 'a' || char == 'A') &&
			jogadorcoluna > 1 {

			jogadorcoluna--
			moveu = true

		} else if (char == 'd' || char == 'D') &&
			jogadorcoluna < colunasMapa-2 {

			jogadorcoluna++
			moveu = true
		}

		if camundongoAtual < 10 &&
			jogadorlinha == camundongolinha[camundongoAtual] &&
			jogadorcoluna == camundongocoluna[camundongoAtual] {

			camundongoAtual++
		}

		if camundongoAtual == 10 &&
			!faseFinal {

			faseFinal = true

			linhasMapa = 25
			colunasMapa = 70

			quantidadeZumbis = 35

			for k := 0; k < quantidadeZumbis; k++ {

				zombielinha[k] =
					rand.Intn(linhasMapa-2) + 1

				zombiecoluna[k] =
					rand.Intn(colunasMapa-2) + 1
			}
		}

		if moveu {

			passos++

			if camundongoAtual < 10 {

				passoRatoLinha :=
					rand.Intn(3) - 1

				passoRatoColuna :=
					rand.Intn(3) - 1

				novaLinhaRato :=
					camundongolinha[camundongoAtual] +
						passoRatoLinha

				novaColunaRato :=
					camundongocoluna[camundongoAtual] +
						passoRatoColuna

				if novaLinhaRato >= 1 &&
					novaLinhaRato <= linhasMapa-2 {

					camundongolinha[camundongoAtual] =
						novaLinhaRato
				}

				if novaColunaRato >= 1 &&
					novaColunaRato <= colunasMapa-2 {

					camundongocoluna[camundongoAtual] =
						novaColunaRato
				}
			}

			for k := 0; k < quantidadeZumbis; k++ {

				passoLinha :=
					rand.Intn(3) - 1

				passoColuna :=
					rand.Intn(3) - 1

				novaLinha :=
					zombielinha[k] +
						passoLinha

				novaColuna :=
					zombiecoluna[k] +
						passoColuna

				if novaLinha >= 1 &&
					novaLinha <= linhasMapa-2 {

					zombielinha[k] =
						novaLinha
				}

				if novaColuna >= 1 &&
					novaColuna <= colunasMapa-2 {

					zombiecoluna[k] =
						novaColuna
				}

				if jogadorlinha == zombielinha[k] &&
					jogadorcoluna == zombiecoluna[k] {

					gameOver = true
				}
			}
		}
	}
}
