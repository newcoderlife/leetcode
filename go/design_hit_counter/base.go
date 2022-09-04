package design_hit_counter

import (
	"container/list"
)

type HitCounter struct {
	buffer *list.List
}

func Constructor() HitCounter {
	return HitCounter{buffer: list.New()}
}

type Item struct {
	Timestamp int
	Count     int
}

func (this *HitCounter) Hit(timestamp int) {
	if this.buffer.Len() == 0 {
		this.buffer.PushBack(&Item{Timestamp: timestamp, Count: 1})
		return
	}

	item := this.buffer.Back().Value.(*Item)
	if item.Timestamp == timestamp {
		item.Count++
		return
	}

	this.buffer.PushBack(&Item{Timestamp: timestamp, Count: 1})
	return
}

func (this *HitCounter) GetHits(timestamp int) int {
	if this.buffer.Len() == 0 {
		return 0
	}

	startTime := timestamp - 299
	start := binarySearch(this.buffer.Front(), nil, timestamp-299)
	end := binarySearch(this.buffer.Front(), nil, timestamp)

	result := 0
	for current := start; current != end.Next(); current = current.Next() {
		if item := current.Value.(*Item); startTime <= item.Timestamp && item.Timestamp <= timestamp {
			result += item.Count
		}
	}
	return result
}

func binarySearch(start, end *list.Element, timestamp int) *list.Element {
	if start == end || start.Next() == end {
		return start
	}

	mid := middle(start, end)
	if mid.Value.(*Item).Timestamp <= timestamp {
		return binarySearch(mid, end, timestamp)
	} else {
		return binarySearch(start, mid, timestamp)
	}
}

func middle(start, end *list.Element) *list.Element {
	fast, slow := start, start
	for fast != end {
		slow = slow.Next()
		if fast = fast.Next(); fast != end {
			fast = fast.Next()
		}
	}
	return slow
}
