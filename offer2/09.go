package main

type CQueue struct {
	a, b []int
}

func Constructor() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.a = append(this.a, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.a) != 0 {
		this.b = append(this.b, this.a...)
		this.a = []int{}
	}
	if len(this.b) != 0 {
		ans := this.b[0]
		this.b = this.b[1:]
		return ans
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
