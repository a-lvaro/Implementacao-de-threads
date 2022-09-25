package main

import (
	"fmt"
)

func main() {

	size := 2
	matrixA := buildRandomMatrix(size)
	matrixB := buildRandomMatrix(size)

	fmt.Println("MATRIX A")
	printMatrix(matrixA)

	fmt.Println()
	fmt.Println()

	fmt.Println("MATRIX B")
	printMatrix(matrixB)

	fmt.Println()
	fmt.Println()

	fmt.Println("SUM")
	printMatrix(sumMatrix(matrixA, matrixB))

	fmt.Println()
	fmt.Println()

	fmt.Println("SUB")
	printMatrix(subMatrix(matrixA, matrixB))

	fmt.Println()
	fmt.Println()

	fmt.Println("MUL")
	printMatrix(multMatrix(matrixA, matrixB))
}
