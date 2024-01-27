package tictactoe

import (
	"errors"
	"fmt"
)

type TicTacToe struct {
	board [][]int
}

func NewTicTacToe() *TicTacToe{
	return &TicTacToe{
		board: [][]int{{0,0,0},{0,0,0},{0,0,0}},
	}
}

func (t *TicTacToe) resetBoard() {
	t.board = [][]int{{0,0,0},{0,0,0},{0,0,0}}
}

func (t *TicTacToe) isGameOver() bool {
	firstRow := t.board[0][0] == t.board[0][1] && t.board[0][1] == t.board[0][2] && t.board[0][0] > 0
	secondRow := t.board[1][0] == t.board[1][1] && t.board[1][1] == t.board[1][2] && t.board[1][0] > 0
	thirdRow := t.board[2][0] == t.board[2][1] && t.board[2][1] == t.board[2][2] && t.board[2][0] > 0

	diagonalLR := t.board[0][0] == t.board[1][1] && t.board[1][1] == t.board[2][2] && t.board[0][0] > 0
	diagonalRL := t.board[0][2] == t.board[1][1] && t.board[1][1] == t.board[2][0] && t.board[1][1] > 0

	firstCol := t.board[0][0] == t.board[1][0] && t.board[1][0] == t.board[2][0] && t.board[0][0] > 0
	secondCol := t.board[0][1] == t.board[1][1] && t.board[1][1] == t.board[2][1] && t.board[1][1] > 0
	thirdCol := t.board[0][2] == t.board[1][2] && t.board[1][2] == t.board[2][2] && t.board[1][2] > 0

	return firstRow || secondRow || thirdRow || diagonalLR || diagonalRL || firstCol || secondCol || thirdCol
}

func (t *TicTacToe) drawBoard() {
	out := ""

	for _, v := range t.board {
		row := ""
		for _, vv := range v {
			if vv == 0 {
				row += "-  "
			} else if vv == 1 {
				row += "o  "
			} else {
				row += "x  "
			}
		}
		out += row + "\n"
	}

	fmt.Print(out)
}

func (t *TicTacToe) updateBoard(choice int, player int) error {

		switch choice {
		case 1:
			if t.board[0][0] != 0 {
				return errors.New("Invalid choice. Spot choosen was already used\n")
			}
			t.board[0][0] = player
			break
		case 2:
			if t.board[0][1] != 0 {
				return errors.New("Invalid choice. Spot choosen was already used\n")
			}
			t.board[0][1] = player
			break
		case 3:
			if t.board[0][2] != 0 {
				return errors.New("Invalid choice. Spot choosen was already used\n")
			}
			t.board[0][2] = player
			break
		case 4:
			if t.board[1][0] != 0 {
				return errors.New("Invalid choice. Spot choosen was already used\n")
			}
			t.board[1][0] = player
			break
		case 5:
			if t.board[1][1] != 0 {
				return errors.New("Invalid choice. Spot choosen was already used\n")
			}
			t.board[1][1] = player
			break
		case 6:
			if t.board[1][2] != 0 {
				return errors.New("Invalid choice. Spot choosen was already used\n")
			}
			t.board[1][2] = player
			break
		case 7:
			if t.board[2][0] != 0 {
				return errors.New("Invalid choice. Spot choosen was already used\n")
			}
			t.board[2][0] = player
			break
		case 8:
			if t.board[2][1] != 0 {
				return errors.New("Invalid choice. Spot choose was already used\n")
			}
			t.board[2][1] = player
			break
		case 9:
			if t.board[2][2] != 0 {
				return errors.New("Invalid choice. Spot choose was already used\n")
			}
			t.board[2][2] = player
			break
		}

		return nil
}

func (t *TicTacToe) readPlayersChoice(player int) (error) {
		player++
		fmt.Printf("Player %d, It's your turn!\nChoose your spot in the board\n", player)
		var choice int
		fmt.Scan(&choice)

		if choice < 1 || choice > 9 {
			return errors.New("Invalid choice. Choose right spot(1-9)\n")
		}
		
		return t.updateBoard(choice, player)
}

func (t * TicTacToe) Start() {
	t.resetBoard() // just to make sure we start with right board values
	t.drawBoard()

	player := 0 // default starts with player 1

	for true{
		for !t.isGameOver() {
			err := t.readPlayersChoice(player)
			if err != nil {
				fmt.Println(err)
				continue
			}
			t.drawBoard()
			player++
			player = player % 2
		}
		break
	}
}