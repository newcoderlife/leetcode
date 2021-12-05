package min_meeting_rooms

import "sort"

type BinaryTreeArray struct {
	buffer []int
	len    int
}

func CreateBinaryTreeArray(n int) *BinaryTreeArray {
	return &BinaryTreeArray{
		buffer: make([]int, n+1),
		len:    n,
	}
}

func lowbit(x int) int {
	return x & -x
}

func (b *BinaryTreeArray) Add(index, value int) {
	index += 1
	for index < b.len {
		b.buffer[index] += value
		index += lowbit(index)
	}
}

func (b *BinaryTreeArray) Get(index int) int {
	index += 1
	result := 0
	for index > 0 {
		result += b.buffer[index]
		index -= lowbit(index)
	}
	return result
}

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	total := intervals[len(intervals)-1][1] + 1
	tree := CreateBinaryTreeArray(total)
	for _, interval := range intervals {
		tree.Add(interval[0], 1)
		tree.Add(interval[1], -1)
	}

	result := 0
	for i := 0; i < total; i++ {
		if val := tree.Get(i); result < val {
			result = val
		}
	}
	return result
}
