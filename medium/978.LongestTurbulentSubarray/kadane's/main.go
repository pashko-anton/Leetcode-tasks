package kadane_s

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

	left, right := 0, 0
	startLeft := false
	max := math.MinInt64
	for i := 0; i < len(nums)-1; i++ {
		if left == 0 && right == 0 {
			if nums[i] > nums[i+1] {
				startLeft = true
				left++
			} else if nums[i] < nums[i+1] {
				right++
			} else {
				if max < 1 {
					max = 1
				}
			}
			continue
		}

		if (left-1 == right || (left == right && !startLeft)) && nums[i] < nums[i+1] {
			right++
			continue
		}

		if (right-1 == left || (left == right && startLeft)) && nums[i] > nums[i+1] {
			left++
			continue
		}

		res := left + right + 1
		if max < res {
			max = res
		}
		i--
		left, right = 0, 0
		startLeft = false
	}

	res := left + right + 1
	if max < res {
		return res
	}

	return max
}
