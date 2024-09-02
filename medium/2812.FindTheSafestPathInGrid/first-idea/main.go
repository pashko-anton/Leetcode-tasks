package first_idea

import "fmt"

/*
You are given a 0-indexed 2D matrix grid of size n x n, where (r, c) represents:

A cell containing a thief if grid[r][c] = 1
An empty cell if grid[r][c] = 0
You are initially positioned at cell (0, 0). In one move, you can move to any adjacent cell in the grid, including cells containing thieves.

The safeness factor of a path on the grid is defined as the minimum manhattan distance from any cell in the path to any thief in the grid.

Return the maximum safeness factor of all paths leading to cell (n - 1, n - 1).

An adjacent cell of cell (r, c), is one of the cells (r, c + 1), (r, c - 1), (r + 1, c) and (r - 1, c) if it exists.

The Manhattan distance between two cells (a, b) and (x, y) is equal to |a - x| + |b - y|, where |val| denotes the absolute value of val.

Example 1:
Input: grid = [[1,0,0],[0,0,0],[0,0,1]]
Output: 0
Explanation: All paths from (0, 0) to (n - 1, n - 1) go through the thieves in cells (0, 0) and (n - 1, n - 1).

Example 2:
Input: grid = [[0,0,1],[0,0,0],[0,0,0]]
Output: 2
Explanation: The path depicted in the picture above has a safeness factor of 2 since:
- The closest cell of the path to the thief at cell (0, 2) is cell (0, 0). The distance between them is | 0 - 0 | + | 0 - 2 | = 2.
It can be shown that there are no other paths with a higher safeness factor.

Example 3:
Input: grid = [[0,0,0,1],[0,0,0,0],[0,0,0,0],[1,0,0,0]]
Output: 2
Explanation: The path depicted in the picture above has a safeness factor of 2 since:
- The closest cell of the path to the thief at cell (0, 3) is cell (1, 2). The distance between them is | 0 - 1 | + | 3 - 2 | = 2.
- The closest cell of the path to the thief at cell (3, 0) is cell (3, 2). The distance between them is | 3 - 3 | + | 0 - 2 | = 2.
It can be shown that there are no other paths with a higher safeness factor.
*/

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	if n == 0 || (n > 0 && (grid[0][0] == 1 || grid[n-1][n-1] == 1)) {
		return 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {

		}
	}

	return maxVal
}

type coordinate struct {
	x, y int
}

type MemoryBlocks map[string]coordinate

func NewMemoryBlock() MemoryBlocks {
	return map[string]coordinate{}
}

func (m MemoryBlocks) Push(x, y int) {
	k := fmt.Sprintf("%d%d", x, y)
	m[k] = coordinate{
		x: x,
		y: y,
	}
}

var roads = map[string][]coordinate{}

func (m MemoryBlocks) Push(x, y int) {
	k := fmt.Sprintf("%d%d", x, y)
	m[k] = coordinate{
		x: x,
		y: y,
	}
}

func moving(blocks MemoryBlocks, x, y, n int, grid [][]int) {
	if x > n || y > n {
		return
	}
	if grid[x][y] == 1 {
		blocks.Push(x, y)
		return
	}

	directions := []coordinate{{x: x + 1, y: y}, {x: x - 1, y: y}, {x: x, y: y + 1}, {x: x, y: y - 1}}
	for _, d := range directions {
		res := moving(c, d.x, d.y, val+sum)
		if res > maxStartSum {
			maxStartSum = res
		}
	}
}
