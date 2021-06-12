package main

import (
	"fmt"
)

func main() {
	board := [9]int{0,0,0,0,0,0,0,0,0}
	gameOver := false
	turn := 1

	showInstructions()

	for gameOver != true {
		displayBoard(board)
		player := 0
		if turn % 2 == 1 {
			fmt.Println("Player 1's turn")
			player = 1
		} else {
			fmt.Println("Player 2's turn")
			player = 2
		}

		currentMove := askToPlay()

		if currentMove == 9 {
			return
		}

		board = executeMove(currentMove, player, board)
		result := checkForWin(board)

		if result != 0 {
			fmt.Printf("Player %d wins\n\n", result)
			displayBoard(board)
			gameOver = true
		} else if turn == 9 {
			fmt.Printf("Tie game!\n\n")
			displayBoard(board)
			gameOver = true
		} else {
			turn++
		}
	}
}

func showInstructions() {
	board := [9]int{2,2,2,2,2,2,2,2,2}
	fmt.Println("Each digit in the boxes represent their play numbers")
	displayBoard(board)
	fmt.Printf("\n\n")
}

func displayBoard(board [9]int) {
	for i, v := range board {
		if v == 0 {
			fmt.Printf(" _ ")
		} else if v == 1 {
			fmt.Printf(" X ")
		} else if v == 10 {
			fmt.Printf(" O ")
		} else {
			fmt.Printf(" %d ", i)
		}

		if i > 0 && (i + 1) % 3 == 0 {
			fmt.Printf("\n")
		} else {
			fmt.Printf("|")
		}
	}
}

func askToPlay() int {
	var mv int
	fmt.Println("Select a move")
	fmt.Scan(&mv)
	return mv
}

func executeMove(move, player int, board [9]int) [9]int {
	 if board[move] != 0 {
		fmt.Println("Please pick an available box")
		move =  askToPlay()
		board = executeMove(move, player, board)
	 } else {
		if player == 1 {
			board[move] = 1
		} else if player == 2 {
			board[move] = 10
		}
	 }

 	for move > 9 {
 		fmt.Println("Please enter a number less than 10")
 		move = askToPlay()
	}

	if player == 1 {
		board[move] = 1
	} else if player == 2 {
		board[move] = 10
	}

	return board
}

func checkForWin(b [9]int) int {
	sums := [8] int {0,0,0,0,0,0,0,0}
	sums[0] = b[2]+b[4]+b[6]
	sums[1] = b[0]+b[3]+b[6]
	sums[2] = b[1]+b[4]+b[7]
	sums[3] = b[2]+b[5]+b[8]
	sums[4] = b[0]+b[4]+b[8]
	sums[5] = b[6]+b[7]+b[8]
	sums[6] = b[3]+b[4]+b[5]
	sums[7] = b[0]+b[1]+b[2]

	for _, v := range sums {
		if v == 3 {
			return 1
		} else if v == 30{
			return 2
		}
	}
	return 0
}
