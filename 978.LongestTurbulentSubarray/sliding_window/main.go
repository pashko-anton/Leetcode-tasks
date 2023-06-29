package longestturbulentsliding

import "math"

/*
Given an integer array arr, return the length of a maximum size turbulent subarray of arr.

A subarray is turbulent if the comparison sign flips between each adjacent pair of elements in the subarray.

More formally, a subarray [arr[i], arr[i + 1], ..., arr[j]] of arr is said to be turbulent if and only if:

For i <= k < j:
arr[k] > arr[k + 1] when k is odd, and
arr[k] < arr[k + 1] when k is even.
Or, for i <= k < j:
arr[k] > arr[k + 1] when k is even, and
arr[k] < arr[k + 1] when k is odd.

Example 1:
Input: arr = [9,4,2,10,7,8,8,1,9]
Output: 5
Explanation: arr[1] > arr[2] < arr[3] > arr[4] < arr[5]

Example 2:
Input: arr = [4,8,12,16]
Output: 2

Example 3:
Input: arr = [100]
Output: 1
*/

func maxTurbulenceSize(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	ptr := 0
	var isContinue bool
	max := math.MinInt64
	for i := 0; i < len(nums)-1; i++ {
		ptr = i + 1
		var next bool
		for ptr < len(nums) {
			if ptr-1 == i && nums[ptr-1] == nums[ptr] {
				isContinue = true
				break
			}

			if ptr-1 == i ||
				(nums[ptr-1] < nums[ptr] && next) ||
				(nums[ptr-1] > nums[ptr] && !next) {
				next = nums[ptr-1] > nums[ptr]
				ptr++
				continue
			}
			break
		}
		res := ptr - i
		if max < res {
			max = res
		}
		if isContinue {
			isContinue = false
			continue
		}
		i = ptr - 2
	}

	return max
}
