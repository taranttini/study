
package main

import (
	"fmt"
)

func main() {
	inputA := []int {1,2,3,4,5}
	inputB := []int {3,2,1,2,6,9}

	solves(inputA, inputB)

	for idx,x := range {1,2,3} {
		fmt.Printf("%v, %v",idx,x)
	}
}

func solves(inputA []int, inputB []int) {
	var player1_wins int
	var player2_wins int

	if len(inputA) != len(inputB) {
		fmt.Print("Vetores de tamanho invalido!/n")
		return
	}

	for idx, _ := range inputA {
		if  inputA[idx] > inputB[idx] { player1_wins++ }
		if  inputA[idx] < inputB[idx] { player2_wins++ }
	}

	fmt.Printf("%v - %v \n", player1_wins, player2_wins)
}
