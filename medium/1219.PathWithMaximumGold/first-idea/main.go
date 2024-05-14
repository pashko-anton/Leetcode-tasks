package first_idea

/*
In a gold mine grid of size m x n, each cell in this mine has an integer representing the amount of gold in that cell, 0 if it is empty.

Return the maximum amount of gold you can collect under the conditions:

Every time you are located in a cell you will collect all the gold in that cell.
From your position, you can walk one step to the left, right, up, or down.
You can't visit the same cell more than once.
Never visit a cell with 0 gold.
You can start and stop collecting gold from any position in the grid that has some gold.

Example 1:
Input: grid = [[0,6,0],[5,8,7],[0,9,0]]
Output: 24
Explanation:
[[0,6,0],
 [5,8,7],
 [0,9,0]]
Path to get the maximum gold, 9 -> 8 -> 7.

Example 2:
Input: grid = [[1,0,7],[2,0,6],[3,4,5],[0,3,0],[9,0,20]]
Output: 28
Explanation:
[[1,0,7],
 [2,0,6],
 [3,4,5],
 [0,3,0],
 [9,0,20]]
Path to get the maximum gold, 1 -> 2 -> 3 -> 4 -> 5 -> 6 -> 7.
*/

type Cursor struct {
	limitX, limitY int
	Grid           [][]int
	Sum            int
}

func (c *Cursor) Clear(x, y int) {
	c.Grid[x][y] = 0
	return
}

func (c *Cursor) Set(x, y, val int) {
	c.Grid[x][y] = val
	return
}

func (c *Cursor) Collect(sum int) int {
	c.Sum = c.Sum + sum
	return c.Sum
}

func (c *Cursor) SumVal() int {
	return c.Sum
}

func (c *Cursor) Val(x, y int) int {
	if (x < 0 || x >= c.limitX) || (y < 0 || y >= c.limitY) {
		return 0
	}

	return c.Grid[x][y]
}

func getMaximumGold(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	maxVal := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				continue
			}
			v := collect(&Cursor{
				limitX: m,
				limitY: n,
				Grid:   grid,
			}, i, j, 0)
			if v > maxVal {
				maxVal = v
			}
		}
	}
	return maxVal
}

func collect(c *Cursor, x, y int, sum int) int {
	if c.Val(x, y) == 0 {
		return 0
	}

	val := c.Val(x, y)
	c.Clear(x, y)

	type coordinate struct {
		x, y int
	}

	maxStartSum := 0
	directions := []coordinate{{x: x + 1, y: y}, {x: x - 1, y: y}, {x: x, y: y + 1}, {x: x, y: y - 1}}
	for _, d := range directions {
		res := collect(c, d.x, d.y, val+sum)
		if res > maxStartSum {
			maxStartSum = res
		}
	}

	c.Set(x, y, val)

	return val + maxStartSum
}
