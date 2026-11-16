package main

import (
	"fmt"
	"math/rand"

	"github.com/eiannone/keyboard"
)

// =========================
// JOGADOR
// =========================

var jogadorlinha int = 10
var jogadorcoluna int = 25

// =========================
// ZUMBIS
// =========================

// O máximo agora é 35, pois a fase final terá 35 zumbis.
var zombielinha [35]int
var zombiecoluna [35]int

var quantidadeZumbis int = 20

var temZombieAqui bool

// =========================
// CAMUNDONGOS
// =========================

var camundongolinha [10]int
var camundongocoluna [10]int

var temCamundongoAqui bool

var camundongoAtual int = 0

// =========================
// JOGO
// =========================

var passos int = 0
var melhorPontuacao int = 0
var partidaAtual int = 1

var gameOver bool = false

// =========================
// FASE FINAL
// =========================

var faseFinal bool = false

// Tamanho inicial do mapa
var linhasMapa int = 20
var colunasMapa int = 50

func main() {

	// =========================
	// POSIÇÃO INICIAL DOS ZUMBIS
	// =========================

	for k := 0; k < quantidadeZumbis; k++ {
		zombielinha[k] = rand.Intn(linhasMapa-2) + 1
		zombiecoluna[k] = rand.Intn(colunasMapa-2) + 1
	}

	// =========================
	// POSIÇÃO INICIAL DOS CAMUNDONGOS
	// =========================

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

		// Limpa a tela
		fmt.Print("\033[H\033[2J")

		// ==================================================
		// GAME OVER
		// ==================================================

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

					// Reinicia a partida
					partidaAtual++
					passos = 0
					camundongoAtual = 0

					jogadorlinha = 10
					jogadorcoluna = 25

					gameOver = false
					faseFinal = false

					// Volta para o tamanho original
					linhasMapa = 20
					colunasMapa = 50

					// Volta para 20 zumbis
					quantidadeZumbis = 20

					// Sorteia novamente os zumbis
					for k := 0; k < quantidadeZumbis; k++ {
						zombielinha[k] =
							rand.Intn(linhasMapa-2) + 1

						zombiecoluna[k] =
							rand.Intn(colunasMapa-2) + 1
					}

					// Sorteia novamente os camundongos
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

		// ==================================================
		// TÍTULO
		// ==================================================

		fmt.Println("\n>>>>> ZOMBIE CURSE CURE <<<<<")
		fmt.Println("         🧟-🐭-🧪           \n")

		// ==================================================
		// AVISO DA FASE FINAL
		// ==================================================

		if faseFinal {
			fmt.Println("==================================================")
			fmt.Println("             A FASE FINAL CHEGOU!")
			fmt.Println("==================================================")
			fmt.Println()
		}

		// ==================================================
		// MELHOR PONTUAÇÃO
		// ==================================================

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

		// ==================================================
		// CAMUNDONGOS CAPTURADOS
		// ==================================================

		fmt.Print("camundongos capturados: ")

		for c := 0; c < camundongoAtual; c++ {
			fmt.Printf("%c ", 'A'+c)
		}

		fmt.Println()

		// ==================================================
		// PARTIDA E PASSOS
		// ==================================================

		textoPartida :=
			fmt.Sprintf("partida atual: %d", partidaAtual)

		fmt.Printf(
			"%-40s %d passos\n",
			textoPartida,
			passos,
		)

		// ==================================================
		// INFORMAÇÕES DA FASE
		// ==================================================

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

		// ==================================================
		// DESENHA O MAPA
		// ==================================================

		for i := 0; i < linhasMapa; i++ {

			for j := 0; j < colunasMapa; j++ {

				// Parede superior e inferior
				if i == 0 || i == linhasMapa-1 {

					fmt.Print("-")

				// Parede esquerda e direita
				} else if j == 0 || j == colunasMapa-1 {

					fmt.Print("|")

				// Jogador
				} else if i == jogadorlinha &&
					j == jogadorcoluna {

					fmt.Print("@")

				} else {

					// ==========================================
					// PROCURA ZUMBI
					// ==========================================

					temZombieAqui = false

					for k := 0; k < quantidadeZumbis; k++ {

						if i == zombielinha[k] &&
							j == zombiecoluna[k] {

							fmt.Print("X")

							temZombieAqui = true

							break
						}
					}

					// ==========================================
					// PROCURA CAMUNDONGO
					// ==========================================

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

						// Se não tiver zumbi nem camundongo
						if !temCamundongoAqui {
							fmt.Print(" ")
						}
					}
				}
			}

			fmt.Println()
		}

		// ==================================================
		// RECEBE A TECLA
		// ==================================================

		char, key, err := keyboard.GetKey()

		if err != nil ||
			key == keyboard.KeyEsc {

			break
		}

		moveu := false

		// ==================================================
		// MOVIMENTO DO JOGADOR
		// ==================================================

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

		// ==================================================
		// CAPTURA DO CAMUNDONGO
		// ==================================================

		if camundongoAtual < 10 &&
			jogadorlinha == camundongolinha[camundongoAtual] &&
			jogadorcoluna == camundongocoluna[camundongoAtual] {

			camundongoAtual++
		}

		// ==================================================
		// ATIVA A FASE FINAL
		// ==================================================

		if camundongoAtual == 10 &&
			!faseFinal {

			faseFinal = true

			// Aumenta o mapa
			linhasMapa = 25
			colunasMapa = 70

			// Aumenta a quantidade de zumbis
			quantidadeZumbis = 35

			// Coloca os novos zumbis
			// em posições aleatórias
			for k := 0; k < quantidadeZumbis; k++ {

				zombielinha[k] =
					rand.Intn(linhasMapa-2) + 1

				zombiecoluna[k] =
					rand.Intn(colunasMapa-2) + 1
			}
		}

		// ==================================================
		// MOVIMENTAÇÃO DOS PERSONAGENS
		// ==================================================

		if moveu {

			passos++

			// ==========================================
			// MOVIMENTO DO CAMUNDONGO
			// ==========================================

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

			// ==========================================
			// MOVIMENTO DOS ZUMBIS
			// ==========================================

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

				// ==========================================
				// VERIFICA COLISÃO COM O JOGADOR
				// ==========================================

				if jogadorlinha == zombielinha[k] &&
					jogadorcoluna == zombiecoluna[k] {

					gameOver = true
				}
			}
		}
	}
}
