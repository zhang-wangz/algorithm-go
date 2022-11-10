package main

// 找出左边第一个大于自己的数字
// 单调栈
type StockSpanner struct {
	nums []int
	left []int
}

func Constructor4() StockSpanner {
	return StockSpanner{nums: []int{}, left: []int{}}
}

func (this *StockSpanner) Next(price int) (ans int) {
	w := 1
	p := this.nums
	l := this.left
	for len(p) > 0 && p[len(p)-1] <= price {
		p = p[:len(p)-1]
		w += l[len(l)-1]
		l = l[:len(l)-1]
	}
	p = append(p, price)
	l = append(l, w)
	this.nums = p
	this.left = l
	return w
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
