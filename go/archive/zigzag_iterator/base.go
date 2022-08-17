package zigzag_iterator

type ZigzagIterator struct {
	Buffer [][]int
	Index  int
}

func Constructor(v1, v2 []int) *ZigzagIterator {
	return &ZigzagIterator{
		Buffer: [][]int{v1, v2},
		Index:  0,
	}
}

func (this *ZigzagIterator) next() int {
	for {
		if len(this.Buffer[this.Index]) > 0 {
			ans := this.Buffer[this.Index][0]
			this.Buffer[this.Index] = this.Buffer[this.Index][1:]
			this.Index = (this.Index + 1) % len(this.Buffer)
			return ans
		}
		this.Index = (this.Index + 1) % len(this.Buffer)
	}
}

func (this *ZigzagIterator) hasNext() bool {
	for _, v := range this.Buffer {
		if len(v) > 0 {
			return true
		}
	}
	return false
}
