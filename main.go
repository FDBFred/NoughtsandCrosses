package main

import "fmt"

const (
	line = iota
	cross
	noughts
	empty
	underscore
)

var player = true

func main() {

	board := [5][5]int{
		{3, 0, 3, 0, 3},
		{4, 4, 4, 4, 4},
		{3, 0, 3, 0, 3},
		{4, 4, 4, 4, 4},
		{3, 0, 3, 0, 3},
	}

	fmt.Println("Welcome to Tic Tac Toe")
	fmt.Println("Player 1 is X")
	fmt.Println("And Player 2 is O")

	win := false
	for win = false; win == false; {

		board = turn(player, board)

		printBoard(board)

		if player == true {
			player = false
		} else {
			player = true
		}

		win = checkWin(board)

	}

}

func checkWin(board [5][5]int) bool {
	if board[0][0] == board[0][2] && board[0][2] == board[0][4] && board[0][0] != empty {
		return true
	} else if board[2][0] == board[2][2] && board[2][4] == board[2][4] && board[2][0] != empty {
		return true
	} else if board[4][0] == board[4][2] && board[4][2] == board[4][4] && board[4][0] != empty {
		return true
	} else if board[0][0] == board[2][0] && board[2][0] == board[4][0] && board[0][0] != empty {
		return true
	} else if board[0][2] == board[2][2] && board[2][2] == board[4][2] && board[0][2] != empty {
		return true
	} else if board[0][4] == board[2][4] && board[2][4] == board[4][4] && board[0][4] != empty {
		return true
	} else if board[0][0] == board[2][2] && board[2][2] == board[4][4] && board[0][0] != empty {
		return true
	} else if board[0][4] == board[2][2] && board[2][2] == board[4][0] && board[0][4] != empty {
		return true
	} else {
		return false
	}
}

func turn(p bool, board [5][5]int) [5][5]int {

	if p == false {
		fmt.Println("Player 2's turn. Enter position as row, col")
	} else {
		fmt.Println("Player 1's turn. Enter position as row, col")
	}

	var x int
	var y int

	scan, err := fmt.Scanf("%d, %d\n", &x, &y)
	if err != nil {
		fmt.Println("Error: ", err)
		fmt.Println(scan)
	}

	fmt.Println(scan, x, y)

	switch x {
	case 0:
		x = 0
	case 1:
		x = 0
	case 2:
		x = 2
	case 3:
		x = 4
	default:
		fmt.Println("Error: Invalid row")
	}
	switch y {
	case 0:
		y = 0
	case 1:
		y = 0
	case 2:
		y = 2
	case 3:
		y = 4
	default:
		fmt.Println("Error: Invalid column")
	}

	if board[y][x] != empty {
		fmt.Println("Error: Space already taken")
		printBoard(board)
		turn(p, board)
	}

	if scan == 2 {
		if p == true {
			board[y][x] = cross
		} else {
			board[y][x] = noughts
		}
	}
	return board
}

func printBoard(board [5][5]int) {
	for rows := 0; rows < len(board); rows++ {
		for cols := 0; cols < len(board); cols++ {

			switch board[rows][cols] {
			case line:
				fmt.Print("|")
			case empty:
				fmt.Print(" ")
			case noughts:
				fmt.Print("O")
			case cross:
				fmt.Print("X")
			case underscore:
				fmt.Print("-")
			}

			if cols == 4 {
				fmt.Println()
			}
		}

	}
}
