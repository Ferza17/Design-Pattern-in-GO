package main

import "fmt"

func PlayGame(
	start, takeTurn func(),
	haveWinner func() bool,
	winningPlayer func() int,
) {
	start()
	for ; !haveWinner(); {
		takeTurn()
	}
	fmt.Printf("Player %d wins. \n", winningPlayer())
}

func main() {
	turn, maxTurn, currentPlayer := 1, 10, 0

	start := func() {
		fmt.Println("Starting game of chess")
	}

	takeTurn := func() {
		turn++
		fmt.Printf("Turn %d taken by player %d\n", turn, currentPlayer)
		currentPlayer = (currentPlayer + 1) % 2
	}

	haveWinner := func() bool {
		return turn == maxTurn
	}

	winningPlayer := func() int {
		return currentPlayer
	}

	PlayGame(start, takeTurn, haveWinner, winningPlayer)

}
