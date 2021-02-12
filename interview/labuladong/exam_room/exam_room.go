// Package exam_room
// @file: exam_room.go
// @date: 2021/2/12
package exam_room

import (
	"github.com/emirpasic/gods/sets/treeset"
)

// 855. 考场就座
type ExamRoom struct {
	Students *treeset.Set
	N        int
}

func Constructor(N int) ExamRoom {
	return ExamRoom{Students: treeset.NewWithIntComparator(), N: N}
}

func (r *ExamRoom) Seat() int {
	var student int // 学生应该坐在哪个位置
	size := r.Students.Size()
	values := r.Students.Values()
	if size > 0 {
		// dist is the distance to the closest student
		dist := 0
		prev := values[0].(int)
		for i := 1; i < size; i++ {
			s := values[i].(int)
			d := (s - prev) / 2
			if d > dist {
				dist = d
				student = prev + d
			}
			prev = s
		}
		if r.N-1-values[size-1].(int) > dist {
			student = r.N - 1
		}
	}
	r.Students.Add(student)
	return student
}

func (r *ExamRoom) Leave(p int) {
	r.Students.Remove(p)
}
