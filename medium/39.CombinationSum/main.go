package main

import (
	"fmt"
	"sort"
)

/*
Given an array of distinct integers candidates and a target integer target, return a list of all unique combinations of candidates where the chosen numbers sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times. Two combinations are unique if the
frequency
 of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target is less than 150 combinations for the given input.

Example 1:
Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.

Example 2:
Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]

Example 3:
Input: candidates = [2], target = 1
Output: []
*/

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	var results [][]int

	sort.Ints(candidates)
	//subset := make([]int, 0)
	var combinator func([]int, int, int)
	combinator = func(currentSubset []int, idx, targetComb int) {
		if targetComb == 0 {
			comb := make([]int, len(currentSubset))
			copy(comb, currentSubset)
			results = append(results, comb)
			return
		} else if targetComb < 0 {
			return
		} else {
			for i := idx; i < len(candidates); i++ {
				combinator(append(currentSubset, candidates[i]), i, targetComb-candidates[i])
			}
		}
	}
	combinator([]int{}, 0, target)

	return results
}
