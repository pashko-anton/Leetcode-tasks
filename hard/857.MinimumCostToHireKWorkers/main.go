package main

import (
	"container/heap"
	"math"
	"sort"
)

/*
There are n workers. You are given two integer arrays quality and wage where quality[i] is the quality of the ith worker and wage[i] is the minimum wage expectation for the ith worker.

We want to hire exactly k workers to form a paid group. To hire a group of k workers, we must pay them according to the following rules:

Every worker in the paid group must be paid at least their minimum wage expectation.
In the group, each worker's pay must be directly proportional to their quality. This means if a worker’s quality is double that of another worker in the group, then they must be paid twice as much as the other worker.
Given the integer k, return the least amount of money needed to form a paid group satisfying the above conditions. Answers within 10-5 of the actual answer will be accepted.



Example 1:

Input: quality = [10,20,5], wage = [70,50,30], k = 2
Output: 105.00000
Explanation: We pay 70 to 0th worker and 35 to 2nd worker.
Example 2:

Input: quality = [3,1,10,10,1], wage = [4,8,2,2,7], k = 3
Output: 30.66667
Explanation: We pay 4 to 0th worker, 13.33333 to 2nd and 3rd workers separately.
*/

type worker struct {
	quality int
	wage    int
}

type priorityQueue []int

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i] > p[j]
}

func (p *priorityQueue) Pop() interface{} {
	v := (*p)[p.Len()-1]
	*p = (*p)[0 : p.Len()-1]
	return v
}

func (p *priorityQueue) Push(v interface{}) {
	*p = append(*p, v.(int))
}

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	workers := make([]worker, len(quality))
	for i := range workers {
		workers[i] = worker{quality[i], wage[i]}
	}

	sort.Slice(workers, func(i, j int) bool {
		return workers[i].wage*workers[j].quality < workers[j].wage*workers[i].quality
	})

	priorQueue, sum, minCost := make(priorityQueue, 0), 0, math.MaxFloat64
	for _, w := range workers {
		heap.Push(&priorQueue, w.quality)
		sum += w.quality
		if priorQueue.Len() > k {
			sum -= heap.Pop(&priorQueue).(int)
		}
		cost := float64(w.wage*sum) / float64(w.quality)
		if priorQueue.Len() == k && cost < minCost {
			minCost = cost
		}
	}

	return minCost
}
