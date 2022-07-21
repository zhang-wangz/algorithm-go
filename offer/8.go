package main

type CQueue struct {
	sk1 []int // 头
	sk2 []int // 尾
}

func Constructor() CQueue {
	return CQueue{sk1: make([]int, 0), sk2: make([]int, 0)}
}

func (this *CQueue) AppendTail(value int) {
	for len(this.sk2) != 0 {
		this.sk1 = append(this.sk1, this.sk2[len(this.sk2)-1])
		this.sk2 = this.sk2[:len(this.sk2)-1]
	}
	this.sk1 = append(this.sk1, value)
}

func (this *CQueue) DeleteHead() int {
	for len(this.sk1) != 0 {
		this.sk2 = append(this.sk2, this.sk1[len(this.sk1)-1])
		this.sk1 = this.sk1[:len(this.sk1)-1]
	}
	var top int
	if len(this.sk2) == 0 {
		return -1
	} else {
		top = this.sk2[len(this.sk2)-1]
		this.sk2 = this.sk2[:len(this.sk2)-1]
	}
	return top
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
