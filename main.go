package main

import (
	"fmt"
	"time"
)

func main() {

	size := 1000

	matrixA := buildRandomMatrix(size)
	matrixB := buildRandomMatrix(size)

	// SHOW MATRIX
	fmt.Printf("\n------- MATRIXA & MATRIX B -------\n")
	// printMatrix(matrixA)
	// fmt.Printf("\n\n\n")
	// printMatrix(matrixB)
	// fmt.Printf("\n\n\n")

	fmt.Printf("\n--------- SUM ---------\n")
	// SEQUENCIAL SUM
	fmt.Println("\nSequential SUM")
	start := time.Now()

	c := make(chan [][]int, size)
	go threadsMatrix(matrixA, matrixB, c, "SUM")
	_ = sumMatrix(matrixA, matrixB)

	elapsed := time.Since(start)
	fmt.Printf("SUM took %s\n", elapsed)

	// CHANEL SUM
	fmt.Println("\n\nCHANNEL SUM")
	// start = time.Now()

	for v := range c {
		_ = v
	}

	elapsed = time.Since(start)
	fmt.Printf("SUM channel took %s\n", elapsed)

	fmt.Printf("\n\n\n--------- SUBTRACTION ---------\n")
	// SEQUENCIAL SUB
	fmt.Printf("\nSEUQUENTIAL SUBTRACTION\n")
	start = time.Now()

	_ = subMatrix(matrixA, matrixB)

	elapsed = time.Since(start)
	fmt.Printf("SUBTRACTION took %s\n", elapsed)

	// CHANEL SUB
	fmt.Printf("\n\nCHANNEL SUBTRACTION\n")
	start = time.Now()

	c = make(chan [][]int, size)
	go threadsMatrix(matrixA, matrixB, c, "SUB")

	for v := range c {
		_ = v
	}

	elapsed = time.Since(start)
	fmt.Printf("SUBTRACTION channel took %s\n", elapsed)

	fmt.Printf("\n\n\n--------- MULTIPLICATION ---------\n")
	// SEQUENTIAL MULT
	fmt.Printf("\nSEQUENTIAL MULTIPLICATION\n")
	start = time.Now()

	_ = multMatrix(matrixA, matrixB)

	elapsed = time.Since(start)
	fmt.Printf("MULTIPLICATION took %s\n", elapsed)
	fmt.Print("\n\n")

	// CHANEL MULT
	fmt.Printf("\n\nCHANEL MULTIPLICATION\n")
	start = time.Now()

	c = make(chan [][]int, size)
	go threadsMatrix(matrixA, matrixB, c, "MULT")

	for v := range c {
		_ = v
	}

	elapsed = time.Since(start)
	fmt.Printf("MULTIPLICATION channel took %s\n", elapsed)

	fmt.Printf("\n\n\n--------- TRANSPOSE ---------\n")
	// SEQUENTIAL TRANSPOSE
	fmt.Printf("\nMATRIX TRANSPOSE\n")
	start = time.Now()

	_ = transpose(matrixA)

	elapsed = time.Since(start)
	fmt.Printf("TRANSPOSE channel took %s\n", elapsed)

	// CHANEL TRANSPOSE
	fmt.Printf("\n\nCHANEL TRANSPOSE\n")
	start = time.Now()

	c = make(chan [][]int, size)
	go threadsMatrix(matrixA, matrixA, c, "TRAN")

	for v := range c {
		_ = v
	}

	elapsed = time.Since(start)
	fmt.Printf("TRANSPOSE channel took %s\n", elapsed)

}
