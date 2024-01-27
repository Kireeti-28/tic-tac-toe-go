package main

import (
	"fmt"

	"github.com/kireeti-28/tic-tac-toe-go/internal/tictactoe"
)


func main() {
	fmt.Println("Hello, Welcome 😊")
	tictactoe := tictactoe.NewTicTacToe()
	tictactoe.Start()
}
