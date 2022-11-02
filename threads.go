package main

func threadsMatrix(matrixA [][]float64, matrixB [][]float64, s chan<- [][]float64, flag string) {

	size := len(matrixA)

	for i := 0; i < size; i++ {

		if flag == "SUM" {
			s <- sumMatrix(matrixA[i:i+1], matrixB[i:i+1])
		} else if flag == "SUB" {
			s <- subMatrix(matrixA[i:i+1], matrixB[i:i+1])
		} else if flag == "MULT" {
			s <- multMatrix(matrixA[i:i+1], matrixB)
		}

	}
	close(s)
}
