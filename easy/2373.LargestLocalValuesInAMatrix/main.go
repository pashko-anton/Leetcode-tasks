package main

import (
	"slices"
)

/*
You are given an n x n integer matrix grid.

Generate an integer matrix maxLocal of size (n - 2) x (n - 2) such that:

maxLocal[i][j] is equal to the largest value of the 3 x 3 matrix in grid centered around row i + 1 and column j + 1.
In other words, we want to find the largest value in every contiguous 3 x 3 matrix in grid.

Return the generated matrix.

	Input: grid = [[9,9,8,1],[5,6,2,6],[8,2,6,4],[6,2,2,2]]
Output: [[9,9],[8,6]]
Explanation: The diagram above shows the original matrix and the generated matrix.
Notice that each value in the generated matrix corresponds to the largest value of a contiguous 3 x 3 matrix in grid.
*/

func largestLocal(grid [][]int) [][]int {
	matrixSize := 3
	rows := len(grid) - matrixSize + 1
	cols := len(grid[0]) - matrixSize + 1
	res := make([][]int, len(grid)-matrixSize+1)

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			maxElem := 0
			for _, subMtrx := range grid[i : i+matrixSize] {
				curMax := slices.Max(subMtrx[j : j+matrixSize])
				if curMax > maxElem {
					maxElem = curMax
				}
			}

			res[i] = append(res[i], maxElem)
		}
	}

	return res
}
