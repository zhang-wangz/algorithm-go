package main

import "container/heap"

// 1801 积压订单中的订单总数
func getNumberOfBacklogOrders(orders [][]int) (ans int) {
	const mod = 1e9 + 7
	buy, sell := hp2{}, hp2{}
	for _, order := range orders {
		price, a, o := order[0], order[1], order[2]
		// buy
		if o == 0 {
			for a > 0 && sell.Len() > 0 && sell[0].price <= price {
				s1 := heap.Pop(&sell).(p)
				if s1.amount < a {
					a -= s1.amount
					s1.amount = 0
				} else {
					s1.amount -= a
					a = 0
				}
				if s1.amount > 0 {
					heap.Push(&sell, p{s1.price, s1.amount})
				}
			}
			if a > 0 {
				heap.Push(&buy, p{-price, a})
			}
		} else {
			for a > 0 && buy.Len() > 0 && -buy[0].price >= price {
				b1 := heap.Pop(&buy).(p)
				if b1.amount < a {
					a -= b1.amount
					b1.amount = 0
				} else {
					b1.amount -= a
					a = 0
				}
				if b1.amount > 0 {
					heap.Push(&buy, p{b1.price, b1.amount})
				}
			}
			if a > 0 {
				heap.Push(&sell, p{price, a})
			}
		}
	}
	for sell.Len() > 0 {
		ans += heap.Pop(&sell).(p).amount
		ans %= mod
	}
	for buy.Len() > 0 {
		ans += heap.Pop(&buy).(p).amount
		ans %= mod
	}
	return
}

type p struct {
	price, amount int
}
type hp2 []p

func (h hp2) Len() int            { return len(h) }
func (h hp2) Less(i, j int) bool  { a, b := h[i], h[j]; return a.price < b.price }
func (h hp2) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp2) Push(v interface{}) { *h = append(*h, v.(p)) }
func (h *hp2) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
