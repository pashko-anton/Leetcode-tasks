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
	return mergeSort(nums)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	m := len(arr) / 2
	left := mergeSort(arr[:m])
	right := mergeSort(arr[m:])

	newArr := merge(left, right)
	return newArr
}

func merge(left, right []int) []int {
	newArr := make([]int, 0)
	var i, j = 0, 0
	for i < len(left) && j < len(right) {
		if left[i] >= right[j] {
			newArr = append(newArr, right[j])
			j++
		} else if left[i] < right[j] {
			newArr = append(newArr, left[i])
			i++
		}
	}

	for i < len(left) {
		newArr = append(newArr, left[i])
		i++
	}

	for j < len(right) {
		newArr = append(newArr, right[j])
		j++
	}

	return newArr
}
