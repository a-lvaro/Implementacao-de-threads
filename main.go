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

	_ = sumMatrix(matrixA, matrixB)
	elapsed := time.Since(start)
	fmt.Printf("SUM took %s\n", elapsed)

	// CHANEL SUM
	fmt.Println("\n\nCHANNEL SUM")

	matrixChannel := make(chan []int, size)
	resultChannel := make(chan []int, size)

	start = time.Now()

	go threadsMatrix(matrixChannel, resultChannel, size, "SUM")
	go threadsMatrix(matrixChannel, resultChannel, size, "SUM")
	go threadsMatrix(matrixChannel, resultChannel, size, "SUM")

	for i := 0; i < size; i++ {
		matrixChannel <- matrixA[i]
		matrixChannel <- matrixB[i]
	}

	closeChannel(matrixChannel)

	for i := 0; i < size; i++ {
		_ = <-resultChannel
	}

	close(resultChannel)

	elapsed = time.Since(start)
	fmt.Printf("SUM channel took %s\n", elapsed)

	fmt.Printf("\n\n\n--------- SUBTRACTION ---------\n")
	// SEQUENCIAL SUB
	fmt.Printf("\nsequenctial subtraction\n")
	start = time.Now()

	_ = subMatrix(matrixA, matrixB)

	elapsed = time.Since(start)
	fmt.Printf("SUBTRACTION took %s\n", elapsed)

	// CHANEL SUB
	fmt.Printf("\n\nchannel subtraction\n")
	start = time.Now()

	matrixChannel = make(chan []int, size)
	resultChannel = make(chan []int, size)

	go threadsMatrix(matrixChannel, resultChannel, size, "SUB")
	go threadsMatrix(matrixChannel, resultChannel, size, "SUB")
	go threadsMatrix(matrixChannel, resultChannel, size, "SUB")

	for i := 0; i < size; i++ {
		matrixChannel <- matrixA[i]
		matrixChannel <- matrixB[i]
	}

	closeChannel(matrixChannel)

	for i := 0; i < size; i++ {
		_ = <-resultChannel
	}

	close(resultChannel)

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
	fmt.Printf("\n\nchannel multiplication\n")
	start = time.Now()

	matrixChannel = make(chan []int)
	resultChannel = make(chan []int, size*size)

	matrixAux := transpose(matrixB)
	go threadsMatrixMult(matrixChannel, resultChannel, matrixAux, size)
	go threadsMatrixMult(matrixChannel, resultChannel, matrixAux, size)
	go threadsMatrixMult(matrixChannel, resultChannel, matrixAux, size)

	for i := 0; i < size; i++ {
		matrixChannel <- matrixA[i]
	}

	close(matrixChannel)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			_ = <-resultChannel
		}
		// fmt.Println()
	}

	elapsed = time.Since(start)
	fmt.Printf("multiplication channel took %s\n", elapsed)
}
