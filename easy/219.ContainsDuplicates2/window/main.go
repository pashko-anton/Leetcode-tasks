package main

/*
Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such that nums[i] == nums[j] and abs(i - j) <= k.

Example 1:
Input: nums = [1,2,3,1], k = 3
Output: true

Example 2:
Input: nums = [1,0,1,1], k = 1
Output: true

Example 3:
Input: nums = [1,2,3,1,2,3], k = 2
Output: false
*/

func main() {
}

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums)-1 < k {
		k = len(nums) - 1
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= k+i && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}
	return false
}
