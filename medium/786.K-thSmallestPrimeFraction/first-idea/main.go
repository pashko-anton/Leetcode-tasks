package main

import "sort"

/*
You are given a sorted integer array arr containing 1 and prime numbers, where all the integers of arr are unique. You are also given an integer k.

For every i and j where 0 <= i < j < arr.length, we consider the fraction arr[i] / arr[j].

Return the kth smallest fraction considered. Return your answer as an array of integers of size 2, where answer[0] == arr[i] and answer[1] == arr[j].
*/

func main() {
}

type fraction struct {
	x   int
	y   int
	res float64
}

func kthSmallestPrimeFraction(arr []int, k int) []int {
	fractions := make([]fraction, 0)
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			fractions = append(fractions, fraction{
				x:   arr[i],
				y:   arr[j],
				res: float64(arr[i]) / float64(arr[j]),
			})
		}
	}

	sort.Slice(fractions, func(i, j int) bool {
		return fractions[i].res < fractions[j].res
	})

	return []int{fractions[k-1].x, fractions[k-1].y}
}
