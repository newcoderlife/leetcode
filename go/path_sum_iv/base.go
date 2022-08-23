package path_sum_iv

func pathSum(nums []int) int {
	exist := make(map[int]int)
	isParent := make(map[int]bool)
	for _, num := range nums {
		depth, rank, value := num/100, (num%100)/10, num%10
		exist[depth*10+rank] = value
		isParent[(depth-1)*10+(rank+1)/2] = true
	}

	result := 0
	for key := range exist {
		if isParent[key] {
			continue
		}

		depth, rank := key/10, key%10
		for depth > 0 {
			result += exist[depth*10+rank]
			depth, rank = depth-1, (rank+1)/2
		}
	}

	return result
}
