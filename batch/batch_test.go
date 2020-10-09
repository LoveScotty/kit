package batch

import "testing"

func TestBatch_Iter(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := New(len(a), 4)
	var start, length int
	t.Log("总共循环次数为: ", b.Times())
	for b.Iter(&start, &length) {
		t.Logf("%v-->%v\n", b.currentTimes, a[start:start+length])
	}
}
