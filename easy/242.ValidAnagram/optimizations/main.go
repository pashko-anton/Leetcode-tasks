package main

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.

Example 1:
Input: s = "anagram", t = "nagaram"
Output: true

Example 2:
Input: s = "rat", t = "car"
Output: false
*/

func main() {
}

func IsAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	alphabet := make([]int, 26)
	for _, sV := range s {
		alphabet[sV-'a']++
	}

	for _, tV := range t {
		if alphabet[tV-'a'] == 0 {
			return false
		} else {
			alphabet[tV-'a']--
		}
	}

	return true
}
