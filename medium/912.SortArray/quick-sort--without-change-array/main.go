package main

/*
Given an array of integers nums, sort the array in ascending order and return it.

You must solve the problem without using any built-in functions in O(nlog(n)) time complexity and with the smallest space complexity possible.


Example 1:

Input: nums = [5,2,3,1]
Output: [1,2,3,5]
Explanation: After sorting the array, the positions of some numbers are not changed (for example, 2 and 3), while the positions of other numbers are changed (for example, 1 and 5).
Example 2:

Input: nums = [5,1,1,2,0,0]
Output: [0,0,1,1,2,5]
Explanation: Note that the values of nums are not necessairly unique.
*/

func main() {
}

func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)

	return nums
}

func quickSort(nums []int, s, e int) {
	if e-s <= 0 {
		return
	}

	pivot := nums[e]
	left := s
	for i := s; i < e; i++ {
		if pivot > nums[i] {
			tmp := nums[left]
			nums[left] = nums[i]
			nums[i] = tmp
			left++
		}
	}

	nums[e] = nums[left]
	nums[left] = pivot

	quickSort(nums, s, left-1)
	quickSort(nums, left+1, e)

	return
}
