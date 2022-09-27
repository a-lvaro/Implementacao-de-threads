package main

import (
	"fmt"
	"time"
)

func main() {

	size := 4

	matrixA := buildRandomMatrix(size)
	matrixB := buildRandomMatrix(size)

	// SHOW MATRIX
	fmt.Printf("\n------- MATRIXA & MATRIX B -------\n")
	printMatrix(matrixA)
	fmt.Printf("\n\n\n")
	printMatrix(matrixB)
	fmt.Printf("\n\n\n")

	fmt.Printf("\n--------- SUM ---------\n")
	// SEQUENCIAL SUM
	fmt.Println("\nSequential SUM")
	start := time.Now()

	printMatrix(sumMatrix(matrixA, matrixB))

	elapsed := time.Since(start)
	fmt.Printf("SUM took %s\n", elapsed)

	// CHANEL SUM
	fmt.Println("\n\nCHANEL SUM")
	start = time.Now()

	c := make(chan [][]int)
	go threadsMatrix(matrixA, matrixB, c, "SUM")

	for v := range c {
		fmt.Println(v)
	}

	elapsed = time.Since(start)
	fmt.Printf("SUM chanel took %s\n", elapsed)

	fmt.Printf("\n\n\n--------- SUB ---------\n")
	// SEQUENCIAL SUB
	fmt.Printf("\nSEUQUENTIAL SUB\n")
	start = time.Now()

	printMatrix(subMatrix(matrixA, matrixB))

	elapsed = time.Since(start)
	fmt.Printf("SUB took %s\n", elapsed)

	// CHANEL SUB
	fmt.Printf("\n\nCHANEL SUB\n")
	start = time.Now()

	c = make(chan [][]int)
	go threadsMatrix(matrixA, matrixB, c, "SUB")

	for v := range c {
		fmt.Println(v)
	}

	elapsed = time.Since(start)
	fmt.Printf("SUB chanel took %s\n", elapsed)

	fmt.Printf("\n\n\n--------- MULT ---------\n")
	// SEQUENTIAL MULT
	fmt.Printf("\nSEQUENTIAL MULT\n")
	start = time.Now()

	printMatrix(multMatrix(matrixA, matrixB))

	elapsed = time.Since(start)
	fmt.Printf("MULT took %s\n", elapsed)
	fmt.Print("\n\n")

	// CHANEL MULT
	fmt.Printf("\n\nCHANEL MULT\n")
	start = time.Now()

	c = make(chan [][]int)
	go threadsMatrix(matrixA, matrixB, c, "MULT")

	for v := range c {
		fmt.Println(v)
	}
}
