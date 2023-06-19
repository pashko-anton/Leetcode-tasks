package main

/*
Given an array nums with n objects colored red, white, or blue, sort them in-place so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers 0, 1, and 2 to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.


Example 1:
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]

Example 2:
Input: nums = [2,0,1]
Output: [0,1,2]
*/
//
//func main() {
//}
//
func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(arr []int, s, e, k int) int {
	pivot := arr[e]
	left := s
	for i := s; i < e; i++ {
		if pivot > arr[i] {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[e], arr[left] = arr[left], arr[e]

	if left > k {
		return quickSelect(arr, s, left-1, k)
	} else if left < k {
		return quickSelect(arr, left+1, e, k)
	} else {
		return pivot
	}
}
