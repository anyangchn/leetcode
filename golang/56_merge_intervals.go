package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort.Sort(IntervalCollection(intervals))
	var result [][]int
	cur := intervals[0]
	for i := 1; i < len(intervals); i++ {
		tmp := intervals[i]
		if tmp[0] <= cur[1] {
			cur = []int{cur[0], max(tmp[1], cur[1])}
		} else {
			result = append(result, cur)
			cur = tmp
		}
	}
	result = append(result, cur)
	return result
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

type IntervalCollection [][]int

func (ic IntervalCollection) Len() int {
	return len(ic)
}

func (ic IntervalCollection) Less(i, j int) bool {
	if ic[i][0] <= ic[j][0] {
		return true
	}
	return false
}

func (ic IntervalCollection) Swap(i, j int) {
	ic[i], ic[j] = ic[j], ic[i]
}
