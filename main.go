package main

import (
	"fmt"
	"math/rand"
	"time"
)

func buildMatrix(tamanho int) [][]int {

	matrix := make([][]int, tamanho)
	for i := range matrix {
		matrix[i] = make([]int, tamanho)
	}

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < tamanho; i++ {
		for j := 0; j < tamanho; j++ {
			matrix[i][j] = r1.Intn(100)
		}
	}

	return matrix
}

// func sumMatrix(matrix01 [][]int, matrix02[][]int) [][]int {
// 	for line := range matrix01
// }

func main() {

	matrixA := buildMatrix(10)

	fmt.Println("a")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d  ", matrixA[i][j])
		}
		fmt.Println()
	}

	// fmt.Println()
	// fmt.Println()

	// fmt.Println("b")
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Printf("%d  ", b[i][j])
	// 	}
	// 	fmt.Println()
	// }

	// fmt.Println("soma")
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 2; j++ {
	// 		fmt.Printf("%d  ", soma[i][j])
	// 	}
	// 	fmt.Println()
	// }
}
