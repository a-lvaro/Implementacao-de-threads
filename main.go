package main

import (
	"fmt"
	"time"
)

func main() {

	size := 2
	matrixA := buildRandomMatrix(size)
	matrixB := buildRandomMatrix(size)

	fmt.Println("MATRIX A")
	printMatrix(matrixA)

	fmt.Print("\n\n")

	fmt.Println("MATRIX B")
	printMatrix(matrixB)

	fmt.Print("\n\n")

	// SUM
	start := time.Now()

	sum := sumMatrix(matrixA, matrixB)

	elapsed := time.Since(start)
	fmt.Printf("\n\nSUM took %s", elapsed)

	fmt.Println("SUM")
	printMatrix(sum)

	fmt.Print("\n\n")

	// SUB
	fmt.Println("SUB")
	printMatrix(subMatrix(matrixA, matrixB))

	fmt.Print("\n\n")

	// MULT
	fmt.Println("MUL")
	printMatrix(multMatrix(matrixA, matrixB))
}
