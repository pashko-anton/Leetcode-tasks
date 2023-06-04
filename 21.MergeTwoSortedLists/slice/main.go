package main

/*
You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]

Example 2:
Input: list1 = [], list2 = []
Output: []

Example 3:
Input: list1 = [], list2 = [0]
Output: [0]

*/

func main() {
	MergeTwoLists([]int{1, 2, 4}, []int{1, 3, 4})
}

func MergeTwoLists(input1, input2 []int) []int {
	if len(input1) == 0 {
		return input2
	}

	if len(input2) == 0 {
		return input1
	}

	mergeLen := len(input1) + len(input2)
	mergeArray := make([]int, 0, mergeLen)

	iLimit := len(input1)
	jLimit := len(input2)

	iFinish, jFinish := false, false
	for i, j := 0, 0; !iFinish || !jFinish; {
		if !iFinish && input1[i] <= input2[j] {
			mergeArray = append(mergeArray, input1[i])
			if i != (iLimit - 1) {
				i++
			} else {
				iFinish = true
			}
			continue
		}

		if !jFinish && input1[i] >= input2[j] {
			mergeArray = append(mergeArray, input2[j])
			if j != (jLimit - 1) {
				j++
			} else {
				jFinish = true
			}
			continue
		}
	}
	return mergeArray
}
