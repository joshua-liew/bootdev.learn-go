package main

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for i := range rows {
		row := []int{}
		for j := range cols {
			row = append(row, i * j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}
