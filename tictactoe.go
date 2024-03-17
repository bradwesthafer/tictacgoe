package main

import (
	"bufio"
	"errors"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	println("Welcome to Tic Tac Goe")

	for {
		print("This game involves a board of x by x squares where x is at least 3. Please enter a value for x: ")
		scanner.Scan()
		size := scanner.Text()
		boardSize, err := checkValidBoardSize(size)
		if err != nil {
			println(err.Error())
		} else {
			playGame(boardSize, scanner)
			print("Would you like to play again? Enter Y or N: ")
			scanner.Scan()
			input := scanner.Text()
			if input[0] != 'Y' && input[0] != 'y' {
				break
			}
		}
	}
}

func checkValidBoardSize(boardSize string) (int, error) {
	ret, err := strconv.Atoi(boardSize)
	if err == nil && ret < 3 {
		return ret, errors.New("Invalid board size")
	}
	return ret, err
}

func playGame(boardSize int, scanner *bufio.Scanner) {
	// Construct board as an array of bytes (alias for uint 8)
	// 0 = empty (automatically initialized to this)
	// 88 = 'X'
	// 79 = 'O'
	// These numbers are ASCII values
	board := make([][]byte, boardSize)
	for i := 0; i < boardSize; i++ {
		board[i] = make([]byte, boardSize)
	}

	print("You can play as X or O. X moves first. You can play as both by entering 2, " +
		"as X by entering X or x or as O by entering anything else. Please enter your choice: ")
	scanner.Scan()
	choice := scanner.Text()
	if choice[0] == '2' {
		println("You are playing as both X and O")
		choice = "2"
	} else if choice[0] == 'X' || choice[0] == 'x' {
		println("You are playing as X")
		choice = "X"
	} else {
		println("You are playing as O")
		choice = "O"
	}
	xMovesNext := true

	for {
		isWinner, isDraw := checkForWinner(board, xMovesNext)
		printBoard(board)

		if isWinner == true && xMovesNext {
			if choice == "2" {
				println("O wins")
			} else if choice == "O" {
				println("Congratulations, you won!")
			} else {
				println("Sorry, you lost. Better luck next time.")
			}
			break
		} else if isWinner {
			if choice == "2" {
				println("X wins")
			} else if choice == "X" {
				println("Congratulations, you won!")
			} else {
				println("Sorry, you lost. Better luck next time.")
			}
			break
		} else if isDraw {
			println("There are no valid moves left. The game is a draw.")
			break
		}

		if (xMovesNext && choice == "O") || (!xMovesNext && choice == "X") {
			row, column := determineComputerMove(board, xMovesNext)
			println("The computer picked " + strconv.Itoa(row) + " " + strconv.Itoa(column) + " as its move.")
			if xMovesNext {
				board[row][column] = 'X'
			} else {
				board[row][column] = 'O'
			}
		} else {
			println("Moves are entered as 2 numbers separated by a space.")
			println("The first number is the row and the second number is the column.")
			println("The top leftmost square of the board is 0 0. The top rightmost square " +
				"is 0 " + strconv.Itoa(boardSize-1) + ".")
			for {
				if xMovesNext {
					print("It is X's turn. Please enter your move: ")
				} else {
					print("It is O's turn. Please enter your move: ")
				}
				scanner.Scan()
				move := scanner.Text()
				row, column, valid := validateMove(board, move)
				if valid {
					if xMovesNext {
						board[row][column] = 'X'
					} else {
						board[row][column] = 'O'
					}
					break
				} else {
					println(move + " is an invalid move.")
				}
			}
		}

		xMovesNext = !xMovesNext
	}
	return
}

// First value returns true if the previous move resulted in a win, false otherwise
// Second value returns true if draw (i.e. no more moves available), false otherwise
func checkForWinner(board [][]byte, xMovesNext bool) (bool, bool) {
	size := len(board)
	// Assume it is a draw until we see the first blank space
	// Assume it isn't a win until proven otherwise.
	// We only check for 'O' wins if xMovesNext or for 'X' wins if !xMovesNext
	// These assumptions simplify the code
	isWin, isDraw := false, true

	var nextMove, lastMove byte = 79, 88
	if xMovesNext {
		nextMove = 88
		lastMove = 79
	}

	// Values are board[i][j]. Need to check every j for a given i, every i for a given j and both diagonals
	// The diagonals are [0][0] to [n][n] for size = n and also [0][n] to [n][0].
	// For the first diagonal, i = j. For the second, j = (n - i - 1)

	// Check rows first. We also check for draws here
	for i := 0; i < size; i++ {
		allMatch := true
		for j := 0; j < size; j++ {
			switch board[i][j] {
			case 0:
				isDraw = false
				allMatch = false
				break
			case nextMove:
				allMatch = false
				if !isDraw {
					break
				}
			}
		}
		if allMatch {
			isWin = true
			break
		}
	}

	// Check columns next. No need to check for draws because we've already done so.
	// Also no need to check if we've already got a winner
	for j := 0; j < size; j++ {
		if isWin {
			break
		}
		allMatch := true
		for i := 0; i < size; i++ {
			if board[i][j] != lastMove {
				allMatch = false
				break
			}
		}
		if allMatch {
			isWin = true
			break
		}
	}

	// Next, do the i = j diagonal
	allMatch := true
	for i := 0; i < size; i++ {
		if isWin {
			break
		}
		if board[i][i] != lastMove {
			allMatch = false
			break
		}
	}
	if allMatch {
		isWin = true
	}

	// Finally, do the j = (n - i - 1) diagonal
	allMatch = true
	for i := 0; i < size; i++ {
		if isWin {
			break
		}
		if board[i][size-i-1] != lastMove {
			allMatch = false
			break
		}
	}
	if allMatch {
		isWin = true
	}

	return isWin, isDraw
}

// Prints the current state of the board to stdout
// The format is an ASCII table with '-' for borders, ' ' for empty spaces, 'X' for X's and 'O' for O's
func printBoard(board [][]byte) {
	size := len(board)
	borderLine := strings.Repeat("-", 1+(size*2))
	for i := 0; i < size; i++ {
		println(borderLine)
		line := "-"
		for j := 0; j < size; j++ {
			switch board[i][j] {
			case 0:
				line += " "
			case 88:
				line += "X"
			case 79:
				line += "O"
			}
			line += "-"
		}
		println(line)
	}
	println(borderLine)
}

// Parses and validates the human move. If valid, the bool is true and the 2 ints contain the row and column.
func validateMove(board [][]byte, move string) (int, int, bool) {
	splitMove := strings.Split(move, " ")
	row, err := strconv.Atoi(splitMove[0])
	if err != nil {
		println(err.Error())
		return -1, -1, false
	}
	col, err2 := strconv.Atoi(splitMove[1])
	if err2 != nil {
		println(err2.Error())
		return -1, -1, false
	}
	return row, col, isValidMove(board, row, col)
}

// This function determines if a move is valid (i.e. is a blank space on the board)
// Returns true if valid, false otherwise
func isValidMove(board [][]byte, row int, col int) bool {
	size := len(board)
	if row >= size || col >= size || board[row][col] != 0 {
		return false
	}
	return true
}

// This function returns the row and column of the computer's next move.
// The algorithm is simple:
// 1) If there is a winning move, it chooses that move.
// 2) If there is a move that blocks the player from winning, it chooses that move
// 3) Otherwise, it generates 2 random numbers corresponding to a move. If the move is valid, it uses that move.
func determineComputerMove(board [][]byte, xMovesNext bool) (int, int) {
	size := len(board)

	// As with the check for winner function, we check each row then each column then both diagonals
	// This time, we need to find size - 1 X's or size - 1 O's along the row/column/diagonal

	// We do this via 2 calls to an external function because step 1 and step 2 are going to be the same with
	// the exception that one checks if there's a winning move for X and the other checks for winning moves for O
	var computer, player byte = 88, 79
	if !xMovesNext {
		computer = 79
		player = 88
	}
	var row, col int
	var winnable bool
	row, col, winnable = canWinOnNextMove(board, computer)
	if winnable {
		return row, col
	} else {
		row, col, winnable = canWinOnNextMove(board, player)
		if winnable {
			return row, col
		}
	}

	for {
		row = rand.IntN(size)
		col = rand.IntN(size)
		if isValidMove(board, row, col) {
			return row, col
		}
	}
}

// This function helps determine the computer's move. Should only be called from there.
func canWinOnNextMove(board [][]byte, player byte) (int, int, bool) {
	size := len(board)

	// rows
	for i := 0; i < size; i++ {
		countPlayer := 0
		countBlank := 0
		blankJ := -1
		for j := 0; j < size; j++ {
			switch board[i][j] {
			case 0:
				countBlank++
				blankJ = j
			case player:
				countPlayer++
			}
			if countBlank > 1 {
				break
			}
		}
		if countBlank == 1 && countPlayer == (size-1) {
			return i, blankJ, true
		}
	}

	// columns
	for j := 0; j < size; j++ {
		countPlayer := 0
		countBlank := 0
		blankI := -1
		for i := 0; i < size; i++ {
			switch board[i][j] {
			case 0:
				countBlank++
				blankI = i
			case player:
				countPlayer++
			}
			if countBlank > 1 {
				break
			}
		}
		if countBlank == 1 && countPlayer == (size-1) {
			return blankI, j, true
		}
	}

	// i = j diagonal
	countPlayer := 0
	countBlank := 0
	blank := -1
	for i := 0; i < size; i++ {
		switch board[i][i] {
		case 0:
			countBlank++
			blank = i
		case player:
			countPlayer++
		}
		if countBlank > 1 {
			break
		}
	}
	if countBlank == 1 && countPlayer == (size-1) {
		return blank, blank, true
	}

	// j = (n - i - 1) diagonal
	countPlayer = 0
	countBlank = 0
	blank = -1
	for i := 0; i < size; i++ {
		switch board[i][size-i-1] {
		case 0:
			countBlank++
			blank = i
		case player:
			countPlayer++
		}
		if countBlank > 1 {
			break
		}
	}
	if countBlank == 1 && countPlayer == (size-1) {
		return blank, (size - blank), true
	}

	return -1, -1, false
}
