// Package interval_problem
// @file: interval.go
// @date: 2021/1/18
package interval_problem

import (
	"fmt"
	"sort"
)

type intervalList [][]int

func (i intervalList) Len() int {
	return len(i)
}

func (i intervalList) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i intervalList) Less(a, b int) bool {
	if i[a][0] == i[b][0] {
		return i[a][1] > i[b][1]
	}
	return i[a][0] < i[b][0]
}

func (i intervalList) String() string {
	res := ""
	for _, v := range i {
		res += fmt.Sprintf("(%d, %d) ", v[0], v[1])
	}
	return res
}

// 1288. 删除被覆盖区间
func removeCoveredIntervals(intervals [][]int) int {
	sort.Sort(intervalList(intervals))

	left, right := intervals[0][0], intervals[0][1]
	res := 0
	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		// 情况1：覆盖区间
		if left <= interval[0] && right >= interval[1] {
			res++
		}
		// 情况2：区间相交
		if left < interval[0] && right < interval[1] {
			right = interval[1]
		}
		// 情况3：区间不相交
		if left < interval[0] && right < interval[0] {
			left = interval[0]
			right = interval[1]
		}
	}
	return len(intervals) - res
}

// 56. 合并区间
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}

	sort.Sort(intervalList(intervals))
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		last := res[len(res)-1]
		current := intervals[i]
		if last[1] >= current[0] {
			if current[1] > last[1] {
				last[1] = current[1]
			}
		} else {
			res = append(res, current)
		}
	}
	return res
}

// 986. 区间列表的交集
func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	intervals := append(firstList, secondList...)
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Sort(intervalList(intervals))
	res := [][]int{}
	right := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		interval := intervals[i]
		if right > interval[1] {
			res = append(res, []int{interval[0], interval[1]})
		} else if right >= interval[0] && right <= interval[1] {
			res = append(res, []int{interval[0], right})
			right = interval[1]
		} else if right < interval[0] {
			right = interval[1]
		}
	}
	return res
}

func intervalIntersection2(firstList [][]int, secondList [][]int) [][]int {
	i, j := 0, 0
	res := [][]int{}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	for i < len(firstList) && j < len(secondList) {
		a1, a2 := firstList[i][0], firstList[i][1]
		b1, b2 := secondList[j][0], secondList[j][1]
		// 存在交集
		if b2 >= a1 && a2 >= b1 {
			res = append(res, []int{max(a1, b1), min(a2, b2)})
		}
		if a2 > b2 {
			j++
		} else {
			i++
		}
	}
	return res
}
