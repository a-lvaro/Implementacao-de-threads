package main

func splitMatrix(matrixA [][]int, matrixB [][]int, s chan<- [][]int) {
	// cpu := runtime.NumCPU()
	size := len(matrixA)

	for i := 0; i < size; i++ {
		// fmt.Println(matrixA[i : i+size/cpu])
		// fmt.Printf("[ %d : %d ] \n", i, i+size/cpu)
		// fmt.Println(matrixA[i : i+size/cpu])

		s <- sumMatrix(matrixA[i:i+1], matrixB[i:i+1])
	}
	close(s)

}

func teste(matrixA [][]int) {
	// size := len(matrixA)

	// for i := 0; i < size; i++ {
	// 	// fmt.Println(matrixA[i : i+size/cpu])
	// 	// fmt.Printf("[ %d : %d ] \n", i, i+size/cpu)
	// 	// fmt.Println(matrixA[i : i+size/cpu])

	// 	s <- sumMatrix(matrixA[i:i+size/cpu], matrixB[i:i+size/cpu])
	// }

}
