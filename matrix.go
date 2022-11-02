package main

import (
	"fmt"
	"math/rand"
	"time"
)

var Teste = 3

func createMatrix(size_row int, size_column int) [][]float64 {

	matrix := make([][]float64, size_row)
	for i := range matrix {
		matrix[i] = make([]float64, size_column)
	}

	return matrix
}

func buildRandomMatrix(size int) [][]float64 {

	matrix := createMatrix(size, size)

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			matrix[i][j] = r1.Float64()
		}
	}

	return matrix
}

func printMatrix(matrix [][]float64) {
	size := len(matrix)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%f  ", matrix[i][j])
		}
		fmt.Println()
	}
}

func transpose(matrix [][]float64) [][]float64 {
	size := len(matrix)
	transpose := createMatrix(size, size)

	size_row := len(matrix)
	size_column := len(matrix[0])
	transpose = createMatrix(size_column, size_row)

	for i := 0; i < size_row; i++ {
		for j := 0; j < size_column; j++ {
			transpose[j][i] = matrix[i][j]
		}
	}
	return transpose
}
