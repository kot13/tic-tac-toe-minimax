package main

import (
	"bufio"
	"fmt"
	"os"
)

const Empty string = ""
const Human string = "1"
const Computer string = "2"

type Move struct {
	score int
	cell  string
}

var gameOver bool = false
var turn int = 0
var currentTurn string = Human
var battleField = map[string]string{
	"A1": Empty,
	"A2": Empty,
	"A3": Empty,
	"B1": Empty,
	"B2": Empty,
	"B3": Empty,
	"C1": Empty,
	"C2": Empty,
	"C3": Empty,
}

func main() {
	fmt.Printf("Tic-tac-toe.\nWho goes first? 1 - player, 2 - computer \n")
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		currentTurn = sc.Text()

		if currentTurn != Human && currentTurn != Computer {
			fmt.Printf("Invalid input. 1 or 2?\n")
			continue
		}

		turn++
		break
	}

	for turn = 1; turn <= 10; turn++ {
		if gameOver {
			break
		}

		if turn == 10 {
			fmt.Printf("Draw!\n")
			break
		}

		if currentTurn == Computer {
			cell := computerTurn()
			battleField[cell] = Computer
			fmt.Printf("Computer: %s\n", cell)

			printBattleField(battleField)

			if turn > 3 {
				if checkVictory(battleField, Computer) {
					fmt.Printf("The computer won!\n")
					break
				}
			}

			currentTurn = Human
			continue
		}

		if currentTurn == Human {
			fmt.Printf("You: ")
			for sc.Scan() {
				cell := sc.Text()
				who, ok := battleField[cell]
				if !ok {
					fmt.Printf("Invalid field! \n")
					continue
				}

				if who != Empty {
					fmt.Printf("Invalid field! \n")
					continue
				}

				battleField[cell] = Human

				printBattleField(battleField)

				if turn > 3 {
					if checkVictory(battleField, Human) {
						fmt.Printf("You won!\n")
						gameOver = true
						break
					}
				}

				currentTurn = Computer
				break
			}
		}
	}
}

func checkVictory(field map[string]string, player string) bool {
	if field["A1"] == player && field["A2"] == player && field["A3"] == player {
		return true
	}
	if field["B1"] == player && field["B2"] == player && field["B3"] == player {
		return true
	}
	if field["C1"] == player && field["C2"] == player && field["C3"] == player {
		return true
	}

	if field["A1"] == player && field["B1"] == player && field["C1"] == player {
		return true
	}
	if field["A2"] == player && field["B2"] == player && field["C2"] == player {
		return true
	}
	if field["A3"] == player && field["B3"] == player && field["C3"] == player {
		return true
	}

	if field["A1"] == player && field["B2"] == player && field["C3"] == player {
		return true
	}
	if field["A3"] == player && field["B2"] == player && field["C1"] == player {
		return true
	}

	return false
}

func computerTurn() string {
	bestMove := minimax(battleField, Computer)

	return bestMove.cell
}

func emptyCell(field map[string]string) (availableCells []string) {
	for cell, player := range field {
		if player == Empty {
			availableCells = append(availableCells, cell)
		}
	}

	return availableCells
}

func printBattleField(field map[string]string) {
	fmt.Printf("%s|%s|%s\n", fmtCell(field["A1"]), fmtCell(field["A2"]), fmtCell(field["A3"]))
	fmt.Printf("------\n")
	fmt.Printf("%s|%s|%s\n", fmtCell(field["B1"]), fmtCell(field["B2"]), fmtCell(field["B3"]))
	fmt.Printf("------\n")
	fmt.Printf("%s|%s|%s\n", fmtCell(field["C1"]), fmtCell(field["C2"]), fmtCell(field["C3"]))
}

func fmtCell(cell string) string {
	if cell == Empty {
		return " "
	}
	if cell == Computer {
		return "X"
	}
	return "O"
}

func minimax(field map[string]string, player string) (bestMove Move) {
	availableCells := emptyCell(field)

	if checkVictory(field, Human) {
		bestMove.score = -1
		return
	}

	if checkVictory(field, Computer) {
		bestMove.score = 1
		return
	}

	if len(availableCells) == 0 {
		bestMove.score = 0
		return
	}

	bestMove.score = -1000
	if player == Human {
		bestMove.score = 1000
	}

	for _, cell := range availableCells {
		field[cell] = player

		opponent := Human
		if player == opponent {
			opponent = Computer
		}

		m := minimax(field, opponent)
		m.cell = cell

		field[cell] = Empty

		if player == Computer {
			if m.score > bestMove.score {
				bestMove = m
			}
		} else {
			if m.score < bestMove.score {
				bestMove = m
			}
		}
	}

	return
}
