package main

import (
	"math"
)

/*
You are given an m x n binary matrix grid.

A move consists of choosing any row or column and toggling each value in that row or column (i.e., changing all 0's to 1's, and all 1's to 0's).

Every row of the matrix is interpreted as a binary number, and the score of the matrix is the sum of these numbers.

Return the highest possible score after making any number of moves (including zero moves).

Example 1:
Input: grid = [[0,0,1,1],[1,0,1,0],[1,1,0,0]]
Output: 39
Explanation: 0b1111 + 0b1001 + 0b1111 = 15 + 9 + 15 = 39

Example 2:
Input: grid = [[0]]
Output: 1
*/

func matrixScore(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// reverse first column without one value
	zerosCnt := make([]int, n)
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			grid[i][0] = 1
			for j := 1; j < n; j++ {
				grid[i][j] ^= 1
				if grid[i][j] == 0 {
					zerosCnt[j]++
				}
			}
			continue
		}

		for j := 1; j < n; j++ {
			if grid[i][j] == 0 {
				zerosCnt[j]++
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			if zerosCnt[j] > m/2 {
				grid[i][j] ^= 1
			}
		}
	}

	score := 0
	for _, rows := range grid {
		score += calcBitToDigit(rows)
	}

	return score
}

func calcBitToDigit(values []int) (score int) {
	for i, pow := len(values)-1, 0.0; i >= 0; i, pow = i-1, pow+1 {
		score += int(math.Pow(2, pow)) * values[i]
	}

	return
}
