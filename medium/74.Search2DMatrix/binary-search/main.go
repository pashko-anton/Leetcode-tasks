package main

/*
You are given an m x n integer matrix matrix with the following two properties:

Each row is sorted in non-decreasing order.
The first integer of each row is greater than the last integer of the previous row.
Given an integer target, return true if target is in matrix or false otherwise.

You must write a solution in O(log(m * n)) time complexity.

Example 1:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3
Output: true

Example 2:
Input: matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 13
Output: false
*/

func searchMatrix(matrix [][]int, target int) bool {
	var nums []int
	for i := range matrix {
		if matrix[i][0] <= target && matrix[i][len(matrix[i])-1] >= target {
			nums = matrix[i]
		}
	}

	start, end := 0, len(nums)-1
	for start <= end {
		mid := (start + end) / 2

		if target > nums[mid] {
			start = mid + 1
		} else if target < nums[mid] {
			end = mid - 1
		} else {
			return true
		}
	}

	return false
}
