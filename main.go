package main

import (
	"fmt"
	"reflect"
)

var board = [9]string{
	"-", "-", "-",
	"-", "-", "-",
	"-", "-", "-"}
var ThreeConsecutiveO = [3]string{
	"O", "O", "O"}
var ThreeConsecutiveX = [3]string{
	"X", "X", "X"}
var playerTurn = "player 1"
var playerMark = map[string]string{"player 1": "O", "player 2": "X"}

func displayBoard() {
	fmt.Println()
	for i := 1; i <= 9; i++ {
		if i%3 == 0 {
			fmt.Printf("%s\n", board[i-1])
		} else {
			fmt.Printf("%s|", board[i-1])
		}
	}
	fmt.Println()
}

func isConsecutiveColumn() bool {
	for i := 0; i < 7; i += 3 {
		if reflect.DeepEqual(board[i:i+3], ThreeConsecutiveO[:3]) || reflect.DeepEqual(board[i:i+3], ThreeConsecutiveX[:3]) {
			return true
		}
	}
	return false
}

func isConsecutiveRow() bool {
	var rowElements []string
	for i := 0; i < 3; i++ {
		rowElements = nil
		rowElements = append(rowElements, board[i], board[i+3], board[i+6])
		if reflect.DeepEqual(rowElements, ThreeConsecutiveO[:3]) || reflect.DeepEqual(rowElements, ThreeConsecutiveX[:3]) {
			return true
		}
	}
	return false
}

func isConsecutiveDiagonal() bool {
	leftDiagonal := []string{board[0], board[4], board[8]}
	rightDiagonal := []string{board[2], board[4], board[6]}
	if reflect.DeepEqual(leftDiagonal, ThreeConsecutiveO[:3]) || reflect.DeepEqual(leftDiagonal, ThreeConsecutiveX[:3]) {
		return true
	} else if reflect.DeepEqual(rightDiagonal, ThreeConsecutiveO[:3]) || reflect.DeepEqual(rightDiagonal, ThreeConsecutiveX[:3]) {
		return true
	}
	return false
}

func thereIsAWinner() bool {
	if isConsecutiveRow() || isConsecutiveColumn() || isConsecutiveDiagonal() {
		fmt.Printf("%s Win!\n", playerTurn)
		return true
	} else {
		return false
	}
}
func thereIsEmptyField() bool {
	for _, e := range board {
		if e == "-" {
			return true
		}
	}
	fmt.Println("Draw")
	return false
}
func playerMove() {
	var idx int
	fmt.Println("Choose the empty field with the number 1-9 : ")
	fmt.Scanln(&idx)
	idx -= 1
	if idx < 0 || idx > 9 {
		fmt.Println("Your input is not valid!")
		playerMove()
	} else if board[idx] != "-" {
		fmt.Println("Field is already filled!")
	} else {
		board[idx] = playerMark[playerTurn]
	}

}

func gameOn() bool {
	return !thereIsAWinner() && thereIsEmptyField()
}
func changeTurn() {
	if playerTurn == "player 1" {
		playerTurn = "player 2"
	} else {
		playerTurn = "player 1"
	}
}
func main() {
	for gameOn() {
		displayBoard()
		playerMove()
		changeTurn()
	}
	displayBoard()
}
