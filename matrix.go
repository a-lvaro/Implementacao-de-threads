package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Teste = 3

func createMatrix(size_row int, size_column int) [][]int {

	matrix := make([][]int, size_row)
	for i := range matrix {
		matrix[i] = make([]int, size_column)
	}

	return matrix
}

func buildRandomMatrix(size int) [][]int {

	matrix := createMatrix(size, size)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = r1.Intn(10)
		}
	}

	return matrix
}

func printMatrix(matrix [][]int) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d  ", matrix[i][j])
		}
		fmt.Println()
	}
}

func transpose(matrix [][]int) [][]int {
	size := len(matrix)
	transpose := createMatrix(size, size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			transpose[j][i] = matrix[i][j]
		}
	}

	return transpose
}
