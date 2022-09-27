package main

func threadsMatrix(matrixA [][]int, matrixB [][]int, s chan<- [][]int, flag string) {
	// cpu := runtime.NumCPU()
	size := len(matrixA)

	for i := 0; i < size; i++ {

		if flag == "SUM" {
			s <- sumMatrix(matrixA[i:i+1], matrixB[i:i+1])
		} else if flag == "SUB" {
			s <- subMatrix(matrixA[i:i+1], matrixB[i:i+1])
		}

	}
	close(s)

}
