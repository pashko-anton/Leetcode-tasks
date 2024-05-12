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

	hash1 := map[uint8]int{}
	hash2 := map[uint8]int{}

	for i := 0; i < len(s); i++ {
		hash1[s[i]] += 1
		hash2[t[i]] += 1
	}

	for k, sV := range hash1 {
		if tV, ok := hash2[k]; !ok || sV != tV {
			return false
		}

	}

	return true
}
