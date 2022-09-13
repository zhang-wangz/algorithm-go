package main

import "strings"

// 2166 位集
type Bitset struct {
	arr           []int
	cnt, reversed int
}

func Constructor3(size int) Bitset {
	return Bitset{make([]int, size), 0, 0}
}

func (this *Bitset) Fix(idx int) {
	if this.arr[idx]^this.reversed == 0 {
		this.arr[idx] ^= 1
		this.cnt++
	}
}

func (this *Bitset) Unfix(idx int) {
	if this.arr[idx]^this.reversed == 1 {
		this.arr[idx] ^= 1
		this.cnt--
	}
}

func (this *Bitset) Flip() {
	this.reversed ^= 1
	this.cnt = len(this.arr) - this.cnt
}

func (this *Bitset) All() bool {
	return this.cnt == len(this.arr)
}

func (this *Bitset) One() bool {
	return this.cnt > 0
}

func (this *Bitset) Count() int {
	return this.cnt
}

func (this *Bitset) ToString() string {
	sb := strings.Builder{}
	for _, bit := range this.arr {
		sb.WriteByte(byte('0' + bit ^ this.reversed))
	}
	return sb.String()
}

/**
 * Your Bitset object will be instantiated and called as such:
 * obj := Constructor(size);
 * obj.Fix(idx);
 * obj.Unfix(idx);
 * obj.Flip();
 * param_4 := obj.All();
 * param_5 := obj.One();
 * param_6 := obj.Count();
 * param_7 := obj.ToString();
 */
