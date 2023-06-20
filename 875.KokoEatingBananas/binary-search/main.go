package main

import "math"

/*
Koko loves to eat bananas. There are n piles of bananas, the ith pile has piles[i] bananas. The guards have gone and will come back in h hours.

Koko can decide her bananas-per-hour eating speed of k. Each hour, she chooses some pile of bananas and eats k bananas from that pile. If the pile has less than k bananas, she eats all of them instead and will not eat any more bananas during this hour.

Koko likes to eat slowly but still wants to finish eating all the bananas before the guards return.

Return the minimum integer k such that she can eat all the bananas within h hours.

Example 1:
Input: piles = [3,6,7,11], h = 8
Output: 4

Example 2:
Input: piles = [30,11,23,4,20], h = 5
Output: 30

Example 3:
Input: piles = [30,11,23,4,20], h = 6
Output: 23
*/

func minEatingSpeed(piles []int, h int) int {
	left, right := 1, max(piles)
	k := 0
	for left <= right {
		mid := (left + right) / 2

		hoursCalc := sumHours(piles, mid)
		if hoursCalc > h {
			left = mid + 1
		} else {
			right = mid - 1
			k = mid
		}
	}

	return k
}

func sumHours(piles []int, k int) int {
	hours := 0
	for _, p := range piles {
		hours += p / k
		if p%k != 0 {
			hours++
		}
	}

	return hours
}

func max(piles []int) int {
	max := math.MinInt
	for _, p := range piles {
		if p > max {
			max = p
		}
	}

	return max
}
