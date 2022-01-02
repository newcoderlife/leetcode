package find_kth_largest

import "container/list"

func insert(rank *list.List, num int) {
	for start := rank.Front(); start != nil; start = start.Next() {
		if start.Value.(int) < num {
			rank.InsertBefore(num, start)
			return
		}
	}
	rank.PushBack(num)
}

func findKthLargest(nums []int, k int) int {
	rank := list.New()
	rank.PushBack(nums[0])

	for i := 1; i < len(nums); i++ {
		insert(rank, nums[i])
		if rank.Len() > k {
			rank.Remove(rank.Back())
		}
	}
	return rank.Back().Value.(int)
}
