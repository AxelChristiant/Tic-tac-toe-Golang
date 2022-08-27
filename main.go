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

func displayBoard() {
	for i := 1; i <= 9; i++ {
		if i%3 == 0 {
			fmt.Printf("%s\n", board[i-1])
		} else {
			fmt.Printf("%s|", board[i-1])
		}
	}
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

func checkingBoard() bool {
	return isConsecutiveRow() || isConsecutiveColumn() || isConsecutiveDiagonal()
}

func main() {
	fmt.Println(checkingBoard())
}
