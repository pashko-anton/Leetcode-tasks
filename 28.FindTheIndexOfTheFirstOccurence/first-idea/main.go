package main

/*
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.


Example 1:
Input: haystack = "sadbutsad", needle = "sad"
Output: 0
Explanation: "sad" occurs at index 0 and 6.
The first occurrence is at index 0, so we return 0.

Example 2:
Input: haystack = "leetcode", needle = "leeto"
Output: -1
Explanation: "leeto" did not occur in "leetcode", so we return -1.
*/

func main() {
}

func StrStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}

	index, finish := 0, false
	for i := 0; i < len(haystack); i++ {
		if finish {
			break
		}
		for j := 0; j < len(needle) && i+j < len(haystack); j++ {
			if haystack[i+j] != needle[j] {
				break
			}

			if j == (len(needle) - 1) {
				finish = true
				index = i
			}
		}
	}

	if !finish {
		return -1
	}

	return index
}
