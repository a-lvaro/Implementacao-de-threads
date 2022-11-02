package main

func threadsMatrix(matrixChannel <-chan []int, result chan<- []int, size int, flag string) {

	if flag == "SUM" {
		for n := range matrixChannel {
			result <- sumThread(n, <-matrixChannel)
		}
	} else if flag == "SUB" {
		for n := range matrixChannel {
			result <- subThread(n, <-matrixChannel)
		}
	}

	for len(result) != size*size {
	}
	close(result)
}

func threadsMatrixMult(matrixChannel <-chan []int, result chan<- []int, matrixB [][]int, size int) {

	for n := range matrixChannel {
		for j := 0; j < size; j++ {
			result <- multThread(n, matrixB[j])
		}
	}

	for len(result) != size*size {
	}
	close(result)
}

func sumThread(matrix01 []int, matrix02 []int) []int {
	result := make([]int, len(matrix01))
	for i := 0; i < len(matrix01); i++ {
		result[i] = matrix01[i] + matrix02[i]
	}
	return result
}

func subThread(matrix01 []int, matrix02 []int) []int {
	result := make([]int, len(matrix01))
	for i := 0; i < len(matrix01); i++ {
		result[i] = matrix01[i] - matrix02[i]
	}
	return result
}

func multThread(matrix01 []int, matrix02 []int) []int {
	result := make([]int, 1)
	size := 4

	for i := 0; i < size; i++ {
		result[0] += matrix01[i] * matrix02[i]
	}
	return result
}

func closeChannel(matrixChannel chan []int) {
	aux := true
	for aux {
		if len(matrixChannel) == 0 {
			aux = false
			close(matrixChannel)
		}
	}
}
