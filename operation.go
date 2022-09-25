package main

func sumMatrix(matrix01 [][]int, matrix02 [][]int) [][]int {

	size := len(matrix01)
	sumMatrix := createMatrix(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			sumMatrix[i][j] = matrix01[i][j] + matrix02[i][j]
		}
	}

	return sumMatrix
}

func subMatrix(matrix01 [][]int, matrix02 [][]int) [][]int {

	size := len(matrix01)
	subMatrix := createMatrix(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			subMatrix[i][j] = matrix01[i][j] - matrix02[i][j]
		}
	}

	return subMatrix
}

func multMatrix(matrix01 [][]int, matrix02 [][]int) [][]int {

	size := len(matrix01)
	multMatrix := createMatrix(size)

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			for k := 0; k < size; k++ {
				multMatrix[i][j] += matrix01[i][k] * matrix02[k][j]
			}
		}
	}

	return multMatrix
}
