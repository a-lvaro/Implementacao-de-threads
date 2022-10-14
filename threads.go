package main

func threadsMatrix(matrixA [][]int, matrixB [][]int, s chan<- [][]int, flag string) {

	size := len(matrixA)

	for i := 0; i < size; i++ {

		if flag == "SUM" {
			s <- sumMatrix(matrixA[i:i+1], matrixB[i:i+1])
		} else if flag == "SUB" {
			s <- subMatrix(matrixA[i:i+1], matrixB[i:i+1])
		} else if flag == "MULT" {
			s <- multMatrix(matrixA[i:i+1], matrixB)
		} else if flag == "TRAN" {
			s <- transpose(matrixA[i : i+1])
		}

	}
	close(s)
}
