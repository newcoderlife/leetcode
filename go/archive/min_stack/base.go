package min_stack

import (
	"container/list"
	"math"
)

type MinStack struct {
	buffer     *list.List
	currentMin int
}

func Constructor() MinStack {
	return MinStack{buffer: list.New(), currentMin: math.MaxInt64}
}

func (this *MinStack) Push(val int) {
	this.buffer.PushBack(val - this.currentMin)
	if val < this.currentMin {
		this.currentMin = val
	}
}

func (this *MinStack) Pop() {
	if this.buffer.Len() == 0 {
		return
	}

	delta := this.buffer.Back().Value.(int)
	if delta < 0 {
		this.currentMin -= delta
	}

	this.buffer.Remove(this.buffer.Back())
}

func (this *MinStack) Top() int {
	if this.buffer.Len() == 0 {
		return 0
	}

	delta := this.buffer.Back().Value.(int)
	if delta < 0 {
		return this.currentMin
	}
	return delta + this.currentMin
}

func (this *MinStack) GetMin() int {
	return this.currentMin
}
