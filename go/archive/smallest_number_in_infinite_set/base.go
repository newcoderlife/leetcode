package smallest_number_in_infinite_set

type SmallestInfiniteSet struct {
	buffer  []int
	len     int
	deleted map[int]int
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{
		buffer:  make([]int, 0),
		deleted: make(map[int]int),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	result := 1
	if this.len > 0 {
		for result = this.buffer[0] - 1; result > 0; result-- {
			if _, ok := this.deleted[result]; !ok {
				break
			}
		}
		if result == 0 {
			for result = this.buffer[0] + 1; ; result++ {
				if _, ok := this.deleted[result]; !ok {
					break
				}
			}
		}
	}

	// insert result into delete heap
	if len(this.buffer) == this.len {
		this.buffer = append(this.buffer, result)
	} else {
		this.buffer[this.len] = result
	}
	this.deleted[result] = this.len
	this.len++

	for current := this.len - 1; current > 0; current /= 2 {
		if this.buffer[current] > this.buffer[current/2] {
			break
		}

		this.buffer[current], this.buffer[current/2] = this.buffer[current/2], this.buffer[current]
		this.deleted[this.buffer[current]] = current
		this.deleted[this.buffer[current/2]] = current / 2
	}

	return result
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if _, ok := this.deleted[num]; !ok {
		return
	}

	index := this.deleted[num]
	delete(this.deleted, num)

	// delete item from heap
	this.buffer[this.len-1], this.buffer[index] = this.buffer[index], this.buffer[this.len-1]
	this.len--
	for index < this.len {
		if index*2 < this.len && this.buffer[index] > this.buffer[index*2] {
			this.buffer[index], this.buffer[index*2] = this.buffer[index*2], this.buffer[index]
			this.deleted[this.buffer[index]] = index
			this.deleted[this.buffer[index*2]] = index * 2
			index = index * 2
		} else if index*2+1 < this.len && this.buffer[index] > this.buffer[index*2+1] {
			this.buffer[index], this.buffer[index*2+1] = this.buffer[index*2+1], this.buffer[index]
			this.deleted[this.buffer[index]] = index
			this.deleted[this.buffer[index*2+1]] = index*2 + 1
			index = index*2 + 1
		} else {
			break
		}
	}
}
