package main

import (
	"fmt"
	"time"
)

func main() {

	size := 10
	matrixA := buildRandomMatrix(size)
	matrixB := buildRandomMatrix(size)

	// teste(matrixA)

	// fmt.Println(matrixA)

	// for i := 0; i < len(matrixA); i++ {
	// 	fmt.Println(matrixA[i : i+1])
	// }

	fmt.Println("\nSequential")
	start := time.Now()

	sumMatrix(matrixA, matrixB)

	elapsed := time.Since(start)
	fmt.Printf("SUM took %s", elapsed)

	// printMatrix(matrixA)
	fmt.Printf("\n\n")
	// printMatrix(matrixB)

	// fmt.Print("\n\n\n")

	fmt.Println("\n\nCHANEL")
	start = time.Now()

	c := make(chan [][]int)
	go splitMatrix(matrixA, matrixB, c)

	// for v := range c {
	// 	fmt.Println(v)
	// }

	elapsed = time.Since(start)
	fmt.Printf("SUM chanel took %s", elapsed)

	fmt.Print("\n\n")

	// ----------------------------------
	// fmt.Println("MATRIX B")
	// printMatrix(matrixB)

	// fmt.Print("\n\n")

	// // SUM
	// start := time.Now()

	// sum := sumMatrix(matrixA, matrixB)

	// elapsed := time.Since(start)
	// fmt.Printf("\n\nSUM took %s", elapsed)

	// fmt.Println("SUM")
	// printMatrix(sum)

	// fmt.Print("\n\n")

	// // SUB
	// fmt.Println("SUB")
	// printMatrix(subMatrix(matrixA, matrixB))

	// fmt.Print("\n\n")

	// // MULT
	// fmt.Println("MUL")
	// printMatrix(multMatrix(matrixA, matrixB))
}
