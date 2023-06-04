package main

import (
	"sort"
)

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

Notice that the solution set must not contain duplicate triplets.



Example 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Example 2:
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

Example 3:
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/

func main() {
}

func ThreeSum(nums []int) [][]int {
	var resp [][]int

	sort.Ints(nums)

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		ptr := len(nums) - 1
		k := i + 1

		for k < ptr {
			sum := nums[i] + nums[ptr] + nums[k]

			if sum == 0 {
				resp = append(resp, []int{nums[i], nums[ptr], nums[k]})
				for k < ptr && nums[k] == nums[k+1] {
					k++
				}
				for ptr > 0 && nums[ptr] == nums[ptr-1] {
					ptr--
				}

				k++
				ptr--
				continue
			}

			if sum < 0 {
				k++
				continue
			}

			if sum > 0 {
				ptr--
				continue
			}
		}
	}

	return resp
}
