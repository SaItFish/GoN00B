// Package others
// @file: others_test.go
// @date: 2021/3/9
package others

import (
	"fmt"
	"testing"
	"time"
)

type eggTest struct {
	in  [2]int
	out int
}

var eggTests = []eggTest{
	{in: [2]int{100, 100000}, out: 17},
	{in: [2]int{2, 100}, out: 14},
	{in: [2]int{3, 14}, out: 4},
}

func TestSuperEggDrop(t *testing.T) {
	for _, et := range eggTests {
		start := time.Now()
		result := superEggDrop(et.in[0], et.in[1])
		end := time.Now()
		delta := end.Sub(start)
		if result != et.out {
			t.Fatalf("get:%v, want false", result)
		}
		fmt.Printf("Spend time: %s\tResult is: %d\n", delta, result)
	}
}
