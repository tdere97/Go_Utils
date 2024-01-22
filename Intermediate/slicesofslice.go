package main

import "fmt"

func createMatrix(rows, cols int) [][]int {
	matrix := make([][]int, 0)

	for i := 0; i < rows; i++ {
		row := make([]int, 0)
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

func sSlice() {
	fmt.Println("===============================sliceofslices.go=================================")

	// SLICE OF SLICES
	// Slices can hold other slices, effectively creating a matrix, or a 2D slice.
	// rows := [][]int{}

	matrix := [][]int{{0, 0, 0, 0, 0}, {1, 1, 1, 1, 1}, {2, 2, 2, 2, 2}, {3, 3, 3, 3, 3}, {4, 4, 4, 4, 4}}
	row := 5
	col := 5

	fmt.Printf("Matrix is :\n")
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Printf("%v\t", matrix[i][j])
		}
		fmt.Printf("\n")
	}
	mat := createMatrix(3, 3)

	fmt.Printf("Matrix is :\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%v\t", mat[i][j])
		}
		fmt.Printf("\n")
	}
	rangeFunc()
}
