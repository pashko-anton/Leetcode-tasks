package main

import "fmt"

/*
Given an integer array nums of unique elements, return all possible
subsets(the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.

Example 1:
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

Example 2:
Input: nums = [0]
Output: [[],[0]]

*/

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	var results [][]int

	subset := make([]int, 0)
	var subcreator func(i int)
	subcreator = func(i int) {
		if i >= len(nums) {
			cp := make([]int, len(subset))
			copy(cp, subset)
			results = append(results, cp)
			return
		}

		subset = append(subset, nums[i])
		subcreator(i + 1)

		subset = subset[:len(subset)-1]
		subcreator(i + 1)
	}
	subcreator(0)

	return results
}
