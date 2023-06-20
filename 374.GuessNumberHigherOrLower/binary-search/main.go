package main

import "fmt"

/*
We are playing the Guess Game. The game is as follows:

I pick a number from 1 to n. You have to guess which number I picked.

Every time you guess wrong, I will tell you whether the number I picked is higher or lower than your guess.

You call a pre-defined API int guess(int num), which returns three possible results:

-1: Your guess is higher than the number I picked (i.e. num > pick).
1: Your guess is lower than the number I picked (i.e. num < pick).
0: your guess is equal to the number I picked (i.e. num == pick).
Return the number that I picked.

Example 1:
Input: n = 10, pick = 6
Output: 6

Example 2:
Input: n = 1, pick = 1
Output: 1

Example 3:
Input: n = 2, pick = 1
Output: 1
*/

func main() {
	fmt.Println(guessNumber(10))
}

var pick = 6

func guess(n int) int {
	if n > pick {
		return -1
	} else if n < pick {
		return 1
	}

	return 0
}

func guessNumber(n int) int {
	start, finish := 1, n

	for start <= finish {
		mid := (start + finish) / 2

		if guess(mid) == -1 {
			finish = mid - 1
		} else if guess(mid) == 1 {
			start = mid + 1
		} else {
			return mid
		}
	}

	return -1
}

//
//func search(nums []int, target int) int {
//	start, end := 0, len(nums)-1
//
//	for start <= end {
//		mid := (start + end) / 2
//
//		if target > nums[mid] {
//			start = mid + 1
//		} else if target < nums[mid] {
//			end = mid - 1
//		} else {
//			return mid
//		}
//	}
//
//	return -1
//}
