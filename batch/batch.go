package batch

import "math"

type Batch struct {
	size         int
	count        int
	roundSize    int
	currentTimes int
	times        int
}

func New(count, size int) *Batch {
	var b Batch
	if count < 0 {
		count = 0
	}
	if size < 1 {
		size = 1
	}
	b.count = count
	b.size = size
	b.times = int(math.Ceil(float64(count) / float64(size)))
	return &b
}

func (b *Batch) Iter(start, length *int) bool {
	if start == nil || length == nil {
		return false
	}
	if b.currentTimes != 0 {
		*start += *length
	}
	roundSize := b.size
	if *start+roundSize > b.count {
		roundSize = b.count - *start
	}
	*length = roundSize
	b.currentTimes++
	return *start < b.count
}

func (b *Batch) Times() int {
	return b.times
}
