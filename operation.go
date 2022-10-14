package main

type stack []int

func sumMatrix(matrix01 [][]int, matrix02 [][]int) [][]int {

	size_row := len(matrix01)
	size_column := len(matrix01[0])
	sumMatrix := createMatrix(size_row, size_column)

	for i := 0; i < size_row; i++ {
		for j := 0; j < size_column; j++ {
			sumMatrix[i][j] = matrix01[i][j] + matrix02[i][j]
		}
	}

	return sumMatrix
}

func subMatrix(matrix01 [][]int, matrix02 [][]int) [][]int {

	size_row := len(matrix01)
	size_column := len(matrix01[0])
	subMatrix := createMatrix(size_row, size_column)

	for i := 0; i < size_row; i++ {
		for j := 0; j < size_column; j++ {
			subMatrix[i][j] = matrix01[i][j] - matrix02[i][j]
		}
	}

	return subMatrix
}

func multMatrix(matrix01 [][]int, matrix02 [][]int) [][]int {

	size_row := len(matrix01)
	size_column := len(matrix01[0])
	multMatrix := createMatrix(size_row, size_column)

	for i := 0; i < size_row; i++ {
		for j := 0; j < size_column; j++ {
			for k := 0; k < size_column; k++ {
				multMatrix[i][j] += matrix01[i][k] * matrix02[k][j]
			}
		}
	}

	return multMatrix
}
